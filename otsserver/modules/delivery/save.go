package delivery

import (
	"fmt"
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/spp"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//SaveQueryResult 保存发货查询获得的结果
func SaveQueryResult(deliveryID int64, resultCode string, returnMsg string, discount types.Decimal, extParams string) (enums.FlowStatus, error) {
	return Save(deliveryID, enums.ResultFromQuery, resultCode, returnMsg, discount, extParams)
}

//SaveNotifyResult 保存供货商通知的结果
func SaveNotifyResult(deliveryID int64, resultCode string, returnMsg string, discount types.Decimal, extParams string) (enums.FlowStatus, error) {
	return Save(deliveryID, enums.ResultFromNotify, resultCode, returnMsg, discount, extParams)
}

//SaveDeliveryResult 保存同步发货返回的结果
func SaveDeliveryResult(deliveryID int64, resultCode string, returnMsg string, discount types.Decimal, extParams string) (enums.FlowStatus, error) {
	return Save(deliveryID, enums.ResultFromDelivery, resultCode, returnMsg, discount, extParams)
}

//SaveAuditResult 保存人工处理的结果
func SaveAuditResult(deliveryID int64, resultCode string, returnMsg string, discount types.Decimal, extParams string) (enums.FlowStatus, error) {
	return Save(deliveryID, enums.ResultAudit, resultCode, returnMsg, discount, extParams)
}

//Save 保存发货结果
func Save(deliveryID int64, source enums.ResultSource, resultCode string, returnMsg string, discount types.Decimal, extParams string) (enums.FlowStatus, error) {

	//查询发货记录
	delivery, err := Get(deliveryID)
	if err != nil {
		return enums.Unknown, err
	}

	//获取处理结果
	result, msg := spp.GetDealCode(
		delivery.GetString(sql.FieldSppNo),
		source,
		delivery.GetInt(sql.FieldPlID),
		resultCode)
	returnMsg = types.GetString(returnMsg, msg)
	fmt.Println("code:", result, msg)
	switch result {
	case enums.Failed:
		err := SaveFailed(
			deliveryID,
			delivery.GetInt64(sql.FieldOrderID),
			delivery.GetInt(sql.FieldTotalFace),
			source.Format(resultCode), returnMsg)
		if err != nil {
			return enums.Unknown, err
		}
		return enums.Failed, nil
	case enums.Success:
		err := SaveSuccess(
			deliveryID,
			discount,
			extParams,
			source.Format(resultCode), returnMsg)
		if err != nil {
			return enums.Unknown, err
		}
		return enums.Success, nil
	default: //结果未知处理
		err := SaveUnknown(deliveryID, source.Format(resultCode), returnMsg)
		return enums.Unknown, err
	}
}

//SaveUnknown 保存未知的发货结果
func SaveUnknown(deliveryID int64, code string, msg string) error {
	row, err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateTradeDeliveryForSaveUnknown, map[string]interface{}{
		sql.FieldDeliveryID: deliveryID,
		sql.FieldResultCode: code,
		sql.FieldReturnMsg:  msg,
	})
	if err != nil {
		return err
	}
	if row == 0 {
		return errs.NewError(http.StatusNoContent, "发货无须处理(save.unknown)")
	}
	return nil
}

//SaveSuccess 保存成功发货结果
func SaveSuccess(deliveryID int64, discount types.Decimal, params string, code string, msg string) error {

	//启动事务修改发货记录、交易订单
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}
	row, err := db.Execute(sql.UpdateTradeDeliveryForSaveSuccess, map[string]interface{}{
		sql.FieldDeliveryID:    deliveryID,
		sql.FieldResultCode:    code,
		sql.FieldReturnMsg:     msg,
		sql.FieldCostDiscount:  discount.String(),
		sql.FieldRequestParams: params,
	})
	if err != nil {
		db.Rollback()
		return err
	}
	if row == 0 {
		db.Rollback()
		return errs.NewError(http.StatusNoContent, "发货无须处理(save.success)")
	}

	data, err := db.Query(sql.SelectTradeDelivery, map[string]interface{}{
		sql.FieldDeliveryID: deliveryID,
	})
	if err != nil {
		db.Rollback()
		return err
	}
	if data.Len() == 0 {
		db.Rollback()
		return fmt.Errorf("查询发货订单失败")
	}

	//修改绑定金额
	row, err = db.Execute(sql.UpdateTradeOrderForDeliverySuccess, map[string]interface{}{
		sql.FieldOrderID:          data.Get(0).GetString(sql.FieldOrderID),
		sql.FieldCostAmount:       data.Get(0).GetString(sql.FieldCostAmount),
		sql.FieldSuccFace:         data.Get(0).GetString(sql.FieldSuccTotalFace),
		sql.FieldSppFeeAmount:     data.Get(0).GetString(sql.FieldSppFeeAmount),
		sql.FieldPaymentFeeAmount: data.Get(0).GetString(sql.FieldPaymentFeeAmount),
		sql.FieldTradeFeeAmount:   data.Get(0).GetString(sql.FieldTradeFeeAmount),
	})
	if err != nil {
		db.Rollback()
		return err
	}
	if row == 0 {
		db.Rollback()
		return fmt.Errorf("无法修改订单(%d)为发货成功", deliveryID)
	}
	db.Commit()
	return nil
}

//SaveFailed 保存失败发货结果
func SaveFailed(deliveryID int64, orderID int64, totalFace int, code string, msg string) error {
	//启动事务修改发货记录、交易订单
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}

	row, err := db.Execute(sql.UpdateTradeDeliveryForFailed, map[string]interface{}{
		sql.FieldDeliveryID: deliveryID,
		sql.FieldResultCode: code,
		sql.FieldReturnMsg:  msg,
	})
	if err != nil {
		db.Rollback()
		return err
	}
	if row == 0 {
		db.Rollback()
		return errs.NewError(http.StatusNoContent, "发货无须处理(save.failed)")
	}

	//修改绑定金额
	row, err = db.Execute(sql.UpdateTradeOrderForDeiveryFailed, map[string]interface{}{
		sql.FieldOrderID: orderID,
		sql.FieldFace:    totalFace,
	})
	if err != nil {
		db.Rollback()
		return err
	}
	if row == 0 {
		db.Rollback()
		return fmt.Errorf("修改订单(%d)绑定金额出错", orderID)
	}
	db.Commit()
	return nil
}
