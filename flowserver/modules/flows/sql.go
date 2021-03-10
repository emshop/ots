package flows

//SelectProductFlowByTag 查询业务流程
const SelectProductFlowByTag = `
select
t.flow_id,
t.flow_tag,
t.pl_id,
t.queue_name,
t.scan_interval,
t.delay,
t.timeout,
t.max_count
from ots_product_flow t
where
(t.pl_id = @pl_id or t.pl_id = 0)
and t.flow_tag = @flow_tag
and t.status = 0
order by t.pl_id desc
limit 1
`

//SelectTradeOrderByOrderID 查询单条数据订单记录
const SelectTradeOrderByOrderID = `
select
t.order_id,
t.mer_no,
t.mer_order_no,
t.mer_product_id,
t.mer_shelf_id,
t.mer_product_no,
t.pl_id,
t.brand_no,
t.province_no,
t.city_no,
t.face,
t.num,
t.total_face,
t.account_name,
t.invoice_type,
t.sell_amount,
t.mer_fee_discount,
t.trade_fee_discount,
t.payment_fee_discount,
t.can_split_order,
t.create_time,
t.order_timeout,
t.payment_timeout,
t.delivery_pause,
t.order_status,
t.payment_status,
t.delivery_status,
t.refund_status,
t.notify_status,
t.is_refund,
t.bind_face
from ots_trade_order t
where
t.order_id = @order_id 
`

//selectDeliveryByDeliveryID 查询单条数据订单发货表
const selectDeliveryByDeliveryID = `
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
t.delivery_id = @delivery_id`
