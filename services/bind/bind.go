package bind

import (
	"net/http"

	"github.com/emshop/ots/modules/bind"
	"github.com/emshop/ots/modules/const/enums"
	"github.com/emshop/ots/modules/const/sql"
	"github.com/emshop/ots/modules/flow"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask"
)

var fields = []string{
	sql.FieldFlowID,
	sql.FieldOrderID,
}

//Bind 绑定上游
type Bind struct {
}

//OrderHandle 处理订单绑定
func (b *Bind) OrderHandle(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理订单绑定----------------------")
	if err := ctx.Request().Check(fields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 处理订单绑定")
	err := bind.Bind(ctx.Request().GetInt64(sql.FieldOrderID))
	if errs.GetCode(err) == http.StatusNoContent {
		qtask.FinishByInput(ctx, ctx.Request())
		return err
	}

	ctx.Log().Info("2. 执行绑定后续流程")
	status := types.DecodeInt(err, nil, enums.Success, enums.Failed)
	err = flow.Next(ctx.Request().GetInt(sql.FieldFlowID), enums.FlowStatus(status), ctx,
		sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID),
		sql.FieldDeliveryID, ctx.Request().GetString(sql.FieldDeliveryID))
	if err != nil {
		ctx.Log().Error("2.1 执行绑定后续流程失败", err)
		return err
	}
	return "success"
}
