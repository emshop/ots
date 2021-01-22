package trade

import (
	"net/http"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//TradeOrderHandler 订单记录处理服务
type TradeOrderHandler struct {
	
	getCheckFileds    map[string]interface{} 
	queryCheckFileds  map[string]interface{} 
	
	
}

func NewTradeOrderHandler() *TradeOrderHandler {
	return &TradeOrderHandler{
		
		getCheckFileds: map[string]interface{}{
			FieldOrderID:"required",
		},
		queryCheckFileds: map[string]interface{}{
			FieldOrderID:"required",
			FieldMerNo:"required",
			FieldMerOrderNo:"required",
			FieldPlID:"required",
			FieldBrandNo:"required",
			FieldProvinceNo:"required",
			FieldCreateTime:"required",
			},
		
		
 }
}


//GetHandle 获取订单记录单条数据
func (u *TradeOrderHandler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取订单记录单条数据--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(u.getCheckFileds); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")
	items, err :=  hydra.C.DB().GetRegularDB().Query(GetTradeOrderByOrderId,ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}

	ctx.Log().Info("3.返回结果")
	return items.Get(0)
}

//QueryHandle  获取订单记录数据列表
func (u *TradeOrderHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取订单记录数据列表--------")

	ctx.Log().Info("1.参数校验")
	if err := ctx.Request().CheckMap(u.queryCheckFileds); err != nil {
		return errs.NewErrorf(http.StatusNotAcceptable, "参数校验错误:%+v", err)
	}

	ctx.Log().Info("2.执行操作")	
	m := ctx.Request().GetMap()
	m["offset"] = (ctx.Request().GetInt("pi") - 1) * ctx.Request().GetInt("ps")
	items, err := hydra.C.DB().GetRegularDB().Query(GetTradeOrderList, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据出错:%+v", err)
	}
	count, err := hydra.C.DB().GetRegularDB().Scalar(GetTradeOrderListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended,"查询数据数量出错:%+v", err)
	}

	ctx.Log().Info("3.返回结果")
	return map[string]interface{}{
		"items": items,
		"count": types.GetInt(count),
	}
}
