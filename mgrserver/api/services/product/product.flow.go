package product

import (
	"net/http"

	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/emshop/ots/mgrserver/api/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"

	"github.com/emshop/ots/mgrserver/api/modules/db"
)

//ProductFlowHandler 业务流程处理服务
type ProductFlowHandler struct {
}

func NewProductFlowHandler() *ProductFlowHandler {
	return &ProductFlowHandler{}
}

//PostHandle 添加业务流程数据
func (u *ProductFlowHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("--------添加业务流程数据--------")

	ctx.Log().Debug("1.参数校验")
	if err := ctx.Request().CheckMap(postProductFlowCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Debug("2.执行操作")
	xdb := hydra.C.DB().GetRegularDB()
	flowID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "业务流程"})
	if err != nil {
		return err
	}
	input := ctx.Request().GetMap()
	input["flow_id"] = flowID
	count, err := xdb.Execute(sql.InsertProductFlow, input)
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Debug("3.返回结果")
	return "success"
}

//GetHandle 获取业务流程单条数据
func (u *ProductFlowHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("--------获取业务流程单条数据--------")

	ctx.Log().Debug("1.参数校验")
	if err := ctx.Request().CheckMap(getProductFlowCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Debug("2.执行操作")
	items, err := hydra.C.DB().GetRegularDB().Query(sql.GetProductFlowByFlowID, ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Debug("3.返回结果")
	return items.Get(0)
}

//QueryHandle  获取业务流程数据列表
func (u *ProductFlowHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("--------获取业务流程数据列表--------")

	ctx.Log().Debug("1.参数校验")
	if err := ctx.Request().CheckMap(queryProductFlowCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Debug("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetProductFlowListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}

	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetProductFlowList, m)
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

//PutHandle 更新业务流程数据
func (u *ProductFlowHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("--------更新业务流程数据--------")

	ctx.Log().Debug("1.参数校验")
	if err := ctx.Request().CheckMap(updateProductFlowCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Debug("2.执行操作")
	_, err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateProductFlowByFlowID, ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "更新数据出错:%+v", err)
	}

	ctx.Log().Debug("3.返回结果")
	return "success"
}

var postProductFlowCheckFields = map[string]interface{}{
	field.FieldFlowTag: "required",
	field.FieldPlID:    "required",
}

var getProductFlowCheckFields = map[string]interface{}{
	field.FieldFlowID: "required",
}

var queryProductFlowCheckFields = map[string]interface{}{
	field.FieldFlowTag: "required",
	field.FieldPlID:    "required",
	field.FieldStatus:  "required",
}

var updateProductFlowCheckFields = map[string]interface{}{
	field.FieldFlowTag: "required",
	field.FieldPlID:    "required",
}
