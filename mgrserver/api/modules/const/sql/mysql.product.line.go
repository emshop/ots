package sql
//InsertProductLine 添加产品线
const InsertProductLine = `
insert into ots_product_line
(
	pl_id,
	pl_name,
	pl_type,
	has_feedback,
	has_logistics,
	status
)
values
(
	@pl_id,
	@pl_name,
	if(isnull(@pl_type)||@pl_type='',0,@pl_type),
	if(isnull(@has_feedback)||@has_feedback='',0,@has_feedback),
	if(isnull(@has_logistics)||@has_logistics='',0,@has_logistics),
	if(isnull(@status)||@status='',0,@status)
)`
//GetProductLineByPlID 查询产品线单条数据
const GetProductLineByPlID = `
select
	t.pl_id,
	t.pl_name,
	t.pl_type,
	t.has_feedback,
	t.has_logistics,
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
	&t.has_feedback
	&t.has_logistics
	&t.status`

//GetProductLineList 查询产品线列表数据
const GetProductLineList = `
select
	t.pl_id,
	t.pl_name,
	t.pl_type,
	t.has_feedback,
	t.has_logistics,
	t.status 
from ots_product_line t
where
	?t.pl_name
	&t.pl_type
	&t.has_feedback
	&t.has_logistics
	&t.status
order by #order_by
limit @ps offset @offset
`
//UpdateProductLineByPlID 更新产品线
const UpdateProductLineByPlID = `
update ots_product_line 
set
	pl_name =	@pl_name,
	has_feedback =	if(isnull(@has_feedback)||@has_feedback='',0,@has_feedback),
	has_logistics =	if(isnull(@has_logistics)||@has_logistics='',0,@has_logistics),
	status =	if(isnull(@status)||@status='',0,@status)
where
	&pl_id`

