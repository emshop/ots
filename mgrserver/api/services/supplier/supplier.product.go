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

//SupplierProductHandler 供货商商品处理服务
type SupplierProductHandler struct {
}

func NewSupplierProductHandler() *SupplierProductHandler {
	return &SupplierProductHandler{}
}

//PostHandle 添加供货商商品数据
func (u *SupplierProductHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加供货商商品数据--------")
	
	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(postSupplierProductCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	xdb := hydra.C.DB().GetRegularDB()
	sppProductID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "供货商商品"})
	if err != nil {
		return err
	}
	input := ctx.Request().GetMap()
	input["spp_product_id"] = sppProductID
	count, err := xdb.Execute(sql.InsertSupplierProduct, input)
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}


//GetHandle 获取供货商商品单条数据
func (u *SupplierProductHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取供货商商品单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getSupplierProductCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetSupplierProductBySppProductID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}


//QueryHandle  获取供货商商品数据列表
func (u *SupplierProductHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取供货商商品数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(querySupplierProductCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}
	

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetSupplierProductListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetSupplierProductList, m)
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

//PutHandle 更新供货商商品数据
func (u *SupplierProductHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------更新供货商商品数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(updateSupplierProductCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	_, err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateSupplierProductBySppProductID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"更新数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}

var postSupplierProductCheckFields = map[string]interface{}{
	field.FieldSppShelfID:"required",
	field.FieldSppNo:"required",
	
	field.FieldPlID:"required",
	field.FieldBrandNo:"required",
	field.FieldProvinceNo:"required",
	field.FieldCityNo:"required",
	field.FieldFace:"required",
	field.FieldCostDiscount:"required",
	field.FieldStatus:"required",
	}

var getSupplierProductCheckFields = map[string]interface{}{
	field.FieldSppProductID:"required",
}


var querySupplierProductCheckFields = map[string]interface{}{
	field.FieldSppShelfID:"required",
	field.FieldSppNo:"required",
	field.FieldPlID:"required",
	field.FieldBrandNo:"required",
	field.FieldProvinceNo:"required",
	}


var updateSupplierProductCheckFields = map[string]interface{}{
	
	field.FieldBrandNo:"required",
	field.FieldCostDiscount:"required",
	field.FieldStatus:"required",
	}


