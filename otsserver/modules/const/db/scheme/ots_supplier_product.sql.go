package scheme
 
//ots_supplier_product 供货商商品
const ots_supplier_product=`
	DROP TABLE IF EXISTS ots_supplier_product;
	CREATE TABLE IF NOT EXISTS ots_supplier_product (
		spp_product_id int  not null auto_increment comment '商品编号' ,
		spp_shelf_id int  not null  comment '货架' ,
		spp_no varchar(32)  not null  comment '供货商' ,
		spp_product_no varchar(32)    comment '供货商商品编号' ,
		pl_id int  not null  comment '产品线' ,
		brand_no varchar(8)  not null  comment '品牌' ,
		province_no varchar(8) default '*'   comment '省份' ,
		city_no varchar(8) default '*'   comment '城市' ,
		face int  not null  comment '面值' ,
		cost_discount decimal(10,5)  not null  comment '成本折扣' ,
		status tinyint default 0 not null  comment '状态' ,
		create_time datetime default current_timestamp   comment '创建时间' 
		,primary key (spp_product_id)
	) ENGINE=InnoDB auto_increment = 60000 DEFAULT CHARSET=utf8mb4 COMMENT='供货商商品'`