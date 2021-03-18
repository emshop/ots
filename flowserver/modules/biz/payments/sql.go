package payments

//dealTimeoutOrder 处理超时订单
var dealTimeoutOrder = []string{
	//检查订单是否超时，超时则修改为订单通知
	`
update ots_trade_order t set
t.payment_status = 90,
t.delivery_status = 90,
t.order_status = 50,
t.notify_status = 20
where
t.order_id = @order_id
and t.order_status = 10
and t.payment_status = 20
and t.delivery_status = 10
and t.notify_status = 10
and (t.payment_timeout < now() or t.order_timeout < now())
`,
}

var dealOrderPayStart = []string{
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
and t.order_status = 10
and t.payment_status = 20`,

	//将订单状态修改为支付成功,发货开始
	`update ots_trade_order t set
	t.order_status = 20,
t.payment_status = 0,
t.delivery_status = 20
where
t.order_id = @order_id
and t.order_status = 10
and t.payment_status = 20
and t.delivery_status = 10
`,
}
