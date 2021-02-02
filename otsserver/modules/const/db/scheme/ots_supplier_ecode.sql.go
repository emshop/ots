package scheme
 
//ots_supplier_ecode 供货商错误码
const ots_supplier_ecode=`
	DROP TABLE IF EXISTS ots_supplier_ecode;
	CREATE TABLE IF NOT EXISTS ots_supplier_ecode (
		id bigint  not null auto_increment comment '编号' ,
		spp_no varchar(32)  not null  comment '商家' ,
		pl_id int  not null  comment '产品线' ,
		category varchar(32)  not null  comment '分类' ,
		deal_code tinyint  not null  comment '处理码' ,
		error_code varchar(32)  not null  comment '错误码' ,
		status tinyint default 0 not null  comment '状态' ,
		error_desc varchar(64)  not null  comment '错误码描述' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (id)
		,unique index unq_spp_error_code(spp_no,pl_id,category,error_code)
	) ENGINE=InnoDB auto_increment = 100 DEFAULT CHARSET=utf8mb4 COMMENT='供货商错误码'`