package bind

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/bind"
	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/qtask"
)

var fields = []string{
	sql.FieldFlowID,
	sql.FieldOrderID,
}

//Bind 绑定上游

//StartHandle 处理订单绑定
func StartHandle(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理订单绑定----------------------")
	if err := ctx.Request().Check(fields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 处理订单绑定")
	qtask.ProcessingByInput(ctx, ctx.Request())
	deliveryID, err := bind.Bind(ctx.Request().GetString(sql.FieldOrderID))
	switch {
	case errs.GetCode(err) == http.StatusNoContent:
		qtask.FinishByInput(ctx, ctx.Request())
	}
	if err != nil {
		return nil
	}

	ctx.Log().Info("2. 执行绑定后续流程")
	flw, err := flow.Next(ctx.Request().GetInt(sql.FieldFlowID), enums.FlowStatus(enums.Success), ctx,
		sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID),
		sql.FieldDeliveryID, deliveryID)
	if err != nil {
		ctx.Log().Error("2.1 执行绑定后续流程失败", err)
	}
	return flw
}
