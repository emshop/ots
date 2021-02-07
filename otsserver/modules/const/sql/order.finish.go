package sql

//UpdateTradeOrderForClose 更新订单记录
const UpdateTradeOrderForClose = `
update ots_trade_order t set
t.finish_time = now(),
t.order_status= case when t.delivery_status = 0 then 0 else 90 end,
t.profit = t.success_sell_amount - t.success_mer_fee - t.success_mer_trade_fee - t.success_mer_payment_fee - t.success_cost_amount + t.success_spp_fee - t.success_spp_trade_fee - t.success_spp_payment_fee
where
t.order_id = @order_id 
and t.notify_status in (0, 11)
and t.refund_status in(10,11,90,0)
and t.payment_status in(0, 90)
and t.order_status = 50
and t.delivery_status in(0,11, 90)
`

//UpdateTradeOrderForOrderTimeout 修改订单为超时关闭(原订单已经支付成功了)
const UpdateTradeOrderForOrderTimeout = `
update ots_trade_order t set
t.delivery_status = 90,
t.refund_status = 20,
t.is_refund = 1,
t.order_status = 70
where
t.order_id = @order_id
and t.order_status = 20
and t.payment_status = 0
and t.success_face = 0
and t.delivery_status in(20, 30)
and t.bind_face = 0
and t.is_refund = 0
and t.refund_status = 10
and t.order_timeout < now()
`

//SelectTradeOrderByNoneedDeal 查询超时订单
const SelectTradeOrderByNoneedDeal = `
select count(0)
from ots_trade_order t
where
t.order_id = @order_id 
and (t.order_timeout < now() or t.order_status in(0,90))
`
