package supplier

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/emshop/ots/mgrserver/api/modules/const/sql"
	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/micro-plat/lib4go/types"
	
	"github.com/emshop/ots/mgrserver/api/modules/db"
)

//SupplierShelfHandler 供货商货架处理服务
type SupplierShelfHandler struct {
}

func NewSupplierShelfHandler() *SupplierShelfHandler {
	return &SupplierShelfHandler{}
}

//PostHandle 添加供货商货架数据
func (u *SupplierShelfHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加供货商货架数据--------")
	
	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(postSupplierShelfCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	xdb := hydra.C.DB().GetRegularDB()
	sppShelfID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "供货商货架"})
	if err != nil {
		return err
	}
	input := ctx.Request().GetMap()
	input["spp_shelf_id"] = sppShelfID
	count, err := xdb.Execute(sql.InsertSupplierShelf, input)
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}


//GetHandle 获取供货商货架单条数据
func (u *SupplierShelfHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取供货商货架单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getSupplierShelfCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetSupplierShelfBySppShelfID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}


//QueryHandle  获取供货商货架数据列表
func (u *SupplierShelfHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取供货商货架数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(querySupplierShelfCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}
	

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetSupplierShelfListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetSupplierShelfList, m)
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

//PutHandle 更新供货商货架数据
func (u *SupplierShelfHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------更新供货商货架数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(updateSupplierShelfCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	_, err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateSupplierShelfBySppShelfID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"更新数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}

var postSupplierShelfCheckFields = map[string]interface{}{
	field.FieldSppShelfID:"required",
	field.FieldSppShelfName:"required",
	field.FieldSppNo:"required",
	field.FieldReqURL:"required",
	
	
	
	
	
	
	
	
	
	
	}

var getSupplierShelfCheckFields = map[string]interface{}{
	field.FieldSppShelfID:"required",
}


var querySupplierShelfCheckFields = map[string]interface{}{
	field.FieldSppShelfName:"required",
	field.FieldSppNo:"required",
	field.FieldStatus:"required",
	}


var updateSupplierShelfCheckFields = map[string]interface{}{
	field.FieldSppShelfName:"required",
	field.FieldReqURL:"required",
	
	
	
	
	
	
	
	
	
	
	}


