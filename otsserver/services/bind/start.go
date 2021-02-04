package bind

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/bind"
	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask"
)

var fields = []string{
	sql.FieldFlowID,
	sql.FieldOrderID,
}

//Binding 处理订单绑定
func Binding(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理订单绑定----------------------")
	if err := ctx.Request().Check(fields...); err != nil {
		return err
	}

	ctx.Log().Infof("1. 处理订单绑定(%s)", ctx.Request().GetString(sql.FieldOrderID))
	qtask.ProcessingByInput(ctx, ctx.Request())
	deliveryID, err := bind.Bind(ctx.Request().GetString(sql.FieldOrderID))
	switch {
	case errs.GetCode(err) == http.StatusAccepted:
		return err
	case errs.GetCode(err) == http.StatusNoContent:
		qtask.FinishByInput(ctx, ctx.Request()) //当订单绑定完成后自动关闭流程
		return nil
	}

	ctx.Log().Info("2. 执行后续流程")
	status := types.DecodeInt(err, nil, enums.Success, enums.Failed)
	flw := flow.Next(ctx.Request().GetInt(sql.FieldFlowID), enums.FlowStatus(status), ctx,
		sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID),
		sql.FieldDeliveryID, deliveryID)
	if err != nil {
		return err
	}
	return flw
}
