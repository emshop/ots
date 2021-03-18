package merchant

import (
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
)

//Create 创建下游商户的基本账户信息
func Create(ctx hydra.IContext) interface{} {

	ctx.Log().Debug("------------创建账户信息--------------")
	if err := ctx.Request().Check(fields.FieldMerNo); err != nil {
		return err
	}

	merNO := ctx.Request().GetString(fields.FieldMerNo)

	//创建基本户
	baccount := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchantMain))
	if _, err := baccount.CreateAccount(nil, merNO, merNO); err != nil {
		return err
	}

	baccount = beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchantFee))
	if _, err := baccount.CreateAccount(nil, merNO, merNO); err != nil {
		return err
	}
	baccount = beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchantTradeFee))
	if _, err := baccount.CreateAccount(nil, merNO, merNO); err != nil {
		return err
	}
	baccount = beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchantPaymentFee))
	if _, err := baccount.CreateAccount(nil, merNO, merNO); err != nil {
		return err
	}
	return nil

}
