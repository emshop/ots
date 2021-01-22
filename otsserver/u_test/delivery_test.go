package ut

import (
	"testing"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/delivery"
	"github.com/micro-plat/lib4go/assert"
	"github.com/micro-plat/lib4go/types"
)

func TestSingleDelivery(t *testing.T) {
	preper()
	err := orderCache.makeBindedOrder(1)
	assert.Equal(t, nil, err)

	order := <-orderCache.bindedOrder

	//获取发货结果
	_, err = delivery.StartNow(order.GetInt64(sql.FieldDeliveryID))
	assert.Equal(t, nil, err, err)

	//保存发货结果
	err = delivery.SaveStart(order.GetInt64(sql.FieldDeliveryID), types.GetDecimal(0), "000", "上游请求成功")
	assert.Equal(t, nil, err, err)

	//保存充值结果
	status, err := delivery.SaveDeliveryResult(order.GetInt64(sql.FieldDeliveryID), "0000", "", types.GetDecimal(0), "")

	assert.Equal(t, nil, err, err)
	assert.Equal(t, enums.Success, status)

}
