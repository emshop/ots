package account

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/emshop/ots/mgrserver/api/modules/const/sql"
	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/micro-plat/lib4go/types"
	
	
)

//AccountRecordHandler 资金变动处理服务
type AccountRecordHandler struct {
}

func NewAccountRecordHandler() *AccountRecordHandler {
	return &AccountRecordHandler{}
}




//GetHandle 获取资金变动单条数据
func (u *AccountRecordHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取资金变动单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getAccountRecordCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetAccountRecordByRecordID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}


//QueryHandle  获取资金变动数据列表
func (u *AccountRecordHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取资金变动数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryAccountRecordCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}
	

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetAccountRecordListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetAccountRecordList, m)
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

//QueryDetailHandle  获取资金变动数据列表
func (u *AccountRecordHandler) QueryDetailHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取资金变动数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryAccountRecordDetailCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetAccountRecordDetailListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetAccountRecordDetailList, m)
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



var getAccountRecordCheckFields = map[string]interface{}{
	field.FieldRecordID:"required",
}


var queryAccountRecordCheckFields = map[string]interface{}{
	field.FieldAccountID:"required",
	field.FieldTradeNo:"required",
	field.FieldTradeType:"required",
	field.FieldChangeType:"required",
	}

var queryAccountRecordDetailCheckFields = map[string]interface{}{
	field.FieldTradeNo:"required",
	
}




