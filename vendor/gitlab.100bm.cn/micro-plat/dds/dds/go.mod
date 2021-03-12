module gitlab.100bm.cn/micro-plat/dds/dds

go 1.15

replace (
	github.com/micro-plat/hydra => ../../../../github.com/micro-plat/hydra
	github.com/micro-plat/lib4go => ../../../../github.com/micro-plat/lib4go
)

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/micro-plat/hydra v1.0.5
	github.com/micro-plat/lib4go v1.0.10
	github.com/micro-plat/qtask v1.10.1
	github.com/zkfy/go-oci8 v1.0.0
)
