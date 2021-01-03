package scheme
 
//dds_dictionary_info 字典表
const dds_dictionary_info=`CREATE TABLE  dds_dictionary_info (
		id int default 100 not null AUTO_INCREMENT comment '编号' ,
		name varchar2(64)  not null  comment '名称' ,
		value varchar2(32)  not null  comment '值' ,
		type varchar2(32)  not null  comment '类型' ,
		sort int default 0 not null  comment '排序值' ,
		status tinyint  not null  comment '状态' 
		,KEY idx_dictionary_info_type(type),PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='字典表'`