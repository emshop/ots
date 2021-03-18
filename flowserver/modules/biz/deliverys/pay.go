package deliverys

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/dbs"
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Pay 处理订单支付
func Pay(deliveryID string) error {
	input := types.XMap{
		fields.FieldDeliveryID: deliveryID,
	}

	//启动事务进行支付处理
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}

	//修改支付状态
	order, err := dbs.Executes(db, input, updateDeliveryForPaying...)
	if errors.Is(err, xerr.ErrNOTEXISTS) {
		db.Rollback()
		return errs.NewErrorf(http.StatusNoContent, "发货编号(%s)无须进行发货支付", deliveryID)
	}
	if err != nil {
		db.Rollback()
		return fmt.Errorf("发货编号(%s)发货支付出错%w", deliveryID, err)
	}

	//------------------------处理记账信息----------------------------------------
	//交易扣款
	account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountSupplierMain))

	rs, err := account.DeductAmount(db,
		order.GetString(fields.FieldSppNo),
		order.GetString(fields.FieldDeliveryID),
		beanpay.AccountTradeType,
		order.GetFloat64(fields.FieldCostAmount),
		"供货商扣款")
	if err != nil {
		db.Rollback()
		return fmt.Errorf("发货支付失败(%s)%v", order.GetString(fields.FieldSppNo), err)
	}
	if rs.GetCode() != beanpay.Success {
		db.Rollback()
		return errs.NewErrorf(rs.GetCode(), "%s扣款失败%s", deliveryID, rs.GetCode())
	}
	//查询发货记录并完成佣金、交易手续费、支付手续费等记账
	if order.GetFloat64(fields.FieldSppFeeAmount) > 0 {
		account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountSupplierFee))
		rs, err := account.DeductAmount(db,
			order.GetString(fields.FieldSppNo),
			order.GetString(fields.FieldDeliveryID),
			beanpay.CommissionTradeType,
			order.GetFloat64(fields.FieldSppFeeAmount),
			"佣金记账")
		if err != nil || rs.GetCode() != beanpay.Success {
			db.Rollback()
			return errs.NewErrorf(int(enums.CodeBalanceLow), "供货商(%s)佣金记账失败,%v", order.GetString(fields.FieldSppNo), err)
		}
	}
	if order.GetFloat64(fields.FieldTradeFeeAmount) > 0 {
		account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountSupplierTradeFee))
		rs, err := account.DeductAmount(db,
			order.GetString(fields.FieldSppNo),
			order.GetString(fields.FieldDeliveryID),
			beanpay.FreeTradeType,
			order.GetFloat64(fields.FieldTradeFeeAmount),
			"交易手续费记账")
		if err != nil || rs.GetCode() != beanpay.Success {
			db.Rollback()
			return errs.NewErrorf(int(enums.CodeBalanceLow), "供货商(%s)交易手续费记账失败,%v", order.GetString(fields.FieldSppNo), err)
		}
	}
	if order.GetFloat64(fields.FieldPaymentFeeAmount) > 0 {
		account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountSupplierPaymentFee))
		rs, err := account.DeductAmount(db,
			order.GetString(fields.FieldSppNo),
			order.GetString(fields.FieldDeliveryID),
			beanpay.FreeTradeType,
			order.GetFloat64(fields.FieldPaymentFeeAmount),
			"支付手续费记账")
		if err != nil || rs.GetCode() != beanpay.Success {
			db.Rollback()
			return errs.NewErrorf(int(enums.CodeBalanceLow), "供货商(%s)支付手续费记账失败,%v", order.GetString(fields.FieldSppNo), err)
		}
	}

	db.Commit()
	return nil
}
