package scheme
 
//ots_refund_apply 退款申请
const ots_refund_apply=`
	DROP TABLE IF EXISTS ots_refund_apply;
	CREATE TABLE IF NOT EXISTS ots_refund_apply (
		apply_id bigint default 100 not null  comment '申请编号' ,
		order_id bigint  not null  comment '订单编号' ,
		mer_no varchar(32)  not null  comment '商户编号' ,
		mer_order_no varchar(32)  not null  comment '商户订单号' ,
		refund_cause int default 10 not null  comment '退款原因' ,
		apply_status tinyint default 10 not null  comment '申请状态' ,
		refund_amount decimal(20,5)  not null  comment '退款金额' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='退款申请'`