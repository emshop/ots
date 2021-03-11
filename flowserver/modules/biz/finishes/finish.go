package finishes

import (
	"errors"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/dbs"
	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Finish 处理订单完结状态（支付成功未发货成功则退款，否则订单处理为成功）
func Finish(orderID string) error {

	//1. 处理正常成功的订单-------------------------------
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}
	_, err = dbs.Executes(db, types.XMap{field.FieldOrderID: orderID}, update2Success...)
	if err == nil {
		db.Commit()
		return nil
	}
	db.Rollback()
	if err != nil && !errors.Is(err, xerr.ErrNOTEXISTS) {
		return err
	}

	//2. 处理未支付的订单-----------------------------------
	db, err = hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}
	_, err = dbs.Executes(db, types.XMap{field.FieldOrderID: orderID}, update2Faild...)
	if err == nil {
		db.Commit()
		return nil
	}
	db.Rollback()
	if err != nil && !errors.Is(err, xerr.ErrNOTEXISTS) {
		return err
	}

	//3. 处理需要退款的订单-----------------------------
	db, err = hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}
	order, err := dbs.Executes(db, types.XMap{field.FieldOrderID: orderID}, update2Refund...)
	if err != nil {
		db.Rollback()
		return err
	}

	//4. 账户扣款
	account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchant))
	rs, err := account.RefundAmount(db,
		order.GetString(fields.FieldMerNo),
		order.GetString(fields.FieldOrderID),
		order.GetString(fields.FieldOrderID),
		beanpay.AccountTradeType,
		order.GetFloat64(fields.FieldSellAmount),
		"订单扣款")
	if err != nil || rs.GetCode() != beanpay.Success {
		db.Rollback()
		return errs.NewErrorf(int(enums.CodeBalanceLow), "商户(%s)退款失败%v", order.GetString(sql.FieldMerNo), err)
	}
	db.Commit()
	return nil
}
