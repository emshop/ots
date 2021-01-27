package sql
//GetMerchantProductByMerProductId 查询单条数据商户商品
const GetMerchantProductByMerProductId = `
select
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
1=1`

//GetMerchantProductList 查询商户商品列表数据
const GetMerchantProductList = `
select
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
1=1`
