package notices

//下游通知成功
var updateNotifyForSuccess = []string{
	//更新通知信息
	`
update ots_notify_info t set
t.notify_status = 0,
t.end_time = now(),
t.notify_msg = @notify_msg 
where
t.order_id = @order_id 
and t.notify_status = 30
`,

	//更新订单信息
	`
update ots_trade_order t set
t.notify_status = 0,
t.order_status = 60
where
t.order_id = @order_id
and t.order_status = 50
and t.notify_status = 10
`,
}

//updateNotifyInfoForUnknown 查询单条数据订单通知表
const updateNotifyInfoForUnknown = `
update ots_notify_info t set
t.notify_msg = @notify_msg 
where
t.order_id = @order_id
and t.notify_status = 30
`
