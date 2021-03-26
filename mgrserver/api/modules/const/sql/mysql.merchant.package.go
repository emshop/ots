package sql
//InsertMerchantPackage 添加组合商品
const InsertMerchantPackage = `
insert into ots_merchant_package
(
	pg_id,
	pg_name,
	pl_id,
	mer_product_id,
	brand_no,
	province_no,
	city_no,
	face,
	num,
	discount,
	status
)
values
(
	@pg_id,
	@pg_name,
	if(isnull(@pl_id)||@pl_id='',0,@pl_id),
	if(isnull(@mer_product_id)||@mer_product_id='',0,@mer_product_id),
	@brand_no,
	@province_no,
	@city_no,
	if(isnull(@face)||@face='',0,@face),
	if(isnull(@num)||@num='',0,@num),
	if(isnull(@discount)||@discount='',0,@discount),
	if(isnull(@status)||@status='',0,@status)
)`

//GetMerchantPackageByPgID 查询组合商品单条数据
const GetMerchantPackageByPgID = `
select
	t.pg_id,
	t.pg_name,
	t.pl_id,
	t.mer_product_id,
	t.brand_no,
	t.province_no,
	t.city_no,
	t.face,
	t.num,
	t.discount,
	t.status,
	t.create_time
from ots_merchant_package t
where
	&pg_id`

//GetMerchantPackageListCount 获取组合商品列表条数
const GetMerchantPackageListCount = `
select count(1)
from ots_merchant_package t
where
	?t.pg_name
	&t.brand_no
	&t.num
	&t.status`

//GetMerchantPackageList 查询组合商品列表数据
const GetMerchantPackageList = `
select
	t.pg_id,
	t.pg_name,
	t.pl_id,
	t.mer_product_id,
	t.brand_no,
	t.province_no,
	t.city_no,
	t.face,
	t.num,
	t.discount,
	t.status 
from ots_merchant_package t
where
	?t.pg_name
	&t.brand_no
	&t.num
	&t.status
order by #order_by
limit @ps offset @offset
`
//GetMerchantPackageDetailListCount 获取组合商品列表条数
const GetMerchantPackageDetailListCount = `
select count(1)
from ots_merchant_package t
where
&mer_product_id`

//GetMerchantPackageDetailList 查询组合商品列表数据
const GetMerchantPackageDetailList = `
select
	t.pg_id,
	t.pg_name,
	t.pl_id,
	t.mer_product_id,
	t.brand_no,
	t.province_no,
	t.city_no,
	t.face,
	t.num,
	t.discount,
	t.status 
from ots_merchant_package t
where
&mer_product_id
limit @ps offset @offset`

//UpdateMerchantPackageByPgID 更新组合商品
const UpdateMerchantPackageByPgID = `
update ots_merchant_package 
set
	pg_name =	@pg_name,
	pl_id =	if(isnull(@pl_id)||@pl_id='',0,@pl_id),
	brand_no =	@brand_no,
	face =	if(isnull(@face)||@face='',0,@face),
	num =	if(isnull(@num)||@num='',0,@num),
	discount =	if(isnull(@discount)||@discount='',0,@discount),
	status =	if(isnull(@status)||@status='',0,@status)
where
	&pg_id`

