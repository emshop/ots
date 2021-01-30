package scheme
 
//ots_trade_delivery 订单发货表
const ots_trade_delivery=`
	DROP TABLE IF EXISTS ots_trade_delivery;
	CREATE TABLE IF NOT EXISTS ots_trade_delivery (
		delivery_id varchar(32) default 20000 not null  comment '发货编号' ,
		order_id varchar(32)  not null  comment '订单编号' ,
		spp_no varchar(32)  not null  comment '供货商' ,
		spp_product_id int  not null  comment '供货商商品编号' ,
		mer_no varchar(32)  not null  comment '商户编号' ,
		mer_product_id int  not null  comment '商户商品编号' ,
		pl_id int  not null  comment '产品线' ,
		brand_no varchar(8)  not null  comment '品牌' ,
		province_no varchar(8)  not null  comment '省份' ,
		city_no varchar(8)  not null  comment '城市' ,
		invoice_type int  not null  comment '开票方式（1.不开发票）' ,
		account_name varchar(64)  not null  comment '用户账户' ,
		delivery_status int default 20 not null  comment '发货状态' ,
		payment_status int default 10 not null  comment '支付状态' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		face int  not null  comment '商品面值' ,
		num int  not null  comment '发货数量' ,
		total_face int  not null  comment '发货总面值' ,
		cost_discount decimal(20,5)  not null  comment '成本折扣' ,
		spp_fee_discount decimal(10,5) default 0 not null  comment '商户佣金' ,
		trade_fee_discount decimal(10,5) default 0 not null  comment '交易服务费' ,
		payment_fee_discount decimal(10,5) default 0 not null  comment '支付手续费' ,
		cost_amount decimal(20,5)  not null  comment '发货成本' ,
		spp_fee_amount decimal(20,5)  not null  comment '供货商佣金' ,
		trade_fee_amount decimal(20,5)  not null  comment '发货服务费' ,
		payment_fee_amount decimal(20,5)  not null  comment '支付服务费' ,
		succ_face int default 0   comment '成功面值' ,
		start_time datetime    comment '开始时间' ,
		end_time datetime    comment '结束时间' ,
		spp_delivery_no varchar(32)    comment '供货商发货编号' ,
		spp_product_no varchar(32)    comment '供货商商品编号' ,
		real_discount decimal(20,5)    comment '供货商实际折扣' ,
		return_msg varchar(256)    comment '发货结果' ,
		request_params varchar(2000)    comment '发货信息参数json' ,
		result_source varchar(32)    comment '发货结果来源（1：通知，2：查询，3：同步返回）' ,
		result_code varchar(32)    comment '发货结果码' ,
		last_update_time datetime default current_timestamp not null  comment '最后更新时间' 
		,primary key (delivery_id)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='订单发货表'`