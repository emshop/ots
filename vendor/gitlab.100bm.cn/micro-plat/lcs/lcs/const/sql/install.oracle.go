// +build oracle

package sql

import (
	"github.com/micro-plat/hydra"
)

func Install() {
	hydra.Installer.DB.AddSQL(`DROP TABLE lcs_life_time;
	create table lcs_life_time(
		id NUMBER(19)  not null ,
		order_no VARCHAR2(32)  not null ,
		batch_no VARCHAR2(32)   ,
		extral_param VARCHAR2(32)   ,
		ip VARCHAR2(32)   ,
		content VARCHAR2(1000)  not null ,
		create_time TIMESTAMP default current_timestamp not null 
		);
	

	comment on table lcs_life_time is '生命周期记录表';
	comment on column lcs_life_time.id is 'id';	
	comment on column lcs_life_time.order_no is '子系统唯一标识';	
	comment on column lcs_life_time.batch_no is '自定义字段';	
	comment on column lcs_life_time.extral_param is '子系统唯一标识';	
	comment on column lcs_life_time.ip is '用户ip';	
	comment on column lcs_life_time.content is '内容';	
	comment on column lcs_life_time.create_time is '创建时间';	
	

 
	alter table lcs_life_time
	add constraint pk_life_time primary key(id);
	create index key_lcs_life_time_order_no on lcs_life_time(order_no);
	
	create sequence seq_life_time_id
	increment by 1
	minvalue 1
	maxvalue 99999999999
	start with 1
	cache 20;`)
}
