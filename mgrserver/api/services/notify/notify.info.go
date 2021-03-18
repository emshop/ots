package notify

import (
	"net/http"

	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/emshop/ots/mgrserver/api/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//NotifyInfoHandler 订单通知表处理服务
type NotifyInfoHandler struct {
}

func NewNotifyInfoHandler() *NotifyInfoHandler {
	return &NotifyInfoHandler{}
}

//GetHandle 获取订单通知表单条数据
func (u *NotifyInfoHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("--------获取订单通知表单条数据--------")

	ctx.Log().Debug("1.参数校验")
	if err := ctx.Request().CheckMap(getNotifyInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Debug("2.执行操作")
	items, err := hydra.C.DB().GetRegularDB().Query(sql.GetNotifyInfoByOrderID, ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Debug("3.返回结果")
	return items.Get(0)
}

//DetailHandle 获取订单通知表详情单条数据
func (u *NotifyInfoHandler) DetailHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("--------获取订单通知表详情单条数据--------")

	ctx.Log().Debug("1.参数校验")
	if err := ctx.Request().CheckMap(getNotifyInfoDetailCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Debug("2.执行操作")
	items, err := hydra.C.DB().GetRegularDB().Query(sql.GetNotifyInfoDetailByOrderID, ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Debug("3.返回结果")
	return items.Get(0)
}

//QueryHandle  获取订单通知表数据列表
func (u *NotifyInfoHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("--------获取订单通知表数据列表--------")

	ctx.Log().Debug("1.参数校验")
	if err := ctx.Request().CheckMap(queryNotifyInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Debug("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetNotifyInfoListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}

	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetNotifyInfoList, m)
		if err != nil {
			return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
		}
	}

	ctx.Log().Debug("3.返回结果")
	return map[string]interface{}{
		"items": items,
		"count": types.GetInt(count),
	}
}

var getNotifyInfoCheckFields = map[string]interface{}{
	field.FieldOrderID: "required",
}

var getNotifyInfoDetailCheckFields = map[string]interface{}{
	field.FieldOrderID: "required",
}

var queryNotifyInfoCheckFields = map[string]interface{}{
	field.FieldMerNo:        "required",
	field.FieldCreateTime:   "required",
	field.FieldNotifyStatus: "required",
}
