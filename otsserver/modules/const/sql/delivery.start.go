package sql

//UpdateTradeDeliveryForStart 查询单条数据订单发货表
const UpdateTradeDeliveryForStart = `
update ots_trade_delivery t set
t.delivery_status = 30,
t.start_time = now(),
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 20
and exists(select 1 from ots_trade_order o where o.order_id = t.order_id and o.delivery_pause = 1)
`

//SelectTradeDeliveryForStart 查询单条数据订单发货表
const SelectTradeDeliveryForStart = `
select
t.delivery_id,
t.order_id,
t.spp_no,
t.spp_product_id,
t.spp_delivery_no,
t.spp_product_no,
t.mer_no,
t.mer_product_id,
t.pl_id,
t.brand_no,
t.province_no,
t.city_no,
t.invoice_type,
t.account_name,
t.delivery_status,
t.payment_status,
t.create_time,
t.face,
t.num,
t.total_face,
t.cost_amount,
t.spp_fee_amount,
t.trade_fee_amount,
t.payment_fee_amount
from ots_trade_delivery t
inner join ots_trade_order o on t.order_id = o.order_id
where
t.delivery_id = @delivery_id 
and t.delivery_id = 30
and o.delivery_pause = 1
`

//UpdateTradeDeliveryForDeliveryingSuccess 查询单条数据订单发货表
const UpdateTradeDeliveryForDeliveryingSuccess = `
update ots_trade_delivery t set
t.delivery_status = 31,
t.cost_amount=case @cost_discount when 0 then t.cost_amount else @cost_discount * t.total_face end,
t.return_msg = @return_msg,
t.result_source = "delivery",
t.result_code = @result_code,
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 30
`

//UpdateTradeDeliveryForDeliveryingFailed 查询单条数据订单发货表
const UpdateTradeDeliveryForDeliveryingFailed = `
update ots_trade_delivery t set
t.delivery_status = 39,
t.return_msg = @return_msg,
t.result_source = "delivery",
t.result_code = @result_code
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 30
`
