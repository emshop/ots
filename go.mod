module github.com/emshop/ots

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/micro-plat/beanpay v1.0.0
	github.com/micro-plat/hydra v1.0.5
	github.com/micro-plat/lib4go v1.0.10
	github.com/micro-plat/qtask v1.10.1
)

replace github.com/micro-plat/hydra => ../../../github.com/micro-plat/hydra

replace github.com/micro-plat/qtask => ../../../github.com/micro-plat/qtask

replace github.com/micro-plat/beanpay => ../../../github.com/micro-plat/beanpay
