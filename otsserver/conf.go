package main

import (
	"github.com/emshop/ots/services/order"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/qtask"
)

func init() {
	hydra.OnReadying(func() {
		qtask.Config(qtask.WithScanInterval(59))
	})
	hydra.OnReady(func() {
		hydra.Conf.Vars().DB().MySQLByConnStr("db", "hydra:123456@tcp(192.168.0.36:3306)/hydra?charset=utf8")
	})
	app.Micro("/order/*", &order.Order{})

}
