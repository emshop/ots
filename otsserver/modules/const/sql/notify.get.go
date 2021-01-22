package sql

//UpdateNotifyInfoForStart 查询单条数据订单通知表
const UpdateNotifyInfoForStart = `
Update ots_notify_info t
case when t.notify_status = 10 then t.notify_status = 20 when t.notify_status = 20 then t.notify_status = 30 else  t.notify_status = t.notify_status end,
case when t.notify_status = 20 then t.start_time = now() else t.start_time = t.start_time end,
case when t.notify_status > 10  then t.notify_count = t.notify_count + 1 else t.notify_count = t.notify_count end,
from ots_notify_info t
where
t.order_id = @order_id 
and t.notify_status in(10, 20)
and t.notify_count < t.max_count
`

//SelectNotifyInfoForStart 查询单条数据订单通知表
const SelectNotifyInfoForStart = `
select
n.notify_url,
t.order_id,
t.mer_no,
t.mer_order_no,
t.account_name,
t.invoice_type,
t.sell_discount,
t.sell_amount,
from ots_notify_info n
inner join ots_trade_order t on t.order_id = n.order_id
where
n.order_id = @order_id 
and n.notify_status in(20,30)
`
