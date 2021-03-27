package sql
//InsertMerchantUpstream 添加商户上游
const InsertMerchantUpstream = `
insert into ots_merchant_upstream
(
	mu_id,
	mer_shelf_id,
	spp_no,
	status
)
values
(
	@mu_id,
	if(isnull(@mer_shelf_id)||@mer_shelf_id='',0,@mer_shelf_id),
	@spp_no,
	if(isnull(@status)||@status='',0,@status)
)`
//GetMerchantUpstreamByMuID 查询商户上游单条数据
const GetMerchantUpstreamByMuID = `
select
	t.mu_id,
	t.mer_shelf_id,
	t.spp_no,
	t.status,
	t.create_time
from ots_merchant_upstream t
where
	&mu_id`

//GetMerchantUpstreamListCount 获取商户上游列表条数
const GetMerchantUpstreamListCount = `
select count(1)
from ots_merchant_upstream t
where
	&t.mer_shelf_id
	&t.spp_no`

//GetMerchantUpstreamList 查询商户上游列表数据
const GetMerchantUpstreamList = `
select
	t.mu_id,
	t.mer_shelf_id,
	t.spp_no,
	t.status,
	t.create_time 
from ots_merchant_upstream t
where
	&t.mer_shelf_id
	&t.spp_no
order by t.mu_id desc
limit @ps offset @offset
`
//GetMerchantUpstreamDetailListCount 获取商户上游列表条数
const GetMerchantUpstreamDetailListCount = `
select count(1)
from ots_merchant_upstream t
where
&mer_shelf_id`

//GetMerchantUpstreamDetailList 查询商户上游列表数据
const GetMerchantUpstreamDetailList = `
select
	t.mu_id,
	t.mer_shelf_id,
	t.spp_no,
	t.status,
	t.create_time 
from ots_merchant_upstream t
where
&mer_shelf_id
limit @ps offset @offset`
//UpdateMerchantUpstreamByMuID 更新商户上游
const UpdateMerchantUpstreamByMuID = `
update ots_merchant_upstream 
set
	status =	if(isnull(@status)||@status='',0,@status)
where
	&mu_id`

