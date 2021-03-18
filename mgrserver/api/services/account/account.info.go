package account

import (
	"net/http"

	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/emshop/ots/mgrserver/api/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"

	"github.com/emshop/ots/mgrserver/api/modules/db"
)

//AccountInfoHandler 账户信息处理服务
type AccountInfoHandler struct {
}

func NewAccountInfoHandler() *AccountInfoHandler {
	return &AccountInfoHandler{}
}

//PostHandle 添加账户信息数据
func (u *AccountInfoHandler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("--------添加账户信息数据--------")

	ctx.Log().Debug("1.参数校验")
	if err := ctx.Request().CheckMap(postAccountInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Debug("2.执行操作")
	xdb := hydra.C.DB().GetRegularDB()
	accountID, err := db.GetNewID(xdb, sql.SQLGetSEQ, map[string]interface{}{"name": "账户信息"})
	if err != nil {
		return err
	}
	input := ctx.Request().GetMap()
	input["account_id"] = accountID
	count, err := xdb.Execute(sql.InsertAccountInfo, input)
	if err != nil || count < 1 {
		return errs.NewErrorf(http.StatusNotExtended, "添加数据出错:%+v", err)
	}

	ctx.Log().Debug("3.返回结果")
	return "success"
}

//GetHandle 获取账户信息单条数据
func (u *AccountInfoHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("--------获取账户信息单条数据--------")

	ctx.Log().Debug("1.参数校验")
	if err := ctx.Request().CheckMap(getAccountInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Debug("2.执行操作")
	items, err := hydra.C.DB().GetRegularDB().Query(sql.GetAccountInfoByAccountID, ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Debug("3.返回结果")
	return items.Get(0)
}

//QueryHandle  获取账户信息数据列表
func (u *AccountInfoHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("--------获取账户信息数据列表--------")

	ctx.Log().Debug("1.参数校验")
	if err := ctx.Request().CheckMap(queryAccountInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Debug("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetAccountInfoListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}

	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetAccountInfoList, m)
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

//PutHandle 更新账户信息数据
func (u *AccountInfoHandler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Debug("--------更新账户信息数据--------")

	ctx.Log().Debug("1.参数校验")
	if err := ctx.Request().CheckMap(updateAccountInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Debug("2.执行操作")
	_, err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateAccountInfoByAccountID, ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "更新数据出错:%+v", err)
	}

	ctx.Log().Debug("3.返回结果")
	return "success"
}

var postAccountInfoCheckFields = map[string]interface{}{
	field.FieldAccountName: "required",
	field.FieldIdent:       "required",
	field.FieldGroups:      "required",
	field.FieldEid:         "required",
	field.FieldCredit:      "required",
	field.FieldStatus:      "required",
}

var getAccountInfoCheckFields = map[string]interface{}{
	field.FieldAccountID: "required",
}

var queryAccountInfoCheckFields = map[string]interface{}{
	field.FieldEid: "required",
}

var updateAccountInfoCheckFields = map[string]interface{}{
	field.FieldAccountName: "required",
	field.FieldCredit:      "required",
	field.FieldStatus:      "required",
}
