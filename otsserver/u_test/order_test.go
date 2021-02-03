package ut

import (
	"testing"

	"github.com/emshop/ots/otsserver/services/order"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra/mock"
	"github.com/micro-plat/lib4go/assert"
)

func TestRequestSuccessHandle(t *testing.T) {
	err := preper()
	assert.Equal(t, nil, err)
	var cases = []struct {
		input  string
		status int
	}{

		{input: "mer_no=qx001&mer_order_no=or0001&mer_product_id=10000&num=1&account_name=15828680877", status: 200},
	}
	for _, c := range cases {

		//构建context
		ctx := mock.NewContext(c.input)

		//构建请求处理------------------------
		rs := order.Request(ctx)
		err := ctx.Response().WriteAny(getAny(rs))

		//校验结果
		assert.Equal(t, nil, err)
		status, _, _ := ctx.Response().GetFinalResponse()
		assert.Equal(t, c.status, status, c.input)
	}
}
func TestRequestHandle(t *testing.T) {

	preper()
	var cases = []struct {
		input  string
		status int
	}{
		{input: "", status: 406},
		{input: "mer_no=&mer_order_no=&mer_product_id=10000&num=1&account_name=", status: 406},
		{input: "mer_no=qx001&mer_order_no=&mer_product_id=10000&num=1&account_name=", status: 406},
		{input: "mer_no=qx001&mer_order_no=mr0002&mer_product_id=&num=1&account_name=", status: 406},
		{input: "mer_no=qx001&mer_order_no=qr0002&mer_product_id=10000&num=1&account_name=", status: 406},
		{input: "mer_no=qx001&mer_order_no=qr0002&mer_product_id=10000&num=1&account_name=", status: 406},
		{input: "mer_no=qx001&mer_order_no=or0001&mer_product_id=10000&num=1&account_name=15828680877", status: 200},
		{input: "mer_no=qx001&mer_order_no=or0002&mer_product_id=10000&num=1&account_name=15828680877", status: 200},
		{input: "mer_no=qx001&mer_order_no=or0003&mer_product_id=10000&num=1&account_name=15828680877", status: 200},
		{input: "mer_no=qx001&mer_order_no=or0004&mer_product_id=10000&num=1&account_name=15828680877", status: 200},
		{input: "mer_no=qx001&mer_order_no=or0013&mer_product_id=10000&num=2&account_name=15828680877", status: 903},
		{input: "mer_no=no002&mer_order_no=or0002&mer_product_id=20000&num=1&account_name=15828680877", status: 901},
		{input: "mer_no=no003&mer_order_no=or0003&mer_product_id=20000&num=1&account_name=15828680877", status: 901},
		{input: "mer_no=qx001&mer_order_no=or0004", status: 406},
		{input: "mer_no=qx001&mer_order_no=", status: 406},
	}
	for _, c := range cases {

		//构建context
		ctx := mock.NewContext(c.input)

		//构建请求处理------------------------
		rs := order.Request(ctx)
		err := ctx.Response().WriteAny(getAny(rs))

		//校验结果
		assert.Equal(t, nil, err)
		status, _, _ := ctx.Response().GetFinalResponse()
		assert.Equal(t, c.status, status, c.input)
	}

}

func TestQueryHandle(t *testing.T) {

	preper()
	var cases = []struct {
		input  string
		status int
	}{
		{input: "", status: 406},
		{input: "mer_no=&mer_order_no=", status: 406},
		{input: "mer_no=qx001&mer_order_no=", status: 406},
		{input: "mer_no=qx001&mer_order_no=or0001", status: 200},
	}
	for _, c := range cases {

		//构建context
		ctx := mock.NewContext(c.input)

		//构建请求处理------------------------
		//验证查询----------------------------
		qs := order.QueryHandle(ctx)
		//写入响应流
		err := ctx.Response().WriteAny(getAny(qs))
		//校验结果
		assert.Equal(t, nil, err)
		status, _, _ := ctx.Response().GetFinalResponse()
		assert.Equal(t, c.status, status, c.input)

		if status != 200 {
			orderCache.order <- ctx.Meta().GetXMap("flow")
		}
	}

}
