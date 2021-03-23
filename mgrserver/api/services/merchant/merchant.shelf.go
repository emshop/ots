package merchant

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/emshop/ots/mgrserver/api/modules/const/sql"
	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/micro-plat/lib4go/types"
	
	"github.com/emshop/ots/mgrserver/api/modules/db"
)

//MerchantShelfHandler 商户货架处理服务
type MerchantShelfHandler struct {
}

func NewMerchantShelfHandler() *MerchantShelfHandler {
	return &MerchantShelfHandler{}
}

//PostHandle 添加商户货架数据
func (u *MerchantShelfHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加商户货架数据--------")
	
	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(postMerchantShelfCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	xdb := hydra.C.DB().GetRegularDB()
	merShelfID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "商户货架"})
	if err != nil {
		return err
	}
	input := ctx.Request().GetMap()
	input["mer_shelf_id"] = merShelfID
	count, err := xdb.Execute(sql.InsertMerchantShelf, input)
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}


//GetHandle 获取商户货架单条数据
func (u *MerchantShelfHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取商户货架单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getMerchantShelfCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetMerchantShelfByMerShelfID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}


//QueryHandle  获取商户货架数据列表
func (u *MerchantShelfHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取商户货架数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryMerchantShelfCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}
	

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetMerchantShelfListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetMerchantShelfList, m)
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

//PutHandle 更新商户货架数据
func (u *MerchantShelfHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------更新商户货架数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(updateMerchantShelfCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	_, err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateMerchantShelfByMerShelfID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"更新数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}

var postMerchantShelfCheckFields = map[string]interface{}{
	field.FieldMerShelfName:"required",
	field.FieldMerNo:"required",
	field.FieldMerFeeDiscount:"required",
	field.FieldTradeFeeDiscount:"required",
	field.FieldPaymentFeeDiscount:"required",
	field.FieldOrderTimeout:"required",
	field.FieldPaymentTimeout:"required",
	field.FieldLimitCount:"required",
	field.FieldInvoiceType:"required",
	field.FieldCanRefund:"required",
	field.FieldCanSplitOrder:"required",
	field.FieldStatus:"required",
	}

var getMerchantShelfCheckFields = map[string]interface{}{
	field.FieldMerShelfID:"required",
}


var queryMerchantShelfCheckFields = map[string]interface{}{
	field.FieldMerShelfName:"required",
	field.FieldMerNo:"required",
	}


var updateMerchantShelfCheckFields = map[string]interface{}{
	field.FieldMerShelfName:"required",
	field.FieldMerFeeDiscount:"required",
	field.FieldTradeFeeDiscount:"required",
	field.FieldPaymentFeeDiscount:"required",
	field.FieldOrderTimeout:"required",
	field.FieldPaymentTimeout:"required",
	field.FieldLimitCount:"required",
	field.FieldInvoiceType:"required",
	field.FieldCanRefund:"required",
	field.FieldCanSplitOrder:"required",
	field.FieldStatus:"required",
	}


