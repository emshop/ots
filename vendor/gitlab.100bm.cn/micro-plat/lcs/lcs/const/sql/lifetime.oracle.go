// +build oracle

package sql

const SaveLifeTimeInfo = `
insert into lcs_life_time
	(id,order_no, batch_no, extral_param,content, ip)
values
	(seq_life_time_id.nextval, @order_no,@batch_no, @extral_param,@content, @ip)
`
