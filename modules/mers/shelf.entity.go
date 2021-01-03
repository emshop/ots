package mers

import (
	"time"

	"github.com/micro-plat/lib4go/types"
)

//MerchantShelf 商户货架
type MerchantShelf struct {

	//MerShelfID 货架编号
	MerShelfID int `json:"mer_shelf_id"`

	//MerShelfName 货架名称
	MerShelfName string `json:"mer_shelf_name"`

	//MerNo 商户编号
	MerNo string `json:"mer_no"`

	//MerFeeDiscount 商户佣金
	MerFeeDiscount types.Decimal `json:"mer_fee_discount"`

	//TradeFeeDiscount 交易服务费
	TradeFeeDiscount types.Decimal `json:"trade_fee_discount"`

	//PaymentFeeDiscount 支付手续费
	PaymentFeeDiscount types.Decimal `json:"payment_fee_discount"`

	//OrderTimeout 订单超时时长
	OrderTimeout int `json:"order_timeout"`

	//PaymentTimeout 支付超时时长
	PaymentTimeout int `json:"payment_timeout"`

	//InvoiceType 开票方式（1.不开发票）
	InvoiceType int `json:"invoice_type"`

	//CanRefund 允许退款(0.是,1否)
	CanRefund int `json:"can_refund"`

	//LimitCount 单次购买数量
	LimitCount int `json:"limit_count"`

	//SplitFace 拆单面值
	SplitFace int `json:"split_face"`

	//Status 状态
	Status int `json:"status"`

	//CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
}
