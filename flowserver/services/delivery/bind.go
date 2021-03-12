package delivery

import (
	"github.com/emshop/ots/flowserver/modules/biz/deliverys"
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/flows"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

var requireFields = []string{
	fields.FieldOrderID,
}

//Binding 处理订单绑定
func Binding(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Info("-------------处理订单绑定----------------------")
	if err := ctx.Request().Check(requireFields...); err != nil {
		return err
	}

	ctx.Log().Infof("1. 处理订单绑定(%s)", ctx.Request().GetString(fields.FieldOrderID))
	defer lcs.New(ctx.Log(), "发货绑定", ctx.Request().GetString(fields.FieldOrderID)).Start("绑定").End(r)
	qtask.ProcessingByInput(ctx, ctx.Request())
	deliveryID, err := deliverys.Bind(ctx.Request().GetString(fields.FieldOrderID))
	switch {
	case errs.NeedStop(err):
		qtask.FinishByInput(ctx, ctx.Request()) //当订单绑定完成后自动关闭流程
		return nil
	case err != nil:
		return err
	}

	ctx.Log().Info("2. 执行后续流程")
	flw := flows.NextByDeliveryID(deliveryID, enums.FlowDelivery, ctx)
	return flw
}
