package services

import (
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/services/delivery"
	"github.com/emshop/ots/flowserver/services/finish"
	"github.com/emshop/ots/flowserver/services/merchant"
	"github.com/emshop/ots/flowserver/services/notify"
	"github.com/emshop/ots/flowserver/services/order"
	"github.com/emshop/ots/flowserver/services/payment"
	"github.com/emshop/ots/flowserver/services/spp"
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

	hydra.S.Micro("/notify/get", notify.GetNotify)
	hydra.S.Micro("/notify/success", notify.SaveSuccess)
	hydra.S.Micro("/notify/unknown", notify.Unknown)

	//测试使用
	hydra.S.Micro("/merchant/create", merchant.Create)
	hydra.S.Micro("/spp/create", spp.Create)
	hydra.S.Micro("/order/mock", order.Mock)

	//内部流程
	hydra.S.MQC("/payment/paying", payment.Paying, string(enums.FlowPaymentStart))
	hydra.S.MQC("/delivery/binding", delivery.Binding, string(enums.FlowDeliveryBind))
	hydra.S.MQC("/delivery/mock", delivery.MockHandle, "delivery:mock")
	hydra.S.MQC("/delivery/paying", delivery.Paying, string(enums.FlowDeliveryPay))
	hydra.S.MQC("/notify/start", notify.Notify, string(enums.FlowNotifyStart))
	hydra.S.MQC("/finish/start", finish.Finish, string(enums.FlowFinishStart))

	hydra.S.CRON("/order/replenish", order.Replenish, "@every 10s")
	hydra.S.CRON("/delivery/replenish", delivery.Replenish, "@every 10s")

}
