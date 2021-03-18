package deliverys

var updateForDeliveryFailed = []string{

	//修改发货记录
	`
update ots_trade_delivery t set
t.delivery_status = 90,
t.payment_status = 11,
t.end_time = now(),
t.return_msg = @return_msg,
t.result_source = @result_source,
t.result_code = @result_code,
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 31
`,
	//查询订单号
	`select t.order_id,t.total_face face from ots_trade_delivery t where 
t.delivery_id = @delivery_id`,

	//更新订单状态
	`update ots_trade_order t set
t.bind_face = t.bind_face - @face
where
t.order_id = @order_id 
and t.order_status = 20
and t.delivery_status = 30
and t.bind_face - @face >= 0`,

	//查询订单信息
	`select t.order_id,t.order_status from ots_trade_order t 	
where
t.order_id = @order_id`,
}

var updateForDeliverySuccess = []string{

	//更新发货系统
	`
update ots_trade_delivery t set
t.delivery_status = 0,
t.payment_status = 20,
t.real_discount = case  when  @cost_discount <= 0 then t.cost_discount else  @cost_discount end,
t.cost_amount = case when @cost_discount <= 0 then t.total_face * t.cost_discount  else @cost_discount * t.total_face end,
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
t.order_id,
t.num,
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
t.notify_status = case t.success_face + @succ_face when  t.total_face then 20 else t.notify_status end,
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
	`select t.order_id,t.order_status from ots_trade_order t 	
	where
	t.order_id = @order_id`,
}

//updateDeliveryForSaveUnknown 查询单条数据订单发货表
var updateDeliveryForSaveUnknown = []string{`
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
t.order_id,
o.order_status
from ots_trade_delivery t
inner join  ots_trade_order o on t.order_id = o.order_id
where
t.delivery_id = @delivery_id 
`,
}
