package sql
//InsertDictionaryInfo 添加字典配置
const InsertDictionaryInfo = `
insert into dds_dictionary_info
(
	id,
	name,
	value,
	type,
	sort_no,
	status
)
values
(
	@id,
	@name,
	@value,
	@type,
	if(isnull(@sort_no)||@sort_no='',0,@sort_no),
	if(isnull(@status)||@status='',0,@status)
)`
//GetDictionaryInfoByID 查询字典配置单条数据
const GetDictionaryInfoByID = `
select
	t.id,
	t.name,
	t.value,
	t.type,
	t.sort_no,
	t.status
from dds_dictionary_info t
where
	&id`

//GetDictionaryInfoListCount 获取字典配置列表条数
const GetDictionaryInfoListCount = `
select count(1)
from dds_dictionary_info t
where
	?t.name
	&t.type
	&t.status`

//GetDictionaryInfoList 查询字典配置列表数据
const GetDictionaryInfoList = `
select
	t.id,
	t.name,
	t.value,
	t.type,
	t.sort_no,
	t.status 
from dds_dictionary_info t
where
	?t.name
	&t.type
	&t.status
order by t.id desc
limit @ps offset @offset
`
//UpdateDictionaryInfoByID 更新字典配置
const UpdateDictionaryInfoByID = `
update dds_dictionary_info 
set
	name =	@name,
	value =	@value,
	type =	@type,
	sort_no =	if(isnull(@sort_no)||@sort_no='',0,@sort_no),
	status =	if(isnull(@status)||@status='',0,@status)
where
	&id`

