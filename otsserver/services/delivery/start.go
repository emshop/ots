package delivery

import (
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/delivery"
	"github.com/micro-plat/hydra"
)

var deliveryStartNowFields = []string{
	sql.FieldDeliveryID,
}
var deliveryStartSaveFields = []string{
	sql.FieldDeliveryID,
	sql.FieldResultCode,
	sql.FieldReturnMsg,
}

//Delivery 发货处理
type Delivery struct {
}

//StartHandle 开始请求
func (o *Delivery) StartHandle(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------开始发货请求----------------------")
	if err := ctx.Request().Check(deliveryStartNowFields...); err != nil {
		return err
	}
	_, err := delivery.StartNow(ctx.Request().GetInt64(sql.FieldDeliveryID))
	return err
}

//SaveSartHandle 保存请求
func (o *Delivery) SaveSartHandle(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存发货请求----------------------")
	if err := ctx.Request().Check(deliveryStartSaveFields...); err != nil {
		return err
	}
	err := delivery.SaveStart(
		ctx.Request().GetInt64(sql.FieldDeliveryID),
		ctx.Request().GetDecimal(sql.FieldRealDiscount),
		ctx.Request().GetString(sql.FieldResultCode),
		ctx.Request().GetString(sql.FieldReturnMsg))
	return err
}
