package delivery

import (
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/spp"
	"github.com/micro-plat/hydra"
)

//PayHandle 上游付款
func PayHandle(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------开始发货付款----------------------")
	if err := ctx.Request().Check(deliveryStartNowFields...); err != nil {
		return err
	}
	err := spp.Pay(ctx.Request().GetString(sql.FieldDeliveryID))
	return err
}
