package scheme
 
//ots_supplier_error_code 供货商错误码
const ots_supplier_error_code=`
	DROP TABLE IF EXISTS ots_supplier_error_code;
	CREATE TABLE IF NOT EXISTS ots_supplier_error_code (
		id bigint  not null auto_increment comment '编号' ,
		spp_no varchar(32)  not null  comment '编号' ,
		pl_id int  not null  comment '产品线' ,
		category varchar(32)  not null  comment '产品线' ,
		deal_code tinyint  not null  comment '处理码' ,
		error_code varchar(32)  not null  comment '错误码' ,
		status tinyint default 0 not null  comment '状态' ,
		error_desc varchar(64)  not null  comment '错误码描述' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,unique index unq_spp_error_code(spp_no,pl_id,category,error_code)
		,primary key (id)
	) ENGINE=InnoDB auto_increment = 100 DEFAULT CHARSET=utf8 COMMENT='供货商错误码'`