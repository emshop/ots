package biz

//SelectMerchantProduct 查询单条数据商户商品
const SelectMerchantProduct = `
select
t.mer_product_id,
t.mer_shelf_id,
t.mer_no,
t.pl_id,
t.brand_no,
t.province_no,
t.city_no,
t.face,
t.mer_product_no,
t.discount
from ots_merchant_product t
inner join ots_merchant_info m on t.mer_no = m.mer_no
where
t.mer_product_id = @mer_product_id 
and t.status = 0 and m.status = 0
`
