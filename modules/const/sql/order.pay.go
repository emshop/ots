package sql

//UpdateTradeOrderForPaying 修改订单为正在支付
const UpdateTradeOrderForPaying = `
Update ots_trade_order t
t.payment_status = 30 
from ots_trade_order t
where
t.order_id = @order_id 
and t.delivery_pause = 1
and t.order_status = 10
and t.payment_status = 20
`

//UpdateTradeOrderForPayTimeout 修改订单为支付超时
const UpdateTradeOrderForPayTimeout = `
Update ots_trade_order t
t.order_status = 90,
t.payment_status = 90,
t.delivery_status = 90,
t.refund_status = 90,
t.notify_status = 10
from ots_trade_order t
where
t.order_id = @order_id
and t.delivery_pause = 1
and t.order_status = 10
and t.payment_status = 20
and t.delivery_status = 10
and t.refund_status = 10
and t.notify_status = 10
and t.payment_timeout < now()
`
