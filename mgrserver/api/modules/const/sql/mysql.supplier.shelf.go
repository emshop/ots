package sql
//InsertSupplierShelf 添加供货商货架
const InsertSupplierShelf = `
insert into ots_supplier_shelf
(
	spp_shelf_id,
	spp_shelf_id,
	spp_shelf_name,
	spp_no,
	req_url,
	query_url,
	notify_url,
	invoice_type,
	spp_fee_discount,
	trade_fee_discount,
	payment_fee_discount,
	can_refund,
	status,
	limit_count
)
values
(
	@spp_shelf_id,
	@spp_shelf_id,
	@spp_shelf_name,
	@spp_no,
	@req_url,
	@query_url,
	@notify_url,
	@invoice_type,
	@spp_fee_discount,
	@trade_fee_discount,
	@payment_fee_discount,
	@can_refund,
	@status,
	@limit_count
)`

//GetSupplierShelfBySppShelfID 查询单条数据供货商货架
const GetSupplierShelfBySppShelfID = `
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
//UpdateSupplierShelfBySppShelfID 更新供货商货架
const UpdateSupplierShelfBySppShelfID = `
update ots_supplier_shelf 
set
	spp_shelf_name = @spp_shelf_name,
	req_url = @req_url,
	query_url = @query_url,
	notify_url = @notify_url,
	invoice_type = @invoice_type,
	spp_fee_discount = @spp_fee_discount,
	trade_fee_discount = @trade_fee_discount,
	payment_fee_discount = @payment_fee_discount,
	can_refund = @can_refund,
	status = @status,
	limit_count = @limit_count
where
	&spp_shelf_id`

