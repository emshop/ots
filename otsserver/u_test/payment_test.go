package ut

import (
	"testing"
	"time"

	"github.com/emshop/ots/otsserver/services/payment"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra/mock"
	"github.com/micro-plat/lib4go/assert"
	"github.com/micro-plat/lib4go/types"
)

func TestSinglePay(t *testing.T) {
	preper()
	err := orderCache.makeOrder(1)
	assert.Equal(t, nil, err)

	order := <-orderCache.order

	//构建context
	ctx := mock.NewContext(string(order.Marshal()))

	//构建请求处理------------------------
	rs := payment.Paying(ctx)
	err = ctx.Response().WriteAny(getAny(rs))
	status, _, _ := ctx.Response().GetFinalResponse()
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, status)

}
func TestPay(t *testing.T) {
	preper()
	var cases = []struct {
		sql    string
		input  types.XMap
		status int
	}{
		{status: 200, input: map[string]interface{}{"order_id": "@OrderID", "flow_id": "1110"}, sql: "INSERT INTO `ots_trade_order`(`order_id`,`mer_no`,`mer_order_no`,`mer_product_id`,`mer_shelf_id`,`mer_product_no`,`pl_id`,`brand_no`,`province_no`,`city_no`,`face`,`num`,`total_face`,`account_name`,`invoice_type`,`sell_discount`,`sell_amount`,`mer_fee_discount`,`trade_fee_discount`,`payment_fee_discount`,`can_split_order`,`create_time`,`order_timeout`,`payment_timeout`,`delivery_pause`)VALUES (@OrderID,'qx001','or0001',10000,100,'',10000,'zsy','*','*',100,1,100,'15828680877',1,0.99000,99.00000,0.00100,0.00200,0.00300,0,'2021-01-19 15:24:27','2021-01-20 15:24:27','2021-01-19 16:24:27',1)"},
		{status: 200, input: map[string]interface{}{"order_id": "@OrderID", "flow_id": "1110"}, sql: "INSERT INTO `ots_trade_order`(`order_status`,`payment_status`,`order_id`,`mer_no`,`mer_order_no`,`mer_product_id`,`mer_shelf_id`,`mer_product_no`,`pl_id`,`brand_no`,`province_no`,`city_no`,`face`,`num`,`total_face`,`account_name`,`invoice_type`,`sell_discount`,`sell_amount`,`mer_fee_discount`,`trade_fee_discount`,`payment_fee_discount`,`can_split_order`,`create_time`,`order_timeout`,`payment_timeout`,`delivery_pause`)VALUES (10,20,@OrderID,'qx001','or0001',10000,100,'',10000,'zsy','*','*',100,1,100,'15828680877',1,0.99000,99.00000,0.00100,0.00200,0.00300,0,'2021-01-19 15:24:27','2021-01-20 15:24:27','2021-01-19 16:24:27',1)"},
		{status: 204, input: map[string]interface{}{"order_id": "@OrderID", "flow_id": "1110"}, sql: "INSERT INTO `ots_trade_order`(`order_status`,`payment_status`,`order_id`,`mer_no`,`mer_order_no`,`mer_product_id`,`mer_shelf_id`,`mer_product_no`,`pl_id`,`brand_no`,`province_no`,`city_no`,`face`,`num`,`total_face`,`account_name`,`invoice_type`,`sell_discount`,`sell_amount`,`mer_fee_discount`,`trade_fee_discount`,`payment_fee_discount`,`can_split_order`,`create_time`,`order_timeout`,`payment_timeout`,`delivery_pause`)VALUES (10,20,@OrderID,'qx001','or0001',10000,100,'',10000,'zsy','*','*',100,1,100,'15828680877',1,0.99000,99.00000,0.00100,0.00200,0.00300,0,'2021-01-19 15:24:27','2021-01-20 15:24:27','2021-01-19 16:24:27',0)"},
		{status: 204, input: map[string]interface{}{"order_id": "@OrderID", "flow_id": "1110"}, sql: "INSERT INTO `ots_trade_order`(`order_status`,`payment_status`,`order_id`,`mer_no`,`mer_order_no`,`mer_product_id`,`mer_shelf_id`,`mer_product_no`,`pl_id`,`brand_no`,`province_no`,`city_no`,`face`,`num`,`total_face`,`account_name`,`invoice_type`,`sell_discount`,`sell_amount`,`mer_fee_discount`,`trade_fee_discount`,`payment_fee_discount`,`can_split_order`,`create_time`,`order_timeout`,`payment_timeout`)VALUES (10,10,@OrderID,'qx001','or0001',10000,100,'',10000,'zsy','*','*',100,1,100,'15828680877',1,0.99000,99.00000,0.00100,0.00200,0.00300,0,'2021-01-19 15:24:27','2021-01-20 15:24:27','2021-01-19 16:24:27')"},
		{status: 204, input: map[string]interface{}{"order_id": "@OrderID", "flow_id": "1110"}, sql: "INSERT INTO `ots_trade_order`(`order_status`,`payment_status`,`order_id`,`mer_no`,`mer_order_no`,`mer_product_id`,`mer_shelf_id`,`mer_product_no`,`pl_id`,`brand_no`,`province_no`,`city_no`,`face`,`num`,`total_face`,`account_name`,`invoice_type`,`sell_discount`,`sell_amount`,`mer_fee_discount`,`trade_fee_discount`,`payment_fee_discount`,`can_split_order`,`create_time`,`order_timeout`,`payment_timeout`)VALUES (10,30,@OrderID,'qx001','or0001',10000,100,'',10000,'zsy','*','*',100,1,100,'15828680877',1,0.99000,99.00000,0.00100,0.00200,0.00300,0,'2021-01-19 15:24:27','2021-01-20 15:24:27','2021-01-19 16:24:27')"},
		{status: 204, input: map[string]interface{}{"order_id": "@OrderID", "flow_id": "1110"}, sql: "INSERT INTO `ots_trade_order`(`order_status`,`payment_status`,`order_id`,`mer_no`,`mer_order_no`,`mer_product_id`,`mer_shelf_id`,`mer_product_no`,`pl_id`,`brand_no`,`province_no`,`city_no`,`face`,`num`,`total_face`,`account_name`,`invoice_type`,`sell_discount`,`sell_amount`,`mer_fee_discount`,`trade_fee_discount`,`payment_fee_discount`,`can_split_order`,`create_time`,`order_timeout`,`payment_timeout`)VALUES (10,00,@OrderID,'qx001','or0001',10000,100,'',10000,'zsy','*','*',100,1,100,'15828680877',1,0.99000,99.00000,0.00100,0.00200,0.00300,0,'2021-01-19 15:24:27','2021-01-20 15:24:27','2021-01-19 16:24:27')"},
		{status: 204, input: map[string]interface{}{"order_id": "@OrderID", "flow_id": "1110"}, sql: "INSERT INTO `ots_trade_order`(`order_status`,`payment_status`,`order_id`,`mer_no`,`mer_order_no`,`mer_product_id`,`mer_shelf_id`,`mer_product_no`,`pl_id`,`brand_no`,`province_no`,`city_no`,`face`,`num`,`total_face`,`account_name`,`invoice_type`,`sell_discount`,`sell_amount`,`mer_fee_discount`,`trade_fee_discount`,`payment_fee_discount`,`can_split_order`,`create_time`,`order_timeout`,`payment_timeout`)VALUES (10,40,@OrderID,'qx001','or0001',10000,100,'',10000,'zsy','*','*',100,1,100,'15828680877',1,0.99000,99.00000,0.00100,0.00200,0.00300,0,'2021-01-19 15:24:27','2021-01-20 15:24:27','2021-01-19 16:24:27')"},
		{status: 204, input: map[string]interface{}{"order_id": "@OrderID", "flow_id": "1110"}, sql: "INSERT INTO `ots_trade_order`(`order_status`,`payment_status`,`order_id`,`mer_no`,`mer_order_no`,`mer_product_id`,`mer_shelf_id`,`mer_product_no`,`pl_id`,`brand_no`,`province_no`,`city_no`,`face`,`num`,`total_face`,`account_name`,`invoice_type`,`sell_discount`,`sell_amount`,`mer_fee_discount`,`trade_fee_discount`,`payment_fee_discount`,`can_split_order`,`create_time`,`order_timeout`,`payment_timeout`)VALUES (20,10,@OrderID,'qx001','or0001',10000,100,'',10000,'zsy','*','*',100,1,100,'15828680877',1,0.99000,99.00000,0.00100,0.00200,0.00300,0,'2021-01-19 15:24:27','2021-01-20 15:24:27','2021-01-19 16:24:27')"},
		{status: 204, input: map[string]interface{}{"order_id": "@OrderID", "flow_id": "1110"}, sql: "INSERT INTO `ots_trade_order`(`order_status`,`payment_status`,`order_id`,`mer_no`,`mer_order_no`,`mer_product_id`,`mer_shelf_id`,`mer_product_no`,`pl_id`,`brand_no`,`province_no`,`city_no`,`face`,`num`,`total_face`,`account_name`,`invoice_type`,`sell_discount`,`sell_amount`,`mer_fee_discount`,`trade_fee_discount`,`payment_fee_discount`,`can_split_order`,`create_time`,`order_timeout`,`payment_timeout`)VALUES (20,20,@OrderID,'qx001','or0001',10000,100,'',10000,'zsy','*','*',100,1,100,'15828680877',1,0.99000,99.00000,0.00100,0.00200,0.00300,0,'2021-01-19 15:24:27','2021-01-20 15:24:27','2021-01-19 16:24:27')"},
		{status: 204, input: map[string]interface{}{"order_id": "@OrderID", "flow_id": "1110"}, sql: "INSERT INTO `ots_trade_order`(`order_status`,`payment_status`,`order_id`,`mer_no`,`mer_order_no`,`mer_product_id`,`mer_shelf_id`,`mer_product_no`,`pl_id`,`brand_no`,`province_no`,`city_no`,`face`,`num`,`total_face`,`account_name`,`invoice_type`,`sell_discount`,`sell_amount`,`mer_fee_discount`,`trade_fee_discount`,`payment_fee_discount`,`can_split_order`,`create_time`,`order_timeout`,`payment_timeout`)VALUES (20,30,@OrderID,'qx001','or0001',10000,100,'',10000,'zsy','*','*',100,1,100,'15828680877',1,0.99000,99.00000,0.00100,0.00200,0.00300,0,'2021-01-19 15:24:27','2021-01-20 15:24:27','2021-01-19 16:24:27')"},
		{status: 204, input: map[string]interface{}{"order_id": "@OrderID", "flow_id": "1110"}, sql: "INSERT INTO `ots_trade_order`(`order_status`,`payment_status`,`order_id`,`mer_no`,`mer_order_no`,`mer_product_id`,`mer_shelf_id`,`mer_product_no`,`pl_id`,`brand_no`,`province_no`,`city_no`,`face`,`num`,`total_face`,`account_name`,`invoice_type`,`sell_discount`,`sell_amount`,`mer_fee_discount`,`trade_fee_discount`,`payment_fee_discount`,`can_split_order`,`create_time`,`order_timeout`,`payment_timeout`)VALUES (20,0,@OrderID,'qx001','or0001',10000,100,'',10000,'zsy','*','*',100,1,100,'15828680877',1,0.99000,99.00000,0.00100,0.00200,0.00300,0,'2021-01-19 15:24:27','2021-01-20 15:24:27','2021-01-19 16:24:27')"},
	}
	var i = 0
	for _, c := range cases {
		i++
		id := time.Now().UnixNano() + int64(i)
		nsql := types.Translate(c.sql, "OrderID", id)
		err := execteSQL(nsql)
		assert.Equal(t, nil, err)

		ctx := mock.NewContext(types.Translate(string(c.input.Marshal()), "OrderID", id))

		//构建请求处理------------------------
		rs := payment.Paying(ctx)

		err = ctx.Response().WriteAny(rs)

		//检查结果
		status, _, _ := ctx.Response().GetFinalResponse()
		assert.Equal(t, nil, err)
		assert.Equal(t, c.status, status)

	}

}
