package services

import (
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/services/delivery"
	"github.com/emshop/ots/flowserver/services/order"
	"github.com/emshop/ots/flowserver/services/payment"
	"github.com/micro-plat/hydra"
)

func init() {

	//外部接口
	hydra.S.Micro("/order/request", order.Request)
	hydra.S.Micro("/order/query", order.Query)
	hydra.S.Micro("/delivery/request/get", delivery.GetRequest)
	hydra.S.Micro("/delivery/request/save", delivery.SaveRequest)
	hydra.S.Micro("/delivery/result/save", delivery.SaveDeliveryResult)
	hydra.S.Micro("/delivery/query/save", delivery.SaveQueryResult)
	hydra.S.Micro("/delivery/notify/save", delivery.SaveNotifyResult)
	hydra.S.Micro("/delivery/audit/save", delivery.SaveAuditResult)

	//内部流程
	hydra.S.MQC("/payment/paying", payment.Paying, string(enums.FlowPaymentStart))
	hydra.S.MQC("/delivery/binding", delivery.Binding, string(enums.FlowDeliveryBind))
	hydra.S.MQC("/delivery/mock", delivery.MockHandle, "delivery:mock")
}
