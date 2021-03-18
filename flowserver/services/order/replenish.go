package order

import (
	"github.com/emshop/ots/flowserver/modules/biz/orders"
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/flows"
	"github.com/micro-plat/hydra"
)

//Query 订单查询
func Replenish(ctx hydra.IContext) interface{} {

	orders, err := orders.QueryReplenish()
	if err != nil {
		return err
	}
	for _, order := range orders {
		orderID := order.GetString(fields.FieldOrderID)
		switch order.GetInt(fields.FieldOrderStatus) {
		case int(enums.OrderPaying):
			flows.NextByOrderNO(orderID, enums.FlowPaymentStart, ctx)
		case int(enums.OrderDelivering):
			flows.NextByOrderNO(orderID, enums.FlowDeliveryBind, ctx)
		case int(enums.OrderNotify):
			flows.NextByOrderNO(orderID, enums.FlowNotifyStart, ctx)
		case int(enums.OrderFinish):
			flows.NextByOrderNO(orderID, enums.FlowFinishStart, ctx)
		}
	}

	return "success"
}
