package main

import (
	"github.com/emshop/ots/mgrserver/api/system"
	"github.com/emshop/ots/mgrserver/api/trade"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/header"
)

func init() {
	hydra.OnReady(func() {
		App.Micro("/trade/order", &trade.TradeOrderHandler{})
		App.Micro("/enums", &system.EnumsHandler{})
		hydra.Conf.Vars().DB().MySQL("db", "hydra", "123456", "222.209.84.37:10036", "hydra")
		hydra.Conf.GetWeb().Header(header.WithCrossDomain())
	})
}
