# docs

###### 1.1 api 绑定
``` go
import (
    "gitlab.100bm.cn/micro-plat/dds/dds"
)

绑定,生成前端可调用的api服务o
dds.Config(dds.WithDBName("dbname"), //数据库配置
    dds.WithPrex("prefix")) //路由地址前缀
```

###### 1.2 前端js调用
``` js
前端: 请看 gitlab.100bm.cn/micro-plat/jspkg/enumbind  项目说明

1.2.1 安装npm 包
npm install qxnw-enumbind

1.2.2 在代码中使用
//下拉框数据集
this.status = this.EnumUtility.Get("inboundstatus")
//inboundstatus 为后台定义的type
<el-select v-model="queryData.status"  placeholder="请选择状态">
    <el-option
        v-for="item in status"
        :key="item.value"
        :label="item.name"
        :value="item.value">
    </el-option>
</el-select>

//列表使用filter
{{scope.row.status | EnumFilter("inboundstatus")}}

```

###### 1.3 api调用接口
获取字典枚举的接口地址: /dds/dictionary/get 枚举值获取 参数:dic_type
获取省的接口地址: /dds/province/get 没有参数
根据省获取市信息的接口地址: /dds/city/get 参数:parent_code(可为空)
获取省市信息的接口地址: /dds/region/get 没有参数
