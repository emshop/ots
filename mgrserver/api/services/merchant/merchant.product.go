package merchant

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/emshop/ots/mgrserver/api/modules/const/sql"
	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/micro-plat/lib4go/types"
	"regexp"
	"github.com/emshop/ots/mgrserver/api/modules/db"
)

//MerchantProductHandler 商户商品处理服务
type MerchantProductHandler struct {
}

func NewMerchantProductHandler() *MerchantProductHandler {
	return &MerchantProductHandler{}
}

//PostHandle 添加商户商品数据
func (u *MerchantProductHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加商户商品数据--------")
	
	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(postMerchantProductCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	xdb := hydra.C.DB().GetRegularDB()
	merProductID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "商户商品"})
	if err != nil {
		return err
	}
	input := ctx.Request().GetMap()
	input["mer_product_id"] = merProductID
	count, err := xdb.Execute(sql.InsertMerchantProduct, input)
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}


//GetHandle 获取商户商品单条数据
func (u *MerchantProductHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取商户商品单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getMerchantProductCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetMerchantProductByMerProductID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}


//QueryHandle  获取商户商品数据列表
func (u *MerchantProductHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取商户商品数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryMerchantProductCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}
	
	orderBy := ctx.Request().GetString("order_by")
	if len(orderBy) > 1 && !regexp.MustCompile("^t.[A-Za-z0-9_,.\\s]+ (asc|desc)$").MatchString(orderBy) {
		return errs.NewErrorf(http.StatusNotAcceptable, "排序参数校验错误!")
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetMerchantProductListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetMerchantProductList, m)
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

//PutHandle 更新商户商品数据
func (u *MerchantProductHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------更新商户商品数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(updateMerchantProductCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	_, err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateMerchantProductByMerProductID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"更新数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}

var postMerchantProductCheckFields = map[string]interface{}{
	field.FieldMerShelfID:"required",
	field.FieldMerNo:"required",
	field.FieldProdName:"required",
	field.FieldPlID:"required",
	field.FieldBrandNo:"required",
	field.FieldProvinceNo:"required",
	field.FieldCityNo:"required",
	field.FieldFace:"required",
	
	field.FieldDiscount:"required",
	field.FieldStatus:"required",
	}

var getMerchantProductCheckFields = map[string]interface{}{
	field.FieldMerProductID:"required",
}


var queryMerchantProductCheckFields = map[string]interface{}{
	field.FieldMerShelfID:"required",
	field.FieldMerNo:"required",
	field.FieldProdName:"required",
	field.FieldPlID:"required",
	field.FieldBrandNo:"required",
	}


var updateMerchantProductCheckFields = map[string]interface{}{
	field.FieldProdName:"required",
	field.FieldBrandNo:"required",
	
	field.FieldDiscount:"required",
	field.FieldStatus:"required",
	}


