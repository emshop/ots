package sql

//UpdateTradeOrderForPayTimeout 修改订单为支付超时
const UpdateTradeOrderForPayTimeout = `
update ots_trade_order t set
t.payment_status = 90,
t.delivery_status = 11,
t.refund_status = 11,
t.order_status = 50
where
t.order_id = @order_id
and t.order_status = 10
and t.payment_status = 20
and t.delivery_status = 10
and t.refund_status = 10
and t.payment_timeout < now()
`
