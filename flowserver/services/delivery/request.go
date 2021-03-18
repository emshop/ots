package delivery

import (
	"net/http"

	"github.com/emshop/ots/flowserver/modules/biz/deliverys"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
)

//GetRequest 获取待发货信息
func GetRequest(ctx hydra.IContext) interface{} {
	ctx.Log().Debug("-------------获取待发货记录----------------------")
	if err := ctx.Request().Check(fields.FieldDeliveryID); err != nil {
		return err
	}
	qtask.ProcessingByInput(ctx, ctx.Request())
	delivery, err := deliverys.Start(ctx.Request().GetString(fields.FieldDeliveryID))
	if errs.GetCode(err) == http.StatusNoContent {
		qtask.FinishByInput(ctx, ctx.Request())
	}
	if err != nil {
		return err
	}
	qtask.FinishByInput(ctx, ctx.Request())
	return delivery
}

//SaveRequest 保存发货请求结果
func SaveRequest(ctx hydra.IContext) interface{} {
	ctx.Log().Debug("-------------保存发货请求结果----------------------")
	if err := ctx.Request().Check(fields.FieldDeliveryID, fields.FieldResultCode, fields.FieldReturnMsg); err != nil {
		return err
	}
	err := deliverys.SaveStart(ctx.Request().GetString(fields.FieldDeliveryID),
		ctx.Request().GetDecimal(fields.FieldRealDiscount),
		ctx.Request().GetString(fields.FieldResultCode),
		ctx.Request().GetString(fields.FieldReturnMsg))
	if err != nil {
		return err
	}
	return "success"

}
