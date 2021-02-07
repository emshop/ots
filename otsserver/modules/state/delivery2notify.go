package state

import (
	"fmt"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//delivery2Notify 订单发货成功后修改为等待通知
type delivery2Notify struct {
}

//Check 检查状态
func (o *delivery2Notify) Check(order types.IXMap) bool {
	orderStatus := order.GetInt(sql.FieldOrderStatus)
	deliveryStatus := order.GetInt(sql.FieldDeliveryStatus)
	return orderStatus == int(enums.OrderDelivering) && deliveryStatus == int(enums.ProcessSuccess)
}

//Update 订单发货成功后修改为等待通知（订单状态:20->50,通知状态:10->20,11->11）
func (o *delivery2Notify) Update(db db.IDBExecuter, orderID string) error {
	row, err := db.Execute(sql.UpdateOrderForNotifyByDeliverySuccess, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil || row == 0 {
		return fmt.Errorf("订单(%s)发货成功后修改为等待通知失败“%w", orderID, err)
	}
	return nil
}
