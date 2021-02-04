package sql
//InsertProductFlow 添加业务流程
const InsertProductFlow = `
insert into ots_product_flow
(
	flow_id,
	flow_name,
	tag_name,
	pl_id,
	success_flow_id,
	failed_flow_id,
	unknown_flow_id,
	queue_name,
	scan_interval,
	delay,
	timeout,
	max_count
)
values
(
	@flow_id,
	@flow_name,
	@tag_name,
	@pl_id,
	@success_flow_id,
	@failed_flow_id,
	@unknown_flow_id,
	@queue_name,
	@scan_interval,
	@delay,
	@timeout,
	@max_count
)`

//GetProductFlowByFlowID 查询单条数据业务流程
const GetProductFlowByFlowID = `
select
	t.flow_id,
	t.flow_name,
	t.tag_name,
	t.pl_id,
	t.success_flow_id,
	t.failed_flow_id,
	t.unknown_flow_id,
	t.queue_name,
	t.scan_interval,
	t.delay,
	t.timeout,
	t.max_count
from ots_product_flow t
where
	&flow_id`

//GetProductFlowListCount 获取业务流程列表条数
const GetProductFlowListCount = `
select count(1)
from ots_product_flow t
where
	?t.flow_name
	?t.tag_name
	&t.pl_id`

//GetProductFlowList 查询业务流程列表数据
const GetProductFlowList = `
select
	t.flow_id,
	t.flow_name,
	t.tag_name,
	t.pl_id,
	t.success_flow_id,
	t.failed_flow_id,
	t.unknown_flow_id 
from ots_product_flow t
where
	?t.flow_name
	?t.tag_name
	&t.pl_id
order by t.flow_id desc
limit @ps offset @offset
`
//UpdateProductFlowByFlowID 更新业务流程
const UpdateProductFlowByFlowID = `
update ots_product_flow 
set
	flow_name = @flow_name,
	tag_name = @tag_name,
	pl_id = @pl_id,
	success_flow_id = @success_flow_id,
	failed_flow_id = @failed_flow_id,
	unknown_flow_id = @unknown_flow_id,
	queue_name = @queue_name,
	scan_interval = @scan_interval,
	delay = @delay,
	timeout = @timeout,
	max_count = @max_count
where
	&flow_id`

