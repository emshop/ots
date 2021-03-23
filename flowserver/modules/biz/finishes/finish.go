package finishes

import (
	"errors"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"

	"github.com/emshop/ots/mgrserver/api/modules/const/field"
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
	orders, err := db.ExecuteBatch(update2Success, types.XMap{field.FieldOrderID: orderID})

	if err != nil {
		db.Rollback()
	}
	if err != nil && !errors.Is(err, errs.ErrNotExist) {
		return err
	}
	order := orders.Get(0)
	if err == nil {
		//查询订单并完成佣金、交易手续费、支付手续费等记账。
		if order.GetFloat64(fields.FieldSuccessMerFee) > 0 {
			account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchantFee))
			rs, err := account.DeductAmount(db,
				order.GetString(fields.FieldMerNo),
				order.GetString(fields.FieldOrderID),
				beanpay.CommissionTradeType,
				order.GetFloat64(fields.FieldSuccessMerFee),
				"佣金记账")
			if err != nil || rs.GetCode() != beanpay.Success {
				db.Rollback()
				return errs.NewErrorf(int(enums.CodeBalanceLow), "商户(%s)佣金记账失败,%v", order.GetString(fields.FieldMerNo), err)
			}
		}
		if order.GetFloat64(fields.FieldSuccessMerTradeFee) > 0 {
			account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchantTradeFee))
			rs, err := account.DeductAmount(db,
				order.GetString(fields.FieldMerNo),
				order.GetString(fields.FieldOrderID),
				beanpay.FreeTradeType,
				order.GetFloat64(fields.FieldSuccessMerTradeFee),
				"交易手续费记账")
			if err != nil || rs.GetCode() != beanpay.Success {
				db.Rollback()
				return errs.NewErrorf(int(enums.CodeBalanceLow), "商户(%s)交易手续费记账失败,%v", order.GetString(fields.FieldMerNo), err)
			}
		}
		if order.GetFloat64(fields.FieldSuccessMerPaymentFee) > 0 {
			account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchantPaymentFee))
			rs, err := account.DeductAmount(db,
				order.GetString(fields.FieldMerNo),
				order.GetString(fields.FieldOrderID),
				beanpay.FreeTradeType,
				order.GetFloat64(fields.FieldSuccessMerPaymentFee),
				"支付手续费记账")
			if err != nil || rs.GetCode() != beanpay.Success {
				db.Rollback()
				return errs.NewErrorf(int(enums.CodeBalanceLow), "商户(%s)支付手续费记账失败,%v", order.GetString(fields.FieldMerNo), err)
			}
		}
		db.Commit()
		return nil
	}

	//2. 处理未支付的订单-----------------------------------
	db, err = hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}
	_, err = db.ExecuteBatch(update2Faild, types.XMap{field.FieldOrderID: orderID})
	if err == nil {
		db.Commit()
		return nil
	}
	db.Rollback()
	if err != nil && !errors.Is(err, errs.ErrNotExist) {
		return err
	}

	//3. 处理需要退款的订单-----------------------------
	db, err = hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}
	orders, err = db.ExecuteBatch(update2Refund, types.XMap{field.FieldOrderID: orderID})
	if err != nil {
		db.Rollback()
		return err
	}

	//4. 订单退款
	order = orders.Get(0)
	account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchantMain))
	rs, err := account.RefundAmount(db,
		order.GetString(fields.FieldMerNo),
		order.GetString(fields.FieldOrderID),
		order.GetString(fields.FieldOrderID),
		beanpay.AccountTradeType,
		order.GetFloat64(fields.FieldSellAmount),
		"订单退款")
	if err != nil || rs.GetCode() != beanpay.Success {
		db.Rollback()
		return errs.NewErrorf(int(enums.CodeBalanceLow), "商户(%s)退款失败%v", order.GetString(fields.FieldMerNo), err)
	}
	db.Commit()
	return nil
}
