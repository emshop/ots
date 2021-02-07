package state

import (
	"fmt"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//Notify2Pay 订单通知完成后结束订单
type Notify2Pay struct {
}

//Check 检查状态
func (o *Notify2Pay) Check(order types.IXMap) bool {
	orderStatus := order.GetInt(sql.FieldOrderStatus)
	notifyStatus := order.GetInt(sql.FieldNotifyStatus)

	return orderStatus == int(enums.OrderNotify) &&
		(notifyStatus == int(enums.Success) ||
			notifyStatus == int(enums.ProcessNoNeed))
}

//Update 订单通知完成后结束订单(订单状态:50->0/90）
func (o *Notify2Pay) Update(db db.IDBExecuter, orderID string) error {
	row, err := db.Execute(sql.UpdateOrderForFinishByNotify, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil || row == 0 {
		return fmt.Errorf("订单(%s)通知成功后修改为完成失败“%w", orderID, err)
	}
	return nil
}
