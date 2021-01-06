package biz

//ProductFlow 业务流程------------------------------------

//FieldIDFlowID 字段流程编号的数据库名称
const FieldIDFlowID = "flow_id"

//FieldFlowIDParentFlowID 字段父级流程编号的数据库名称
const FieldFlowIDParentFlowID = "parent_flow_id"

//FieldFlowIDSuccessFlowID 字段成功后续流程的数据库名称
const FieldFlowIDSuccessFlowID = "success_flow_id"

//FieldFlowIDFailedFlowID 字段失败后续流程的数据库名称
const FieldFlowIDFailedFlowID = "failed_flow_id"

//FieldFlowIDUnknownFlowID 字段未知后续流程的数据库名称
const FieldFlowIDUnknownFlowID = "unknown_flow_id"

//SelectProductFlowByPlID 根据产品线获取业务流程配置
const SelectProductFlowByPlID = `
select
t.flow_id,
t.parent_flow_id,
t.success_flow_id,
t.failed_flow_id,
t.unknown_flow_id
from ots_product_flow t
where
t.pl_id = @pl_id
and t.parent_flow_id=@parent_flow_id
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
