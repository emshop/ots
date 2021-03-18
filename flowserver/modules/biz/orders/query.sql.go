package orders

//SelectTradeOrderByMerchant 通过mer_no,mer_order_no获取订单
const SelectTradeOrderByMerchant = `
select
t.order_id,
t.mer_no,
t.mer_order_no,
t.account_name,
t.face,
t.num,
FORMAT(t.sell_amount,2) 'amount',
t.total_face,
case t.order_status when 0 then "SUCCESS" when 91 then "SUCCESS" when 90 then 'FAILED' else 'UNDERWAY' end 'order_status'
from ots_trade_order t
where t.mer_no = @mer_no and t.mer_order_no=@mer_order_no limit 1`

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