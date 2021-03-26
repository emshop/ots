package deliverys

//dealOrderTimeout 处理超时订单
var dealOrderTimeout = []string{
	`select t.order_id from ots_trade_order t
where
t.order_id = @order_id
and t.order_status = 20
and t.payment_status = 0
and t.delivery_status in(20,30)
and t.notify_status = 10
and t.order_timeout < now()
and t.bind_face = 0
`,
	//检查订单是否超时，超时则修改为订单通知
	`
update ots_trade_order t set
t.delivery_status = 90,
t.order_status = 50,
t.notify_status = 20
where
t.order_id = @order_id
and t.order_status = 20
and t.payment_status = 0
and t.delivery_status in(20,30)
and t.notify_status = 10
and t.order_timeout < now()
and t.bind_face = 0
`,
}

//查询订单是否需要绑定
var checkOrderForBind = `select
t.order_id,
t.mer_no,
t.mer_order_no,
t.mer_product_id,
t.mer_shelf_id,
t.pl_id,
t.num,
t.brand_no,
t.province_no,
t.city_no,
t.account_name,
t.invoice_type,
t.can_split_order,
t.sell_discount,
t.sell_discount discount,
t.order_timeout,
t.payment_timeout,
t.delivery_pause,
t.order_status,
t.payment_status,
t.delivery_status,
p.face,
'0' pg_id,
'1' pnum,
l.pl_type
	from ots_trade_order t
	inner join ots_merchant_product p on p.mer_product_id = t.mer_product_id
	inner join ots_product_line l on l.pl_id = t.pl_id
	where
	t.order_id = @order_id
	and t.order_status = 20
	and t.delivery_status in(20,30)
	and t.total_face - t.bind_face > 0
	`

//查询订单是否需要绑定
var checkOrderForBindPackage = `select
t.order_id,
t.mer_no,
t.mer_order_no,
t.mer_product_id,
t.mer_shelf_id,
p.pl_id,
p.brand_no,
p.province_no,
p.city_no,
p.pg_id,
t.account_name,
t.invoice_type,
t.can_split_order,
t.sell_discount,
p.discount,
t.order_timeout,
t.payment_timeout,
t.delivery_pause,
t.order_status,
t.payment_status,
t.delivery_status,
t.num,
p.num pnum,
p.face,
l.pl_type
	from ots_trade_order t
	inner join ots_merchant_package p on t.mer_product_id = p.mer_product_id
	inner join ots_product_line l on l.pl_id = p.pl_id
	where
	t.order_id = @order_id
	and t.order_status = 20
	and (select IFNULL(sum(d.num),0) from ots_trade_delivery d where d.order_id = t.order_id and d.pg_id = p.pg_id and d.delivery_status!=90) < t.num * p.num
	and t.delivery_status in(20,30)
	and t.total_face - t.bind_face >= p.face
	limit 1
	`

//查询符合条件的产品信息
var querySppProducts = []string{
	//查询符合条件的商品信息
	`
select
t.spp_product_id,
t.spp_shelf_id,
t.spp_no,
t.spp_product_no,
t.face,
t.cost_discount,
f.spp_shelf_name,
f.req_url,
f.query_url,
f.notify_url,
f.invoice_type,
f.spp_fee_discount,
f.trade_fee_discount,
f.payment_fee_discount
from ots_supplier_product t
inner join ots_supplier_info s on t.spp_no = s.spp_no
inner join ots_supplier_shelf f on f.spp_shelf_id= t.spp_shelf_id
where
t.pl_id = @pl_id
and t.brand_no = @brand_no
and (t.province_no = @province_no or t.province_no = '*') 
and (t.city_no = @city_no or t.city_no = '*')
and t.face = @face
and t.cost_discount <= @sell_discount
and t.cost_discount <= @discount
and f.invoice_type = @invoice_type
and t.status = 0
and s.status = 0
and f.status = 0
and t.spp_no not in(select d.spp_no from ots_trade_delivery d where d.order_id = @order_id and d.delivery_status = 90
and (select IFNULL(sum(d.num),0) from ots_trade_delivery d where d.order_id = @order_id and d.pg_id = @pg_id and d.delivery_status!=90) < @num * @pnum
and d.end_time > date_add(now(),interval -1 minute))
order by t.cost_discount asc,t.province_no asc,t.city_no asc
`,
}

//创建绑定记录
var binds = []string{
	//更新订单信息
	`
	update ots_trade_order t set
	t.delivery_status = 30,
	t.bind_face = t.bind_face + @face
	where
	t.order_id = @order_id 
	and t.order_status = 20
	and t.delivery_pause = 1
	and t.delivery_status in(20, 30)
	and t.bind_face + @face <= t.total_face
	and (select IFNULL(sum(d.num),0) from ots_trade_delivery d where d.order_id = t.order_id and d.pg_id = @pg_id and d.delivery_status!=90) < t.num * @pnum
	
	`,

	//添加发货记录
	`
insert into ots_trade_delivery(
delivery_id,
order_id,
spp_no,
spp_product_id,
spp_product_no,
mer_no,
mer_product_id,
pg_id,
pl_id,
brand_no,
province_no,
city_no,
invoice_type,
account_name,
delivery_status,
payment_status,
face,
num,
total_face,
cost_discount,
spp_fee_discount,
trade_fee_discount,
payment_fee_discount
)select
@delivery_id,
t.order_id,
@spp_no,
@spp_product_id,
@spp_product_no,
@mer_no,
@mer_product_id,
@pg_id,
@pl_id,
@brand_no,
@province_no,
@city_no,
@invoice_type,
@account_name,
20,
10,
@face,
1,
@face,
@cost_discount,
@spp_fee_discount,
@trade_fee_discount,
@payment_fee_discount
from ots_trade_order t 
where t.order_id = @order_id
and t.order_status = 20
and t.delivery_pause = 1
and t.delivery_status in(20,30)
and (t.total_face - t.bind_face) >= 0
and (select IFNULL(sum(d.num),0) from ots_trade_delivery d where d.order_id = t.order_id and d.pg_id = @pg_id and d.delivery_status!=90) < t.num * @pnum

`,
}
