package notify

import (
	"github.com/emshop/ots/flowserver/modules/biz/notices"
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/flows"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
)

var notifyFields = []string{
	fields.FieldOrderID,
}

//GetNotify 获取订单通知信息
func GetNotify(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------获取订单通知信息----------------------")
	if err := ctx.Request().Check(notifyFields...); err != nil {
		return err
	}

	//获取订单通知数据
	qtask.ProcessingByInput(ctx, ctx.Request())
	notify, err := notices.Start(ctx.Request().GetString(fields.FieldOrderID))
	switch {
	case errs.NeedStop(err):
		qtask.FinishByInput(ctx, ctx.Request()) //当订单绑定完成后自动关闭流程
		flows.NextByOrderNO(ctx.Request().GetString(fields.FieldOrderID), enums.FlowFinishStart, ctx)
		return err
	case err != nil:
		return err
	}
	return notify
}

//SaveSuccess 保存成功通知结果
func SaveSuccess(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存通知成功结果----------------------")
	if err := ctx.Request().Check(notifyFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 保存通知结果")
	order, err := notices.Save(ctx.Request().GetString(fields.FieldOrderID),
		enums.Success, ctx.Request().GetString(fields.FieldNotifyMsg))

	switch {
	case errs.NeedStop(err):
		qtask.FinishByInput(ctx, ctx.Request())
		return err
	case err != nil:
		return err
	}
	return flows.NextByOrderNO(order.GetString(fields.FieldOrderID), enums.FlowFinishStart, ctx)
}

//Unknown 保存失败通知结果
func Unknown(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存通知失败结果----------------------")
	if err := ctx.Request().Check(notifyFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 保存通知结果")
	order, err := notices.Save(ctx.Request().GetString(fields.FieldOrderID),
		enums.Unknown, ctx.Request().GetString(fields.FieldNotifyMsg))

	switch {
	case errs.NeedStop(err):
		qtask.FinishByInput(ctx, ctx.Request())
		return err
	case err != nil:
		return err
	}
	return flows.NextByOrderNO(order.GetString(fields.FieldOrderID), enums.FlowFinishStart, ctx)
}
