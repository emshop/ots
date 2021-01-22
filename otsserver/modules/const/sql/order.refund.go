package sql

//UpdateTradeOrderForRefund 查询单条数据订单记录
const UpdateTradeOrderForRefund = `
update ots_trade_order t set
t.refund_status = 0,
t.is_refund = 0
t.order_status = 50
t.notify_status = 20
where
t.order_id = @order_id
and t.bind_face = 0
and t.success_face = 0
and t.payment_status = 0
and is_refund = 1
and t.order_status = 70
and t.delivery_status = 90
and t.refund_status = 20,
and t.order_timeout < now()
`
