package database
 const ots_audit_info=`CREATE TABLE  ots_audit_info (
		delivery_id number(20)    comment '发货编号' ,
		order_id number(20)  not null  comment '订单编号' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		delivery_status number(2)  not null  comment '发货结果' ,
		end_order number(1)  not null  comment '是否终结订单' ,
		add_to_blacklist number(1)  not null  comment '是否加入黑名单' ,
		audit_status number(3)  not null  comment '审核状态' ,
		audit_by number(10)    comment '审核人' ,
		audit_time DATETIME    comment '审核时间' ,
		audit_msg varchar2(256)    comment '审核信息' 
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='发货人工审核表';`