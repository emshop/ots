package lcs

import "gitlab.100bm.cn/micro-plat/lcs/lcs/const/conf"

//Config 配置数据库节点名称等
func Config(opts ...ConfOption) {
	for _, opt := range opts {
		opt()
	}
}

//GetDBName 获取数据库节点名称
func GetDBName() string {
	return conf.DBName
}

//ConfOption 配置选项
type ConfOption func()

//WithDBName 设置db数据节点
func WithDBName(dbName string) ConfOption {
	return func() {
		conf.DBName = dbName
	}
}
