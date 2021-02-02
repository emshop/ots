package sql

//UpdateNotifyInfoForSuccess 查询单条数据订单通知表
const UpdateNotifyInfoForSuccess = `
update ots_notify_info t set
t.notify_status = 0,
t.end_time = now(),
t.notify_msg = @notify_msg 
where
t.order_id = @order_id 
and t.notify_status = 30
`

//UpdateNotifyInfoForUnknown 查询单条数据订单通知表
const UpdateNotifyInfoForUnknown = `
update ots_notify_info t set
t.notify_msg = @notify_msg 
where
t.order_id = @order_id
and t.notify_status = 30
`
