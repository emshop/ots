package payments

import (
	"errors"
	"net/http"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"

	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Pay 处理订单支付
func Pay(orderID string) error {

	//---------处理订单超时-------------------------------
	ok, err := timeout(orderID)
	if err != nil {
		return err
	}
	if ok {
		return errs.NewErrorf(http.StatusNoContent, "订单(%s)已超时%w", orderID, xerr.ErrOrderTimeout)
	}

	//---------正常的订单扣款支付----------------------------
	input := types.NewXMap()
	input.SetValue(fields.FieldOrderID, orderID)

	//构建事务
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}

	//执行数据库逻辑处理
	orders, err := db.ExecuteBatch(dealOrderPayStart, input)
	if errors.Is(err, errs.ErrNotExist) {
		db.Rollback()
		return errs.NewErrorf(200, "订单(%s)无须处理%w", orderID, errs.ErrNotExist)
	}
	if err != nil {
		db.Rollback()
		return err
	}

	//账户扣款
	order := orders.Get(0)
	account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchantMain))
	rs, err := account.DeductAmount(db,
		order.GetString(fields.FieldMerNo),
		order.GetString(fields.FieldOrderID),
		beanpay.AccountTradeType,
		order.GetFloat64(fields.FieldSellAmount),
		"订单扣款")
	if err != nil || rs.GetCode() != beanpay.Success {
		db.Rollback()
		return errs.NewErrorf(int(enums.CodeBalanceLow), "商户(%s)扣款失败%v", order.GetString(fields.FieldMerNo), err)
	}
	db.Commit()
	return nil

}
