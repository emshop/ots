package sql
//GetMerchantProductByMerProductID 查询单条数据商户商品
const GetMerchantProductByMerProductID = `
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
	t.discount,
	t.status,
	t.create_time
from ots_merchant_product t
where
	&mer_product_id`

//GetMerchantProductListCount 获取商户商品列表条数
const GetMerchantProductListCount = `
select count(1)
from ots_merchant_product t
where
	&t.mer_shelf_id
	&t.mer_no
	&t.pl_id
	&t.brand_no`

//GetMerchantProductList 查询商户商品列表数据
const GetMerchantProductList = `
select
	t.mer_product_id,
	t.mer_shelf_id,
	t.mer_no,
	t.pl_id,
	t.brand_no,
	t.province_no,
	t.city_no,
	t.face,
	t.discount,
	t.status,
	t.create_time 
from ots_merchant_product t
where
	&t.mer_shelf_id
	&t.mer_no
	&t.pl_id
	&t.brand_no
order by t.mer_product_id desc
limit @ps offset @offset
`
//UpdateMerchantProductByMerProductID 更新商户商品
const UpdateMerchantProductByMerProductID = `
update ots_merchant_product 
set
	face = @face,
	mer_product_no = @mer_product_no,
	discount = @discount,
	status = @status
where
	&mer_product_id`

