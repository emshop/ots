package main

import (
	"github.com/emshop/ots/tserver/api/trade"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/header"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	hydra.Conf.Web("8089", api.WithTrace()).Header(header.WithCrossDomain())
	hydra.S.Micro("/trade/order", trade.NewTradeOrderHandler())

}
