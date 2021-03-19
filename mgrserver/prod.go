// +build prod

package main

import (
	"embed"
	"fmt"
	_ "github.com/emshop/ots/mgrserver/api"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/conf/server/auth/basic"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/server/processor"
	"github.com/micro-plat/hydra/conf/server/static"
)

//go:embed web/dist
var fs embed.FS

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {

	//设置配置参数
	hydra.Conf.Web("8060").Header(header.WithCrossDomain()).
		Static(static.WithAutoRewrite(), static.WithEmbed("web/dist", fs)).
		Processor(processor.WithServicePrefix("api")).Basic(basic.WithUP("yanglei", "1qaz2wsx"))

	hydra.Conf.Vars().DB().MySQLByConnStr("db", "ots:a1qaz2wsxA@tcp(rm-abp1jncd360df9dqvy.mysql.rds.aliyuncs.com:3306)/ots?charset=utf8")

	//启动时参数配置检查
	App.OnStarting(func(appConf app.IAPPConf) error {
		if _, err := hydra.C.DB().GetDB(); err != nil {
			return fmt.Errorf("db数据库配置错误,err:%v", err)
		}
		return nil
	})
}
