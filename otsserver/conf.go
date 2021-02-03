package main

import (
	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/services/bind"
	"github.com/emshop/ots/otsserver/services/delivery"
	"github.com/emshop/ots/otsserver/services/order"
	"github.com/emshop/ots/otsserver/services/payment"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/vars/queue/lmq"
	"github.com/micro-plat/qtask"
)

func init() {
	hydra.OnReadying(func() {
		qtask.Config(qtask.WithScanInterval(59))
	})
	hydra.OnReady(func() {
		hydra.Conf.API("8081")
		hydra.Conf.MQC(lmq.MQ)
		hydra.Conf.Vars().Queue().LMQ("queue")
		// hydra.Conf.Vars().DB().MySQLByConnStr("db", "hydra:123456@tcp(192.168.0.36:3306)/hydra?charset=utf8")
		hydra.Conf.Vars().DB().MySQLByConnStr("db", "hydra:123456@tcp(222.209.84.37:10036)/hydra?charset=utf8")

		qtask.BindFlow()

	})
	app.Micro("/order/request", order.RequestHandle)
	app.MQC("/order/pay", payment.PayHandle, string(enums.FlowOrderPay))
	app.MQC("/order/pay/timeout", payment.TimeoutHandle, string(enums.FlowOrderPayTimeout))
	app.MQC("/order/bind", bind.StartHandle, string(enums.FlowOrderBind))
	app.MQC("/delivery/pay", delivery.PayHandle, string(enums.FlowDeliveryPay))
	app.MQC("/order/notify", order.NotifyHandle, string(enums.FlowOrderNotify))
	app.MQC("/order/finish", order.FinishHandle, string(enums.FlowOrderFinish))

}
