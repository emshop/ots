package sql
//GetAccountRecordByRecordID 查询单条数据账户余额变动信息
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

//GetAccountRecordListCount 获取账户余额变动信息列表条数
const GetAccountRecordListCount = `
select count(1)
from beanpay_account_record t
where
	&t.account_id
	&t.trade_no
	&t.trade_type
	&t.change_type`

//GetAccountRecordList 查询账户余额变动信息列表数据
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
