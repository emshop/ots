package api

import (
	"github.com/emshop/ots/mgrserver/api/services/system"
	"github.com/emshop/ots/mgrserver/api/services/trade"
	"github.com/micro-plat/hydra"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	hydra.OnReady(func() {
		hydra.S.Web("/system/enums", system.NewSystemEnumsHandler())
		hydra.S.Web("/trade/order", trade.NewTradeOrderHandler())
	})
}
