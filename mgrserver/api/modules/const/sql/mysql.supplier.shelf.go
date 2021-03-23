
// +build mysql

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
	spp_fee_discount,
	trade_fee_discount,
	payment_fee_discount,
	limit_count,
	invoice_type,
	can_refund,
	status
)
values
(
	@spp_shelf_id,
	if(isnull(@spp_shelf_id)||@spp_shelf_id='',0,@spp_shelf_id),
	@spp_shelf_name,
	@spp_no,
	@req_url,
	@query_url,
	@notify_url,
	if(isnull(@spp_fee_discount)||@spp_fee_discount='',0,@spp_fee_discount),
	if(isnull(@trade_fee_discount)||@trade_fee_discount='',0,@trade_fee_discount),
	if(isnull(@payment_fee_discount)||@payment_fee_discount='',0,@payment_fee_discount),
	if(isnull(@limit_count)||@limit_count='',0,@limit_count),
	if(isnull(@invoice_type)||@invoice_type='',0,@invoice_type),
	if(isnull(@can_refund)||@can_refund='',0,@can_refund),
	if(isnull(@status)||@status='',0,@status)
)`

//GetSupplierShelfBySppShelfID 查询供货商货架单条数据
const GetSupplierShelfBySppShelfID = `
select
	t.spp_shelf_id,
	t.spp_shelf_name,
	t.spp_no,
	t.req_url,
	t.query_url,
	t.notify_url,
	t.spp_fee_discount,
	t.trade_fee_discount,
	t.payment_fee_discount,
	t.limit_count,
	t.invoice_type,
	t.can_refund,
	t.create_time,
	t.status
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
	t.spp_fee_discount,
	t.trade_fee_discount,
	t.payment_fee_discount,
	t.invoice_type,
	t.can_refund,
	t.status 
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
	spp_shelf_name =	@spp_shelf_name,
	req_url =	@req_url,
	query_url =	@query_url,
	notify_url =	@notify_url,
	spp_fee_discount =	if(isnull(@spp_fee_discount)||@spp_fee_discount='',0,@spp_fee_discount),
	trade_fee_discount =	if(isnull(@trade_fee_discount)||@trade_fee_discount='',0,@trade_fee_discount),
	payment_fee_discount =	if(isnull(@payment_fee_discount)||@payment_fee_discount='',0,@payment_fee_discount),
	limit_count =	if(isnull(@limit_count)||@limit_count='',0,@limit_count),
	invoice_type =	if(isnull(@invoice_type)||@invoice_type='',0,@invoice_type),
	can_refund =	if(isnull(@can_refund)||@can_refund='',0,@can_refund),
	status =	if(isnull(@status)||@status='',0,@status)
where
	&spp_shelf_id`

