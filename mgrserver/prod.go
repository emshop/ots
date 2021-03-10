// +build prod

package main

//go:embed web/dist
var fs embed.FS

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {

	//设置配置参数
	hydra.Conf.Web("8090").Header(header.WithCrossDomain())
	Static(static.WithAutoRewrite(), static.WithEmbed("web/dist", fs)).
		Processor(processor.WithServicePrefix("api"))

	// hydra.Conf.Vars().DB().MySQL("db", "hydra", "123456", "192.168.0.36:3306", "hydra", db.WithConnect(20, 10, 600))
	hydra.Conf.Vars().DB().MySQLByConnStr("db", "hydra:123456@tcp(222.209.84.37:10036)/hydra?charset=utf8")
	//启动时参数配置检查
	App.OnStarting(func(appConf app.IAPPConf) error {

		if _, err := hydra.C.DB().GetDB(); err != nil {
			return fmt.Errorf("db数据库配置错误,err:%v", err)
		}

		return nil
	})
}