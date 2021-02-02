package sql
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
