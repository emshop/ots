package delivery

import (
	"github.com/emshop/ots/flowserver/modules/biz/deliverys"
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/flows"
	"github.com/micro-plat/hydra"
)

//Query 订单查询
func Replenish(ctx hydra.IContext) interface{} {

	ctx.Log().Debug("-------------查询发货后补数据----------------------")
	deliverys, err := deliverys.QueryReplenish()
	if err != nil {
		return err
	}
	ctx.Log().Debugf("1.查询到数据%d条", deliverys.Len())
	for _, delivery := range deliverys.Maps() {
		deliveryID := delivery.GetString(fields.FieldDeliveryID)
		switch delivery.GetInt(fields.FieldDeliveryStatus) {
		case int(enums.ProcessWaiting):
			flows.NextByDeliveryID(deliveryID, enums.FlowDelivery, ctx)
		case int(enums.ProcessHandling):
			flows.NextByDeliveryID(deliveryID, enums.FlowQueryStart, ctx)
		case int(enums.ProcessRequested):
			flows.NextByDeliveryID(deliveryID, enums.FlowQueryStart, ctx)
		case int(enums.ProcessSuccess):
			if delivery.GetInt(fields.FieldPaymentStatus) == int(enums.ProcessWaiting) {
				flows.NextByDeliveryID(deliveryID, enums.FlowDeliveryPay, ctx)
			}
		}
	}

	return "success"
}
