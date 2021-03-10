package delivery

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
)

//MockHandle 上游付款
func MockHandle(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------打桩发货处理----------------------")
	ctx.Log().Info("1. 获取发货信息")
	qtask.ProcessingByInput(ctx, ctx.Request())
	rpns := GetRequest(ctx)
	if err := errs.GetError(rpns); err != nil {
		return err

	}

	ctx.Log().Info("2. 保存为已提交到上游")
	rpns = SaveRequest(ctx)
	if err := errs.GetError(rpns); err != nil {
		return err
	}

	ctx.Log().Info("3. 保存为发货成功")
	rpns = SaveDeliveryResult(ctx)
	if err := errs.GetError(rpns); err != nil {
		return err
	}
	qtask.FinishByInput(ctx, ctx.Request())
	return rpns
}
