package supplier

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"github.com/emshop/ots/mgrserver/api/modules/const/sql"
	"github.com/emshop/ots/mgrserver/api/modules/const/field"
)

//SupplierEcodeHandler 供货商错误码处理服务
type SupplierEcodeHandler struct {
}

func NewSupplierEcodeHandler() *SupplierEcodeHandler {
	return &SupplierEcodeHandler{}
}




//GetHandle 获取供货商错误码单条数据
func (u *SupplierEcodeHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取供货商错误码单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getSupplierEcodeCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetSupplierEcodeById,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}

//QueryHandle  获取供货商错误码数据列表
func (u *SupplierEcodeHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取供货商错误码数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(querySupplierEcodeCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetSupplierEcodeListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetSupplierEcodeList, m)
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



var getSupplierEcodeCheckFields = map[string]interface{}{
	field.FieldID:"required",
}

var querySupplierEcodeCheckFields = map[string]interface{}{
	field.FieldSppNo:"required",
	field.FieldPlID:"required",
	field.FieldCategory:"required",
	field.FieldStatus:"required",
	}




