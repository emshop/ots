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

	deliverys, err := deliverys.QueryReplenish()
	if err != nil {
		return err
	}
	for _, delivery := range deliverys {
		deliveryID := delivery.GetString(fields.FieldDeliveryID)
		switch delivery.GetInt(fields.FieldDeliveryStatus) {
		case int(enums.ProcessWaiting):
			flows.NextByDeliveryID(deliveryID, enums.FlowDelivery, ctx)
		case int(enums.ProcessHandling):
			return nil
		case int(enums.ProcessRequested):
			return nil
		}
	}

	return "success"
}
