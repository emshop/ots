package sql

//UpdateTradeOrderForClose 更新订单记录
const UpdateTradeOrderForClose = `
update ots_trade_order t set
t.finish_time = now(),
t.order_status= case when t.delivery_status =0 then 0 else 90 end,
t.profit = t.success_sell_amount - t.success_mer_fee - t.success_mer_trade_fee - t.success_mer_payment_fee - t.success_cost_amount + t.success_spp_fee - t.success_spp_trade_fee - t.success_spp_payment_fee
where
t.order_id = @order_id 
and t.notify_status in (0, 11)
and t.refund_status = 10
and t.order_status = 50
and t.delivery_status in(0, 90)
`
