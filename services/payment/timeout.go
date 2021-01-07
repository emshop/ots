package payment

import (
	"net/http"

	"github.com/emshop/ots/modules/const/sql"
	"github.com/emshop/ots/modules/payment"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
)

//TimeoutHandle 支付超时(异步流程)
func (p *Payment) TimeoutHandle(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理支付超时----------------------")
	if err := ctx.Request().Check(fields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 订单超时则关闭支付")
	err := payment.Timeout(ctx.Request().GetInt64(sql.FieldOrderID))
	if errs.GetCode(err) == http.StatusNoContent {
		qtask.FinishByInput(ctx, ctx.Request())
	}
	return err
}
