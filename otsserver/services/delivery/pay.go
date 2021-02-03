package delivery

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/emshop/ots/otsserver/modules/spp"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask"
)

//Paying 上游付款
func Paying(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------开始供货商订单付款----------------------")
	if err := ctx.Request().Check(deliveryStartNowFields...); err != nil {
		return err
	}

	ctx.Log().Infof("1. 处理供货商订单付款(%s)", ctx.Request().GetString(sql.FieldDeliveryID))
	qtask.ProcessingByInput(ctx, ctx.Request())
	err := spp.Pay(ctx.Request().GetString(sql.FieldDeliveryID))
	switch {
	case errs.GetCode(err) == http.StatusNoContent || err == nil:
		qtask.FinishByInput(ctx, ctx.Request()) //当无须处理或成功后关闭流程
	}

	ctx.Log().Info("2. 执行后续流程")
	status := types.DecodeInt(err, nil, enums.Success, enums.Failed)
	flw := flow.Next(ctx.Request().GetInt(sql.FieldFlowID), enums.FlowStatus(status), ctx,
		sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID),
		sql.FieldDeliveryID, ctx.Request().GetString(sql.FieldDeliveryID))
	if err != nil {
		return err
	}
	return flw
}
