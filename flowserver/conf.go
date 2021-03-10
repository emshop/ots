package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/vars/queue/lmq"
	"github.com/micro-plat/qtask"

	_ "github.com/emshop/ots/flowserver/modules/const/db/data"
	_ "github.com/emshop/ots/flowserver/modules/const/db/scheme"
	_ "github.com/emshop/ots/flowserver/services"
)

func init() {
	hydra.OnReadying(func() {
		qtask.Config(qtask.WithScanInterval(1))
	})
	hydra.OnReady(func() {
		hydra.Conf.API("8081")
		hydra.Conf.MQC(lmq.MQ)
		hydra.Conf.Vars().Queue().LMQ("queue")
		hydra.Conf.Vars().DB().MySQLByConnStr("db", "hydra:123456@tcp(222.209.84.37:10036)/hydra?charset=utf8")
		qtask.BindFlow()

	})

}