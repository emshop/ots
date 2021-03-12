// +build !oracle

package sql

const SaveLifeTimeInfo = `
insert into lcs_life_time
	(order_no, batch_no, extral_param,content, ip)
values
	(@order_no, @batch_no, @extral_param,@content,@ip);
`
