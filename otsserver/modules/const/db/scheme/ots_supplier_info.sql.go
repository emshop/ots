package scheme
 
//ots_supplier_info 供货商信息
const ots_supplier_info=`
	DROP TABLE IF EXISTS ots_supplier_info;
	CREATE TABLE IF NOT EXISTS ots_supplier_info (
		spp_no varchar(32)  not null  comment '编号' ,
		spp_name varchar(64)  not null  comment '名称' ,
		mer_crop varchar(64)    comment '公司' ,
		bd_uid bigint default 0   comment '商务人员' ,
		status tinyint default 0   comment '状态' ,
		create_time datetime default current_timestamp   comment '创建时间' 
		,primary key (spp_no)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='供货商信息'`