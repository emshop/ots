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

//TimeoutHandle 支付超时(异步流程)
func TimeoutHandle(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理支付超时----------------------")
	if err := ctx.Request().Check(fields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 订单超时则关闭支付")
	qtask.ProcessingByInput(ctx, ctx.Request())
	err := payment.Timeout(ctx.Request().GetString(sql.FieldOrderID))
	switch {
	case errs.GetCode(err) == http.StatusNoContent:
		qtask.FinishByInput(ctx, ctx.Request())
		return nil
	case err == nil:
		qtask.FinishByInput(ctx, ctx.Request()) //无须处理或处理成功则关闭流程
	}

	ctx.Log().Info("2. 获取超时后续流程")
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
