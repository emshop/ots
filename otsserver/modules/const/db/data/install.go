package data

import "github.com/micro-plat/hydra"

func init() {
	//注册服务包
	hydra.OnReady(func() error {
		hydra.Installer.DB.AddSQL(
			DataList...,
		)
		return nil
	})
}
