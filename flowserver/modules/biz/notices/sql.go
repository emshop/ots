package notices

var updateNoNeedNotices = []string{

	//查询通知订单是否为无须通知
	`
	select
	t.order_id
	from ots_notify_info t
	where
	t.order_id = @order_id 
	and t.notify_status = 11
	`,
	//更新通知信息
	`
update ots_notify_info t set
t.start_time = now(),
t.end_time = now(),
t.notify_msg = '无须通知' 
where
t.order_id = @order_id 
and t.notify_status = 11
`,

	//修改订单状态
	`
update ots_trade_order t set
t.notify_status = 11,
t.order_status = 60
where
t.order_id = @order_id
and t.order_status = 50
and t.notify_status = 20
`,
}

var startNotices = []string{

	//更新发货记录
	`
update ots_notify_info t set
t.notify_status = case t.notify_status when 10 then 30 else t.notify_status = t.notify_status end,
t.start_time = case t.notify_status when 10 then now() else t.start_time end,
t.notify_count = case t.notify_status when 10 then 1 else t.notify_count + 1 end
where
t.order_id = @order_id 
and t.notify_status in(10,30)
and t.notify_count < t.max_count
and exists(select 1 from ots_trade_order t where t.order_id = @order_id and t.notify_status = 10 and t.order_status = 50)
`,

	//查询通知信息
	`
select
n.notify_url,
t.order_id,
t.mer_no,
t.mer_order_no,
t.account_name,
t.invoice_type,
t.sell_discount,
t.sell_amount
from ots_notify_info n
inner join ots_trade_order t on t.order_id = n.order_id
where
n.order_id = @order_id 
and n.notify_status = 30
`,
}

//UpdateNotifyInfoForSuccess 查询单条数据订单通知表
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
