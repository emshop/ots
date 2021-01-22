package sql

//UpdateNotifyInfoForSuccess 查询单条数据订单通知表
const UpdateNotifyInfoForSuccess = `
Update ots_notify_info t
t.notify_status = 0,
t.end_time = now(),
t.notify_msg = @notify_msg 
from ots_notify_info t
where
t.order_id = @order_id 
and t.notify_status = 30
`

//UpdateNotifyInfoForUnknown 查询单条数据订单通知表
const UpdateNotifyInfoForUnknown = `
Update ots_notify_info t
t.notify_msg = @notify_msg 
from ots_notify_info t
where
t.order_id = @order_id
and t.notify_status = 30
`
