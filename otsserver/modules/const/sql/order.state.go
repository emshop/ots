package sql

//UpdateOrderForWaitPaying 将订单状态修改为等待支付
const UpdateOrderForWaitPaying = `
update ots_trade_order t set
t.payment_status = 20
where
t.order_id = @order_id
and t.order_status = 10
and t.payment_status = 10
`

//UpdateOrderForPaySuccess 将订单状态修改为等待支付
const UpdateOrderForPaySuccess = `
update ots_trade_order t set
t.payment_status = 0
where
t.order_id = @order_id
and t.order_status = 10
and t.payment_status = 20
`

//UpdateOrderForPayFailed 将订单状态修改为支付失败
const UpdateOrderForPayFailed = `
update ots_trade_order t set
t.payment_status = 90
where
t.order_id = @order_id
and t.order_status = 10
and t.payment_status = 20
`

//UpdateOrderForWaitDelivery 将订单状态修改为等待发货
const UpdateOrderForWaitDelivery = `
update ots_trade_order t set
t.order_status = 20,
t.delivery_status =20
where
t.order_id = @order_id
and t.order_status = 10
and t.payment_status = 0
and t.delivery_status = 10
`

//UpdateOrderForNotifyByPayFailed 支付失败改为等待通知
const UpdateOrderForNotifyByPayFailed = `
update ots_trade_order t set
t.order_status = 50,
t.delivery_status =11,
t.notify_status = case when t.notify_status = 11 then 11 else 20 end
where
t.order_id = @order_id
and t.order_status = 10
and t.payment_status = 90
and t.delivery_status = 10
and t.notify_status in(10,11)
`

//UpdateOrderForNotifyByDeliverySuccess 将发货成功的订单修改为等待通知
const UpdateOrderForNotifyByDeliverySuccess = `
update ots_trade_order t set
t.order_status = 50,
t.notify_status = case when t.notify_status = 11 then 11 else 10 end
where
t.order_id = @order_id
and t.order_status = 20
and t.delivery_status = 0
and t.notify_status in(10,11)
`

//UpdateOrderForFinishByNotify 支付失败改为等待通知
const UpdateOrderForFinishByNotify = `
update ots_trade_order t set
t.finish_time = now(),
t.order_status= case when t.delivery_status = 0 then 0 else 90 end,
t.profit = t.success_sell_amount - t.success_mer_fee - t.success_mer_trade_fee - t.success_mer_payment_fee - t.success_cost_amount + t.success_spp_fee - t.success_spp_trade_fee - t.success_spp_payment_fee
where
t.order_id = @order_id
and t.order_status = 50
and t.notify_status in(0,11)
`
