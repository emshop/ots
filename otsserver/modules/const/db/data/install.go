package data

import "github.com/micro-plat/hydra"

func init() {
	hydra.OnReady(func() error {
		hydra.Installer.DB.AddSQL(
			DataList...,
		)
		return nil
	})
}
