package database
 const ots_merchant_info=`CREATE TABLE  ots_merchant_info (
		mer_no varchar2(32)  not null AUTO_INCREMENT comment '编号' ,
		mer_name varchar2(64)  not null  comment '名称' ,
		status number(1) default 0 not null  comment '状态' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (mer_no)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='商户信息';`