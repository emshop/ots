package notify

import (
	"net/http"

	"github.com/emshop/ots/modules/const/enums"
	"github.com/emshop/ots/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

//Save 保存通知结果
func Save(orderID int64, status enums.FlowStatus, msg string) error {
	//启动事务进行支付处理
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}
	var row int64
	switch status {
	case enums.Success:

		//修改通知记录
		row, err = db.Execute(sql.UpdateNotifyInfoForSuccess, map[string]interface{}{
			sql.FieldOrderID:   orderID,
			sql.FieldNotifyMsg: msg,
		})
		if err != nil || row == 0 {
			db.Rollback()
			return err
		}

		//修改订单状态
		row, err = db.Execute(sql.UpdateTradeOrderForFinish, map[string]interface{}{
			sql.FieldOrderID: orderID,
		})

	default:
		//保存通知消息
		row, err = db.Execute(sql.UpdateNotifyInfoForUnknown, map[string]interface{}{
			sql.FieldOrderID:   orderID,
			sql.FieldNotifyMsg: msg,
		})
	}

	if err != nil {
		db.Rollback()
		return err
	}
	if row == 0 {
		db.Rollback()
		return errs.NewError(http.StatusNoContent, "订单无须处理")
	}
	db.Commit()
	return nil
}
