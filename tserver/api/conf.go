package main

import (
	"github.com/micro-plat/hydra"
)

func init() {
	hydra.Conf.Vars().DB().MySQL("db", "hydra", "123456", "222.209.84.37:10036", "hydra")
	hydra.OnReady(func() error {
		return nil
	})
}
