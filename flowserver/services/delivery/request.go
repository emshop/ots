package delivery

import (
	"net/http"

	"github.com/emshop/ots/flowserver/modules/biz/deliverys"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
)

//GetRequest 获取待发货信息
func GetRequest(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------获取待发货记录----------------------")
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
	ctx.Log().Info("-------------保存发货请求结果----------------------")
	if err := ctx.Request().Check(sql.FieldDeliveryID, sql.FieldResultCode, sql.FieldReturnMsg); err != nil {
		return err
	}
	return deliverys.SaveStart(ctx.Request().GetString(sql.FieldDeliveryID),
		ctx.Request().GetDecimal(sql.FieldRealDiscount),
		ctx.Request().GetString(sql.FieldResultCode),
		ctx.Request().GetString(sql.FieldReturnMsg))

}
