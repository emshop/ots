package deliverys

//查询订单是否需要绑定
var checkOrderForBind = `select
t.order_id,
t.mer_no,
t.mer_order_no,
t.mer_product_id,
t.mer_shelf_id,
t.pl_id,
t.brand_no,
t.province_no,
t.city_no,
case when t.can_split_order = 0 then t.total_face-t.bind_face else t.face end 'face',
t.account_name,
t.invoice_type,
t.can_split_order,
t.sell_discount,
t.order_timeout,
t.payment_timeout,
t.delivery_pause,
t.order_status,
t.payment_status,
t.delivery_status,
t.total_face,
t.bind_face
	from ots_trade_order t
	where
	t.order_id = @order_id
	and t.order_status = 20
	and t.delivery_status = 20
	and t.total_face - t.bind_face > 0
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
t.pl_id = @pl_id and t.pl_id !=0
and t.brand_no = @brand_no
and t.province_no = @province_no
and t.city_no = @city_no
and case when @can_split_order = 0 then t.face <= @face else t.face = @face end
and t.cost_discount <= @sell_discount
and f.invoice_type = @invoice_type
and t.status = 0
and s.status = 0
order by t.cost_discount asc
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
)values(
@delivery_id,
@order_id,
@spp_no,
@spp_product_id,
@spp_product_no,
@mer_no,
@mer_product_id,
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
)`,
}
