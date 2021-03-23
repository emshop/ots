package audit

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/emshop/ots/mgrserver/api/modules/const/sql"
	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/micro-plat/lib4go/types"
	
	
)

//AuditInfoHandler 发货人工审核表处理服务
type AuditInfoHandler struct {
}

func NewAuditInfoHandler() *AuditInfoHandler {
	return &AuditInfoHandler{}
}




//GetHandle 获取发货人工审核表单条数据
func (u *AuditInfoHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取发货人工审核表单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(getAuditInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(sql.GetAuditInfoByDeliveryID,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}


//QueryHandle  获取发货人工审核表数据列表
func (u *AuditInfoHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取发货人工审核表数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryAuditInfoCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}
	

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetAuditInfoListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetAuditInfoList, m)
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




var getAuditInfoCheckFields = map[string]interface{}{
	field.FieldDeliveryID:"required",
}


var queryAuditInfoCheckFields = map[string]interface{}{
	field.FieldOrderID:"required",
	field.FieldAuditStatus:"required",
	}





