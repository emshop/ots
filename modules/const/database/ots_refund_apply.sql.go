package database
 const ots_refund_apply=`CREATE TABLE  ots_refund_apply (
		apply_id number(20)  not null  comment '申请编号' ,
		order_id number(20)  not null  comment '订单编号' ,
		mer_no varchar2(32)  not null  comment '商户编号' ,
		mer_order_no varchar2(32)  not null  comment '商户订单号' ,
		refund_cause number(3) default 10 not null  comment '退款原因' ,
		apply_status number(2) default 10 not null  comment '申请状态' ,
		refund_amount number(20,5)  not null  comment '退款金额' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' 
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='退款申请';`