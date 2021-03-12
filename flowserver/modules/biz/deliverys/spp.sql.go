package deliverys

//queryDealCode 查询错误码
const queryDealCode = `
select
t.id,
t.spp_no,
t.pl_id,
t.deal_code,
t.error_code,
t.error_desc
from ots_supplier_ecode t
inner join ots_trade_delivery d on t.spp_no = d.spp_no
where
d.delivery_id = @delivery_id 
and t.category = @category
and (t.pl_id = d.pl_id or t.pl_id = 0)
and t.error_code = @error_code
and t.status = 0
order by t.pl_id desc
limit 1
`
