package spp

import (
	"net/http"

	"github.com/emshop/ots/modules/const/enums"
	"github.com/emshop/ots/modules/const/sql"
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/lib4go/errs"
)

//Pay 处理订单支付
func Pay(deliveryID int64) error {

	//启动事务进行支付处理
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}

	//修改支付状态
	row, err := db.Execute(sql.UpdateTradeDeliveryForPaying, map[string]interface{}{
		sql.FieldDeliveryID: deliveryID,
	})
	if err != nil {
		db.Rollback()
		return err
	}
	if row == 0 {
		db.Rollback()
		return errs.NewError(http.StatusNoContent, "订单无须处理")
	}

	//查询发货信息
	orders, err := db.Query(sql.SelectTradeDelivery, map[string]interface{}{
		sql.FieldDeliveryID: deliveryID,
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
	account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountSupplier))

	//处理交易扣款
	_, err = account.DeductAmount(db,
		order.GetString(sql.FieldSppNo),
		order.GetString(sql.FieldDeliveryID),
		beanpay.AccountTradeType,
		order.GetFloat64(sql.FieldCostAmount),
		"供货商扣款")
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}
