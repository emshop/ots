package finish

import (
	"github.com/emshop/ots/flowserver/modules/biz/finishes"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/qtask"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

//Finish 完结订单处理
func Finish(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Debug("-------------处理订单完结---------------------")
	if err := ctx.Request().Check(fields.FieldOrderID); err != nil {
		return err
	}
	qtask.ProcessingByInput(ctx, ctx.Request())
	defer lcs.New(ctx.Log(), "订单完结", ctx.Request().GetString(fields.FieldOrderID)).Start("结束订单").End(&r)
	err := finishes.Finish(ctx.Request().GetString(fields.FieldOrderID))
	if err != nil {
		return err
	}
	qtask.FinishByInput(ctx, ctx.Request())
	return "success"
}
