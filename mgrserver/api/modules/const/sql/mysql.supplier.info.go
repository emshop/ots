package sql
//InsertSupplierInfo 添加供货商信息
const InsertSupplierInfo = `
insert into ots_supplier_info
(
	
	spp_no,
	spp_name,
	mer_crop,
	bd_uid,
	status
)
values
(
	
	@spp_no,
	@spp_name,
	@mer_crop,
	if(isnull(@bd_uid)||@bd_uid='',0,@bd_uid),
	if(isnull(@status)||@status='',0,@status)
)`

//GetSupplierInfoBySppNo 查询供货商信息单条数据
const GetSupplierInfoBySppNo = `
select
	t.spp_no,
	t.spp_name,
	t.mer_crop,
	t.bd_uid,
	t.status,
	t.create_time
from ots_supplier_info t
where
	&spp_no`

//GetSupplierInfoListCount 获取供货商信息列表条数
const GetSupplierInfoListCount = `
select count(1)
from ots_supplier_info t
where
	?t.spp_name
	&t.status`

//GetSupplierInfoList 查询供货商信息列表数据
const GetSupplierInfoList = `
select
	t.spp_no,
	t.spp_name,
	t.mer_crop,
	t.status,
	t.create_time 
from ots_supplier_info t
where
	?t.spp_name
	&t.status
order by t.spp_no desc
limit @ps offset @offset
`

//UpdateSupplierInfoBySppNo 更新供货商信息
const UpdateSupplierInfoBySppNo = `
update ots_supplier_info 
set
	spp_name =	@spp_name,
	mer_crop =	@mer_crop,
	bd_uid =	if(isnull(@bd_uid)||@bd_uid='',0,@bd_uid),
	status =	if(isnull(@status)||@status='',0,@status)
where
	&spp_no`

