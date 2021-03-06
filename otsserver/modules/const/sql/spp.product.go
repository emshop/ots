package sql

//SelectSupplierProduct 根据产品线，省，市获取上游产品信息
const SelectSupplierProduct = `
select
t.spp_product_id,
t.spp_shelf_id,
t.spp_no,
t.spp_product_no,
t.pl_id,
t.brand_no,
t.province_no,
t.city_no,
t.face,
t.cost_discount,
t.status
from ots_supplier_product t
inner join ots_supplier_info s on t.spp_no = s.spp_no
where
t.pl_id = @pl_id and t.pl_id !=0
and t.brand_no = @brand_no
and t.province_no = @province_no
and t.city_no = @city_no
and case when @can_split_order = 0 then t.face <= @face else t.face = @face end
and t.cost_discount <= @cost_discount
and t.status = 0
and s.status = 0
order by t.cost_discount asc
`
