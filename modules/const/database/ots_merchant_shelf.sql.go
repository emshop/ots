package database
 const ots_merchant_shelf=`CREATE TABLE  ots_merchant_shelf (
		mer_shelf_id number(10)  not null AUTO_INCREMENT comment '货架编号' ,
		mer_shelf_name varchar2(64)  not null  comment '货架名称' ,
		mer_no varchar2(32)  not null  comment '商户编号' ,
		can_refund number(1)  not null  comment '支持退款(0.是,1否)' ,
		limit_count number(10) default 1 not null  comment '单次最大购买数量' ,
		can_split number(1)  not null  comment '是否拆单(0.是,1.否)' ,
		split_face number(20,5)  not null  comment '拆单面值' ,
		order_timeout number(10)  not null  comment '订单超时时长' ,
		refund_timeout number(10)  not null  comment '退款超时时长' ,
		status number(1) default 0 not null  comment '状态' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (mer_shelf_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='商户货架';`