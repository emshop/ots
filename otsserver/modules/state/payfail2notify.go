package state

import (
	"fmt"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//PayFail2Notify 支付失败改为等待通知
type PayFail2Notify struct {
}

//Check 检查状态
func (o PayFail2Notify) Check(order types.IXMap) bool {
	orderStatus := order.GetInt(sql.FieldOrderStatus)
	payStatus := order.GetInt(sql.FieldPaymentStatus)
	deliveryStatus := order.GetInt(sql.FieldDeliveryStatus)
	notifyStatus := order.GetInt(sql.FieldNotifyStatus)

	return orderStatus == int(enums.OrderPaying) &&
		payStatus == int(enums.ProcessFailed) &&
		deliveryStatus == int(enums.ProcessNotStart) &&
		(notifyStatus == int(enums.ProcessNoNeed) ||
			notifyStatus == int(enums.ProcessNotStart))
}

//Update 支付失败改为等待通知(订单状态10->50,发货状态10->11,通知状态:11->11,10->20 )
func (o PayFail2Notify) Update(db db.IDBExecuter, orderID string) error {
	row, err := db.Execute(sql.UpdateOrderForNotifyByPayFailed, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil || row == 0 {
		return fmt.Errorf("订单(%s)支付失败后修改为等待通知失败“%w", orderID, err)
	}
	return nil
}
