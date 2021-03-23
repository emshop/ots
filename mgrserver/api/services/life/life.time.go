package life

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/emshop/ots/mgrserver/api/modules/const/sql"
	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/micro-plat/lib4go/types"
	
	
)

//LifeTimeHandler 生命周期处理服务
type LifeTimeHandler struct {
}

func NewLifeTimeHandler() *LifeTimeHandler {
	return &LifeTimeHandler{}
}






//QueryHandle  获取生命周期数据列表
func (u *LifeTimeHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取生命周期数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryLifeTimeCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}
	

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetLifeTimeListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetLifeTimeList, m)
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

//QueryDetailHandle  获取生命周期数据列表
func (u *LifeTimeHandler) QueryDetailHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取生命周期数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(queryLifeTimeDetailCheckFields); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")

	count, err := hydra.C.DB().GetRegularDB().Scalar(sql.GetLifeTimeDetailListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}
	
	var items types.XMaps
	if types.GetInt(count) > 0 {
		items, err = hydra.C.DB().GetRegularDB().Query(sql.GetLifeTimeDetailList, m)
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





var queryLifeTimeCheckFields = map[string]interface{}{
	field.FieldOrderNo:"required",
	}

var queryLifeTimeDetailCheckFields = map[string]interface{}{
	field.FieldOrderNo:"required",
	
}




