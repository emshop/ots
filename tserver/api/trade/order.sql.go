package trade

//GetTradeOrderByOrderId查询单条数据订单记录
const GetTradeOrderByOrderId = `
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
t.sell_discount,
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
t.bind_face,
t.success_face,
t.success_sell_amount,
t.success_mer_fee,
t.success_mer_trade_fee,
t.success_mer_payment_fee,
t.success_cost_amount,
t.success_spp_fee,
t.success_spp_trade_fee,
t.success_spp_payment_fee,
t.profit
from ots_trade_order t
where
&order_id`

//GetTradeOrderListCount 获取订单记录列表条数
const GetTradeOrderListCount = `
select count(1)
from ots_trade_order t
where
&t.order_id
&t.mer_no
and if(isnull(@mer_order_no)||@mer_order_no='',1=1,t.mer_order_no like CONCAT('%',@mer_order_no,'%'))
&t.pl_id
&t.brand_no
&t.province_no

`

//and t.create_time>=@create_time and t.create_time<date_add(@create_time, interval 1 day)

//GetTradeOrderList 查询订单记录列表数据
const GetTradeOrderList = `
select
t.order_id,
t.mer_no,
t.pl_id,
t.brand_no,
t.province_no,
t.face,
t.num,
t.account_name,
t.sell_discount,
t.create_time,
t.order_status 
from ots_trade_order t
where
&t.order_id
&t.mer_no
and if(isnull(@mer_order_no)||@mer_order_no='',1=1,t.mer_order_no like CONCAT('%',@mer_order_no,'%'))
&t.pl_id
&t.brand_no
&t.province_no
order by t.order_id desc
limit @ps offset @offset
`

//and t.create_time>=@create_time and t.create_time<date_add(@create_time, interval 1 day)
