package deliverys

//updateDeliveryForPaying 查询单条数据订单发货表
var updateDeliveryForPaying = []string{
	//更新发货信息
	`
update ots_trade_delivery t set
t.payment_status = 0,
t.last_update_time = now() 
where
t.delivery_id = @delivery_id 
and t.delivery_status = 0
and t.payment_status = 20
`,
	//查询支付数据
	`
select
t.delivery_id,
t.spp_no,
t.cost_amount
from ots_trade_delivery t
where
t.delivery_id = @delivery_id 
`,
}
