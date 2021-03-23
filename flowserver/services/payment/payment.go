package payment

import (
	"errors"

	"github.com/emshop/ots/flowserver/modules/biz/payments"
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/flows"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

var requireFields = []string{
	fields.FieldOrderID,
}

//Paying 处理订单支付(异步流程)
func Paying(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Debug("-------------处理订单支付----------------------")
	if err := ctx.Request().Check(requireFields...); err != nil {
		return err
	}

	ctx.Log().Debug("1. 处理订单支付")
	defer lcs.New(ctx.Log(), "订单支付", ctx.Request().GetString(fields.FieldOrderID)).Start("支付").End(&r)
	qtask.ProcessingByInput(ctx, ctx.Request())
	err := payments.Pay(ctx.Request().GetString(fields.FieldOrderID))
	switch {
	case errors.Is(err, errs.ErrNotExist): //不存在，订单状态错误可能已完成处理
		qtask.FinishByInput(ctx, ctx.Request())
		return err
	case errors.Is(err, xerr.ErrOrderTimeout): //订单超时，关闭流程进行后续处理
		qtask.FinishByInput(ctx, ctx.Request())
		flows.NextByOrderNO(ctx.Request().GetString(fields.FieldOrderID), enums.FlowNotifyStart, ctx)
		return err
	case err != nil: //支付过程出错
		return err
	case err == nil: //支付成功
		qtask.FinishByInput(ctx, ctx.Request())
	}

	ctx.Log().Debug("2. 处理订单后续流程")
	flows.NextByOrderNO(ctx.Request().GetString(fields.FieldOrderID), enums.FlowDeliveryBind, ctx)
	return "success"
}
