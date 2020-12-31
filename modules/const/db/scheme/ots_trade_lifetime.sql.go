package scheme
 
//ots_trade_lifetime 生命周期记录表
const ots_trade_lifetime=`CREATE TABLE  ots_trade_lifetime (
		id number(20) default 100000 not null  comment '编号' ,
		order_id number(20)  not null  comment '订单编号' ,
		delivery_id varchar2(30)    comment '发货编号' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		ip varchar2(20)    comment '服务器ip' ,
		content varchar2(1000)  not null  comment '操作内容' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='生命周期记录表';`