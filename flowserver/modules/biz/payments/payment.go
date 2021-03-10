package payments

import (
	"errors"
	"fmt"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/dbs"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Pay 处理订单支付
func Pay(orderID string) error {
	input := types.NewXMap()
	input.SetValue(fields.FieldOrderID, orderID)

	//构建事务
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}

	//执行数据库逻辑处理
	order, err := dbs.Executes(db, input, dealOrderPayStart...)
	if errors.Is(err, xerr.ErrNOTEXISTS) {
		db.Rollback()
		return errs.NewStop(200, fmt.Errorf("订单(%s)无须处理(err:%w)", orderID, err))
	}
	if err != nil {
		db.Rollback()
		return err
	}

	//账户扣款
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
	db.Commit()
	return nil

}
