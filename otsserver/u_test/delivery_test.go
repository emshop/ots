package ut

import (
	"testing"

	"github.com/emshop/ots/otsserver/services/delivery"
	"github.com/micro-plat/hydra/mock"
	"github.com/micro-plat/lib4go/assert"
	"github.com/micro-plat/lib4go/errs"
)

func TestSingleDelivery(t *testing.T) {
	preper()
	err := orderCache.makeBindedOrder(1)
	assert.Equal(t, nil, err)

	//开始发货--------------------
	order := <-orderCache.bindedOrder
	input := string(order.Marshal())
	ctx := mock.NewContext(input)

	rs := delivery.StartRequest(ctx)

	assert.Equal(t, nil, errs.GetError(rs), rs)

	//保存请求结果---------------------------
	ctx.Request().SetValue("result_code", "000")
	ctx.Request().SetValue("return_msg", "上游请求成功")
	rs = delivery.SaveRequest(ctx)
	assert.Equal(t, nil, errs.GetError(rs), rs)

	//保存发货结果---------------------------
	ctx.Request().SetValue("result_code", "000")
	ctx.Request().SetValue("return_msg", "发货成功000")
	rs = delivery.SaveResult(ctx)
	assert.Equal(t, nil, errs.GetError(rs), rs)

}

func TestSingleDeliveryFailed(t *testing.T) {
	preper()
	err := orderCache.makeBindedOrder(1)
	assert.Equal(t, nil, err)

	//开始发货--------------------
	order := <-orderCache.bindedOrder
	input := string(order.Marshal())
	ctx := mock.NewContext(input)

	rs := delivery.StartRequest(ctx)

	assert.Equal(t, nil, errs.GetError(rs), rs)

	//保存请求结果---------------------------
	ctx.Request().SetValue("result_code", "000")
	ctx.Request().SetValue("return_msg", "上游请求成功")
	rs = delivery.SaveRequest(ctx)
	assert.Equal(t, nil, errs.GetError(rs), rs)

	//保存发货结果---------------------------
	ctx.Request().SetValue("result_code", "1111")
	ctx.Request().SetValue("return_msg", "发货失败")
	rs = delivery.SaveResult(ctx)
	assert.Equal(t, nil, errs.GetError(rs), rs)

}
