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
t.discount
from ots_merchant_product t
inner join ots_merchant_info m on t.mer_no = m.mer_no
where
t.mer_product_id = @mer_product_id 
and t.mer_no = @mer_no
and t.status = 0 
and m.status = 0`,

	//货架信息
	`
select
t.mer_shelf_id,
t.mer_shelf_name,
t.mer_no,
t.mer_fee_discount,
t.trade_fee_discount,
t.payment_fee_discount,
t.order_timeout,
t.payment_timeout,
t.invoice_type,
t.can_refund,
t.limit_count,
t.can_split_order
from ots_merchant_shelf t
inner join ots_merchant_info m on t.mer_no = m.mer_no
where
t.mer_shelf_id = @mer_shelf_id 
and t.mer_no = @mer_no
and t.status = 0
and m.status = 0
`,

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
10
)
`,
}

//SelectTradeOrderByMerchant 通过mer_no,mer_order_no获取订单
const SelectTradeOrderByMerchant = `
select
t.order_id,
t.mer_no,
t.mer_order_no,
t.account_name,
t.face,
t.num,
FORMAT(t.sell_amount,2) 'amount',
t.total_face,
case t.order_status when 0 then "SUCCESS" when 91 then "SUCCESS" when 90 then 'FAILED' else 'UNDERWAY' end 'order_status'
from ots_trade_order t
where t.mer_no = @mer_no and t.mer_order_no=@mer_order_no limit 1`

//SelectTradeOrderByOrderID 查询单条数据订单记录
const SelectTradeOrderByOrderID = `
select
t.order_id,
t.mer_no,
t.mer_order_no,
t.mer_product_id,
t.mer_shelf_id,
t.mer_product_no,
t.pl_id,
t.brand_no,
t.province_no,
t.city_no,
t.face,
t.num,
t.total_face,
t.account_name,
t.invoice_type,
t.sell_amount,
t.mer_fee_discount,
t.trade_fee_discount,
t.payment_fee_discount,
t.can_split_order,
t.create_time,
t.order_timeout,
t.payment_timeout,
t.delivery_pause,
t.order_status,
t.payment_status,
t.delivery_status,
t.refund_status,
t.notify_status,
t.is_refund,
t.bind_face
from ots_trade_order t
where
t.order_id = @order_id 
`
