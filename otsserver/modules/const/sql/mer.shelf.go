package sql

//SelectMerchantShelf 查询单条数据商户货架
const SelectMerchantShelf = `
select
t.mer_shelf_id,
t.mer_shelf_name,
t.mer_no,
t.mer_fee_discount,
t.trade_fee_discount,
t.payment_fee_discount,
t.order_timeout,
t.payment_timeout,
t.invoice_type,
t.can_refund,
t.limit_count,
t.can_split_order
from ots_merchant_shelf t
inner join ots_merchant_info m on t.mer_no = m.mer_no
where
t.mer_shelf_id = @mer_shelf_id 
and t.mer_no = @mer_no
and t.status = 0
and m.status = 0
`
