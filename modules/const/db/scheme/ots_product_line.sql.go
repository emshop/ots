package scheme
 
//ots_product_line 产品线
const ots_product_line=`CREATE TABLE  ots_product_line (
		pl_id number(10)  not null AUTO_INCREMENT comment '产品线编号' ,
		pl_name varchar2(64)  not null  comment '产品线名称' ,
		status number(1) default 0 not null  comment '状态' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (pl_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='产品线';`