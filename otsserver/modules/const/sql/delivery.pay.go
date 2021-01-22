package sql

//UpdateTradeDeliveryForPaying 查询单条数据订单发货表
const UpdateTradeDeliveryForPaying = `
update ots_trade_delivery t set
t.payment_status = 0
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 0
and t.payment_status = 20
`
