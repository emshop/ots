package order

import (
	"github.com/emshop/ots/modules/const/db/biz"
	"github.com/emshop/ots/modules/const/enums"
	"github.com/emshop/ots/modules/mers"
	"github.com/emshop/ots/modules/orders"
	"github.com/emshop/ots/modules/pkg"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

var orderRequestFields = []string{
	biz.FieldTradeOrderMerNo,
	biz.FieldTradeOrderMerOrderNo,
	biz.FieldTradeOrderMerProductID,
	biz.FieldTradeOrderNum,
	biz.FieldNameAccountName,
}

//Order 订单处理
type Order struct {
}

//QueryHandle 订单查询
func (o *Order) QueryHandle(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理订单查询----------------------")
	if err := ctx.Request().Check(biz.FieldTradeOrderMerNo, biz.FieldTradeOrderMerOrderNo); err != nil {
		return err
	}

	ctx.Log().Info("1. 查询订单信息")
	order, err := orders.Query(ctx.Request().GetString(biz.FieldTradeOrderMerNo),
		ctx.Request().GetString(biz.FieldTradeOrderMerOrderNo))
	if err == nil && order.Len() > 0 {
		return order
	}
	ctx.Log().Info("2. 订单不存在")
	return errs.NewError(int(enums.OrderNotExists), "订单不存在")

}

//RequestHandle 下单处理
func (o *Order) RequestHandle(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------处理订单收单----------------------")
	if err := ctx.Request().Check(orderRequestFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 查询订单信息")
	order, err := orders.Query(ctx.Request().GetString(biz.FieldTradeOrderMerNo),
		ctx.Request().GetString(biz.FieldTradeOrderMerOrderNo))
	if err == nil && order.Len() > 0 {
		return order
	}

	ctx.Log().Info("2. 查询产品配置")
	product, err := mers.GetProductByID(ctx.Request().GetInt(biz.FieldTradeOrderMerProductID))
	if err != nil || product.Len() == 0 {
		return errs.NewError(int(enums.ParamNoSetting), "商户产品未配置")
	}

	ctx.Log().Info("3. 查询产品对应的货架配置")
	shelf, err := mers.GetShelfByID(product.GetInt(biz.FieldTradeOrderMerShelfID))
	if err != nil || shelf.Len() == 0 {
		return errs.NewError(int(enums.ParamNoSetting), "商户产品未配置")
	}

	ctx.Log().Info("4. 创建订单信息")
	order, ok, err := orders.Create(
		ctx.Request().GetString(biz.FieldTradeOrderMerNo),
		ctx.Request().GetString(biz.FieldTradeOrderMerOrderNo),
		shelf, product, ctx.Request().GetInt(biz.FieldTradeOrderNum))
	if err != nil {
		return errs.NewErrorf(int(enums.UnknownErr), "订单创建失败:%w", err)
	}
	if !ok {
		return order
	}

	ctx.Log().Info("5. 处理订单后续流程")
	err = pkg.NextFlow(product.GetInt(biz.FieldTradeOrderPlID), enums.Success,
		ctx, biz.FieldTradeOrderOrderID,
		order.GetString(biz.FieldTradeOrderOrderID))
	if err != nil {
		ctx.Log().Error(err)
		return order
	}
	return order
}
