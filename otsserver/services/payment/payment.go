package payment

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

var fields = []string{
	sql.FieldFlowID,
	sql.FieldOrderID,
}

//Payment 订单支付流程
type Payment struct {
}

//PayHandle 处理订单支付(异步流程)
func (p *Payment) PayHandle(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理订单支付----------------------")
	if err := ctx.Request().Check(fields...); err != nil {
		return err
	}

	ctx.Log().Infof("1. 处理订单支付(%d)", ctx.Request().GetInt64(sql.FieldOrderID))
	err := payment.Pay(ctx.Request().GetInt64(sql.FieldOrderID))
	if errs.GetCode(err) == http.StatusNoContent {
		qtask.FinishByInput(ctx, ctx.Request())
		return err
	}
	if err != nil {
		return err
	}

	ctx.Log().Info("2. 获取支付后续流程")
	status := types.DecodeInt(err, nil, enums.Success, enums.Failed)
	flw, err := flow.Next(ctx.Request().GetInt(sql.FieldFlowID),
		enums.FlowStatus(status), ctx,
		sql.FieldOrderID,
		ctx.Request().GetString(sql.FieldOrderID))
	if err != nil {
		ctx.Log().Error("执行支付后续流程失败", err)
		return err
	}
	return flw
}
