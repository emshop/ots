package order

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/emshop/ots/otsserver/modules/payment"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask"
)

//Tiemout 订单完结处理
func Tiemout(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------处理订单超时关闭----------------------")
	if err := ctx.Request().Check(finishFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 处理订单完结")
	qtask.ProcessingByInput(ctx, ctx.Request())
	err := payment.Refund(ctx.Request().GetString(sql.FieldOrderID))
	switch {
	case errs.GetCode(err) == http.StatusAccepted:
		return err
	case errs.GetCode(err) == http.StatusNoContent:
		qtask.FinishByInput(ctx, ctx.Request()) //无须处理或处理成功则关闭流程
		return err
	case err == nil:
		qtask.FinishByInput(ctx, ctx.Request()) //无须处理或处理成功则关闭流程
	}

	ctx.Log().Info("2. 处理后续流程")
	status := types.DecodeInt(err, nil, int(enums.Success), int(enums.Failed))
	flw := flow.Next(ctx.Request().GetInt(sql.FieldFlowID), enums.FlowStatus(status), ctx,
		sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID))
	return flw
}
