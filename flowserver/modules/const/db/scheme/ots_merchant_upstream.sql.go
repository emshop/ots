package scheme
 
//ots_merchant_upstream 商户上游
const ots_merchant_upstream=`
	DROP TABLE IF EXISTS ots_merchant_upstream;
	CREATE TABLE IF NOT EXISTS ots_merchant_upstream (
		mu_id int  not null auto_increment comment '编号' ,
		mer_shelf_id int default 0 not null  comment '货架' ,
		spp_no varchar(32)  not null  comment '供货商' ,
		status tinyint default 0 not null  comment '状态' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (mu_id)
	) ENGINE=InnoDB auto_increment = 100 DEFAULT CHARSET=utf8mb4 COMMENT='商户上游'`