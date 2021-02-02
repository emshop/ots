package ut

import (
	"testing"

	"github.com/emshop/ots/otsserver/services/order"
	"github.com/micro-plat/hydra/mock"
	"github.com/micro-plat/lib4go/assert"
	"github.com/micro-plat/lib4go/errs"
)

func TestFinish(t *testing.T) {
	preper()
	err := orderCache.makeNotifyOrder(1)
	assert.Equal(t, nil, err)

	rd := <-orderCache.notifyOrder

	input := string(rd.Marshal())
	ctx := mock.NewContext(input)

	//开始
	ntf := &order.Order{}
	rs := ntf.FinishHandle(ctx)
	assert.Equal(t, nil, errs.GetError(rs))

}
