package delivery

import (
	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/delivery"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

var deliverySaveResultFields = []string{
	sql.FieldDeliveryID,
	sql.FieldResultCode,
	sql.FieldReturnMsg,
}

//SaveResultHandle 开始请求
func SaveResultHandle(ctx hydra.IContext) interface{} {

	ctx.Log().Info("-------------保存发货结果----------------------")
	if err := ctx.Request().Check(deliverySaveResultFields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 保存发货结果")
	_, err := delivery.SaveDeliveryResult(
		ctx.Request().GetString(sql.FieldDeliveryID),
		ctx.Request().GetString(sql.FieldResultCode),
		ctx.Request().GetString(sql.FieldReturnMsg),
		ctx.Request().GetDecimal(sql.FieldRealDiscount),
		ctx.Request().GetString(sql.FieldRequestParams),
	)

	ctx.Log().Info("2. 执行后续流程")
	delivery, err1 := delivery.Get(ctx.Request().GetString(sql.FieldDeliveryID))
	if err1 != nil {
		ctx.Log().Error("2.1 执行后续流程获取发货记录失败", err1)
		return nil
	}
	status := types.DecodeInt(err, nil, enums.Success, enums.Failed)
	flw, err1 := flow.NextByFirst(enums.FlowFlagDeliveryFinish, delivery.GetInt(sql.FieldPlID), enums.FlowStatus(status), ctx,
		sql.FieldOrderID, ctx.Request().GetString(sql.FieldOrderID),
		sql.FieldDeliveryID, ctx.Request().GetString(sql.FieldDeliveryID))
	if err1 != nil {
		ctx.Log().Error("2.2 执行绑定后续流程失败", err1)
	}
	if err != nil {
		return err
	}
	return flw
}
