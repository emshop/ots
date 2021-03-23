package deliverys

import (
	"fmt"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//SaveQueryResult 保存发货查询获得的结果
func SaveQueryResult(deliveryID string, resultCode string, returnMsg string, discount types.Decimal, extParams string) (enums.FlowStatus, types.IXMap, error) {
	return Save(deliveryID, enums.ResultFromQuery, resultCode, returnMsg, discount, extParams)
}

//SaveNotifyResult 保存供货商通知的结果
func SaveNotifyResult(deliveryID string, resultCode string, returnMsg string, discount types.Decimal, extParams string) (enums.FlowStatus, types.IXMap, error) {
	return Save(deliveryID, enums.ResultFromNotify, resultCode, returnMsg, discount, extParams)
}

//SaveDeliveryResult 保存同步发货返回的结果
func SaveDeliveryResult(deliveryID string, resultCode string, returnMsg string, discount types.Decimal, extParams string) (enums.FlowStatus, types.IXMap, error) {
	return Save(deliveryID, enums.ResultFromDelivery, resultCode, returnMsg, discount, extParams)
}

//SaveAuditResult 保存人工处理的结果
func SaveAuditResult(deliveryID string, resultCode string, returnMsg string, discount types.Decimal, extParams string) (enums.FlowStatus, types.IXMap, error) {
	return Save(deliveryID, enums.ResultAudit, resultCode, returnMsg, discount, extParams)
}

//Save 保存发货结果
func Save(deliveryID string, source enums.ResultSource, resultCode string, returnMsg string, discount types.Decimal, extParams string) (enums.FlowStatus, types.IXMap, error) {

	//获取处理结果
	result, msg := GetDealCode(deliveryID, source, resultCode)
	returnMsg = fmt.Sprintf("%s(%s)", returnMsg, msg)
	switch result {
	case enums.Failed:
		r, err := SaveFailed(
			deliveryID,
			source,
			source.Format(resultCode), returnMsg)
		if err != nil {
			return enums.Unknown, r, err
		}
		return enums.Failed, r, nil
	case enums.Success:
		r, err := SaveSuccess(
			deliveryID,
			discount,
			source,
			extParams,
			source.Format(resultCode), returnMsg)
		if err != nil {
			return enums.Unknown, r, err
		}
		return enums.Success, r, nil
	default: //结果未知处理
		r, err := SaveUnknown(deliveryID, source.Format(resultCode), returnMsg)
		return enums.Unknown, r, err
	}
}

//SaveUnknown 保存未知的发货结果
func SaveUnknown(deliveryID string, code string, msg string) (types.IXMap, error) {
	db := hydra.C.DB().GetRegularDB()
	r, err := db.ExecuteBatch(updateDeliveryForSaveUnknown, types.XMap{
		fields.FieldDeliveryID: deliveryID,
		fields.FieldResultCode: code,
		fields.FieldReturnMsg:  msg,
	})
	if err != nil {
		return nil, err
	}
	return r.Get(0), nil
}

//SaveSuccess 保存成功发货结果
func SaveSuccess(deliveryID string, discount types.Decimal, source enums.ResultSource, params string, code string, msg string) (types.IXMap, error) {

	input := types.XMap{
		fields.FieldDeliveryID:    deliveryID,
		fields.FieldResultCode:    code,
		fields.FieldResultSource:  source,
		fields.FieldReturnMsg:     msg,
		fields.FieldCostDiscount:  discount.String(),
		fields.FieldRequestParams: params,
	}

	//启动事务修改发货记录、交易订单
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}

	r, err := db.ExecuteBatch(updateForDeliverySuccess, input)
	if err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return r.Get(0), nil
}

//SaveFailed 保存失败发货结果
func SaveFailed(deliveryID string, source enums.ResultSource, code string, msg string) (types.IXMap, error) {
	input := types.XMap{
		fields.FieldDeliveryID:   deliveryID,
		fields.FieldResultCode:   code,
		fields.FieldResultSource: source,
		fields.FieldReturnMsg:    msg,
	}
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}
	r, err := db.ExecuteBatch(updateForDeliveryFailed, input)
	if err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return r.Get(0), nil
}
