package scheme
 
//ots_trade_order 订单记录
const ots_trade_order=`
	DROP TABLE IF EXISTS ots_trade_order;
	CREATE TABLE IF NOT EXISTS ots_trade_order (
		order_id bigint default 1100000000 not null  comment '订单编号' ,
		mer_no varchar(32)  not null  comment '商户编号' ,
		mer_order_no varchar(64)  not null  comment '商户订单编号' ,
		mer_product_id int default 300 not null  comment '商品编号' ,
		mer_shelf_id int  not null  comment '货架编号' ,
		mer_product_no varchar(32)    comment '外部商品编号' ,
		pl_id int  not null  comment '产品线' ,
		brand_no varchar(8)  not null  comment '品牌' ,
		province_no varchar(8)  not null  comment '省份' ,
		city_no varchar(8)  not null  comment '城市' ,
		face decimal(20,5)  not null  comment '商品面值' ,
		num int  not null  comment '商品数量' ,
		total_face decimal(20,5)  not null  comment '商品总面值' ,
		account_name varchar(64)  not null  comment '用户账户信息' ,
		invoice_type tinyint  not null  comment '开票方式（1.不开发票）' ,
		sell_discount decimal(20,5)  not null  comment '销售折扣' ,
		sell_amount decimal(20,5)  not null  comment '总销售金额' ,
		mer_fee_amount decimal(20,5)  not null  comment '商户佣金金额' ,
		trade_fee_amount decimal(20,5)  not null  comment '交易服务费金额' ,
		payment_fee_amount decimal(20,5)  not null  comment '支付手续费金额' ,
		can_split_order tinyint default 1 not null  comment '是否拆单（0.是，1否）' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		order_timeout datetime  not null  comment '订单超时时间' ,
		payment_timeout datetime  not null  comment '支付超时时间' ,
		delivery_pause tinyint default 1 not null  comment '发货暂停（0.是，1否）' ,
		order_status int default 10 not null  comment '订单状态' ,
		payment_status int default 20 not null  comment '支付状态' ,
		delivery_status int default 10 not null  comment '发货状态' ,
		refund_status int default 10 not null  comment '退款状态' ,
		notify_status int default 10 not null  comment '通知状态' ,
		is_refund tinyint default 1 not null  comment '用户退款（0.是，1否）' ,
		bind_face decimal(20,5) default 0 not null  comment '成功绑定总面值' ,
		success_face decimal(20,5) default 0 not null  comment '实际成功总面值' ,
		success_user_amount decimal(20,5) default 0 not null  comment '实际成功总销售金额 （1）' ,
		success_mer_fee decimal(20,5) default 0 not null  comment '实际成功总佣金金额 （2）' ,
		success_trade_fee decimal(20,5) default 0 not null  comment '实际成功总服务费金额 （3）' ,
		success_payment_fee decimal(20,5) default 0 not null  comment '实际成功总手续费金额 （4）' ,
		success_cost_amount decimal(20,5) default 0 not null  comment '实际发货成功总成本 （5）' ,
		success_spp_fee decimal(20,5) default 0 not null  comment '实际发货成功总供货商佣金 （6）' ,
		success_spp_trade_fee decimal(20,5) default 0 not null  comment '实际发货成功总供货商服务费 （7）' ,
		profit decimal(20,5) default 0 not null  comment '利润（1-2-3-4-5add6-7）' 
		,primary key (order_id)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='订单记录'`