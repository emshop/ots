package sql
//InsertSupplierProduct 添加供货商商品
const InsertSupplierProduct = `
insert into ots_supplier_product
(
	spp_product_id,
	spp_product_id,
	spp_shelf_id,
	spp_no,
	spp_product_no,
	pl_id,
	brand_no,
	province_no,
	city_no,
	face,
	cost_discount,
	status
)
values
(
	@spp_product_id,
	@spp_product_id,
	@spp_shelf_id,
	@spp_no,
	@spp_product_no,
	@pl_id,
	@brand_no,
	@province_no,
	@city_no,
	@face,
	@cost_discount,
	@status
)`

//GetSupplierProductBySppProductID 查询单条数据供货商商品
const GetSupplierProductBySppProductID = `
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
//UpdateSupplierProductBySppProductID 更新供货商商品
const UpdateSupplierProductBySppProductID = `
update ots_supplier_product 
set
	spp_product_no = @spp_product_no
where
	&spp_product_id`

