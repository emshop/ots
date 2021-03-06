package sql

//SelectTradeDelivery 查询单条数据订单发货表
const SelectTradeDelivery = `
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
t.succ_face,
t.cost_discount,
t.cost_amount,
t.spp_fee_amount,
t.trade_fee_amount,
t.payment_fee_amount
from ots_trade_delivery t
where
t.delivery_id = @delivery_id 
`

//InsertTradeDelivery 插入单条数据订单发货表
const InsertTradeDelivery = `
insert into ots_trade_delivery(
delivery_id,
order_id,
spp_no,
spp_product_id,
spp_product_no,
mer_no,
mer_product_id,
pl_id,
brand_no,
province_no,
city_no,
invoice_type,
account_name,
delivery_status,
payment_status,
face,
num,
total_face,
cost_discount,
spp_fee_discount,
trade_fee_discount,
payment_fee_discount
)values(
@delivery_id,
@order_id,
@spp_no,
@spp_product_id,
@spp_product_no,
@mer_no,
@mer_product_id,
@pl_id,
@brand_no,
@province_no,
@city_no,
@invoice_type,
@account_name,
20,
10,
@face,
1,
@face,
@cost_discount,
@spp_fee_discount,
@trade_fee_discount,
@payment_fee_discount
)`
