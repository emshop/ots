package sql
//InsertTradeDelivery 添加订单发货表
const InsertTradeDelivery = `
insert into ots_trade_delivery
(
	
	is_mf
)
values
(
	
	if(isnull(@is_mf)||@is_mf='',0,@is_mf)
)`
//GetTradeDeliveryByDeliveryID 查询订单发货表单条数据
const GetTradeDeliveryByDeliveryID = `
select
	t.delivery_id,
	t.order_id,
	t.spp_no,
	t.spp_product_id,
	t.mer_product_id,
	t.pg_id,
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
	t.cost_discount,
	t.real_discount,
	t.spp_fee_discount,
	t.trade_fee_discount,
	t.payment_fee_discount,
	t.is_mf,
	t.cost_amount,
	t.spp_fee_amount,
	t.trade_fee_amount,
	t.payment_fee_amount,
	t.succ_face,
	t.start_time,
	t.end_time,
	t.spp_delivery_no,
	t.spp_product_no,
	t.return_msg,
	t.request_params,
	t.result_source,
	t.result_code,
	t.last_update_time,
	t.batch_id
from ots_trade_delivery t
where
	&delivery_id`

//GetTradeDeliveryListCount 获取订单发货表列表条数
const GetTradeDeliveryListCount = `
select count(1)
from ots_trade_delivery t
where
	&t.order_id
	&t.spp_no
	&t.pl_id
	&t.delivery_status
	and t.create_time >= @create_time 
	and t.create_time < date_add(@create_time, interval 1 day)`

//GetTradeDeliveryList 查询订单发货表列表数据
const GetTradeDeliveryList = `
select
	t.delivery_id,
	t.order_id,
	t.spp_no,
	t.pg_id,
	t.pl_id,
	t.brand_no,
	t.province_no,
	t.account_name,
	t.delivery_status,
	t.create_time,
	t.face,
	t.end_time,
	t.return_msg 
from ots_trade_delivery t
where
	&t.order_id
	&t.spp_no
	&t.pl_id
	&t.delivery_status
	and t.create_time >= @create_time 
	and t.create_time < date_add(@create_time, interval 1 day)
order by t.delivery_id desc
limit @ps offset @offset
`
//GetTradeDeliveryDetailListCount 获取订单发货表列表条数
const GetTradeDeliveryDetailListCount = `
select count(1)
from ots_trade_delivery t
where
&order_id`

//GetTradeDeliveryDetailList 查询订单发货表列表数据
const GetTradeDeliveryDetailList = `
select
	t.delivery_id,
	t.order_id,
	t.spp_no,
	t.pg_id,
	t.pl_id,
	t.brand_no,
	t.province_no,
	t.account_name,
	t.delivery_status,
	t.create_time,
	t.face,
	t.end_time,
	t.return_msg 
from ots_trade_delivery t
where
&order_id
limit @ps offset @offset`
//UpdateTradeDeliveryByDeliveryID 更新订单发货表
const UpdateTradeDeliveryByDeliveryID = `
update ots_trade_delivery 
set
	is_mf =	if(isnull(@is_mf)||@is_mf='',0,@is_mf)
where
	&delivery_id`

