package scheme
 
//ots_product_line 产品线
const ots_product_line=`
	DROP TABLE IF EXISTS ots_product_line;
	CREATE TABLE IF NOT EXISTS ots_product_line (
		pl_id int  not null auto_increment comment '产品线编号' ,
		pl_name varchar(64)  not null  comment '产品线名称' ,
		pl_type tinyint default 0 not null  comment '类型' ,
		status tinyint default 0 not null  comment '状态' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (pl_id)
	) ENGINE=InnoDB auto_increment = 100 DEFAULT CHARSET=utf8mb4 COMMENT='产品线'`