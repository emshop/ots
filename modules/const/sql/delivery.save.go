package sql

//UpdateTradeDeliveryForSaveSuccess 查询单条数据订单发货表
const UpdateTradeDeliveryForSaveSuccess = `
Update ots_trade_delivery t
t.delivery_status = 0,
t.payment_status = 20,
case when @cost_discount = 0 then t.real_discount = t.cost_discount else t.real_discount = @cost_discount end,
case when @cost_discount = 0 then t.cost_amount = t.total_face * t.cost_discount  else t.cost_amount = @cost_discount * t.total_face end,
t.succ_face = t.total_face,
t.spp_fee_amount = t.total_face * t.spp_fee_discount,
t.trade_fee_amount =  t.total_face * t.trade_fee_discount,
t.payment_fee_amount = t.total_face * t.payment_fee_discount,
t.end_time = now(),
t.return_msg = @return_msg,
t.request_params = @request_params,
t.result_source = @result_source,
t.result_code = @result_code,
t.last_update_time = now() 
from ots_trade_delivery t
where
t.delivery_id = @delivery_id 
and t.delivery_status = 31
`

//UpdateTradeDeliveryForFailed 查询单条数据订单发货表
const UpdateTradeDeliveryForFailed = `
Update ots_trade_delivery t
t.delivery_status = 90,
t.payment_status = 10,
t.succ_face = 0,
t.end_time = now(),
t.return_msg = @return_msg,
t.result_source = @result_source,
t.result_code = @result_code,
t.last_update_time = now() 
from ots_trade_delivery t
where
t.delivery_id = @delivery_id 
and t.delivery_status = 31
`

//UpdateTradeDeliveryForSaveUnkown 查询单条数据订单发货表
const UpdateTradeDeliveryForSaveUnkown = `
Update ots_trade_delivery t
t.return_msg = @return_msg,
t.result_source = @result_source,
t.result_code = @result_code,
t.last_update_time = now() 
from ots_trade_delivery t
where
t.delivery_id = @delivery_id 
and t.delivery_status = 31
`
