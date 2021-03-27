package scheme
 
//ots_merchant_info 商户信息
const ots_merchant_info=`
	DROP TABLE IF EXISTS ots_merchant_info;
	CREATE TABLE IF NOT EXISTS ots_merchant_info (
		mer_no varchar(32)  not null  comment '编号' ,
		mer_name varchar(64)  not null  comment '商户名称' ,
		mer_crop varchar(64)    comment '所属公司' ,
		mer_type tinyint  not null  comment '商户分类' ,
		bd_uid bigint default 0   comment '商务人员' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		status tinyint default 0 not null  comment '状态' 
		,primary key (mer_no)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='商户信息'`