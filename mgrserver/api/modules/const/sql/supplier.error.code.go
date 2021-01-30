package sql
//GetSupplierErrorCodeById 查询单条数据供货商错误码
const GetSupplierErrorCodeById = `
select
	t.id,
	t.spp_no,
	t.pl_id,
	t.category,
	t.deal_code,
	t.error_code,
	t.status,
	t.error_desc,
	t.create_time
from ots_supplier_error_code t
where
	&id`

//GetSupplierErrorCodeListCount 获取供货商错误码列表条数
const GetSupplierErrorCodeListCount = `
select count(1)
from ots_supplier_error_code t
where
1=1`

//GetSupplierErrorCodeList 查询供货商错误码列表数据
const GetSupplierErrorCodeList = `
select
	t.id,
	t.spp_no,
	t.pl_id,
	t.category,
	t.deal_code,
	t.error_code,
	t.status,
	t.error_desc,
	t.create_time 
from ots_supplier_error_code t
where
1=1`
