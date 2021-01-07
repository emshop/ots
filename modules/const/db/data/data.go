package data

//DataList 初始需导入数据
var DataList = []string{
	"INSERT INTO `ots_merchant_info` (`mer_no`,`mer_name`,`mer_crop`,`mer_type`,`bd_uid`,`status`,`create_time`) VALUES ('qx001','千行惠捷','惠捷能源',0,0,0,'2021-01-06 16:02:44');",
	"INSERT INTO `ots_merchant_shelf` (`mer_shelf_id`,`mer_shelf_name`,`mer_no`,`mer_fee_discount`,`trade_fee_discount`,`payment_fee_discount`,`order_timeout`,`payment_timeout`,`invoice_type`,`can_refund`,`limit_count`,`can_split_order`,`status`,`create_time`) VALUES (100,'充值类商品','qx001',0.00100,0.00200,0.00300,86400,3600,1,1,1,0,0,'2021-01-06 16:05:16');",
	"INSERT INTO `ots_merchant_product` (`mer_product_id`,`mer_shelf_id`,`mer_no`,`pl_id`,`brand_no`,`province_no`,`city_no`,`face`,`mer_product_no`,`discount`,`status`,`create_time`) VALUES (10000,100,'qx001',10000,'zsy','*','*',100.00000,'',0.99000,0,'2021-01-06 16:06:52');",
	"INSERT INTO `ots_product_line` (`pl_id`,`pl_name`,`status`,`create_time`) VALUES (10000,'直充类商品',0,'2021-01-06 16:07:33');",
	"INSERT INTO `ots_product_flow` (`flow_id`,`flow_Name`,`pl_id`,`parent_flow_id`,`success_flow_id`,`failed_flow_id`,`unknown_flow_id`,`queue_name`,`scan_interval`,`delay`,`timeout`,`max_count`) VALUES (1100,'收单流程',10000,0,'1110','0','0','-',0,0,0,0),(1110,'支付流程',10000,1100,'1120','1119','0','order_pay',10,0,86400,100),(1119,'支付超时',10000,1110,'1160','0','0','pay_timeout',30,0,86400,100),(1120,'绑定流程',10000,1110,'1130','1129','0','order_bind',30,0,86400,10000),(1130,'发货流程',10000,1120,'1140','1139','1138','order_delivery',120,0,86400,10000);",
}
