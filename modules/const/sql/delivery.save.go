package sql

//UpdateTradeDelivery 查询单条数据订单发货表
const UpdateTradeDelivery = `
Update ots_trade_delivery t
t.delivery_id = @delivery_id,
t.order_id = @order_id,
t.spp_no = @spp_no,
t.spp_product_id = @spp_product_id,
t.spp_delivery_no = @spp_delivery_no,
t.spp_product_no = @spp_product_no,
t.mer_no = @mer_no,
t.mer_product_id = @mer_product_id,
t.pl_id = @pl_id,
t.brand_no = @brand_no,
t.province_no = @province_no,
t.city_no = @city_no,
t.invoice_type = @invoice_type,
t.account_name = @account_name,
t.delivery_status = @delivery_status,
t.payment_status = @payment_status,
t.create_time = @create_time,
t.face = @face,
t.num = @num,
t.total_face = @total_face,
t.cost_amount = @cost_amount,
t.spp_fee_amount = @spp_fee_amount,
t.trade_fee_amount = @trade_fee_amount,
t.payment_fee_amount = @payment_fee_amount,
t.succ_face = @succ_face,
t.start_time = @start_time,
t.end_time = @end_time,
t.return_msg = @return_msg,
t.request_params = @request_params,
t.result_source = @result_source,
t.result_code = @result_code,
t.spp_product_cost = @spp_product_cost,
t.last_update_time = @last_update_time 
from ots_trade_delivery t
where
t.delivery_id = @delivery_id 
and t.delivery_status=30
`
