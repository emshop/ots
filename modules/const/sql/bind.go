package sql

//UpdateTradeOrderForBind 查询单条数据订单记录
const UpdateTradeOrderForBind = `
Update ots_trade_order t
t.delivery_status = 30,
t.bind_face = t.bind_face + @face
from ots_trade_order t
where
t.order_id = @order_id 
and t.order_status = 20
and t.delivery_pause = 1
and (t.delivery_status = 30 or t.delivery_status = 20)
and t.bind_face + @face <= t.total_face
`
