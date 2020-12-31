package scheme
 
//ots_merchant_info 商户信息
const ots_merchant_info=`CREATE TABLE  ots_merchant_info (
		mer_no varchar2(32)  not null AUTO_INCREMENT comment '编号' ,
		mer_name varchar2(64)  not null  comment '名称' ,
		mer_crop varchar2(64)  not null  comment '公司' ,
		mer_type number(1)  not null  comment '类型' ,
		bd_uid number(20)  not null  comment '商务人员' ,
		status number(1) default 0 not null  comment '状态' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (mer_no)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='商户信息';`