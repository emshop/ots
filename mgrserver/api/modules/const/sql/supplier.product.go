package sql
//GetSupplierProductBySppProductId 查询单条数据供货商商品
const GetSupplierProductBySppProductId = `
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
	t.status,
	t.create_time
from ots_supplier_product t
where
	&spp_product_id`

//GetSupplierProductListCount 获取供货商商品列表条数
const GetSupplierProductListCount = `
select count(1)
from ots_supplier_product t
where
	&t.spp_shelf_id
	&t.spp_no
	&t.pl_id
	&t.brand_no
	&t.province_no`

//GetSupplierProductList 查询供货商商品列表数据
const GetSupplierProductList = `
select
	t.spp_product_id,
	t.spp_shelf_id,
	t.spp_no,
	t.pl_id,
	t.brand_no,
	t.province_no,
	t.city_no,
	t.face,
	t.cost_discount,
	t.status,
	t.create_time 
from ots_supplier_product t
where
	&t.spp_shelf_id
	&t.spp_no
	&t.pl_id
	&t.brand_no
	&t.province_no
order by t.spp_product_id desc
limit @ps offset @offset
`
