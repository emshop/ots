package sql//GetNotifyInfoByOrderID 查询订单通知表单条数据
const GetNotifyInfoByOrderID = `
select
	t.order_id,
	t.mer_no,
	t.mer_order_no,
	t.notify_url,
	t.create_time,
	t.start_time,
	t.end_time,
	t.max_count,
	t.notify_count,
	t.notify_status,
	t.notify_msg
from ots_notify_info t
where
	&order_id`
//GetNotifyInfoDetail 查询订单通知表单条详情数据
const GetNotifyInfoDetail= `
select
	t.order_id,
	t.mer_no,
	t.mer_order_no,
	t.notify_url,
	t.create_time,
	t.start_time,
	t.end_time,
	t.max_count,
	t.notify_count,
	t.notify_status,
	t.notify_msg
from ots_notify_info t
where
	&order_id
`

//GetNotifyInfoListCount 获取订单通知表列表条数
const GetNotifyInfoListCount = `
select count(1)
from ots_notify_info t
where
	&t.mer_no
	and t.create_time >= @create_time 
	and t.create_time < date_add(@create_time, interval 1 day)
	&t.notify_status`

//GetNotifyInfoList 查询订单通知表列表数据
const GetNotifyInfoList = `
select
	t.order_id,
	t.mer_no,
	t.mer_order_no,
	t.create_time,
	t.start_time,
	t.end_time,
	t.notify_count,
	t.notify_status,
	t.notify_msg 
from ots_notify_info t
where
	&t.mer_no
	and t.create_time >= @create_time 
	and t.create_time < date_add(@create_time, interval 1 day)
	&t.notify_status
order by t.order_id desc
limit @ps offset @offset
`
