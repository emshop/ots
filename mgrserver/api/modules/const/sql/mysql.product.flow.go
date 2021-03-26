package sql
//InsertProductFlow 添加业务流程
const InsertProductFlow = `
insert into ots_product_flow
(
	flow_id,
	flow_tag,
	pl_id,
	queue_name,
	scan_interval,
	delay,
	timeout,
	max_count,
	status
)
values
(
	@flow_id,
	@flow_tag,
	if(isnull(@pl_id)||@pl_id='',0,@pl_id),
	@queue_name,
	if(isnull(@scan_interval)||@scan_interval='',0,@scan_interval),
	if(isnull(@delay)||@delay='',0,@delay),
	if(isnull(@timeout)||@timeout='',0,@timeout),
	if(isnull(@max_count)||@max_count='',0,@max_count),
	if(isnull(@status)||@status='',0,@status)
)`
//GetProductFlowByFlowID 查询业务流程单条数据
const GetProductFlowByFlowID = `
select
	t.flow_id,
	t.flow_tag,
	t.pl_id,
	t.queue_name,
	t.scan_interval,
	t.delay,
	t.timeout,
	t.max_count,
	t.status
from ots_product_flow t
where
	&flow_id`

//GetProductFlowListCount 获取业务流程列表条数
const GetProductFlowListCount = `
select count(1)
from ots_product_flow t
where
	&t.flow_tag
	&t.pl_id
	&t.status`

//GetProductFlowList 查询业务流程列表数据
const GetProductFlowList = `
select
	t.flow_id,
	t.flow_tag,
	t.pl_id,
	t.queue_name,
	t.scan_interval,
	t.delay,
	t.timeout,
	t.max_count,
	t.status 
from ots_product_flow t
where
	&t.flow_tag
	&t.pl_id
	&t.status
order by t.flow_id desc
limit @ps offset @offset
`
//UpdateProductFlowByFlowID 更新业务流程
const UpdateProductFlowByFlowID = `
update ots_product_flow 
set
	flow_tag =	@flow_tag,
	pl_id =	if(isnull(@pl_id)||@pl_id='',0,@pl_id),
	queue_name =	@queue_name,
	scan_interval =	if(isnull(@scan_interval)||@scan_interval='',0,@scan_interval),
	delay =	if(isnull(@delay)||@delay='',0,@delay),
	timeout =	if(isnull(@timeout)||@timeout='',0,@timeout),
	max_count =	if(isnull(@max_count)||@max_count='',0,@max_count),
	status =	if(isnull(@status)||@status='',0,@status)
where
	&flow_id`

