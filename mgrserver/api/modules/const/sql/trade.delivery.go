package sql
//GetTradeDeliveryByDeliveryId 查询单条数据订单发货表
const GetTradeDeliveryByDeliveryId = `
select
	t.pl_id,
	t.invoice_type,
	t.face,
	t.num,
	t.total_face,
	t.cost_discount,
	t.spp_fee_discount,
	t.trade_fee_discount,
	t.payment_fee_discount,
	t.cost_amount,
	t.spp_fee_amount,
	t.trade_fee_amount,
	t.payment_fee_amount,
	t.succ_face,
	t.start_time,
	t.end_time,
	t.spp_delivery_no,
	t.spp_product_no,
	t.real_discount,
	t.return_msg,
	t.request_params,
	t.result_source,
	t.result_code,
	t.last_update_time
from ots_trade_delivery t
where
	&delivery_id`

//GetTradeDeliveryListCount 获取订单发货表列表条数
const GetTradeDeliveryListCount = `
select count(1)
from ots_trade_delivery t
where
	&t.pl_id`

//GetTradeDeliveryList 查询订单发货表列表数据
const GetTradeDeliveryList = `
select
	t.delivery_id,
	t.pl_id,
	t.brand_no,
	t.province_no,
	t.city_no,
	t.account_name,
	t.delivery_status,
	t.payment_status,
	t.create_time 
from ots_trade_delivery t
where
	&t.pl_id
order by t.delivery_id desc
limit @ps offset @offset
`
