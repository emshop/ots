package sql

//UpdateTradeOrderForPaying 修改订单为正在支付
const UpdateTradeOrderForPaying = `
update ots_trade_order t set
t.payment_status = 0,
t.order_status = 20,
t.delivery_status = 20
where
t.order_id = @order_id 
and t.order_status = 10
and t.payment_status = 20
and t.delivery_status = 10
`

//UpdateTradeOrderForPayTimeout 修改订单为支付超时
const UpdateTradeOrderForPayTimeout = `
update ots_trade_order t set
t.order_status = 90,
t.payment_status = 90,
t.delivery_status = 90,
t.refund_status = 90,
t.notify_status = 10
where
t.order_id = @order_id
and t.order_status = 10
and t.payment_status = 20
and t.delivery_status = 10
and t.refund_status = 10
and t.notify_status = 10
and t.payment_timeout < now()
`
