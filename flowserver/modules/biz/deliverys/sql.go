package deliverys

//查询订单是否需要绑定
var checkOrderForBind = `select
t.order_id,
t.mer_no,
t.mer_order_no,
t.mer_product_id,
t.mer_shelf_id,
t.pl_id,
t.brand_no,
t.province_no,
t.city_no,
case when t.can_split_order = 0 then t.total_face-t.bind_face else t.face end 'face',
t.account_name,
t.invoice_type,
t.can_split_order,
t.sell_discount,
t.order_timeout,
t.payment_timeout,
t.delivery_pause,
t.order_status,
t.payment_status,
t.delivery_status,
t.total_face,
t.bind_face
	from ots_trade_order t
	where
	t.order_id = @order_id
	and t.order_status = 20
	and t.delivery_status = 20
	and t.total_face - t.bind_face > 0
	`

//查询符合条件的产品信息
var querySppProducts = []string{

	//查询待绑定订单及面值
	`select
	t.order_id,
	t.mer_no,
	t.mer_order_no,
	t.mer_product_id,
	t.mer_shelf_id,
	t.pl_id,
	t.brand_no,
	t.province_no,
	t.city_no,
	case when t.can_split_order = 0 then t.total_face-t.bind_face else t.face end 'face',
	t.account_name,
	t.invoice_type,
	t.can_split_order,
	t.sell_discount,
	t.order_timeout,
	t.payment_timeout,
	t.delivery_pause,
	t.order_status,
	t.payment_status,
	t.delivery_status,
	t.total_face,
	t.bind_face
	from ots_trade_order t
	where
	t.order_id = @order_id
	and t.order_status = 20
	and t.delivery_status = 20
	and t.total_face - t.bind_face > 0
	`,

	//查询符合条件的商品信息
	`
select
t.spp_product_id,
t.spp_shelf_id,
t.spp_no,
t.spp_product_no,
t.face,
t.cost_discount
from ots_supplier_product t
inner join ots_supplier_info s on t.spp_no = s.spp_no
where
t.pl_id = @pl_id and t.pl_id !=0
and t.brand_no = @brand_no
and t.province_no = @province_no
and t.city_no = @city_no
and case when @can_split_order = 0 then t.face <= @face else t.face = @face end
and t.cost_discount <= @sell_discount
and t.status = 0
and s.status = 0
order by t.cost_discount asc
`,
}

//创建绑定记录
var binds = []string{
	//查询SPP信息
	`
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
	t.payment_fee_discount
	from ots_supplier_shelf t
	where
	t.spp_shelf_id = @spp_shelf_id 
	and t.status = 0
	`,

	//更新订单信息
	`
	update ots_trade_order t set
	t.delivery_status = 30,
	t.bind_face = t.bind_face + @face
	where
	t.order_id = @order_id 
	and t.order_status = 20
	and t.delivery_pause = 1
	and (t.delivery_status = 30 or t.delivery_status = 20)
	and t.bind_face + @face <= t.total_face
	`,

	//添加发货记录
	`
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
)`,
}

//处理开始发货
var deliveryStartNow = []string{
	//更改发货状态为正在发货
	`
update ots_trade_delivery t set
t.delivery_status = 30,
t.start_time = now(),
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 20
and exists(select 1 from ots_trade_order o where o.order_id = t.order_id and o.delivery_pause = 1)
`,
	//查询发货信息
	`
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
t.face,
t.num,
t.total_face,
t.cost_amount,
t.spp_fee_amount,
t.trade_fee_amount,
t.payment_fee_amount
from ots_trade_delivery t
inner join ots_trade_order o on t.order_id = o.order_id
where
t.delivery_id = @delivery_id 
and t.delivery_status = 30
and o.delivery_pause = 1
`,
}

//queryDealCode 查询错误码
const queryDealCode = `
select
t.id,
t.spp_no,
t.pl_id,
t.deal_code,
t.error_code,
t.error_desc
from ots_supplier_ecode t
inner join ots_trade_delivery d on t.spp_no = d.spp_no
where
d.delivery_id = @delivery_id 
and t.category = @category
and (t.pl_id = d.pl_id or t.pl_id = 0)
and t.error_code = @error_code
and t.status = 0
order by t.pl_id desc
limit 1
`

//UpdateTradeDeliveryForDeliveryingFailed 查询单条数据订单发货表
const UpdateTradeDeliveryForDeliveryingFailed = `
update ots_trade_delivery t set
t.delivery_status = 39,
t.return_msg = @return_msg,
t.result_source = "delivery",
t.result_code = @result_code
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 30
`

