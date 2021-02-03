package order

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/emshop/ots/otsserver/modules/notify"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
)

var fields = []string{
	sql.FieldFlowID,
	sql.FieldOrderID,
}

//Notify 订单通知
func Notify(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理订单通知----------------------")
	if err := ctx.Request().Check(fields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 开始订单通知")
	qtask.ProcessingByInput(ctx, ctx.Request())
	ntf, err := notify.Start(ctx.Request().GetString(sql.FieldOrderID))
	switch {
	case errs.GetCode(err) == http.StatusNoContent:
		qtask.FinishByInput(ctx, ctx.Request())
	case err == nil:
		//处理通知
		ctx.Log().Info("1.1 下游订单通知", ntf)

		ctx.Log().Info("1.2. 保存通知结果")
		err = notify.Save(ctx.Request().GetString(sql.FieldOrderID), enums.Success, "sucess")
		if err != nil {
			return err
		}
	default:
		return err
	}

	ctx.Log().Info("3. 获取通知后续流程")
	qtask.FinishByInput(ctx, ctx.Request())
	flw := flow.Next(ctx.Request().GetInt(sql.FieldFlowID),
		enums.Success, ctx,
		sql.FieldOrderID,
		ctx.Request().GetString(sql.FieldOrderID))
	return flw
}
