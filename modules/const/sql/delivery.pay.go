package sql

//UpdateTradeDeliveryForPaying 查询单条数据订单发货表
const UpdateTradeDeliveryForPaying = `
Update ots_trade_delivery t
t.payment_status = 0
t.last_update_time = now() 
from ots_trade_delivery t
where
t.delivery_id = @delivery_id 
and t.delivery_status = 0
and t.payment_status = 20
`
