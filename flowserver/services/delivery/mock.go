package delivery

import (
	"errors"
	"fmt"

	"github.com/emshop/ots/flowserver/modules/biz/deliverys"
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

//MockRequestHandle 上游发货
func MockRequestHandle(ctx hydra.IContext) (r interface{}) {

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
	ctx.Request().SetValue(fields.FieldRequestParams, `{"id":100}`)

	rpns = SaveDeliveryResult(ctx)
	if err := errs.GetError(rpns); err != nil {
		return err
	}
	qtask.FinishByInput(ctx, ctx.Request())
	return fmt.Sprintf("%s|%s", code, msg)
}

//MockQueryHandle 上游查询
func MockQueryHandle(ctx hydra.IContext) (r interface{}) {
	ctx.Log().Debug("-------------打桩查询处理----------------------")
	if err := ctx.Request().Check(fields.FieldDeliveryID); err != nil {
		return err
	}

	ctx.Log().Debug("1. 获取查询信息")
	defer lcs.New(ctx.Log(), "交易查询", ctx.Request().GetString(fields.FieldOrderID)).Start("查询").End(&r)
	delivery, err := deliverys.GetQuery(ctx.Request().GetString(fields.FieldDeliveryID))
	if err != nil && errors.Is(err, errs.ErrNotExist) {
		qtask.FinishByInput(ctx, ctx.Request())
		return err
	}
	if err != nil {
		return err
	}

	ctx.Log().Debug("2. 根据状态处理业务")
	switch delivery.GetInt(fields.FieldDeliveryStatus) {
	case int(enums.ProcessWaiting):
		return MockRequestHandle(ctx)

	case int(enums.ProcessHandling):
		qtask.ProcessingByInput(ctx, ctx.Request())
		ctx.Request().SetValue(fields.FieldResultCode, "0100")
		ctx.Request().SetValue(fields.FieldReturnMsg, "上游请求结果未知")
		rpns := SaveRequest(ctx)
		if err := errs.GetError(rpns); err != nil {
			return err
		}
		code := "000"
		msg := "发货成功"

		ctx.Request().SetValue(fields.FieldResultCode, code)
		ctx.Request().SetValue(fields.FieldReturnMsg, msg)
		ctx.Request().SetValue(fields.FieldRequestParams, `{"id":100}`)

		rpns = SaveDeliveryResult(ctx)
		if err := errs.GetError(rpns); err != nil {
			return err
		}
		qtask.FinishByInput(ctx, ctx.Request())
		return fmt.Sprintf("%s|%s", code, msg)
	case int(enums.ProcessRequested):
		qtask.ProcessingByInput(ctx, ctx.Request())
		code := "000"
		msg := "发货成功"

		ctx.Request().SetValue(fields.FieldResultCode, code)
		ctx.Request().SetValue(fields.FieldReturnMsg, msg)

		rpns := SaveQueryResult(ctx)
		if err := errs.GetError(rpns); err != nil {
			return err
		}
		qtask.FinishByInput(ctx, ctx.Request())
		return fmt.Sprintf("%s|%s", code, msg)
	}
	return "success"
}
