package sql

//UpdateTradeOrderForDeliverySuccess 发货成功后修改订单金额
const UpdateTradeOrderForDeliverySuccess = `
Update ots_trade_order t
case when t.success_face + @success_face = t.total_face then t.order_status = 40 else t.order_status=t.order_status end,
case when t.success_face + @success_face = t.total_face then t.delivery_status = 0 else t.delivery_status = t.delivery_status end,
t.success_face = t.success_face + @success_face,
t.success_sell_amount =(t.success_face + @success_face) * t.sell_discount,
t.success_mer_fee = (t.success_face + @success_face) * t.mer_fee_discount,
t.success_mer_trade_fee = (t.success_face + @success_face) * t.trade_fee_discount,
t.success_mer_payment_fee = (t.success_face + @success_face) * t.payment_fee_discount,
t.success_cost_amount = t.success_cost_amount + @cost_amount,
t.success_spp_fee = t.success_spp_fe + @spp_fee_amount,
t.success_spp_trade_fee = t.success_spp_trade_fee + @trade_fee_amount,
t.success_spp_payment_fee = t.success_spp_payment_fee + @payment_fee_amount
from ots_trade_order t
where
t.order_id = @order_id
and t.order_status = 30
and t.delivery_status = 30
`
