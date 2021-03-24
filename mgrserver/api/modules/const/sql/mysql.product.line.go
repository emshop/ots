package sql
//InsertProductLine 添加产品线
const InsertProductLine = `
insert into ots_product_line
(
	pl_id,
	pl_name,
	pl_type,
	status
)
values
(
	@pl_id,
	@pl_name,
	if(isnull(@pl_type)||@pl_type='',0,@pl_type),
	if(isnull(@status)||@status='',0,@status)
)`

//GetProductLineByPlID 查询产品线单条数据
const GetProductLineByPlID = `
select
	t.pl_id,
	t.pl_name,
	t.pl_type,
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
	&t.pl_type
	&t.status`

//GetProductLineList 查询产品线列表数据
const GetProductLineList = `
select
	t.pl_id,
	t.pl_name,
	t.pl_type,
	t.status 
from ots_product_line t
where
	?t.pl_name
	&t.pl_type
	&t.status
order by #order_by
limit @ps offset @offset
`

//UpdateProductLineByPlID 更新产品线
const UpdateProductLineByPlID = `
update ots_product_line 
set
	pl_name =	@pl_name,
	status =	if(isnull(@status)||@status='',0,@status)
where
	&pl_id`

