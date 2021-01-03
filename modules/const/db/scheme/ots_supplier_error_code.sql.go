package scheme
 
//ots_supplier_error_code 供货商错误码
const ots_supplier_error_code=`CREATE TABLE  ots_supplier_error_code (
		id bigint default 100 not null AUTO_INCREMENT comment '编号' ,
		spp_no varchar2(32)  not null  comment '编号' ,
		pl_id int  not null  comment '产品线' ,
		deal_code tinyint  not null  comment '处理码' ,
		error_code varchar2(32)  not null  comment '错误码' ,
		error_desc varchar2(64)  not null  comment '错误码描述' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='供货商错误码'`