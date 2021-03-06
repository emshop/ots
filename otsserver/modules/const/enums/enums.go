package enums

import "fmt"

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

	//CodeOverLimit 超过数量限制
	CodeOverLimit OrderCode = 903

	//CodeBalanceLow 余额不足
	CodeBalanceLow OrderCode = 801
)

//OrderStatus 订单状态码
type OrderStatus int

const (

	//OrderPaying 支付中
	OrderPaying OrderStatus = 10

	//OrderDelivering 发货中
	OrderDelivering OrderStatus = 20

	//OrderSPPPaying 上游支付
	OrderSPPPaying OrderStatus = 40

	//OrderNotify 下游通知中
	OrderNotify OrderStatus = 50

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

	//ProcessNoNeed 无须处理
	ProcessNoNeed ProcessStatus = 11

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

	//AccountSupplier 供货商
	AccountSupplier AccountType = "supplier"
)

//ResultSource 结果来源
type ResultSource string

const (
	//ResultFromRequest 来源于发货请求
	ResultFromRequest ResultSource = "request"

	//ResultFromDelivery 来源于发货结果
	ResultFromDelivery ResultSource = "delivery"

	//ResultFromNotify 来源于通知
	ResultFromNotify ResultSource = "notify"

	//ResultFromQuery 来源于查询
	ResultFromQuery ResultSource = "query"

	//ResultAudit 来源于人工处理
	ResultAudit ResultSource = "audit"
)

//Format 格式化错误码
func (r ResultSource) Format(code string) string {
	return fmt.Sprintf("%s_%s", r, code)
}

//FlowTag 流程标签
type FlowTag string

const (
	//FlowFlagRecvOrderByPID 来源于发货请求
	FlowFlagRecvOrderByPID FlowTag = "recv_by_pid"

	//FlowFlagOrderNotify 通知
	FlowFlagOrderNotify FlowTag = "order_notify"

	//FlowFlagDeliveryFinish 发货完成流程
	FlowFlagDeliveryFinish FlowTag = "delivery_finish"
)

//FlowName 流程名称
type FlowName string

const (

	//FlowOrderPay 订单支付流程
	FlowOrderPay FlowName = "order_pay"

	//FlowOrderPayTimeout 订单支付超时流程
	FlowOrderPayTimeout FlowName = "pay_timeout"

	//FlowOrderBind 订单绑定
	FlowOrderBind FlowName = "order_bind"

	//FlowDeliveryPay 发货支付
	FlowDeliveryPay FlowName = "spp_payment"

	//FlowMockDelivery 发货流程
	FlowMockDelivery FlowName = "mock_deliery"

	//FlowOrderNotify 订单通知
	FlowOrderNotify FlowName = "order_notify"

	//FlowOrderTimeout 订单超时
	FlowOrderTimeout FlowName = "order_timeout"

	//FlowOrderFinish 订单完结
	FlowOrderFinish FlowName = "order_finish"
)
