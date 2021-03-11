package finishes

//订单支付成功、发货成功 => 订单状态为成功,退款关闭
var update2Success = []string{

	`update ots_trade_order t set
	t.order_status = 0,
	t.refund_status = 11,
	t.finish_time = now()
	where
	t.order_id = @order_id
	and t.order_status = 60
	and t.delivery_status = 0
	and t.payment_status = 0
	and t.refund_status = 10`,
}

//订单支付失败，未发货 => 订单状态为失败,退款关闭
var update2Faild = []string{
	`update ots_trade_order t set
	t.order_status = 90,
	t.refund_status = 11,
	t.finish_time = now()
	where
	t.order_id = @order_id
	and t.order_status = 60
	and t.delivery_status = 90
	and t.payment_status = 90
	and t.refund_status = 10`,
}

//支付成功，发货失败 => 退款，订单失败
var update2Refund = []string{
	//查询订单信息
	`select
	 t.order_id,
	 t.mer_no,
	 t.mer_order_no,
	 t.face,
	 t.num,
	 t.total_face,
	 t.sell_amount
	 from ots_trade_order t
	 where
	 t.order_id = @order_id 
	 and t.order_status = 60
	 and t.payment_status = 0
	 and t.delivery_status = 90`,

	//将订单状态修改为成功
	`update ots_trade_order t set
	t.refund_status = 0,
	t.order_status = 90,
	t.finish_time = now()
	where
	t.order_id = @order_id
	and t.order_status = 60
	and t.delivery_status = 90
	and t.success_face = 0
	and t.payment_status = 0
	and t.refund_status = 10`,
}

var forceRefund = []string{

	//查询订单信息
	`select
	 t.order_id,
	 t.mer_no,
	 t.mer_order_no,
	 t.face,
	 t.num,
	 t.total_face,
	 t.sell_amount
	 from ots_trade_order t
	 where
	 t.order_id = @order_id 
	 and t.order_status = 60`,

	//将订单状态修改为成功
	`update ots_trade_order t set
	t.refund_status = 0,
	t.success_sell_amount = 0,
	t.success_mer_fee = 0,
	t.success_mer_trade_fee = 0,
	t.success_mer_payment_fee = 0,
	t.profit=-1*(t.success_cost_amount-t.success_spp_fee+t.success_spp_trade_fee+t.success_spp_payment_fee)
	where
	t.order_id = @order_id
	and t.order_status = 60
	and t.delivery_status = 90
	and t.success_face = 0
	and t.payment_status = 0
	and t.refund_status = 10`,
}
