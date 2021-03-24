package sql
//GetAccountRecordByRecordID 查询资金变动单条数据
const GetAccountRecordByRecordID = `
select
	t.record_id,
	t.account_id,
	t.trade_no,
	t.ext_no,
	t.trade_type,
	t.change_type,
	t.amount,
	t.balance,
	t.create_time,
	t.memo,
	t.ext
from beanpay_account_record t
where
	&record_id`

//GetAccountRecordListCount 获取资金变动列表条数
const GetAccountRecordListCount = `
select count(1)
from beanpay_account_record t
where
	&t.account_id
	&t.trade_no
	&t.trade_type
	&t.change_type`

//GetAccountRecordList 查询资金变动列表数据
const GetAccountRecordList = `
select
	t.record_id,
	t.account_id,
	t.trade_no,
	t.trade_type,
	t.change_type,
	t.amount,
	t.balance,
	t.create_time 
from beanpay_account_record t
where
	&t.account_id
	&t.trade_no
	&t.trade_type
	&t.change_type
order by t.record_id desc
limit @ps offset @offset
`
//GetAccountRecordDetailListCount 获取资金变动列表条数
const GetAccountRecordDetailListCount = `
select count(1)
from beanpay_account_record t
where
&trade_no`

//GetAccountRecordDetailList 查询资金变动列表数据
const GetAccountRecordDetailList = `
select
	t.record_id,
	t.account_id,
	t.trade_no,
	t.trade_type,
	t.change_type,
	t.amount,
	t.balance,
	t.create_time 
from beanpay_account_record t
where
&trade_no
limit @ps offset @offset`

