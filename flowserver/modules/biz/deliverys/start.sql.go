package deliverys

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
and exists(select 1 from ots_trade_order o where o.order_id = t.order_id and o.order_status = 20 and  o.delivery_pause = 1)
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
and o.order_status = 20
and o.delivery_status in(20,30)
and o.delivery_pause = 1
`,
}

//updateDeliveryForDeliveryingFailed 查询单条数据订单发货表
var updateDeliveryForDeliveryingFailed = []string{`
update ots_trade_delivery t set
t.delivery_status = 90,
t.return_msg = @return_msg,
t.result_source = "delivery",
t.result_code = @result_code,
t.end_time = now(),
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 30
`,
	//查询发货状态
	`
select
t.total_face
from ots_trade_delivery t
where
t.delivery_id = @delivery_id
and t.delivery_status = 90`,

	//更改订单状态
	`
update ots_trade_order t set
t.bind_face = t.bind_face - @total_face
where
t.order_id = @order_id 
and t.order_status = 20
and t.delivery_status = 30
and t.bind_face > @total_face
`,
}

//updateDeliveryForDeliveryingSuccess 查询单条数据订单发货表
const updateDeliveryForDeliveryingSuccess = `
update ots_trade_delivery t set
t.delivery_status = 31,
t.cost_amount=case when @cost_discount <=0 then t.cost_amount else @cost_discount * t.total_face end,
t.return_msg = @return_msg,
t.result_source = "delivery",
t.result_code = @result_code,
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 30
`
