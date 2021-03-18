package orders

var reps = []string{

	`update ots_trade_order t
	set t.batch_id=@batch_id,
	t.last_update_time = now()
	where t.last_update_time > DATE_ADD(now(),INTERVAL -1 hour)
	and  t.last_update_time < DATE_ADD(now(),INTERVAL -30 minute)
	and (t.order_status in(10,50,60) or ((t.total_face - t.bind_face)>0 and t.order_status = 20))
	limit 20`,

	`select
	t.order_id,
	t.order_status
	from ots_trade_order t
	where t.last_update_time > DATE_ADD(now(),INTERVAL -1 hour)
	and  t.last_update_time < DATE_ADD(now(),INTERVAL -30 minute)
	and (t.order_status in(10,50,60) or ((t.total_face - t.bind_face)>0 and t.order_status = 20))
	and t.batch_id=@batch_id
	limit 20
`,
}
