package sql

//UpdateTradeDeliveryForSaveSuccess 查询单条数据订单发货表
const UpdateTradeDeliveryForSaveSuccess = `
update ots_trade_delivery t set
t.delivery_status = 0,
t.payment_status = 20,
t.real_discount = case @cost_discount when 0 then t.cost_discount else  @cost_discount end,
t.cost_amount = case @cost_discount when 0 then t.total_face * t.cost_discount  else @cost_discount * t.total_face end,
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
where
t.delivery_id = @delivery_id 
and t.delivery_status = 31
`

//UpdateTradeDeliveryForFailed 查询单条数据订单发货表
const UpdateTradeDeliveryForFailed = `
update ots_trade_delivery t set
t.delivery_status = 90,
t.payment_status = 11,
t.succ_face = 0,
t.end_time = now(),
t.return_msg = @return_msg,
t.result_source = @result_source,
t.result_code = @result_code,
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 31
and t.payment_status = 10
`

//UpdateTradeDeliveryForSaveUnknown 查询单条数据订单发货表
const UpdateTradeDeliveryForSaveUnknown = `
update ots_trade_delivery t set
t.return_msg = @return_msg,
t.result_source = @result_source,
t.result_code = @result_code,
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 31
`
