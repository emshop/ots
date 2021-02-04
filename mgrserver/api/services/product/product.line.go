package product

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/emshop/ots/mgrserver/api/modules/const/sql"
	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/emshop/ots/mgrserver/api/modules/db"
)

//ProductLineHandler 产品线处理服务
type ProductLineHandler struct {
}

func NewProductLineHandler() *ProductLineHandler {
	return &ProductLineHandler{}
}

//PostHandle 添加产品线数据
func (u *ProductLineHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加产品线数据--------")
	
	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(postProductLineCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	xdb := hydra.C.DB().GetRegularDB()
	plID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "产品线"})
	if err != nil {
		return err
	}
	input := ctx.Request().GetMap()
	input["pl_id"] = plID
	count, err := xdb.Execute(sql.InsertProductLine, input)
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}


//GetHandle 获取产品线单条数据
func (u *ProductLineHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取产品线单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getProductLineCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetProductLineByPlID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}

//QueryHandle  获取产品线数据列表
func (u *ProductLineHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取产品线数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryProductLineCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetProductLineListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetProductLineList, m)
		if err != nil {
			return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
		}
	}

	ctx.Log().Info("3.返回结果")
	return map[string]interface{}{
		"items": items,
		"count": types.GetInt(count),
	}
}
//PutHandle 更新产品线数据
func (u *ProductLineHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------更新产品线数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(updateProductLineCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	count,err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateProductLineByPlID,ctx.Request().GetMap())
	if err != nil||count<1 {
		return errs.NewErrorf(http.StatusNotExtended,"更新数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}

var postProductLineCheckFields = map[string]interface{}{
	field.FieldPlName:"required",
	field.FieldStatus:"required",
	}

var getProductLineCheckFields = map[string]interface{}{
	field.FieldPlID:"required",
}

var queryProductLineCheckFields = map[string]interface{}{
	field.FieldPlName:"required",
	field.FieldStatus:"required",
	}

var updateProductLineCheckFields = map[string]interface{}{
	field.FieldPlName:"required",
	field.FieldStatus:"required",
	}


