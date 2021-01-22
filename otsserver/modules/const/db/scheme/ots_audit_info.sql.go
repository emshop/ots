package scheme
 
//ots_audit_info 发货人工审核表
const ots_audit_info=`
	DROP TABLE IF EXISTS ots_audit_info;
	CREATE TABLE IF NOT EXISTS ots_audit_info (
		delivery_id bigint    comment '发货编号' ,
		order_id bigint  not null  comment '订单编号' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		delivery_status tinyint  not null  comment '发货结果' ,
		end_order tinyint  not null  comment '是否终结订单' ,
		add_to_blacklist tinyint  not null  comment '是否加入黑名单' ,
		audit_status int default 10 not null  comment '审核状态' ,
		audit_by int    comment '审核人' ,
		audit_time datetime    comment '审核时间' ,
		audit_msg varchar(256)    comment '审核信息' 
		
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='发货人工审核表'`