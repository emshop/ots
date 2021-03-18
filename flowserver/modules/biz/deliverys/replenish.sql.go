package deliverys

var reps = []string{

	`update ots_trade_delivery t
	set t.batch_id = @batch_id,
	t.last_update_time = now()
	where t.last_update_time > DATE_ADD(now(),INTERVAL -1 hour)
	and  t.last_update_time < DATE_ADD(now(),INTERVAL -30 minute)
	and t.delivery_status in(20,30,31)
	limit 20`,

	`select
	t.delivery_id,
	t.order_id,
	t.delivery_status
	from ots_trade_delivery t
	where t.last_update_time > DATE_ADD(now(),INTERVAL -1 minute)
	and t.delivery_status in(20,30,31)
	and t.batch_id=@batch_id
	limit 20
`,
}
