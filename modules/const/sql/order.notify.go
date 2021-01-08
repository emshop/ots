package sql

//UpdateTradeOrderForFinish 查询单条数据订单记录
const UpdateTradeOrderForFinish = `
Update ots_trade_order t
t.notify_status = 0
from ots_trade_order t
where
t.order_id = @order_id
and t.order_status = 50
and t.notify_status = 20
`
