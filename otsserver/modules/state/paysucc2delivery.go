package state

import (
	"fmt"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//PaySucc2Delivery 订单支付成功的订单修改为开始发货
type PaySucc2Delivery struct {
}

//Check 检查状态
func (o *PaySucc2Delivery) Check(order types.IXMap) bool {
	orderStatus := order.GetInt(sql.FieldOrderStatus)
	payStatus := order.GetInt(sql.FieldPaymentStatus)
	deliveryStatus := order.GetInt(sql.FieldDeliveryStatus)

	return orderStatus == int(enums.OrderPaying) &&
		payStatus == int(enums.ProcessSuccess) &&
		deliveryStatus == int(enums.ProcessNotStart)
}

//Update 订单支付成功的订单修改为开始发货（订单状态:10->20,发货状态:10->20）
func (o *PaySucc2Delivery) Update(db db.IDBExecuter, orderID string) error {
	row, err := db.Execute(sql.UpdateOrderForWaitDelivery, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil || row == 0 {
		return fmt.Errorf("订单(%s)支付成功后修改为开始发货失败“%w", orderID, err)
	}
	return nil
}
