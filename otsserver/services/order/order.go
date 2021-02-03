package order

import (
	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/emshop/ots/otsserver/modules/mers"
	"github.com/emshop/ots/otsserver/modules/orders"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

var orderRequestFields = []string{
	sql.FieldMerNo,
	sql.FieldMerOrderNo,
	sql.FieldMerProductID,
	sql.FieldNum,
	sql.FieldAccountName,
}

//RequestHandle 下单处理
func RequestHandle(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------处理订单收单----------------------")
	if err := ctx.Request().Check(orderRequestFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 查询订单信息")
	order, err := orders.Query(ctx.Request().GetString(sql.FieldMerNo),
		ctx.Request().GetString(sql.FieldMerOrderNo))
	if err == nil && order.Len() > 0 {
		return order
	}

	ctx.Log().Info("2. 查询产品配置")
	product, err := mers.GetProductByID(ctx.Request().GetString(sql.FieldMerNo),
		ctx.Request().GetInt(sql.FieldMerProductID))
	if err != nil || product.Len() == 0 {
		return errs.NewError(int(enums.CodeParamNoSetting), "商户产品未配置")
	}

	ctx.Log().Info("3. 查询产品对应的货架配置")
	shelf, err := mers.GetShelfByID(
		ctx.Request().GetString(sql.FieldMerNo),
		product.GetInt(sql.FieldMerShelfID),
		ctx.Request().GetInt(sql.FieldNum))
	if err != nil {
		return err
	}

	ctx.Log().Info("4. 创建订单信息")
	order, ok, err := orders.Create(
		ctx.Request().GetString(sql.FieldMerNo),
		ctx.Request().GetString(sql.FieldMerOrderNo),
		ctx.Request().GetString(sql.FieldAccountName),
		ctx.Request().GetInt(sql.FieldNum),
		ctx.Request().GetString(sql.FieldNotifyUrl),
		shelf, product)
	if err != nil {
		return errs.NewErrorf(int(enums.CodeUnknownErr), "订单创建失败:%v", err)
	}
	if !ok {
		return order
	}

	ctx.Log().Info("5. 处理订单后续流程")
	flw, err := flow.NextByFirst(enums.FlowFlagRecvOrderByPID, product.GetInt(sql.FieldPlID), enums.Success,
		ctx, sql.FieldOrderID,
		order.GetString(sql.FieldOrderID))
	if err != nil {
		ctx.Log().Error("5.1", err)
	}
	ctx.Meta().SetValue("flow", flw)
	return order
}
