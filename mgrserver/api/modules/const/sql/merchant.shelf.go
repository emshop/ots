package sql
//GetMerchantShelfByMerShelfId 查询单条数据商户货架
const GetMerchantShelfByMerShelfId = `
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
	t.can_split_order,
	t.status,
	t.create_time
from ots_merchant_shelf t
where
	&mer_shelf_id`

//GetMerchantShelfListCount 获取商户货架列表条数
const GetMerchantShelfListCount = `
select count(1)
from ots_merchant_shelf t
where
	?t.mer_shelf_name
	&t.mer_no`

//GetMerchantShelfList 查询商户货架列表数据
const GetMerchantShelfList = `
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
	t.can_split_order,
	t.status,
	t.create_time 
from ots_merchant_shelf t
where
	?t.mer_shelf_name
	&t.mer_no
order by t.mer_shelf_id desc
limit @ps offset @offset
`
