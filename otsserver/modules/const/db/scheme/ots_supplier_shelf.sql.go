package scheme
 
//ots_supplier_shelf 供货商货架
const ots_supplier_shelf=`
	DROP TABLE IF EXISTS ots_supplier_shelf;
	CREATE TABLE IF NOT EXISTS ots_supplier_shelf (
		spp_shelf_id int  not null auto_increment comment '货架编号' ,
		spp_shelf_name varchar(64)  not null  comment '货架名称' ,
		spp_no varchar(32)  not null  comment '供货商' ,
		req_url varchar(128)  not null  comment '请求地址' ,
		query_url varchar(128)    comment '查询地址' ,
		notify_url varchar(128)    comment '回调地址' ,
		invoice_type tinyint default 1   comment '开发票' ,
		spp_fee_discount decimal(10,5) default 0   comment '商户佣金' ,
		trade_fee_discount decimal(10,5) default 0   comment '交易服务费' ,
		payment_fee_discount decimal(10,5) default 0   comment '支付手续费' ,
		can_refund tinyint default 1   comment '支持退货 (0.是,1.否)' ,
		status tinyint default 0   comment '货架状态' ,
		limit_count int default 1   comment '单次最大发货数量' ,
		create_time datetime default current_timestamp   comment '创建时间' 
		,primary key (spp_shelf_id)
	) ENGINE=InnoDB auto_increment = 600 DEFAULT CHARSET=utf8mb4 COMMENT='供货商货架'`