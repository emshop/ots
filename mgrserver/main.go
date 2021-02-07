package main

import (
	_ "github.com/emshop/ots/mgrserver/api"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var App = hydra.NewApp(
	hydra.WithPlatName("emshop", "充值交易平台"),
	hydra.WithSystemName("mgrserver", "运营系统"),
	hydra.WithServerTypes(http.Web),
)

func main() {
	App.Start()
}
