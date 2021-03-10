package delivery

import (
	"github.com/emshop/ots/flowserver/modules/biz/deliverys"
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/flows"
	"github.com/micro-plat/hydra"
)

var deliverySaveResultFields = []string{
	fields.FieldDeliveryID,
	fields.FieldResultCode,
	fields.FieldReturnMsg,
}

//SaveDeliveryResult 开始请求
func SaveDeliveryResult(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存发货结果----------------------")
	if err := ctx.Request().Check(deliverySaveResultFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 保存发货结果")
	s, r, err := deliverys.SaveDeliveryResult(
		ctx.Request().GetString(fields.FieldDeliveryID),
		ctx.Request().GetString(fields.FieldResultCode),
		ctx.Request().GetString(fields.FieldReturnMsg),
		ctx.Request().GetDecimal(fields.FieldRealDiscount),
		ctx.Request().GetString(fields.FieldRequestParams),
	)
	if err != nil {
		return err
	}
	if s == enums.Success && r.GetInt(fields.FieldOrderStatus) == int(enums.OrderNotify) {
		ctx.Log().Info("2. 执行后续流程")
		flw := flows.NextByOrderNO(r.GetString(fields.FieldOrderID), enums.FlowOrderNotifyStart, ctx)
		return flw
	}
	return nil
}

//SaveQueryResult 保存查询结果
func SaveQueryResult(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存查询结果----------------------")
	if err := ctx.Request().Check(deliverySaveResultFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 保存发货结果")
	s, r, err := deliverys.SaveQueryResult(
		ctx.Request().GetString(fields.FieldDeliveryID),
		ctx.Request().GetString(fields.FieldResultCode),
		ctx.Request().GetString(fields.FieldReturnMsg),
		ctx.Request().GetDecimal(fields.FieldRealDiscount),
		ctx.Request().GetString(fields.FieldRequestParams),
	)
	if err != nil {
		return err
	}

	if s == enums.Success && r.GetInt(fields.FieldOrderStatus) == int(enums.OrderNotify) {
		ctx.Log().Info("2. 执行后续流程")
		flw := flows.NextByOrderNO(r.GetString(fields.FieldOrderID), enums.FlowOrderNotifyStart, ctx)
		return flw
	}
	return nil
}

//SaveNotifyResult 保存通知结果
func SaveNotifyResult(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存通知结果----------------------")
	if err := ctx.Request().Check(deliverySaveResultFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 保存发货结果")
	s, r, err := deliverys.SaveNotifyResult(
		ctx.Request().GetString(fields.FieldDeliveryID),
		ctx.Request().GetString(fields.FieldResultCode),
		ctx.Request().GetString(fields.FieldReturnMsg),
		ctx.Request().GetDecimal(fields.FieldRealDiscount),
		ctx.Request().GetString(fields.FieldRequestParams),
	)
	if err != nil {
		return err
	}

	if s == enums.Success && r.GetInt(fields.FieldOrderStatus) == int(enums.OrderNotify) {
		ctx.Log().Info("2. 执行后续流程")
		flw := flows.NextByOrderNO(r.GetString(fields.FieldOrderID), enums.FlowOrderNotifyStart, ctx)
		return flw
	}
	return nil
}

//SaveAuditResult 保存人工审核结果
func SaveAuditResult(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存人工审核结果----------------------")
	if err := ctx.Request().Check(deliverySaveResultFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 保存发货结果")
	s, r, err := deliverys.SaveAuditResult(
		ctx.Request().GetString(fields.FieldDeliveryID),
		ctx.Request().GetString(fields.FieldResultCode),
		ctx.Request().GetString(fields.FieldReturnMsg),
		ctx.Request().GetDecimal(fields.FieldRealDiscount),
		ctx.Request().GetString(fields.FieldRequestParams),
	)
	if err != nil {
		return err
	}

	if s == enums.Success && r.GetInt(fields.FieldOrderStatus) == int(enums.OrderNotify) {
		ctx.Log().Info("2. 执行后续流程")
		flw := flows.NextByOrderNO(r.GetString(fields.FieldOrderID), enums.FlowOrderNotifyStart, ctx)
		return flw
	}
	return nil
}