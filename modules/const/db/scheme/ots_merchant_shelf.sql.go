package scheme
 
//ots_merchant_shelf 商户货架
const ots_merchant_shelf=`CREATE TABLE  ots_merchant_shelf (
		mer_shelf_id number(10)  not null AUTO_INCREMENT comment '货架编号' ,
		mer_shelf_name varchar2(64)  not null  comment '货架名称' ,
		mer_no varchar2(32)  not null  comment '商户编号' ,
		mer_fee_discount number(10,5) default 0 not null  comment '商户佣金' ,
		trade_fee_discount number(10,5) default 0 not null  comment '交易服务费' ,
		payment_fee_discount number(10,5) default 0 not null  comment '支付手续费' ,
		order_timeout number(10)  not null  comment '订单超时时长' ,
		payment_timeout number(10)  not null  comment '支付超时时长' ,
		invoice_type number(2) default 1 not null  comment '开票方式（1.不开发票）' ,
		can_refund number(1) default 1 not null  comment '允许退款(0.是,1否)' ,
		limit_count number(10) default 1 not null  comment '单次购买数量' ,
		split_face number(20) default 0 not null  comment '拆单面值' ,
		status number(1) default 0 not null  comment '状态' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (mer_shelf_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='商户货架';`