package trade

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
)

//TradeOrderHandler 订单记录处理服务
type TradeOrderHandler struct {
}

//QueryHandle 订单记录查询服务
func (o *TradeOrderHandler) QueryHandle(ctx hydra.IContext) interface{} {
	items, err := hydra.C.DB().GetRegularDB().Query(sqlTradeOrderQuery, ctx.Request().GetMap())
	if err != nil {
		return err
	}
	count, err := hydra.C.DB().GetRegularDB().Scalar(sqlTradeOrderCount, ctx.Request().GetMap())
	return map[string]interface{}{
		"items": items,
		"count": count,
	}
}

//GetHandle 订单记录查询服务
func (o *TradeOrderHandler) GetHandle(ctx hydra.IContext) interface{} {
	items, err := hydra.C.DB().GetRegularDB().Query(sqlTradeOrderGet, ctx.Request().GetMap())
	if err != nil {
		return err
	}
	return items.Get(0)
}

const sqlTradeOrderQuery = `
select t.* from ots_trade_order t where 1 = 1
&t.order_id
&t.mer_no
&t.mer_order_no
&t.pl_id
&t.brand_no
&t.province_no`
const sqlTradeOrderCount = `
select count(1) from ots_trade_order t where 1 = 1
&t.order_id
&t.mer_no
&t.mer_order_no
&t.pl_id
&t.brand_no
&t.province_no`

const sqlTradeOrderGet = `
select t.* from ots_trade_order t where t.order_id=@order_id`
