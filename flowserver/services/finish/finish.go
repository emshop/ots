package finish

import (
	"github.com/emshop/ots/flowserver/modules/biz/finishes"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/qtask"
)

//Finish 完结订单处理
func Finish(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理订单完结---------------------")
	if err := ctx.Request().Check(fields.FieldOrderID); err != nil {
		return err
	}
	qtask.ProcessingByInput(ctx, ctx.Request())
	err := finishes.Finish(ctx.Request().GetString(fields.FieldOrderID))
	if err != nil {
		return err
	}
	qtask.FinishByInput(ctx, ctx.Request())
	return nil
}
