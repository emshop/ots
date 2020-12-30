package database
 const ots_supplier_shelf=`CREATE TABLE  ots_supplier_shelf (
		spp_shelf_id number(10)  not null AUTO_INCREMENT comment '货架编号' ,
		spp_shelf_name varchar2(64)  not null  comment '货架名称' ,
		spp_no varchar2(32)  not null  comment '供货商编号' ,
		req_url varchar2(128)  not null  comment '请求地址' ,
		query_url varchar2(128)    comment '查询地址' ,
		notify_url varchar2(128)  not null  comment '回调地址' ,
		can_refund number(1)  not null  comment '支持退货 (0.是,1.否)' ,
		request_replenish_time number(10)  not null  comment '发货后补间隔时间' ,
		can_query number(1) default 0 not null  comment '是否支持查询' ,
		first_query_time number(10)    comment '首次查询时间' ,
		query_replenish_time number(10)    comment '查询后补间隔时间' ,
		status number(1) default 0 not null  comment '货架状态' ,
		delivery_timeout number(10) default 300 not null  comment '发货超时时间' ,
		return_timeout number(10) default 300 not null  comment '退货超时时间' ,
		limit_count number(10) default 1 not null  comment '单次最大发货数量' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (spp_shelf_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='供货商货架';`