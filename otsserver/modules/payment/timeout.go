package payment

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

//Timeout 处理支付超时
func Timeout(orderID int64) error {
	//启动事务进行支付处理
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}

	//修改支付状态
	row, err := db.Execute(sql.UpdateTradeOrderForPayTimeout, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		db.Rollback()
		return err
	}
	if row > 0 {
		db.Commit()
		return nil
	}

	//查询订单，检查无法处理的原因
	db.Rollback()
	order, err := db.Query(sql.SelectTradeOrderByOrderID, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		return err
	}
	if order.Get(0).GetInt(sql.FieldPaymentStatus) != int(enums.ProcessWaiting) {
		return errs.NewError(http.StatusNoContent, "无须关闭订单")
	}
	return errs.NewError(http.StatusAccepted, "订单暂时不用处理")
}
