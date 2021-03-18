package delivery

import (
	"errors"

	"github.com/emshop/ots/flowserver/modules/biz/deliverys"
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/flows"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/qtask"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

var requireFields = []string{
	fields.FieldOrderID,
}

//Binding 处理订单绑定
func Binding(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Debug("-------------处理订单绑定----------------------")
	if err := ctx.Request().Check(requireFields...); err != nil {
		return err
	}

	ctx.Log().Debugf("1. 处理订单绑定(%s)", ctx.Request().GetString(fields.FieldOrderID))
	defer lcs.New(ctx.Log(), "发货绑定", ctx.Request().GetString(fields.FieldOrderID)).Start("绑定").End(&r)
	qtask.ProcessingByInput(ctx, ctx.Request())
	deliveryID, err := deliverys.Bind(ctx.Request().GetString(fields.FieldOrderID))
	switch {
	case errors.Is(err, xerr.ErrNOTEXISTS): //订单无须绑定
		qtask.FinishByInput(ctx, ctx.Request())
		return err
	case errors.Is(err, xerr.ErrOrderTimeout):
		qtask.FinishByInput(ctx, ctx.Request())
		flows.NextByOrderNO(ctx.Request().GetString(fields.FieldOrderID), enums.FlowNotifyStart, ctx)
		return err
	case err != nil:
		return err
	}

	ctx.Log().Debug("2. 执行后续流程")
	flows.NextByDeliveryID(deliveryID, enums.FlowDelivery, ctx, fields.FieldOrderID, ctx.Request().GetString(fields.FieldOrderID))
	return "success"
}
