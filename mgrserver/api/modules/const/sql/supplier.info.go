package sql
//GetSupplierInfoBySppNo 查询单条数据供货商信息
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
