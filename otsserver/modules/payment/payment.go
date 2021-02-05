package payment

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Pay 处理订单支付
func Pay(orderID string) error {

	//检查订单是否超时
	r, err := hydra.C.DB().GetRegularDB().Scalar(sql.SelectTradeOrderByNoneedPay, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		return err
	}
	if types.GetInt(r) > 0 {
		return errs.NewError(http.StatusNoContent, "订单已超时，无须支付")
	}

	//启动事务进行支付处理
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}

	//修改支付状态
	row, err := db.Execute(sql.UpdateTradeOrderForPaying, map[string]interface{}{
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

	//处理交易扣款
	rs, err := account.DeductAmount(db,
		order.GetString(sql.FieldMerNo),
		order.GetString(sql.FieldOrderID),
		beanpay.AccountTradeType,
		order.GetFloat64(sql.FieldSellAmount),
		"订单扣款")
	if err != nil {
		db.Rollback()
		return errs.NewErrorf(int(enums.CodeBalanceLow), "商户(%s)扣款失败%v", order.GetString(sql.FieldMerNo), err)
	}
	if rs.GetCode() != beanpay.Success {
		db.Rollback()
		return errs.NewErrorf(rs.GetCode(), "商户(%s)扣款失败%d", order.GetString(sql.FieldMerNo), rs.GetCode())
	}
	db.Commit()
	return nil
}
