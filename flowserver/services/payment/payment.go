package payment

import (
	"github.com/emshop/ots/flowserver/modules/biz/payments"
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/flows"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
)

var requireFields = []string{
	fields.FieldOrderID,
}

//Paying 处理订单支付(异步流程)
func Paying(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理订单支付----------------------")
	if err := ctx.Request().Check(requireFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 处理订单支付")
	qtask.ProcessingByInput(ctx, ctx.Request())
	err := payments.Pay(ctx.Request().GetString(fields.FieldOrderID))
	switch {
	case errs.NeedStop(err): //订单不符合支付条件,不能进行支付
		qtask.FinishByInput(ctx, ctx.Request())
		return err
	case err != nil: //支付过程出错
		return err
	case err == nil: //支付成功
		qtask.FinishByInput(ctx, ctx.Request())
	}

	ctx.Log().Info("2. 处理订单后续流程")
	flw := flows.NextByOrderNO(ctx.Request().GetString(fields.FieldOrderID), enums.FlowDeliveryBind, ctx)
	return flw
}
