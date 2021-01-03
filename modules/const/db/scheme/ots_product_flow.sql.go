package scheme
 
//ots_product_flow 业务流程
const ots_product_flow=`CREATE TABLE  ots_product_flow (
		flow_id int default 200 not null AUTO_INCREMENT comment '流程编号' ,
		pl_id int  not null  comment '产品线编号' ,
		parent_flow_id int default 0 not null  comment '父级流程编号' ,
		success_flow_id varchar2(32) default '-' not null  comment '成功后续流程' ,
		failed_flow_id varchar2(32) default '-' not null  comment '失败后续流程' ,
		unknown_flow_id varchar2(32) default '-' not null  comment '未知后续流程' ,
		queue_name varchar2(64) default '-' not null  comment '队列名称' ,
		delay int default 0 not null  comment '延后处理时长' ,
		timeout int  not null  comment '超时时长' ,
		max_count int  not null  comment '最大执行次数' 
		,PRIMARY KEY (flow_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='业务流程'`