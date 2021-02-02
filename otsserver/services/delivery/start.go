package delivery

import (
	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/delivery"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
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
	_, err := delivery.StartNow(ctx.Request().GetString(sql.FieldDeliveryID))
	return err
}

//SaveSartHandle 保存请求
func (o *Delivery) SaveSartHandle(ctx hydra.IContext) interface{} {

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

	ctx.Log().Info("2. 执行后续流程")
	status := types.DecodeInt(err, nil, enums.Success, enums.Failed)
	flw, err := flow.Next(ctx.Request().GetInt(sql.FieldFlowID), enums.FlowStatus(status), ctx,
		sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID),
		sql.FieldDeliveryID, ctx.Request().GetString(sql.FieldDeliveryID))
	if err != nil {
		ctx.Log().Error("2.1 执行绑定后续流程失败", err)
	}
	return flw
}
