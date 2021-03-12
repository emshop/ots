package lifetime

import (
	"github.com/micro-plat/lib4go/logger"
)

//Create 新增生命周期 数据(由于要链式调用,不能返回错误)
func Create(info *LifeTimeInfo, log logger.ILogger) {
	defer func() {
		if v := recover(); v != nil {
			log.Infof("保存生命周期数据时出错:%+v", v)
		}
	}()
	if err := info.Valid(); err != nil {
		panic(err)
	}

	dbCreate(info)
}
