package deliverys

var reps = []string{

	`update ots_trade_delivery t
	set t.batch_id = @batch_id,
	t.last_update_time = now()
	where t.last_update_time > DATE_ADD(now(),INTERVAL -3 hour)
	and  t.last_update_time < DATE_ADD(now(),INTERVAL -1 minute)
	and (t.delivery_status in(20,30,31) or (t.delivery_status = 0 and t.payment_status = 20))
	limit 200`,

	`select
	t.delivery_id,
	t.order_id,
	t.delivery_status,
	t.payment_status
	from ots_trade_delivery t
	where t.last_update_time > DATE_ADD(now(),INTERVAL -1 minute)
	and (t.delivery_status in(20,30,31) or (t.delivery_status = 0 and t.payment_status = 20))
	and t.batch_id=@batch_id
	limit 200
`,
}
