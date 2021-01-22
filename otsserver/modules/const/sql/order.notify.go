package sql

//UpdateTradeOrderForFinish 查询单条数据订单记录
const UpdateTradeOrderForFinish = `
Update ots_trade_order t set
t.notify_status = 0
where
t.order_id = @order_id
and t.order_status = 50
and t.notify_status = 20
`
