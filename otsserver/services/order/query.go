package order

import (
	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/orders"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

//QueryHandle 订单查询
func QueryHandle(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理订单查询----------------------")
	if err := ctx.Request().Check(sql.FieldMerNo, sql.FieldMerOrderNo); err != nil {
		return err
	}

	ctx.Log().Info("1. 查询订单信息")
	order, err := orders.Query(ctx.Request().GetString(sql.FieldMerNo),
		ctx.Request().GetString(sql.FieldMerOrderNo))
	if err == nil && order.Len() > 0 {
		return order
	}

	ctx.Log().Info("2. 订单不存在")
	return errs.NewError(int(enums.CodeOrderNotExists), "订单不存在")

}
