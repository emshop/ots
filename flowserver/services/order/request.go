package order

import (
	"github.com/emshop/ots/flowserver/modules/biz/orders"
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/flows"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

//Request 下单处理
func Request(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("-------------处理订单收单----------------------")
	if err := ctx.Request().Check(orderRequestFields...); err != nil {
		return err
	}

	ctx.Log().Debug("1. 查询订单信息")

	order, err := orders.Query(ctx.Request().GetString(fields.FieldMerNo),
		ctx.Request().GetString(fields.FieldMerOrderNo))
	if err == nil && order.Len() > 0 {
		return order
	}
	ctx.Log().Debug("2. 创建订单信息")
	order, ok, err := orders.Create(
		ctx.Request().GetString(fields.FieldMerNo),
		ctx.Request().GetString(fields.FieldMerOrderNo),
		ctx.Request().GetInt(fields.FieldMerProductID),
		ctx.Request().GetString(fields.FieldAccountName),
		ctx.Request().GetInt(fields.FieldNum),
		ctx.Request().GetString(fields.FieldNotifyUrl))
	if err != nil && errs.GetCode(err) != 0 {
		return err
	}
	if err != nil {
		return errs.NewErrorf(int(enums.CodeUnknownErr), "订单创建失败:%v", err)
	}
	if !ok {
		return order
	}
	defer lcs.New(ctx.Log(), "创建订单", order.GetString(fields.FieldOrderID)).Start("生成订单").End(&r)
	ctx.Log().Debug("3. 处理订单后续流程")
	flows.NextByOrderNO(order.GetString(fields.FieldOrderID), enums.FlowPaymentStart, ctx)
	return order
}

var orderRequestFields = []string{
	fields.FieldMerNo,
	fields.FieldMerOrderNo,
	fields.FieldMerProductID,
	fields.FieldNum,
	fields.FieldAccountName,
}
