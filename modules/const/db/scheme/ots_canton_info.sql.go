package scheme
 
//ots_canton_info 省市信息
const ots_canton_info=`CREATE TABLE  ots_canton_info (
		canton_code varchar2(32)  not null  comment '区域编号' ,
		chinese_name varchar2(32)  not null  comment '中文名称' ,
		spell varchar2(64)  not null  comment '英文或全拼' ,
		grade number(1)  not null  comment '行政级别' ,
		parent varchar2(32)  not null  comment '父级' ,
		simple_spell varchar2(8)  not null  comment '简拼' ,
		area_code varchar2(8)  not null  comment '区号' ,
		standard_code number(6)  not null  comment '行政编码' ,
		PRIMARY KEY (canton_code)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='省市信息';`