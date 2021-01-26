package sql

//SelectChildFlowByPlID 根据产品线获取业务流程配置
const SelectChildFlowByPlID = `
select
t.flow_id,
t.pl_id,
t.success_flow_id,
t.failed_flow_id,
t.unknown_flow_id
from ots_product_flow t
where
t.pl_id = @pl_id
`

//SelectFlowByTag 根据产品线获取业务流程配置
const SelectFlowByTag = `
select
t.flow_id,
t.pl_id,
t.success_flow_id,
t.failed_flow_id,
t.unknown_flow_id
from ots_product_flow t
where
t.pl_id = @pl_id
and t.tag_name=@tag_name
`

//SelectChildFlowByFlowID 根据产品线获取业务流程配置
const SelectChildFlowByFlowID = `
select
t.flow_id,
t.pl_id,
t.success_flow_id,
t.failed_flow_id,
t.unknown_flow_id
from ots_product_flow t
where
t.flow_id=@flow_id
`

//SelectProductFlowByFlowID 根据流程编号获取流程配置信息
const SelectProductFlowByFlowID = `
select
t.flow_id,
t.flow_Name,
t.pl_id,
t.queue_name,
t.scan_interval,
t.delay,
t.timeout,
t.max_count 
from ots_product_flow t
where
t.pl_id = @pl_id
and t.flow_id=@flow_id
`
