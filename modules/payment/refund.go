package payment

import (
	"net/http"

	"github.com/emshop/ots/modules/const/enums"
	"github.com/emshop/ots/modules/const/sql"
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/lib4go/errs"
)

//Refund 订单超时退款
func Refund(orderID int64) error {
	//启动事务进行支付处理
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}

	//修改超时订单为退款中
	row, err := db.Execute(sql.UpdateTradeOrderForRefund, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		db.Rollback()
		return err
	}
	if row == 0 {
		db.Rollback()
		return errs.NewError(http.StatusNoContent, "订单无须处理")
	}

	//查询订单信息
	orders, err := db.Query(sql.SelectTradeOrderByOrderID, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		db.Rollback()
		return err
	}
	if orders.Len() == 0 {
		db.Rollback()
		return errs.NewError(http.StatusNoContent, "订单无须处理")
	}

	//获取账户信息
	order := orders.Get(0)
	account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchant))

	//处理交易退款
	_, err = account.RefundAmount(db,
		order.GetString(sql.FieldMerNo),
		order.GetString(sql.FieldOrderID),
		order.GetString(sql.FieldOrderID),
		beanpay.AccountTradeType,
		order.GetFloat64(sql.FieldSellAmount),
		"订单退款")
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}
