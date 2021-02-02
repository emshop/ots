package ut

import (
	"testing"

	"github.com/emshop/ots/otsserver/services/notify"
	"github.com/micro-plat/hydra/mock"
	"github.com/micro-plat/lib4go/assert"
	"github.com/micro-plat/lib4go/errs"
)

func TestSingleNotifySuccess(t *testing.T) {
	preper()
	err := orderCache.makeDeliveryOrder(1)
	assert.Equal(t, nil, err)

	order := <-orderCache.deliveriededOrder

	input := string(order.Marshal())
	ctx := mock.NewContext(input)

	//开始通知
	ntf := &notify.Notify{}
	rs := ntf.StartHandle(ctx)
	assert.Equal(t, nil, errs.GetError(rs), rs)

	//保存通知结果
	rs = ntf.SuccessHandle(ctx)
	assert.Equal(t, nil, errs.GetError(rs), rs)

}
