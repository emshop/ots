// +build !oracle

package sql

import (
	"github.com/micro-plat/hydra"
)

func Install() {
	hydra.Installer.DB.AddSQL(`DROP TABLE IF EXISTS  lcs_life_time;
		CREATE TABLE  lcs_life_time (
			id bigint  not null AUTO_INCREMENT comment 'id' ,
			order_no varchar(32)  not null  comment '子系统唯一标识' ,
			batch_no varchar(32)    comment '自定义字段' ,
			extral_param varchar(32)    comment '子系统唯一标识' ,
			ip varchar(32)    comment '用户ip' ,
			content varchar(1000)  not null  comment '内容' ,
			create_time datetime default current_timestamp not null  comment '创建时间' ,
			PRIMARY KEY (id),
			KEY key_lcs_life_time_order_no (order_no)
			) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='生命周期记录表';`)
}
