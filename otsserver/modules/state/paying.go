package state

import (
	"fmt"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//Paying 将订单状态修改为支付成功
type Paying struct {
}

//IsTimeout 检查支付是否超时
func (o *Paying) IsTimeout(order types.IXMap) bool {
	return true
}

//Check 检查状态
func (o *Paying) Check(order types.IXMap) bool {
	orderStatus := order.GetInt(sql.FieldOrderStatus)
	payStatus := order.GetInt(sql.FieldPaymentStatus)
	return orderStatus == int(enums.OrderPaying) && payStatus == int(enums.ProcessWaiting)
}

//Update2Success 修改订单状态为等待支付（支付状态修改为0）
func (o *Paying) Update2Success(db db.IDBExecuter, orderID string) error {
	row, err := db.Execute(sql.UpdateOrderForPaySuccess, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil || row == 0 {
		return fmt.Errorf("将订单(%s)支付状态修改为支付成功时失败“%w", orderID, err)
	}
	return nil
}

//Update2Failed 修改订单状态为支付失败（支付状态修改为0）
func (o *Paying) Update2Failed(db db.IDBExecuter, orderID string) error {
	row, err := db.Execute(sql.UpdateOrderForPayFailed, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil || row == 0 {
		return fmt.Errorf("将订单(%s)支付状态修改为支付失败时失败“%w", orderID, err)
	}
	return nil
}
