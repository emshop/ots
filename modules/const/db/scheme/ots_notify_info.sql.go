package scheme
 
//ots_notify_info 订单通知表
const ots_notify_info=`CREATE TABLE  ots_notify_info (
		order_id bigint  not null  comment '订单编号' ,
		mer_no varchar2(32)  not null  comment '商户编号' ,
		mer_order_no varchar2(64)  not null  comment '商户订单编号' ,
		mer_product_id int  not null  comment '商户商品编号' ,
		notify_url varchar2(128)  not null  comment '通知地址' ,
		order_status int default 10 not null  comment '订单状态' ,
		notify_status int default 10 not null  comment '通知状态（0成功,10未开始,20等待通知,30正在通知）' ,
		max_count int  not null  comment '最大通知次数' ,
		notify_count int default 0 not null  comment '通知次数' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		start_time datetime    comment '开始时间' ,
		end_time datetime    comment '结束时间' ,
		notify_msg varchar2(256)    comment '通知结果' 
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='订单通知表'`