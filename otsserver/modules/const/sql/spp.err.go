package sql

//SelectSupplierErrorCode 查询单条数据供货商错误码
const SelectSupplierErrorCode = `
select
t.id,
t.spp_no,
t.pl_id,
t.deal_code,
t.error_code,
t.error_desc
from ots_supplier_ecode t
where
t.spp_no = @spp_no
and  (t.pl_id = @pl_id or t.pl_id = 0)
and t.category = @category
and t.error_code = @error_code
and t.status = 0
order by t.pl_id desc
`
