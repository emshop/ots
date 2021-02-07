package payment

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/orders"
	"github.com/emshop/ots/otsserver/modules/state"
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/lib4go/errs"
)

//Pay 处理订单支付
func Pay(orderID string) error {
	//1. 查询订单信息
	order, err := orders.QueryByOrderID(orderID)
	if err != nil {
		return err
	}

	//2. 检查订单状态是否允许支付
	var paying = state.Paying{}

	if !paying.Check(order) {
		return errs.NewStop(http.StatusNoContent, "订单支付状态错误，无法支付")
	}

	//3. 启动事务进行支付处理
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}

	//4. 检查是否支付超时
	if paying.IsTimeout(order) {
		if err := paying.Update2Failed(db, orderID); err != nil {
			db.Rollback()
			return err
		}
		notify := state.PayFail2Notify{}
		if err := notify.Update(db, orderID); err != nil {
			db.Rollback()
			return err
		}
		db.Commit()
		return errs.NewStop(http.StatusNoContent, "订单已支付超时")
	}

	//5. 更新支付状态为支付成功
	err = paying.Update2Success(db, orderID)
	if err != nil {
		db.Rollback()
		return err
	}

	//6. 处理交易扣款
	account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchant))
	rs, err := account.DeductAmount(db,
		order.GetString(sql.FieldMerNo),
		order.GetString(sql.FieldOrderID),
		beanpay.AccountTradeType,
		order.GetFloat64(sql.FieldSellAmount),
		"订单扣款")
	if err != nil || rs.GetCode() != beanpay.Success {
		db.Rollback()
		return errs.NewErrorf(int(enums.CodeBalanceLow), "商户(%s)扣款失败%v", order.GetString(sql.FieldMerNo), err)
	}

	//7. 订单支付成功后修改为等待发货
	var delivery = state.PaySucc2Delivery{}
	if err := delivery.Update(db, orderID); err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}
