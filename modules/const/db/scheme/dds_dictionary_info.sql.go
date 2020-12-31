package scheme
 
//dds_dictionary_info 字典表
const dds_dictionary_info=`CREATE TABLE  dds_dictionary_info (
		id number(10)  not null AUTO_INCREMENT comment '编号' ,
		name varchar2(64)  not null  comment '名称' ,
		value varchar2(32)  not null  comment '值' ,
		type varchar2(32)  not null  comment '类型' ,
		sort number(3) default 0 not null  comment '排序值' ,
		status number(1)  not null  comment '状态' ,
		PRIMARY KEY (id),
		KEY IDX_DICTIONARY_INFO_TYPE (type)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='字典表';`