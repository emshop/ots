package sql

//GetLifeTimeListCount 获取生命周期列表条数
const GetLifeTimeListCount = `
select count(1)
from lcs_life_time t
where
	&t.order_no`

//GetLifeTimeList 查询生命周期列表数据
const GetLifeTimeList = `
select
	t.id,
	t.order_no,
	t.ip,
	t.content,
	t.create_time 
from lcs_life_time t
where
	&t.order_no
order by t.id desc
limit @ps offset @offset
`
//GetLifeTimeDetailListCount 获取生命周期列表条数
const GetLifeTimeDetailListCount = `
select count(1)
from lcs_life_time t
where
&order_no`

//GetLifeTimeDetailList 查询生命周期列表数据
const GetLifeTimeDetailList = `
select
	t.id,
	t.order_no,
	t.ip,
	t.content,
	t.create_time 
from lcs_life_time t
where
&order_no
limit @ps offset @offset`
