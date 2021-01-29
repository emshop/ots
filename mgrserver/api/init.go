package main

import (
	"github.com/micro-plat/hydra"
	"github.com/emshop/ots/mgrserver/api/services/trade"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	hydra.OnReady(func() {
		hydra.S.Web("/trade/order", trade.NewTradeOrderHandler())
		hydra.S.Web("/trade/delivery", trade.NewTradeDeliveryHandler())
	})
}

