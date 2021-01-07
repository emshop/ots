package sql

//SelectSupplierShelf 查询单条数据供货商货架
const SelectSupplierShelf = `
select
t.spp_shelf_id,
t.spp_shelf_name,
t.spp_no,
t.req_url,
t.query_url,
t.notify_url,
t.invoice_type,
t.spp_fee_discount,
t.trade_fee_discount,
t.payment_fee_discount
from ots_supplier_shelf t
where
t.spp_shelf_id = @spp_shelf_id 
and t.status = 0
`
