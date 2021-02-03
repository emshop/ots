package notify

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/emshop/ots/otsserver/modules/notify"
	"github.com/emshop/ots/otsserver/modules/orders"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

var notifyFields = []string{
	sql.FieldOrderID,
}

//StartHandle 开始订单通知信息
func StartHandle(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------开始订单通知----------------------")
	if err := ctx.Request().Check(notifyFields...); err != nil {
		return err
	}
	ctx.Log().Info("1. 启动订单通知")
	notify, err := notify.Start(ctx.Request().GetString(sql.FieldOrderID))
	if errs.GetCode(err) == http.StatusNoContent {
		ctx.Log().Info("2. 执行后续流程")
		flw, err := flow.Next(ctx.Request().GetInt(sql.FieldFlowID), enums.Success, ctx,
			sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID))
		if err != nil {
			ctx.Log().Error("2.1 执行绑定后续流程失败", err)
		}
		return flw
	}
	if err != nil {
		return err
	}
	return notify
}

//SuccessHandle 保存成功通知结果
func SuccessHandle(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存通知成功结果----------------------")
	if err := ctx.Request().Check(notifyFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 保存通知结果")
	err := notify.Save(ctx.Request().GetString(sql.FieldOrderID),
		enums.Success, ctx.Request().GetString(sql.FieldNotifyMsg))
	if err != nil {
		return err
	}

	ctx.Log().Info("2. 执行后续流程")
	orders, err := orders.QueryByOrderID(ctx.Request().GetString(sql.FieldOrderID))
	flw, err1 := flow.NextByFirst(enums.FlowFlagOrderNotify,
		orders.GetInt(sql.FieldPlID), enums.Success, ctx,
		sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID))
	if err1 != nil {
		ctx.Log().Error("2.1 执行绑定后续流程失败", err)
	}
	if err != nil {
		return err
	}
	return flw
}

//FailedHandle 保存失败通知结果
func FailedHandle(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存通知失败结果----------------------")
	if err := ctx.Request().Check(notifyFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 保存通知结果")
	err := notify.Save(ctx.Request().GetString(sql.FieldOrderID),
		enums.Success, ctx.Request().GetString(sql.FieldNotifyMsg))
	if err != nil {
		return err
	}

	ctx.Log().Info("2. 执行后续流程")
	if errs.GetCode(err) == http.StatusNoContent {
		flw, err := flow.Next(ctx.Request().GetInt(sql.FieldFlowID), enums.Success, ctx,
			sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID))
		if err != nil {
			ctx.Log().Error("2.1 执行绑定后续流程失败", err)
		}
		return flw
	}

	flw, err := flow.Next(ctx.Request().GetInt(sql.FieldFlowID), enums.Failed, ctx,
		sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID))
	if err != nil {
		ctx.Log().Error("2.1 执行绑定后续流程失败", err)
	}
	return flw
}
