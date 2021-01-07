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
	//CodeParamNoSetting 参数未配置
	CodeParamNoSetting OrderCode = 901

	//CodeUnknownErr 未知错误
	CodeUnknownErr OrderCode = 999

	//CodeOrderNotExists 订单不存在
	CodeOrderNotExists OrderCode = 902
)

//OrderStatus 订单状态码
type OrderStatus int

const (

	//OrderPaying 支付中
	OrderPaying OrderStatus = 10

	//OrderDelivering 发货中
	OrderDelivering OrderStatus = 20

	//OrderRefunding 退款中
	OrderRefunding OrderStatus = 70

	//OrderSuccess 成功
	OrderSuccess OrderStatus = 0

	//OrderFailed 失败
	OrderFailed OrderStatus = 90
)

//ProcessStatus 处理状态
type ProcessStatus int

const (

	//ProcessNotStart 未开始
	ProcessNotStart ProcessStatus = 10

	//ProcessWaiting 等待处理
	ProcessWaiting ProcessStatus = 20

	//ProcessHandling 处理中
	ProcessHandling ProcessStatus = 30

	//ProcessSuccess 成功
	ProcessSuccess ProcessStatus = 0

	//ProcessFailed 失败
	ProcessFailed ProcessStatus = 90
)

//AccountType 账户类型
type AccountType string

const (
	//AccountMerchant 商户账户
	AccountMerchant AccountType = "merchant"
)

//TradeType 账户类型
type TradeType string
