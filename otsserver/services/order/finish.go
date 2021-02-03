package order

import (
	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/emshop/ots/otsserver/modules/orders"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

var finishFields = []string{
	sql.FieldFlowID,
	sql.FieldOrderID,
}

//FinishHandle 订单完结处理
func FinishHandle(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------处理订单完结----------------------")
	if err := ctx.Request().Check(finishFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 处理订单完结")
	err := orders.Finish(ctx.Request().GetString(sql.FieldOrderID))
	if err != nil {
		return err
	}

	ctx.Log().Info("2. 处理订单后续流程")
	status := types.DecodeInt(err, nil, enums.Success, enums.Failed)
	flw, err := flow.Next(ctx.Request().GetInt(sql.FieldFlowID), enums.FlowStatus(status), ctx,
		sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID))
	if err != nil {
		ctx.Log().Error("2.1 执行绑定后续流程失败", err)
	}
	return flw
}
