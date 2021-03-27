package scheme
 
//ots_notify_info 订单通知表
const ots_notify_info=`
	DROP TABLE IF EXISTS ots_notify_info;
	CREATE TABLE IF NOT EXISTS ots_notify_info (
		order_id varchar(32)  not null  comment '订单编号' ,
		mer_no varchar(32)  not null  comment '商户名称' ,
		mer_order_no varchar(64)  not null  comment '订单编号' ,
		notify_url varchar(128)  not null  comment '通知地址' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		start_time datetime    comment '开始时间' ,
		end_time datetime    comment '结束时间' ,
		max_count int default 10 not null  comment '最大通知次数' ,
		notify_count int default 0 not null  comment '通知次数' ,
		notify_status int default 10 not null  comment '通知状态（0成功,10未开始,20等待通知,30正在通知）' ,
		notify_msg varchar(256)    comment '通知结果' 
		,primary key (order_id)
		,unique index mer_order(mer_order_no)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='订单通知表'`