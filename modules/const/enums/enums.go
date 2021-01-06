package enums

//FlowStatus 流程状态
type FlowStatus int

const (
	//Success 成功状态
	Success FlowStatus = 0

	//Failed 失败
	Failed FlowStatus = 1

	//Unknown 未知
	Unknown FlowStatus = 2
)

//OrderCode 订单错误码
type OrderCode int

const (
	//ParamNoSetting 参数未配置
	ParamNoSetting OrderCode = 901
)