//UpdateTradeDeliveryForDeliveryingSuccess 查询单条数据订单发货表
const UpdateTradeDeliveryForDeliveryingSuccess = `
update ots_trade_delivery t set
t.delivery_status = 31,
t.cost_amount=case @cost_discount when 0 then t.cost_amount else @cost_discount * t.total_face end,
t.return_msg = @return_msg,
t.result_source = "delivery",
t.result_code = @result_code,
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 30
`

var updateForDeliveryFailed = []string{

	//修改发货记录
	`
update ots_trade_delivery t set
t.delivery_status = 90,
t.payment_status = 11,
t.succ_face = 0,
t.end_time = now(),
t.return_msg = @return_msg,
t.result_source = @result_source,
t.result_code = @result_code,
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 31
and t.payment_status = 10
`,
	//查询订单信息
	`select t.order_id from ots_trade_delivery t where 
t.delivery_id = @delivery_id`,

	//更新订单状态
	`update ots_trade_order t set
t.refund_status = case when t.bind_face - @face = 0 and t.order_timeout < now() and t.success_face = 0 then 20 else t.refund_status end,
t.order_status = case when t.bind_face - @face = 0 and t.order_timeout < now() and t.success_face = 0 then 70 else t.order_status end,
t.delivery_status =case when t.bind_face - @face = 0 and t.order_timeout < now() and t.success_face = 0 then 90 else t.delivery_status end,
t.bind_face = t.bind_face - @face
where
t.order_id = @order_id 
and t.order_status = 20
and t.delivery_status = 30
and t.bind_face - @face >= 0`,
}

var updateForDeliverySuccess = []string{

	//更新发货系统
	`
update ots_trade_delivery t set
t.delivery_status = 0,
t.payment_status = 20,
t.real_discount = case @cost_discount when 0 then t.cost_discount else  @cost_discount end,
t.cost_amount = case @cost_discount when 0 then t.total_face * t.cost_discount  else @cost_discount * t.total_face end,
t.succ_face = t.total_face,
t.spp_fee_amount = t.total_face * t.spp_fee_discount,
t.trade_fee_amount =  t.total_face * t.trade_fee_discount,
t.payment_fee_amount = t.total_face * t.payment_fee_discount,
t.end_time = now(),
t.return_msg = @return_msg,
t.request_params = @request_params,
t.result_source = @result_source,
t.result_code = @result_code,
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 31
`,
	//查询发货信息
	`
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
`,

	//更新订单信息
	`
update ots_trade_order t set
t.order_status = case t.success_face + @succ_face when t.total_face then 50 else t.order_status end,
t.notify_status = case when (t.success_face + @succ_face = t.total_face and t.notify_status = 10) then 20 else t.notify_status end,
t.delivery_status = case t.success_face + @succ_face when t.total_face then 0 else t.delivery_status end,
t.success_face = t.success_face + @succ_face,
t.success_sell_amount = (t.success_sell_amount + @succ_face) * t.sell_discount,
t.success_mer_fee = t.success_mer_fee +  @succ_face * t.mer_fee_discount,
t.success_mer_trade_fee = t.success_mer_trade_fee + @succ_face * t.trade_fee_discount,
t.success_mer_payment_fee = t.success_mer_payment_fee + @succ_face * t.payment_fee_discount,
t.success_cost_amount = t.success_cost_amount + @cost_amount,
t.success_spp_fee = t.success_spp_fee + @spp_fee_amount,
t.success_spp_trade_fee = t.success_spp_trade_fee + @trade_fee_amount,
t.success_spp_payment_fee = t.success_spp_payment_fee + @payment_fee_amount
where
t.order_id = @order_id
and t.order_status = 20
and t.delivery_status = 30
`,
	//查询订单信息
	`select t.order_status from ots_trade_order t 	
	where
	t.order_id = @order_id`,
}

//updateTradeDeliveryForSaveUnknown 查询单条数据订单发货表
var updateTradeDeliveryForSaveUnknown = []string{`
update ots_trade_delivery t set
t.return_msg = @return_msg,
t.result_source = @result_source,
t.result_code = @result_code,
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 31
`,

	//查询发货基本信息
	`
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
t.city_no
from ots_trade_delivery t
where
t.delivery_id = @delivery_id 
`,
}
