package mers

//MerchantShelf 商户货架
type MerchantShelf struct {
	//MerShelfID 货架编号
	MerShelfID int `json:"mer_shelf_id"  valid:"required"`

	//MerShelfName 货架名称
	MerShelfName string `json:"mer_shelf_name"  valid:"required"`

	//MerNo 商户编号
	MerNo string `json:"mer_no"  valid:"required"`

	//MerFeeDiscount 商户佣金
	MerFeeDiscount float32 `json:"mer_fee_discount"  valid:"required"`

	//TradeFeeDiscount 交易服务费
	TradeFeeDiscount float32 `json:"trade_fee_discount"  valid:"required"`

	//PaymentFeeDiscount 支付手续费
	PaymentFeeDiscount float32 `json:"payment_fee_discount"  valid:"required"`

	//OrderTimeout 订单超时时长
	OrderTimeout int `json:"order_timeout"  valid:"required"`

	//PaymentTimeout 支付超时时长
	PaymentTimeout int `json:"payment_timeout"  valid:"required"`

	//InvoiceType 开票方式（1.不开发票）
	InvoiceType int `json:"invoice_type"  valid:"required"`

	//CanRefund 允许退款
	CanRefund int `json:"can_refund"  valid:"required"`

	//LimitCount 单次购买数量
	LimitCount int `json:"limit_count"  valid:"required"`

	//SplitFace 拆单面值
	SplitFace int `json:"split_face"  valid:"required"`
}
