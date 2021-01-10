package main

import (
	"github.com/emshop/ots/services/bind"
	"github.com/emshop/ots/services/order"
	"github.com/emshop/ots/services/payment"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/vars/queue/lmq"
	"github.com/micro-plat/qtask"
)

func init() {
	hydra.OnReadying(func() {
		qtask.Config(qtask.WithScanInterval(59))
	})
	hydra.OnReady(func() {
		hydra.Conf.MQC(lmq.MQ)
		hydra.Conf.Vars().Queue().LMQ("queue")
		// hydra.Conf.Vars().DB().MySQLByConnStr("db", "hydra:123456@tcp(192.168.0.36:3306)/hydra?charset=utf8")
		hydra.Conf.Vars().DB().MySQLByConnStr("db", "hydra:123456@tcp(222.209.84.37:10036)/hydra?charset=utf8")
	})

	app.Micro("/order/*", &order.Order{})
	app.MQC("/payment/*", &payment.Payment{}, "order_pay")
	app.MQC("/bind/*", &bind.Bind{}, "order_bind")
	// hydra.MQC.Add("order_pay", "/payment/order")
	// hydra.MQC.Add("pay_timeout", "/payment/timeout")
	// hydra.MQC.Add("order_bind", "/bind/order")

}
