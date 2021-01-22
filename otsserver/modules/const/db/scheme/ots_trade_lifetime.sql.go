package scheme
 
//ots_trade_lifetime 生命周期记录表
const ots_trade_lifetime=`
	DROP TABLE IF EXISTS ots_trade_lifetime;
	CREATE TABLE IF NOT EXISTS ots_trade_lifetime (
		id bigint default 100000 not null  comment '编号' ,
		order_id bigint  not null  comment '订单编号' ,
		delivery_id varchar(30)    comment '发货编号' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		ip varchar(20)    comment '服务器ip' ,
		content varchar(1000)  not null  comment '操作内容' 
		,primary key (id)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='生命周期记录表'`