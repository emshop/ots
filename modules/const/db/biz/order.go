package biz

//TradeOrder 订单记录

//FieldTradeOrderOrderID 字段订单记录的数据库名称
const FieldTradeOrderOrderID = "order_id"

//FieldTradeOrderMerNo 字段订单记录的数据库名称
const FieldTradeOrderMerNo = "mer_no"

//FieldTradeOrderMerOrderNo 字段订单记录的数据库名称
const FieldTradeOrderMerOrderNo = "mer_order_no"

//FieldTradeOrderMerProductID 字段订单记录的数据库名称
const FieldTradeOrderMerProductID = "mer_product_id"

//FieldTradeOrderMerShelfID 字段订单记录的数据库名称
const FieldTradeOrderMerShelfID = "mer_shelf_id"

//FieldTradeOrderMerProductNo 字段订单记录的数据库名称
const FieldTradeOrderMerProductNo = "mer_product_no"

//FieldTradeOrderPlID 字段订单记录的数据库名称
const FieldTradeOrderPlID = "pl_id"

//FieldTradeOrderBrandNo 字段订单记录的数据库名称
const FieldTradeOrderBrandNo = "brand_no"

//FieldTradeOrderProvinceNo 字段订单记录的数据库名称
const FieldTradeOrderProvinceNo = "province_no"

//FieldTradeOrderCityNo 字段订单记录的数据库名称
const FieldTradeOrderCityNo = "city_no"

//FieldTradeOrderFace 字段订单记录的数据库名称
const FieldTradeOrderFace = "face"

//FieldTradeOrderNum 字段订单记录的数据库名称
const FieldTradeOrderNum = "num"

//FieldTradeOrderTotalFace 字段订单记录的数据库名称
const FieldTradeOrderTotalFace = "total_face"

//FieldNameAccountName 字段用户账户信息的数据库名称
const FieldNameAccountName = "account_name"

//FieldTradeOrderInvoiceType 字段订单记录的数据库名称
const FieldTradeOrderInvoiceType = "invoice_type"

//FieldTradeOrderSellAmount 字段订单记录的数据库名称
const FieldTradeOrderSellAmount = "sell_amount"

//FieldTradeOrderMerFeeAmount 字段订单记录的数据库名称
const FieldTradeOrderMerFeeAmount = "mer_fee_amount"

//FieldTradeOrderTradeFeeAmount 字段订单记录的数据库名称
const FieldTradeOrderTradeFeeAmount = "trade_fee_amount"

//FieldTradeOrderPaymentFeeAmount 字段订单记录的数据库名称
const FieldTradeOrderPaymentFeeAmount = "payment_fee_amount"

//FieldTradeOrderCanSplitOrder 字段订单记录的数据库名称
const FieldTradeOrderCanSplitOrder = "can_split_order"

//FieldTradeOrderSplitOrderFace 字段订单记录的数据库名称
const FieldTradeOrderSplitOrderFace = "split_order_face"

//FieldTradeOrderCreateTime 字段订单记录的数据库名称
const FieldTradeOrderCreateTime = "create_time"

//FieldTradeOrderOrderTimeout 字段订单记录的数据库名称
const FieldTradeOrderOrderTimeout = "order_timeout"

//FieldTradeOrderPaymentTimeout 字段订单记录的数据库名称
const FieldTradeOrderPaymentTimeout = "payment_timeout"

//FieldTradeOrderDeliveryPause 字段订单记录的数据库名称
const FieldTradeOrderDeliveryPause = "delivery_pause"

//FieldTradeOrderOrderStatus 字段订单记录的数据库名称
const FieldTradeOrderOrderStatus = "order_status"

//FieldTradeOrderPaymentStatus 字段订单记录的数据库名称
const FieldTradeOrderPaymentStatus = "payment_status"

//FieldTradeOrderDeliveryStatus 字段订单记录的数据库名称
const FieldTradeOrderDeliveryStatus = "delivery_status"

//FieldTradeOrderRefundStatus 字段订单记录的数据库名称
const FieldTradeOrderRefundStatus = "refund_status"

//FieldTradeOrderNotifyStatus 字段订单记录的数据库名称
const FieldTradeOrderNotifyStatus = "notify_status"

//FieldTradeOrderIsRefund 字段订单记录的数据库名称
const FieldTradeOrderIsRefund = "is_refund"

//FieldTradeOrderBindFace 字段订单记录的数据库名称
const FieldTradeOrderBindFace = "bind_face"

//FieldTradeOrderSuccessFace 字段订单记录的数据库名称
const FieldTradeOrderSuccessFace = "success_face"

//FieldTradeOrderSuccessUserAmount 字段订单记录的数据库名称
const FieldTradeOrderSuccessUserAmount = "success_user_amount"

//FieldTradeOrderSuccessMerFee 字段订单记录的数据库名称
const FieldTradeOrderSuccessMerFee = "success_mer_fee"

//FieldTradeOrderSuccessTradeFee 字段订单记录的数据库名称
const FieldTradeOrderSuccessTradeFee = "success_trade_fee"

//FieldTradeOrderSuccessPaymentFee 字段订单记录的数据库名称
const FieldTradeOrderSuccessPaymentFee = "success_payment_fee"

//FieldTradeOrderSuccessCostAmount 字段订单记录的数据库名称
const FieldTradeOrderSuccessCostAmount = "success_cost_amount"

//FieldTradeOrderSuccessSppFee 字段订单记录的数据库名称
const FieldTradeOrderSuccessSppFee = "success_spp_fee"

//FieldTradeOrderSuccessSppTradeFee 字段订单记录的数据库名称
const FieldTradeOrderSuccessSppTradeFee = "success_spp_trade_fee"

//FieldTradeOrderProfit 字段订单记录的数据库名称
const FieldTradeOrderProfit = "profit"

//FieldName 字段名称
const FieldName = "name"

//SelectTradeOrderByMerchant 通过mer_no,mer_order_no获取订单
const SelectTradeOrderByMerchant = `
select t.order_id,t.mer_no,t.mer_order_no,t.mer_product_id,
t.pl_id,t.brand_no,t.province_no,t.city_no,t.face,t.num,
t.total_face,t.invoice_type from ots_trade_order t
where t.mer_no = @mer_no and t.mer_order_no=@mer_order_no limit 1
`

//InsertTradeOrder 插入单条数据订单记录
const InsertTradeOrder = `
insert into ots_trade_order(
order_id,
mer_no,
mer_order_no,
mer_product_id,
mer_shelf_id,
mer_product_no,
pl_id,
brand_no,
province_no,
city_no,
face,
num,
total_face,
account_name,
invoice_type,
sell_amount,
mer_fee_amount,
trade_fee_amount,
payment_fee_amount,
can_split_order,
split_order_face,
order_timeout,
payment_timeout
)values(
@order_id,
@mer_no,
@mer_order_no,
@mer_product_id,
@mer_shelf_id,
@mer_product_no,
@pl_id,
@brand_no,
@province_no,
@city_no,
@face,
@num,
@face*@num,
@account_name,
@invoice_type,
@face*@num*@discount,
@mer_fee_discount*@face*@num,
@trade_fee_discount*@face*@nu,
@payment_fee_discount*@face*@nu,
case when @split_face='0' then 1 else 0 end,
@split_face,
DATE_ADD(now(),INTERVAL @order_timeout SECOND),
DATE_ADD(now(),INTERVAL @payment_timeout SECOND)
)
`
