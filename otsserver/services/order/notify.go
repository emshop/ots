package order

import (
	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/flow"
	"github.com/emshop/ots/otsserver/modules/notify"
	"github.com/micro-plat/hydra"
)

var fields = []string{
	sql.FieldFlowID,
	sql.FieldOrderID,
}

//NotifyHandle 订单通知
func (o *Order) NotifyHandle(ctx hydra.IContext) interface{} {
	ctx.Log().Info("-------------处理订单查询----------------------")
	if err := ctx.Request().Check(fields...); err != nil {
		return err
	}

	ctx.Log().Info("1. 开始订单通知")
	ntf, err := notify.Start(ctx.Request().GetString(sql.FieldOrderID))
	if err != nil {
		return err
	}

	//处理通知
	ctx.Log().Info("2. 下游订单通知", ntf)

	ctx.Log().Info("3. 保存通知结果")
	err = notify.Save(ctx.Request().GetString(sql.FieldOrderID), enums.Success, "sucess")
	if err != nil {
		return err
	}

	ctx.Log().Info("4. 获取通知后续流程")
	flw, err := flow.Next(ctx.Request().GetInt(sql.FieldFlowID),
		enums.Success, ctx,
		sql.FieldOrderID,
		ctx.Request().GetString(sql.FieldOrderID))
	if err != nil {
		ctx.Log().Error("执行支付后续流程失败", err)
		return err
	}
	return flw
}
