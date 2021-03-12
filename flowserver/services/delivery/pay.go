package delivery

import (
	"net/http"

	"github.com/emshop/ots/flowserver/modules/biz/deliverys"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

//Paying 上游订单支付记账
func Paying(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------------处理发货支付----------------------")
	if err := ctx.Request().Check(fields.FieldDeliveryID); err != nil {
		return err
	}
	defer lcs.New(ctx.Log(), "发货支付", ctx.Request().GetString(fields.FieldOrderID)).Start("支付").End(r)
	qtask.ProcessingByInput(ctx, ctx.Request())
	err := deliverys.Pay(ctx.Request().GetString(fields.FieldDeliveryID))
	if err != nil && errs.GetCode(err) != http.StatusNoContent {
		return err
	}
	qtask.FinishByInput(ctx, ctx.Request())
	return err
}
