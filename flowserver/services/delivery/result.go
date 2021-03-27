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

	ctx.Log().Debug("-------------保存发货结果----------------------")
	if err := ctx.Request().Check(deliverySaveResultFields...); err != nil {
		return err
	}

	ctx.Log().Debug("1. 保存发货结果")
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
		ctx.Log().Debug("2. 执行后续流程")
		flows.NextByDeliveryID(ctx.Request().GetString(fields.FieldDeliveryID), enums.FlowDeliveryPay, ctx, fields.FieldOrderID, r.GetString(fields.FieldOrderID))
		flows.NextByOrderNO(r.GetString(fields.FieldOrderID), enums.FlowNotifyStart, ctx)
		return "success"
	}
	if s == enums.Failed {
		ctx.Log().Debug("2. 执行后续流程")
		flows.NextByOrderNO(r.GetString(fields.FieldOrderID), enums.FlowDeliveryBind, ctx)
		return "success"
	}
	return "success"
}

//SaveQueryResult 保存查询结果
func SaveQueryResult(ctx hydra.IContext) interface{} {

	ctx.Log().Debug("-------------保存查询结果----------------------")
	if err := ctx.Request().Check(deliverySaveResultFields...); err != nil {
		return err
	}

	ctx.Log().Debug("1. 保存发货结果")
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
		ctx.Log().Debug("2. 执行后续流程")
		flows.NextByDeliveryID(ctx.Request().GetString(fields.FieldDeliveryID), enums.FlowDeliveryPay, ctx, fields.FieldOrderID, r.GetString(fields.FieldOrderID))
		flows.NextByOrderNO(r.GetString(fields.FieldOrderID), enums.FlowNotifyStart, ctx)
		return "success"
	}
	if s == enums.Failed {
		ctx.Log().Debug("2. 执行后续流程")
		flows.NextByOrderNO(r.GetString(fields.FieldOrderID), enums.FlowDeliveryBind, ctx)
		return "success"
	}
	return "success"
}

//SaveNotifyResult 保存通知结果
func SaveNotifyResult(ctx hydra.IContext) interface{} {

	ctx.Log().Debug("-------------保存通知结果----------------------")
	if err := ctx.Request().Check(deliverySaveResultFields...); err != nil {
		return err
	}

	ctx.Log().Debug("1. 保存发货结果")
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
		ctx.Log().Debug("2. 执行后续流程")
		flows.NextByDeliveryID(ctx.Request().GetString(fields.FieldDeliveryID), enums.FlowDeliveryPay, ctx, fields.FieldOrderID, r.GetString(fields.FieldOrderID))
		flows.NextByOrderNO(r.GetString(fields.FieldOrderID), enums.FlowNotifyStart, ctx)
		return "success"
	}
	if s == enums.Failed {
		ctx.Log().Debug("2. 执行后续流程")
		flows.NextByOrderNO(r.GetString(fields.FieldOrderID), enums.FlowDeliveryBind, ctx)
		return "success"
	}
	return "success"
}

//SaveAuditResult 保存人工审核结果
func SaveAuditResult(ctx hydra.IContext) interface{} {

	ctx.Log().Debug("-------------保存人工审核结果----------------------")
	if err := ctx.Request().Check(deliverySaveResultFields...); err != nil {
		return err
	}

	ctx.Log().Debug("1. 保存发货结果")
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
		ctx.Log().Debug("2. 执行后续流程")
		flows.NextByDeliveryID(ctx.Request().GetString(fields.FieldDeliveryID), enums.FlowDeliveryPay, ctx, fields.FieldOrderID, r.GetString(fields.FieldOrderID))
		flows.NextByOrderNO(r.GetString(fields.FieldOrderID), enums.FlowNotifyStart, ctx)
		return "success"
	}
	if s == enums.Failed {
		ctx.Log().Debug("2. 执行后续流程")
		flows.NextByOrderNO(r.GetString(fields.FieldOrderID), enums.FlowDeliveryBind, ctx)
		return "success"
	}
	return "success"
}
