package sql

//UpdateNotifyInfoForStart 查询单条数据订单通知表
const UpdateNotifyInfoForStart = `
update ots_notify_info t set
t.notify_status = case t.notify_status when 10 then 30 else t.notify_status = t.notify_status end,
t.start_time = case t.notify_status when 10 then now() else t.start_time end,
t.notify_count = case t.notify_status when 10 then 1 else t.notify_count + 1 end
where
t.order_id = @order_id 
and t.notify_status in(10,30)
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
and n.notify_status = 30
`
