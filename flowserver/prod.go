// +build prod

package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/vars/queue/lmq"
	"github.com/micro-plat/qtask"

	_ "github.com/emshop/ots/flowserver/modules/const/db/data"
	_ "github.com/emshop/ots/flowserver/modules/const/db/scheme"
	_ "github.com/emshop/ots/flowserver/services"
	_ "gitlab.100bm.cn/micro-plat/dds/dds"
)

func init() {
	hydra.OnReadying(func() {
		qtask.Config(qtask.WithScanInterval(1))
	})
	hydra.OnReady(func() {
		hydra.Conf.API("8068")
		hydra.Conf.MQC(lmq.MQ)
		hydra.Conf.Vars().Queue().LMQ("queue")

		hydra.Conf.Vars().DB().MySQLByConnStr("db", "ots:a1qaz2wsxA@tcp(rm-abp1jncd360df9dqvy.mysql.rds.aliyuncs.com:3306)/ots?charset=utf8")

		qtask.BindFlow()

	})

}
