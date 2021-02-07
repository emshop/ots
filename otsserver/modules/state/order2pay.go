package state

import (
	"fmt"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//Order2Pay 将订单状态置为等待支付
type Order2Pay struct {
}

//Check 检查状态
func (o Order2Pay) Check(order types.IXMap) bool {
	orderStatus := order.GetInt(sql.FieldOrderStatus)
	payStatus := order.GetInt(sql.FieldPaymentStatus)
	return orderStatus == int(enums.OrderPaying) && payStatus == int(enums.ProcessNotStart)
}

//Update 修改订单状态为等待支付（订单状态为10不变，支付状态修改为20）
func (o Order2Pay) Update(db db.IDBExecuter, orderID string) error {
	row, err := db.Execute(sql.UpdateOrderForWaitPaying, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil || row == 0 {
		return fmt.Errorf("将订单(%s)状态修改为等待支付失败“%w", orderID, err)
	}
	return nil
}
