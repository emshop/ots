module github.com/emshop/ots/tserver/api

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/micro-plat/hydra v1.0.5
	github.com/micro-plat/lib4go v1.0.10
)

replace github.com/micro-plat/hydra => ../../../../../github.com/micro-plat/hydra

replace github.com/micro-plat/lib4go => ../../../../../github.com/micro-plat/lib4go