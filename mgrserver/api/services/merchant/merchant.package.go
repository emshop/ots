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

//MerchantPackageHandler 组合商品处理服务
type MerchantPackageHandler struct {
}

func NewMerchantPackageHandler() *MerchantPackageHandler {
	return &MerchantPackageHandler{}
}

//PostHandle 添加组合商品数据
func (u *MerchantPackageHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加组合商品数据--------")
	
	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(postMerchantPackageCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	xdb := hydra.C.DB().GetRegularDB()
	pgID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "组合商品"})
	if err != nil {
		return err
	}
	input := ctx.Request().GetMap()
	input["pg_id"] = pgID
	count, err := xdb.Execute(sql.InsertMerchantPackage, input)
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}


//GetHandle 获取组合商品单条数据
func (u *MerchantPackageHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取组合商品单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getMerchantPackageCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetMerchantPackageByPgID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}


//QueryHandle  获取组合商品数据列表
func (u *MerchantPackageHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取组合商品数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryMerchantPackageCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}
	
	orderBy := ctx.Request().GetString("order_by")
	if len(orderBy) > 1 && !regexp.MustCompile("^t.[A-Za-z0-9_,.\\s]+ (asc|desc)$").MatchString(orderBy) {
		return errs.NewErrorf(http.StatusNotAcceptable, "排序参数校验错误!")
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetMerchantPackageListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetMerchantPackageList, m)
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

//QueryDetailHandle  获取组合商品数据列表
func (u *MerchantPackageHandler) QueryDetailHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取组合商品数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryMerchantPackageDetailCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetMerchantPackageDetailListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetMerchantPackageDetailList, m)
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
//PutHandle 更新组合商品数据
func (u *MerchantPackageHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------更新组合商品数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(updateMerchantPackageCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	_, err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateMerchantPackageByPgID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"更新数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}

var postMerchantPackageCheckFields = map[string]interface{}{
	field.FieldPgName:"required",
	field.FieldPlID:"required",
	field.FieldMerProductID:"required",
	field.FieldBrandNo:"required",
	field.FieldProvinceNo:"required",
	field.FieldCityNo:"required",
	field.FieldFace:"required",
	field.FieldNum:"required",
	field.FieldDiscount:"required",
	field.FieldStatus:"required",
	}

var getMerchantPackageCheckFields = map[string]interface{}{
	field.FieldPgID:"required",
}


var queryMerchantPackageCheckFields = map[string]interface{}{
	field.FieldPgName:"required",
	field.FieldBrandNo:"required",
	field.FieldNum:"required",
	field.FieldStatus:"required",
	}

var queryMerchantPackageDetailCheckFields = map[string]interface{}{
	field.FieldMerProductID:"required",
	
}

var updateMerchantPackageCheckFields = map[string]interface{}{
	field.FieldPgName:"required",
	field.FieldPlID:"required",
	field.FieldBrandNo:"required",
	field.FieldFace:"required",
	field.FieldNum:"required",
	field.FieldDiscount:"required",
	field.FieldStatus:"required",
	}


