package main

import (
	_ "github.com/emshop/ots/mgrserver/api"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var App = hydra.NewApp(
	hydra.WithPlatName("平台名", "平台中文名"),
	hydra.WithSystemName("系统名", "系统中文名"),
	hydra.WithServerTypes(http.Web),
)

func main() {
	App.Start()
}
