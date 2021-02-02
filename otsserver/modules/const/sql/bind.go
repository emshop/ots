package sql

//UpdateTradeOrderForBind 查询单条数据订单记录
const UpdateTradeOrderForBind = `
update ots_trade_order t set
t.delivery_status = 30,
t.bind_face = t.bind_face + @face
where
t.order_id = @order_id 
and t.order_status = 20
and t.delivery_pause = 1
and (t.delivery_status = 30 or t.delivery_status = 20)
and t.bind_face + @face <= t.total_face
`

//UpdateTradeOrderForDeiveryFailed 查询单条数据订单记录
const UpdateTradeOrderForDeiveryFailed = `
update ots_trade_order t set
t.refund_status = case when t.bind_face - @face = 0 and t.order_timeout < now() and t.success_face = 0 then 20 else t.refund_status end,
t.order_status = case when t.bind_face - @face = 0 and t.order_timeout < now() and t.success_face = 0 then 70 else t.order_status end,
t.delivery_status =case when t.bind_face - @face = 0 and t.order_timeout < now() and t.success_face = 0 then 90 else t.delivery_status end,
t.bind_face = t.bind_face - @face
where
t.order_id = @order_id 
and t.order_status = 20
and t.delivery_status = 30
and t.bind_face - @face >= 0
`
