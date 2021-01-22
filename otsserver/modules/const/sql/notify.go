package sql

//InsertNotifyInfo 插入订单通知表
const InsertNotifyInfo = `
insert into ots_notify_info(
        order_id,
        mer_no,
        mer_order_no,
        notify_url,
        notify_status,
        max_count
)values(
        @order_id,
        @mer_no,
        @mer_order_no,
        @notify_url,
        10,
        10
)
`
