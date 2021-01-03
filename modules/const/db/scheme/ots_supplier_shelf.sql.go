package scheme
 
//ots_supplier_shelf 供货商货架
const ots_supplier_shelf=`CREATE TABLE  ots_supplier_shelf (
		spp_shelf_id int default 600 not null AUTO_INCREMENT comment '货架编号' ,
		spp_shelf_name varchar2(64)  not null  comment '货架名称' ,
		spp_no varchar2(32)  not null  comment '供货商编号' ,
		req_url varchar2(128)  not null  comment '请求地址' ,
		query_url varchar2(128)    comment '查询地址' ,
		notify_url varchar2(128)  not null  comment '回调地址' ,
		invoice_type tinyint default 0 not null  comment '开票方式（0.不开发票）' ,
		spp_fee_discount decimal(10,5) default 0 not null  comment '商户佣金' ,
		trade_fee_discount decimal(10,5) default 0 not null  comment '交易服务费' ,
		payment_fee_discount decimal(10,5) default 0 not null  comment '支付手续费' ,
		can_refund tinyint default 1 not null  comment '支持退货 (0.是,1.否)' ,
		status tinyint default 0 not null  comment '货架状态' ,
		limit_count int default 1 not null  comment '单次最大发货数量' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,PRIMARY KEY (spp_shelf_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='供货商货架'`