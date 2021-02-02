package sql
//GetMerchantShelfByMerShelfID 查询单条数据商户货架
const GetMerchantShelfByMerShelfID = `
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
//UpdateMerchantShelfByMerShelfID 更新商户货架
const UpdateMerchantShelfByMerShelfID = `
update ots_merchant_shelf 
set
	mer_shelf_name = @mer_shelf_name,
	mer_fee_discount = @mer_fee_discount,
	trade_fee_discount = @trade_fee_discount,
	payment_fee_discount = @payment_fee_discount,
	order_timeout = @order_timeout,
	payment_timeout = @payment_timeout,
	invoice_type = @invoice_type,
	can_refund = @can_refund,
	limit_count = @limit_count,
	can_split_order = @can_split_order,
	status = @status
where
	&mer_shelf_id`

