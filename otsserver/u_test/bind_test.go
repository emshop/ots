package ut

import (
	"testing"

	"github.com/emshop/ots/otsserver/services/bind"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra/mock"
	"github.com/micro-plat/lib4go/assert"
	"github.com/micro-plat/lib4go/errs"
)

func TestSingleBind(t *testing.T) {

	preper()
	err := orderCache.makePaidOrder(1)
	assert.Equal(t, nil, err)

	order := <-orderCache.paidOrder

	//构建context
	input := string(order.Marshal())
	ctx := mock.NewContext(input)

	//构建请求处理-----------------

	rs := bind.Binding(ctx)

	assert.Equal(t, nil, errs.GetError(rs), rs)

	err = ctx.Response().WriteAny(getAny(rs))
	status, _, _ := ctx.Response().GetFinalResponse()

	assert.Equal(t, nil, err, input)
	assert.Equal(t, 200, status, input)
}
