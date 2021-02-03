package delivery

import (
	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/delivery"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/micro-plat/hydra"
)

var deliverySaveResultFields = []string{
	sql.FieldDeliveryID,
	sql.FieldResultCode,
	sql.FieldReturnMsg,
}

//SaveResult 开始请求
func SaveResult(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存发货结果----------------------")
	if err := ctx.Request().Check(deliverySaveResultFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 保存发货结果")
	_, err := delivery.SaveDeliveryResult(
		ctx.Request().GetString(sql.FieldDeliveryID),
		ctx.Request().GetString(sql.FieldResultCode),
		ctx.Request().GetString(sql.FieldReturnMsg),
		ctx.Request().GetDecimal(sql.FieldRealDiscount),
		ctx.Request().GetString(sql.FieldRequestParams),
	)

	ctx.Log().Info("2. 执行后续流程")
	delivery, err := delivery.Get(ctx.Request().GetString(sql.FieldDeliveryID))
	if err != nil {
		return err
	}

	flw := flow.NextByFirst(enums.FlowFlagDeliveryFinish, delivery.GetInt(sql.FieldPlID), enums.Success, ctx,
		sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID),
		sql.FieldDeliveryID, ctx.Request().GetString(sql.FieldDeliveryID))
	return flw
}
