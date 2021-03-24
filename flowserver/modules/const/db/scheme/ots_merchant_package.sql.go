package scheme
 
//ots_merchant_package 组合商品
const ots_merchant_package=`
	DROP TABLE IF EXISTS ots_merchant_package;
	CREATE TABLE IF NOT EXISTS ots_merchant_package (
		pg_id int  not null auto_increment comment '包编号' ,
		pg_name varchar(64)  not null  comment '包名称' ,
		pl_id int default 100 not null  comment '产品线' ,
		mer_product_id int  not null  comment '商品编号' ,
		brand_no varchar(8)  not null  comment '品牌' ,
		province_no varchar(8) default '*' not null  comment '省份' ,
		city_no varchar(8) default '*' not null  comment '城市' ,
		face int  not null  comment '商品面值' ,
		num tinyint default 1 not null  comment '数量' ,
		status tinyint default 0 not null  comment '状态' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (pg_id)
		,unique index unq_pkg_prod(pl_id,brand_no,province_no,city_no,face)
	) ENGINE=InnoDB auto_increment = 100 DEFAULT CHARSET=utf8mb4 COMMENT='组合商品'`