package order

import (
	"github.com/emshop/ots/flowserver/modules/biz/orders"
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

//Query 订单查询
func Query(ctx hydra.IContext) interface{} {
	ctx.Log().Debug("-------------处理订单查询----------------------")
	if err := ctx.Request().Check(fields.FieldMerNo, fields.FieldMerOrderNo); err != nil {
		return err
	}

	ctx.Log().Debug("1. 查询订单信息")
	order, err := orders.QueryDetail(ctx.Request().GetString(fields.FieldMerNo),
		ctx.Request().GetString(fields.FieldMerOrderNo))
	if err == nil && order.Len() > 0 {
		return order
	}

	ctx.Log().Debug("2. 订单不存在")
	return errs.NewError(int(enums.CodeOrderNotExists), "订单不存在")

}
