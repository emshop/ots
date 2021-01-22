package scheme
 
//ots_merchant_info 商户信息
const ots_merchant_info=`
	DROP TABLE IF EXISTS ots_merchant_info;
	CREATE TABLE IF NOT EXISTS ots_merchant_info (
		mer_no varchar(32)  not null  comment '编号' ,
		mer_name varchar(64)  not null  comment '名称' ,
		mer_crop varchar(64)  not null  comment '公司' ,
		mer_type tinyint  not null  comment '类型' ,
		bd_uid bigint  not null  comment '商务人员' ,
		status tinyint default 0 not null  comment '状态' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (mer_no)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='商户信息'`