// +build oracle

package sql

import (
	_ "gitlab.100bm.cn/micro-plat/dds/dds/internal/modules/const/sql/oracle"
)

//Get 获取某个类型下的字典信息
const Get = `
select 
	id,
	name,
	value,
	type,
	sort_no,
	group_code,
	status
from dds_dictionary_info
where type = @type and
	  status = 0
order by sort_no,id
`

//GetProvince 获取第一级省市
const GetProvince = `
select 
  'province' type,
	canton_code value,
	chinese_name name,
	parent_code pid,
	grade,
	full_spell,
	simple_spell
from dds_area_info
where &parent_code
and grade='1'
order by sort_id
`

//GetCityByProvinceID 根据省获取市信息
const GetCityByProvinceID = `
select 
 'city' type,
	canton_code value,
	chinese_name name,
	parent_code pid,
	grade,
	full_spell,
	simple_spell
from dds_area_info
where &parent_code
and grade='2'
order by sort_id
`

//GetAll 根据省获取市信息
const GetAll = `
select 
  'region' type,
	canton_code value,
	chinese_name name,
	parent_code pid,
	grade,
	full_spell,
	simple_spell
from dds_area_info
order by sort_id
`
