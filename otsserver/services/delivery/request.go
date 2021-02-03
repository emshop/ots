package delivery

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/delivery"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
)

var deliveryStartNowFields = []string{
	sql.FieldDeliveryID,
}
var deliveryStartSaveFields = []string{
	sql.FieldDeliveryID,
	sql.FieldResultCode,
	sql.FieldReturnMsg,
}

//StartRequest 开始请求
func StartRequest(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------开始发货请求----------------------")
	if err := ctx.Request().Check(deliveryStartNowFields...); err != nil {
		return err
	}
	_, err := delivery.StartNow(ctx.Request().GetString(sql.FieldDeliveryID))
	switch { //无须处理或处理成功则关闭流程
	case errs.GetCode(err) == http.StatusNoContent || err == nil:
		qtask.FinishByInput(ctx, ctx.Request())
		return nil
	case err == nil:
		qtask.FinishByInput(ctx, ctx.Request())
	}
	return err
}

//SaveRequest 保存请求
func SaveRequest(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存发货请求----------------------")
	if err := ctx.Request().Check(deliveryStartSaveFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 保存请求结果")
	err := delivery.SaveStart(
		ctx.Request().GetString(sql.FieldDeliveryID),
		ctx.Request().GetDecimal(sql.FieldRealDiscount),
		ctx.Request().GetString(sql.FieldResultCode),
		ctx.Request().GetString(sql.FieldReturnMsg))
	if err != nil {
		return err
	}
	return nil
}
