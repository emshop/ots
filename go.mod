module github.com/emshop/ots

go 1.16

replace github.com/micro-plat/beanpay => ../../../github.com/micro-plat/beanpay

replace github.com/micro-plat/qtask => ../../../github.com/micro-plat/qtask

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/micro-plat/beanpay v1.0.0
	github.com/micro-plat/hydra v1.1.1
	github.com/micro-plat/lib4go v1.1.0
	github.com/micro-plat/qtask v1.10.1
	gitlab.100bm.cn/micro-plat/dds/dds v1.0.3
	gitlab.100bm.cn/micro-plat/lcs/lcs v0.0.4
)
