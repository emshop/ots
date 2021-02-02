package sql
//GetProductLineByPlID 查询单条数据产品线
const GetProductLineByPlID = `
select
	t.pl_id,
	t.pl_name,
	t.status,
	t.create_time
from ots_product_line t
where
	&pl_id`

//GetProductLineListCount 获取产品线列表条数
const GetProductLineListCount = `
select count(1)
from ots_product_line t
where
	?t.pl_name
	&t.status`

//GetProductLineList 查询产品线列表数据
const GetProductLineList = `
select
	t.pl_id,
	t.pl_name,
	t.status,
	t.create_time 
from ots_product_line t
where
	?t.pl_name
	&t.status
order by t.pl_id desc
limit @ps offset @offset
`
