package dictionary

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/emshop/ots/mgrserver/api/modules/const/sql"
	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/micro-plat/lib4go/types"
	"github.com/emshop/ots/mgrserver/api/modules/db"
)

//DictionaryInfoHandler 字典配置处理服务
type DictionaryInfoHandler struct {
}

func NewDictionaryInfoHandler() *DictionaryInfoHandler {
	return &DictionaryInfoHandler{}
}

//PostHandle 添加字典配置数据
func (u *DictionaryInfoHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------添加字典配置数据--------")
	
	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(postDictionaryInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	xdb := hydra.C.DB().GetRegularDB()
	ID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "字典配置"})
	if err != nil {
		return err
	}
	input := ctx.Request().GetMap()
	input["id"] = ID
	count, err := xdb.Execute(sql.InsertDictionaryInfo, input)
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}


//GetHandle 获取字典配置单条数据
func (u *DictionaryInfoHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取字典配置单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getDictionaryInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetDictionaryInfoByID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}


//QueryHandle  获取字典配置数据列表
func (u *DictionaryInfoHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取字典配置数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryDictionaryInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetDictionaryInfoListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetDictionaryInfoList, m)
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

//PutHandle 更新字典配置数据
func (u *DictionaryInfoHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------更新字典配置数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(updateDictionaryInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	_, err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateDictionaryInfoByID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"更新数据出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return "success"
}

var postDictionaryInfoCheckFields = map[string]interface{}{
	field.FieldName:"required",
	field.FieldValue:"required",
	field.FieldType:"required",
	field.FieldStatus:"required",
	field.FieldSortNo:"required",
	}

var getDictionaryInfoCheckFields = map[string]interface{}{
	field.FieldID:"required",
}


var queryDictionaryInfoCheckFields = map[string]interface{}{
	field.FieldName:"required",
	field.FieldType:"required",
	field.FieldStatus:"required",
	}


var updateDictionaryInfoCheckFields = map[string]interface{}{
	field.FieldName:"required",
	field.FieldValue:"required",
	field.FieldType:"required",
	field.FieldStatus:"required",
	field.FieldSortNo:"required",
	}


