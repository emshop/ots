package delivery

import (
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/delivery"
	"github.com/micro-plat/hydra"
)

var deliverySaveResultFields = []string{
	sql.FieldDeliveryID,
	sql.FieldResultCode,
	sql.FieldReturnMsg,
}

//SaveResultHandle 开始请求
func (o *Delivery) SaveResultHandle(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存发货结果----------------------")
	if err := ctx.Request().Check(deliverySaveResultFields...); err != nil {
		return err
	}

	//保存充值结果
	_, err := delivery.SaveDeliveryResult(
		ctx.Request().GetInt64(sql.FieldDeliveryID),
		ctx.Request().GetString(sql.FieldResultCode),
		ctx.Request().GetString(sql.FieldReturnMsg),
		ctx.Request().GetDecimal(sql.FieldRealDiscount),
		ctx.Request().GetString(sql.FieldRequestParams),
	)

	return err
}
