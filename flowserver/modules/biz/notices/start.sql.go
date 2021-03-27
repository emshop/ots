package notices

//处理无须处理订单
var updateNoNeedNotices = []string{

	//修改订单状态
	`
update ots_trade_order t set
t.order_status = 60
where
t.order_id = @order_id
and t.order_status = 50
and t.notify_status = 11
`,
}

var startNotices = []string{

	//更新通知信息
	`
update ots_notify_info t set
t.start_time = case t.notify_status when 10 then now() else t.start_time end,
t.notify_status = case t.notify_status when 10 then 30 else t.notify_status = t.notify_status end,
t.notify_count = t.notify_count + 1
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
t.delivery_status status,
t.mer_no,
t.mer_order_no,
t.account_name,
t.sell_amount
from ots_notify_info n
inner join ots_trade_order t on t.order_id = n.order_id
where
n.order_id = @order_id 
and n.notify_status = 30
and t.order_status = 50
and t.delivery_status in(0, 90)
`,
}
