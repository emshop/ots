package sql
//InsertMerchantShelf 添加商户货架
const InsertMerchantShelf = `
insert into ots_merchant_shelf
(
	mer_shelf_id,
	mer_shelf_name,
	mer_no,
	mer_fee_discount,
	trade_fee_discount,
	payment_fee_discount,
	order_timeout,
	payment_timeout,
	invoice_type,
	can_refund,
	limit_count,
	can_split_order,
	status
)
values
(
	@mer_shelf_id,
	@mer_shelf_name,
	@mer_no,
	if(isnull(@mer_fee_discount)||@mer_fee_discount='',0,@mer_fee_discount),
	if(isnull(@trade_fee_discount)||@trade_fee_discount='',0,@trade_fee_discount),
	if(isnull(@payment_fee_discount)||@payment_fee_discount='',0,@payment_fee_discount),
	if(isnull(@order_timeout)||@order_timeout='',0,@order_timeout),
	if(isnull(@payment_timeout)||@payment_timeout='',0,@payment_timeout),
	if(isnull(@invoice_type)||@invoice_type='',0,@invoice_type),
	if(isnull(@can_refund)||@can_refund='',0,@can_refund),
	if(isnull(@limit_count)||@limit_count='',0,@limit_count),
	if(isnull(@can_split_order)||@can_split_order='',0,@can_split_order),
	if(isnull(@status)||@status='',0,@status),
)`

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
	mer_shelf_name =	@mer_shelf_name,
	mer_fee_discount =	if(isnull(@mer_fee_discount)||@mer_fee_discount='',0,@mer_fee_discount),
	trade_fee_discount =	if(isnull(@trade_fee_discount)||@trade_fee_discount='',0,@trade_fee_discount),
	payment_fee_discount =	if(isnull(@payment_fee_discount)||@payment_fee_discount='',0,@payment_fee_discount),
	order_timeout =	if(isnull(@order_timeout)||@order_timeout='',0,@order_timeout),
	payment_timeout =	if(isnull(@payment_timeout)||@payment_timeout='',0,@payment_timeout),
	invoice_type =	if(isnull(@invoice_type)||@invoice_type='',0,@invoice_type),
	can_refund =	if(isnull(@can_refund)||@can_refund='',0,@can_refund),
	limit_count =	if(isnull(@limit_count)||@limit_count='',0,@limit_count),
	can_split_order =	if(isnull(@can_split_order)||@can_split_order='',0,@can_split_order),
	status =	if(isnull(@status)||@status='',0,@status),
where
	&mer_shelf_id`

