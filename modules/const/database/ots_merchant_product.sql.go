package database
 const ots_merchant_product=`CREATE TABLE  ots_merchant_product (
		mer_product_id number(10)  not null AUTO_INCREMENT comment '商品编号' ,
		mer_shelf_id number(10)  not null  comment '货架编号' ,
		pl_id number(10)  not null  comment '产品线' ,
		mer_product_no varchar2(32)    comment '商户商品编号' ,
		carrier_no varchar2(8)  not null  comment '运营商' ,
		province_no varchar2(8) default '-' not null  comment '省份' ,
		city_no varchar2(8) default '-' not null  comment '城市' ,
		invoice_type number(2)  not null  comment '开票方式（1.不开发票）' ,
		face number(20,5)  not null  comment '面值' ,
		user_discount number(10,5)  not null  comment '用户折扣（以面值算）' ,
		mer_fee_discount number(10,5)  not null  comment '商户佣金' ,
		trade_fee_discount number(10,5)  not null  comment '交易服务费' ,
		payment_fee_discount number(10,5)  not null  comment '支付手续费' ,
		status number(1) default 0 not null  comment '状态(0.是,1.否)' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (mer_product_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='商户商品';`