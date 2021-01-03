package scheme
 
//ots_supplier_product 供货商商品
const ots_supplier_product=`CREATE TABLE  ots_supplier_product (
		spp_product_id int default 60000 not null AUTO_INCREMENT comment '商品编号' ,
		spp_shelf_id int  not null  comment '货架编号' ,
		spp_no varchar2(32)  not null  comment '供货商编号' ,
		spp_product_no varchar2(32)    comment '供货商商品编号' ,
		pl_id int  not null  comment '产品线' ,
		brand_no varchar2(8)  not null  comment '品牌' ,
		province_no varchar2(8) default '-' not null  comment '省份' ,
		city_no varchar2(8) default '-' not null  comment '城市' ,
		face decimal(20,5)  not null  comment '面值' ,
		cost_discount decimal(10,5)  not null  comment '成本折扣' ,
		status tinyint default 0 not null  comment '状态' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,PRIMARY KEY (spp_product_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='供货商商品'`