# 生命周期记录系统 文档和数据字典, 以及使用说明

使用说明

## 1. 在所需项目中引入lcs项目包

``` GO
import "gitlab.100bm.cn/micro-plat/lcs/lcs"

```

## 2. 设置数据库配置节点

``` GO
//如果数据库配置节点不是db,需要调用SetDB初始化的时候进行设置
import "gitlab.100bm.cn/micro-plat/lcs/lcs"
lcs.Conf(lcs.WithDBName("lcsdb"))   //不设置默认为:db
```

## 3. 需要创建的表

``` 

根据不同的数据库类型,程序编译时直接通过指定oralce/mysql标识进行编译即可;如:go build -tags "oracle"
所有的建表语句sql已经打包到代码中,引用方只需要引入代码包,编译后按照正常的hydra数据库安装命令进行安装即可:./xxxx db install

```

## 4. 使用

``` Go

import (
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

//1: 成对数据(会生成两条数据)
defer lcs.New(log, "创建订单", "20190401123").Start("开始创单").End(err)

//2: 创建一条数据
lcs.New(log, "创建订单", "147258", "369").Create("测试创建") //可以单个创建

```

## 5. 数据库设计发生变更

重新运行脚本sql.sh更新代码中的建表语句:(oracle和mysql都要更新)   
./sql.sh mysql   
./sql.sh oracle
