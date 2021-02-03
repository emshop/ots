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
	"github.com/micro-plat/qtask"
)

var notifyFields = []string{
	sql.FieldOrderID,
}

//Start 开始订单通知信息
func Start(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------获取订单通知信息----------------------")
	if err := ctx.Request().Check(notifyFields...); err != nil {
		return err
	}

	//获取订单通知数据
	notify, err := notify.Start(ctx.Request().GetString(sql.FieldOrderID))
	switch {
	case errs.GetCode(err) == http.StatusNoContent:
		qtask.FinishByInput(ctx, ctx.Request()) //当订单绑定完成后自动关闭流程
		return nil
	case err != nil:
		return err
	}
	return notify
}

//Success 保存成功通知结果
func Success(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存通知成功结果----------------------")
	if err := ctx.Request().Check(notifyFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 保存通知结果")
	err := notify.Save(ctx.Request().GetString(sql.FieldOrderID),
		enums.Success, ctx.Request().GetString(sql.FieldNotifyMsg))
	switch {
	case errs.GetCode(err) == http.StatusNoContent || err == nil:
		qtask.FinishByInput(ctx, ctx.Request()) //无须处理或处理成功则关闭流程
	}

	ctx.Log().Info("2. 执行后续流程")
	orders, err := orders.QueryByOrderID(ctx.Request().GetString(sql.FieldOrderID))
	if err != nil {
		return err
	}

	flw := flow.NextByFirst(enums.FlowFlagOrderNotify,
		orders.GetInt(sql.FieldPlID), enums.Success, ctx,
		sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID))

	return flw
}

//Unknown 保存失败通知结果
func Unknown(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存通知失败结果----------------------")
	if err := ctx.Request().Check(notifyFields...); err != nil {
		return err
	}

	ctx.Log().Infof("1. 保存通知结果为未知%s", ctx.Request().GetString(sql.FieldOrderID))
	err := notify.Save(ctx.Request().GetString(sql.FieldOrderID),
		enums.Unknown, ctx.Request().GetString(sql.FieldNotifyMsg))

	return err
}
