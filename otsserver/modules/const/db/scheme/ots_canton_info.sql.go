package scheme
 
//ots_canton_info 省市信息
const ots_canton_info=`
	DROP TABLE IF EXISTS ots_canton_info;
	CREATE TABLE IF NOT EXISTS ots_canton_info (
		canton_code varchar(32)  not null  comment '区域编号' ,
		chinese_name varchar(32)  not null  comment '中文名称' ,
		spell varchar(64)  not null  comment '英文或全拼' ,
		grade tinyint  not null  comment '行政级别' ,
		parent varchar(32)  not null  comment '父级' ,
		simple_spell varchar(8)  not null  comment '简拼' ,
		area_code varchar(8)  not null  comment '区号' ,
		standard_code int  not null  comment '行政编码' 
		,primary key (canton_code)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='省市信息'`