package spp

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
	if err := ctx.Request().Check(fields.FieldSppNo); err != nil {
		return err
	}

	sppNO := ctx.Request().GetString(fields.FieldSppNo)

	//创建基本户
	baccount := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountSupplierMain))
	if _, err := baccount.CreateAccount(nil, sppNO, sppNO); err != nil {
		return err
	}

	baccount = beanpay.GetAccount(global.Def.PlatName, string(enums.AccountSupplierFee))
	if _, err := baccount.CreateAccount(nil, sppNO, sppNO); err != nil {
		return err
	}
	baccount = beanpay.GetAccount(global.Def.PlatName, string(enums.AccountSupplierTradeFee))
	if _, err := baccount.CreateAccount(nil, sppNO, sppNO); err != nil {
		return err
	}
	baccount = beanpay.GetAccount(global.Def.PlatName, string(enums.AccountSupplierPaymentFee))
	if _, err := baccount.CreateAccount(nil, sppNO, sppNO); err != nil {
		return err
	}
	return nil

}
