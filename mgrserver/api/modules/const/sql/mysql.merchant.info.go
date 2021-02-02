package sql
//InsertMerchantInfo 添加商户信息
const InsertMerchantInfo = `
insert into ots_merchant_info
(
	
	mer_no,
	mer_name,
	mer_crop,
	mer_type,
	bd_uid,
	status
)
values
(
	
	@mer_no,
	@mer_name,
	@mer_crop,
	@mer_type,
	@bd_uid,
	@status
)`

//GetMerchantInfoByMerNo 查询单条数据商户信息
const GetMerchantInfoByMerNo = `
select
	t.mer_no,
	t.mer_name,
	t.mer_crop,
	t.mer_type,
	t.bd_uid,
	t.status
from ots_merchant_info t
where
	&mer_no`

//GetMerchantInfoListCount 获取商户信息列表条数
const GetMerchantInfoListCount = `
select count(1)
from ots_merchant_info t
where
	?t.mer_name
	&t.status`

//GetMerchantInfoList 查询商户信息列表数据
const GetMerchantInfoList = `
select
	t.mer_no,
	t.mer_name,
	t.mer_crop,
	t.mer_type,
	t.status 
from ots_merchant_info t
where
	?t.mer_name
	&t.status
order by t.mer_no desc
limit @ps offset @offset
`
//UpdateMerchantInfoByMerNo 更新商户信息
const UpdateMerchantInfoByMerNo = `
update ots_merchant_info 
set
	mer_name = @mer_name,
	mer_crop = @mer_crop,
	mer_type = @mer_type,
	bd_uid = @bd_uid,
	status = @status
where
	&mer_no`

