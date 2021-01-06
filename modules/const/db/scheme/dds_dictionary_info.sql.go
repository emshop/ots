package scheme
 
//dds_dictionary_info 字典表
const dds_dictionary_info=`
	DROP TABLE IF EXISTS dds_dictionary_info;
	CREATE TABLE IF NOT EXISTS dds_dictionary_info (
		id int  not null auto_increment comment '编号' ,
		name varchar(64)  not null  comment '名称' ,
		value varchar(32)  not null  comment '值' ,
		type varchar(32)  not null  comment '类型' ,
		sort int default 0 not null  comment '排序值' ,
		status tinyint  not null  comment '状态' 
		,primary key (id)
		,index idx_dictionary_info_type(type)
	) ENGINE=InnoDB auto_increment = 100 DEFAULT CHARSET=utf8 COMMENT='字典表'`