package database
 const ots_supplier_info=`CREATE TABLE  ots_supplier_info (
		spp_no varchar2(32)  not null AUTO_INCREMENT comment '编号' ,
		spp_name varchar2(64)  not null  comment '名称' ,
		status number(1) default 0 not null  comment '状态' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (spp_no)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='供货商信息';`