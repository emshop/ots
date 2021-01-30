package sql
//GetSupplierShelfBySppShelfId 查询单条数据供货商货架
const GetSupplierShelfBySppShelfId = `
select
	t.spp_shelf_id,
	t.spp_shelf_name,
	t.spp_no,
	t.req_url,
	t.query_url,
	t.notify_url,
	t.invoice_type,
	t.spp_fee_discount,
	t.trade_fee_discount,
	t.payment_fee_discount,
	t.can_refund,
	t.status,
	t.limit_count,
	t.create_time
from ots_supplier_shelf t
where
	&spp_shelf_id`

//GetSupplierShelfListCount 获取供货商货架列表条数
const GetSupplierShelfListCount = `
select count(1)
from ots_supplier_shelf t
where
	?t.spp_shelf_name
	&t.spp_no
	&t.status`

//GetSupplierShelfList 查询供货商货架列表数据
const GetSupplierShelfList = `
select
	t.spp_shelf_id,
	t.spp_shelf_name,
	t.spp_no,
	t.invoice_type,
	t.spp_fee_discount,
	t.trade_fee_discount,
	t.payment_fee_discount,
	t.can_refund,
	t.status,
	t.create_time 
from ots_supplier_shelf t
where
	?t.spp_shelf_name
	&t.spp_no
	&t.status
order by t.spp_shelf_id desc
limit @ps offset @offset
`
