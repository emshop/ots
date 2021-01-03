package scheme
 
//ots_merchant_product 商户商品
const ots_merchant_product=`CREATE TABLE  ots_merchant_product (
		mer_product_id int default 10000 not null AUTO_INCREMENT comment '商品编号' ,
		mer_shelf_id int  not null  comment '货架编号' ,
		mer_no varchar2(32)  not null  comment '商户编号' ,
		pl_id int  not null  comment '产品线' ,
		brand_no varchar2(8)  not null  comment '品牌' ,
		province_no varchar2(8) default '-' not null  comment '省份' ,
		city_no varchar2(8) default '-' not null  comment '城市' ,
		face decimal(10,5)  not null  comment '面值' ,
		mer_product_no varchar2(32)    comment '商户商品编号' ,
		discount decimal(10,5)  not null  comment '销售折扣（以面值算）' ,
		status tinyint default 0 not null  comment '状态(0.是,1.否)' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,PRIMARY KEY (mer_product_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='商户商品'`