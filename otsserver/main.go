package main

import (
	_ "github.com/emshop/ots/modules/const/db/data"
	_ "github.com/emshop/ots/modules/const/db/scheme"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/hydra/hydra/servers/mqc"
)

var app = hydra.NewApp(
	hydra.WithPlatName("ots"),
	hydra.WithSystemName("ot"),
	hydra.WithServerTypes(http.API, cron.CRON, mqc.MQC),
	hydra.WithClusterName("prod"),
)

func main() {
	app.Start()
}
