package delivery

import (
	"fmt"

	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

//MockHandle 上游付款
func MockHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("-------------打桩发货处理----------------------")
	defer lcs.New(ctx.Log(), "交易发货", ctx.Request().GetString(fields.FieldOrderID)).Start("发货").End(&r)
	ctx.Log().Debug("1. 获取发货信息")
	rpns := GetRequest(ctx)
	if err := errs.GetError(rpns); err != nil {
		return err

	}

	ctx.Log().Debug("2. 保存为已提交到上游")
	ctx.Request().SetValue(fields.FieldResultCode, "000")
	ctx.Request().SetValue(fields.FieldReturnMsg, "上游请求成功")
	rpns = SaveRequest(ctx)
	if err := errs.GetError(rpns); err != nil {
		return err
	}

	ctx.Log().Debug("3. 保存为发货成功")

	// deliveryID := ctx.Request().GetInt(fields.FieldDeliveryID)
	// code := types.DecodeString(deliveryID%2, 1, "000", "1111")
	// msg := types.DecodeString(deliveryID%5, 1, "发货成功", "发货失败")
	code := "000"
	msg := "发货成功"

	ctx.Request().SetValue(fields.FieldResultCode, code)
	ctx.Request().SetValue(fields.FieldReturnMsg, msg)

	rpns = SaveDeliveryResult(ctx)
	if err := errs.GetError(rpns); err != nil {
		return err
	}
	qtask.FinishByInput(ctx, ctx.Request())
	return fmt.Sprintf("%s|%s", code, msg)
}
