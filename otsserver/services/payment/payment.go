package payment

import (
	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/emshop/ots/otsserver/modules/payment"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask"
)

var fields = []string{
	sql.FieldFlowID,
	sql.FieldOrderID,
}

//Paying 处理订单支付(异步流程)
func Paying(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理订单支付----------------------")
	if err := ctx.Request().Check(fields...); err != nil {
		return err
	}

	ctx.Log().Infof("1. 处理订单支付(%s)", ctx.Request().GetString(sql.FieldOrderID))
	qtask.ProcessingByInput(ctx, ctx.Request())
	err := payment.Pay(ctx.Request().GetString(sql.FieldOrderID))
	switch {
	case errs.GetStop(err):
		qtask.FinishByInput(ctx, ctx.Request())
		return err
	case err == nil:
		qtask.FinishByInput(ctx, ctx.Request())
	}

	ctx.Log().Info("2. 获取支付后续流程", err)
	status := types.DecodeInt(err, nil, int(enums.Success), int(enums.Failed))
	flw := flow.Next(ctx.Request().GetInt(sql.FieldFlowID),
		enums.FlowStatus(status), ctx,
		sql.FieldOrderID,
		ctx.Request().GetString(sql.FieldOrderID))

	if err != nil {
		return err
	}

	return flw
}
