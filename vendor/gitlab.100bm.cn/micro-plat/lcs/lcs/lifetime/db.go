package lifetime

import (
	"fmt"

	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/lcs/lcs/const/conf"
	"gitlab.100bm.cn/micro-plat/lcs/lcs/const/sql"
)

//DBCreate 新增生命周期数据
func dbCreate(info *LifeTimeInfo) {

	if len(info.Content) >= 1000 {
		info.Content = info.Content[0:1000]
	}
	arg, err := types.Struct2Map(info)
	if err != nil {
		panic(fmt.Errorf("新增生命周期数据:struct to map发生错误(err:%v)", err))
	}

	dbe := components.Def.DB().GetRegularDB(conf.DBName)
	if _, err := dbe.Execute(sql.SaveLifeTimeInfo, arg); err != nil {
		panic(fmt.Errorf("新增生命周期数据错误err:%v", err))
	}
}
