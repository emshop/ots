package sql
//GetAuditInfoByDeliveryID 查询单条数据发货人工审核表
const GetAuditInfoByDeliveryID = `
select
	t.delivery_id,
	t.order_id,
	t.create_time,
	t.delivery_status,
	t.end_order,
	t.audit_status,
	t.audit_by,
	t.audit_time,
	t.audit_msg
from ots_audit_info t
where
	&delivery_id`

//GetAuditInfoListCount 获取发货人工审核表列表条数
const GetAuditInfoListCount = `
select count(1)
from ots_audit_info t
where
	&t.order_id
	&t.audit_status`

//GetAuditInfoList 查询发货人工审核表列表数据
const GetAuditInfoList = `
select
	t.delivery_id,
	t.order_id,
	t.create_time,
	t.delivery_status,
	t.audit_status,
	t.audit_by,
	t.audit_time,
	t.audit_msg 
from ots_audit_info t
where
	&t.order_id
	&t.audit_status
order by t.delivery_id desc
limit @ps offset @offset
`
