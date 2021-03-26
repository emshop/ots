package sql
//InsertSupplierEcode 添加供货商错误码
const InsertSupplierEcode = `
insert into ots_supplier_ecode
(
	id,
	id,
	spp_no,
	pl_id,
	category,
	deal_code,
	error_code,
	error_desc,
	status
)
values
(
	@id,
	if(isnull(@id)||@id='',0,@id),
	@spp_no,
	if(isnull(@pl_id)||@pl_id='',0,@pl_id),
	@category,
	if(isnull(@deal_code)||@deal_code='',0,@deal_code),
	@error_code,
	@error_desc,
	if(isnull(@status)||@status='',0,@status)
)`
//GetSupplierEcodeByID 查询供货商错误码单条数据
const GetSupplierEcodeByID = `
select
	t.id,
	t.spp_no,
	t.pl_id,
	t.category,
	t.deal_code,
	t.error_code,
	t.error_desc,
	t.create_time,
	t.status
from ots_supplier_ecode t
where
	&id`

//GetSupplierEcodeListCount 获取供货商错误码列表条数
const GetSupplierEcodeListCount = `
select count(1)
from ots_supplier_ecode t
where
	&t.spp_no
	&t.pl_id
	&t.category
	&t.status`

//GetSupplierEcodeList 查询供货商错误码列表数据
const GetSupplierEcodeList = `
select
	t.id,
	t.spp_no,
	t.pl_id,
	t.category,
	t.deal_code,
	t.error_code,
	t.error_desc,
	t.status 
from ots_supplier_ecode t
where
	&t.spp_no
	&t.pl_id
	&t.category
	&t.status
order by t.id desc
limit @ps offset @offset
`
//UpdateSupplierEcodeByID 更新供货商错误码
const UpdateSupplierEcodeByID = `
update ots_supplier_ecode 
set
	id =	if(isnull(@id)||@id='',0,@id),
	spp_no =	@spp_no,
	pl_id =	if(isnull(@pl_id)||@pl_id='',0,@pl_id),
	category =	@category,
	deal_code =	if(isnull(@deal_code)||@deal_code='',0,@deal_code),
	error_code =	@error_code,
	error_desc =	@error_desc,
	status =	if(isnull(@status)||@status='',0,@status)
where
	&id`

