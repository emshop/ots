package scheme
 
//ots_product_flow 业务流程
const ots_product_flow=`
	DROP TABLE IF EXISTS ots_product_flow;
	CREATE TABLE IF NOT EXISTS ots_product_flow (
		flow_id int  not null auto_increment comment '流程编号' ,
		flow_tag varchar(32)  not null  comment '流程名称' ,
		pl_id int  not null  comment '产品线' ,
		queue_name varchar(64) default '-'   comment '队列名称' ,
		scan_interval int default 0   comment '执行间隔' ,
		delay int default 0   comment '延后时长' ,
		timeout int default 0   comment '超时时长' ,
		max_count int default 0   comment '最大次数' ,
		status tinyint default 0   comment '状态' 
		,primary key (flow_id)
		,unique index unq_flow_tag(flow_tag,pl_id)
	) ENGINE=InnoDB auto_increment = 200 DEFAULT CHARSET=utf8mb4 COMMENT='业务流程'`