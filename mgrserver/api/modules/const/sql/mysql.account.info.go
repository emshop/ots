package sql
//InsertAccountInfo 添加账户信息
const InsertAccountInfo = `
insert into beanpay_account_info
(
	account_id,
	account_name,
	ident,
	groups,
	eid,
	credit,
	status
)
values
(
	@account_id,
	@account_name,
	@ident,
	@groups,
	@eid,
	@credit,
	@status
)`

//GetAccountInfoByAccountID 查询单条数据账户信息
const GetAccountInfoByAccountID = `
select
	t.account_id,
	t.account_name,
	t.ident,
	t.groups,
	t.eid,
	t.balance,
	t.credit,
	t.status,
	t.create_time
from beanpay_account_info t
where
	&account_id`

//GetAccountInfoListCount 获取账户信息列表条数
const GetAccountInfoListCount = `
select count(1)
from beanpay_account_info t
where
	&t.eid`

//GetAccountInfoList 查询账户信息列表数据
const GetAccountInfoList = `
select
	t.account_id,
	t.account_name,
	t.groups,
	t.eid,
	t.status,
	t.create_time 
from beanpay_account_info t
where
	&t.eid
order by t.account_id desc
limit @ps offset @offset
`
//UpdateAccountInfoByAccountID 更新账户信息
const UpdateAccountInfoByAccountID = `
update beanpay_account_info 
set
	account_name = @account_name,
	credit = @credit,
	status = @status
where
	&account_id`

