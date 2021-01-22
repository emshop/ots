package sql

//UpdateTradeOrderForDeliverySuccess 发货成功后修改订单金额
const UpdateTradeOrderForDeliverySuccess = `
update ots_trade_order t set
t.order_status=case t.success_face + @success_face when  t.total_face then 40 else t.order_status end,
t.delivery_status=case t.success_face + @success_face when t.total_face then 0 else t.delivery_status end,
t.success_face = t.success_face + @success_face,
t.success_sell_amount =(t.success_face + @success_face) * t.sell_discount,
t.success_mer_fee = (t.success_face + @success_face) * t.mer_fee_discount,
t.success_mer_trade_fee = (t.success_face + @success_face) * t.trade_fee_discount,
t.success_mer_payment_fee = (t.success_face + @success_face) * t.payment_fee_discount,
t.success_cost_amount = t.success_cost_amount + @cost_amount,
t.success_spp_fee = t.success_spp_fee + @spp_fee_amount,
t.success_spp_trade_fee = t.success_spp_trade_fee + @trade_fee_amount,
t.success_spp_payment_fee = t.success_spp_payment_fee + @payment_fee_amount
where
t.order_id = @order_id
and t.order_status = 20
and t.delivery_status = 30
`
