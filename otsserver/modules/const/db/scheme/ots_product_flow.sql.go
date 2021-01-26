package scheme
 
//ots_product_flow 业务流程
const ots_product_flow=`
	DROP TABLE IF EXISTS ots_product_flow;
	CREATE TABLE IF NOT EXISTS ots_product_flow (
		flow_id int  not null auto_increment comment '流程编号' ,
		flow_name varchar(64)  not null  comment '流程名称' ,
		tag_name varchar(64)  not null  comment 'tag标签' ,
		pl_id int  not null  comment '产品线编号' ,
		parent_flow_id int default 0 not null  comment '父级流程编号' ,
		success_flow_id varchar(32) default '-' not null  comment '成功后续流程' ,
		failed_flow_id varchar(32) default '-' not null  comment '失败后续流程' ,
		unknown_flow_id varchar(32) default '-' not null  comment '未知后续流程' ,
		queue_name varchar(64) default '-' not null  comment '队列名称' ,
		scan_interval int  not null  comment '超时时长' ,
		delay int default 0 not null  comment '延后处理时长' ,
		timeout int  not null  comment '超时时长' ,
		max_count int  not null  comment '最大执行次数' 
		,primary key (flow_id)
		,unique index unq_flow_tag(tag_name,pl_id)
	) ENGINE=InnoDB auto_increment = 200 DEFAULT CHARSET=utf8mb4 COMMENT='业务流程'`