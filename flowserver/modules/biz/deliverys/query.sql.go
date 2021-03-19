package deliverys

var deliveryQuery = []string{
	`
	select
	t.delivery_id,
	t.order_id,
	t.spp_no,
	t.spp_product_id,
	t.spp_delivery_no,
	t.spp_product_no,
	t.mer_no,
	t.mer_product_id,
	t.pl_id,
	t.brand_no,
	t.province_no,
	t.city_no,
	t.invoice_type,
	t.account_name,
	t.delivery_status,
	t.face,
	t.num,
	t.total_face
	from ots_trade_delivery t
	where
	t.delivery_id = @delivery_id 
	and t.delivery_status in(30,31)`,

	`update ots_trade_delivery t
	set t.last_update_time = now()
	where t.last_update_time < DATE_ADD(now(),INTERVAL -2 second)
	and t.delivery_id = @delivery_id 
	and t.delivery_status in(30,31)`,
}
