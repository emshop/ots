package scheme
 
//ots_merchant_shelf 商户货架
const ots_merchant_shelf=`
	DROP TABLE IF EXISTS ots_merchant_shelf;
	CREATE TABLE IF NOT EXISTS ots_merchant_shelf (
		mer_shelf_id int  not null auto_increment comment '货架编号' ,
		mer_shelf_name varchar(64)  not null  comment '货架名称' ,
		mer_no varchar(32)  not null  comment '商户编号' ,
		mer_fee_discount decimal(10,5) default 0 not null  comment '商户佣金' ,
		trade_fee_discount decimal(10,5) default 0 not null  comment '交易服务费' ,
		payment_fee_discount decimal(10,5) default 0 not null  comment '支付手续费' ,
		order_timeout int  not null  comment '订单超时时长' ,
		payment_timeout int  not null  comment '支付超时时长' ,
		invoice_type tinyint default 1 not null  comment '开票方式（1.不开发票）' ,
		can_refund tinyint default 1 not null  comment '允许退款(0.是,1否)' ,
		limit_count int default 1 not null  comment '单次购买数量' ,
		can_split_order tinyint default 0 not null  comment '允许拆单' ,
		status tinyint default 0 not null  comment '状态' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (mer_shelf_id)
	) ENGINE=InnoDB auto_increment = 100 DEFAULT CHARSET=utf8mb4 COMMENT='商户货架'`