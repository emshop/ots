package notify

import (
	"errors"

	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

//Notify 获取订单通知信息
func Notify(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Debug("-------------处理下游订单通知----------------------")
	if err := ctx.Request().Check(notifyFields...); err != nil {
		return err
	}
	ctx.Log().Debug("1. 获取通知信息")
	defer lcs.New(ctx.Log(), "订单通知", ctx.Request().GetString(fields.FieldOrderID)).Start("通知").End(&r)
	rpns := GetNotify(ctx)
	if err := errs.GetError(rpns); err != nil {
		if errors.Is(err, errs.ErrNotExist) {
			qtask.FinishByInput(ctx, ctx.Request())
			return err
		}
		return err
	}

	//通知到下游

	ctx.Log().Debug("2. 保存通知结果")
	ctx.Request().SetValue(fields.FieldNotifyMsg, "success")
	rpns = SaveSuccess(ctx)
	if err := errs.GetError(rpns); err != nil {
		return err

	}
	qtask.FinishByInput(ctx, ctx.Request())
	return "success"
}
