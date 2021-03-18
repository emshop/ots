package orders

var createOrder = []string{

	//查询产品信息
	`
	select
t.mer_product_id,
t.mer_shelf_id,
t.mer_no,
t.pl_id,
t.brand_no,
t.province_no,
t.city_no,
t.face,
t.mer_product_no,
t.discount,
s.mer_shelf_name,
s.mer_fee_discount,
s.trade_fee_discount,
s.payment_fee_discount,
s.order_timeout,
s.payment_timeout,
s.invoice_type,
s.can_refund,
s.limit_count,
s.can_split_order
from ots_merchant_product t
inner join ots_merchant_info m on t.mer_no = m.mer_no
inner join ots_merchant_shelf s on s.mer_shelf_id = t.mer_shelf_id and s.mer_no = m.mer_no
where
t.mer_product_id = @mer_product_id 
and t.mer_no = @mer_no
and s.limit_count >= @num
and t.status = 0 
and m.status = 0
and s.status = 0`,

	//添加通知记录
	`insert into ots_notify_info(
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
	case @notify_url when '' then '-' else @notify_url end,
	case @notify_url when '' then 11 else 10 end,
	10
)`,

	//添加订单信息
	`insert into ots_trade_order(
order_id,
mer_no,
mer_order_no,
mer_product_id,
mer_shelf_id,
mer_product_no,
pl_id,
brand_no,
province_no,
city_no,
face,
num,
total_face,
account_name,
invoice_type,
sell_discount,
sell_amount,
mer_fee_discount,
trade_fee_discount,
payment_fee_discount,
can_split_order,
order_timeout,
payment_timeout,
order_status,
payment_status,
delivery_status,
refund_status,
notify_status
)values(
@order_id,
@mer_no,
@mer_order_no,
@mer_product_id,
@mer_shelf_id,
@mer_product_no,
@pl_id,
@brand_no,
@province_no,
@city_no,
@face,
@num,
@face*@num,
@account_name,
@invoice_type,
@discount,
@face*@num*@discount,
@mer_fee_discount,
@trade_fee_discount,
@payment_fee_discount,
@can_split_order,
DATE_ADD(now(),INTERVAL @order_timeout SECOND),
DATE_ADD(now(),INTERVAL @payment_timeout SECOND),
10,20,10,10,10
)
`,
}
