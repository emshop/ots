package sql
//GetRefundApplyByApplyID 查询单条数据退款申请
const GetRefundApplyByApplyID = `
select
	t.apply_id,
	t.order_id,
	t.mer_no,
	t.mer_order_no,
	t.refund_cause,
	t.apply_status,
	t.refund_amount,
	t.create_time
from ots_refund_apply t
where
	&apply_id`

//GetRefundApplyListCount 获取退款申请列表条数
const GetRefundApplyListCount = `
select count(1)
from ots_refund_apply t
where
	&t.order_id
	&t.mer_no
	&t.mer_order_no`

//GetRefundApplyList 查询退款申请列表数据
const GetRefundApplyList = `
select
	t.apply_id,
	t.order_id,
	t.mer_no,
	t.mer_order_no,
	t.refund_cause,
	t.refund_amount,
	t.create_time 
from ots_refund_apply t
where
	&t.order_id
	&t.mer_no
	&t.mer_order_no
order by t.apply_id desc
limit @ps offset @offset
`
