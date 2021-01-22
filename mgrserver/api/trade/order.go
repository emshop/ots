package trade

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
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

//SingleHandle 查询单条数据
func (o *TradeOrderHandler) SingleHandle(ctx hydra.IContext) interface{} {
	items, err := hydra.C.DB().GetRegularDB().Query(sqlSignleTradeOrder, ctx.Request().GetMap())
	if err != nil {
		return err
	}
	if items.Len() == 0 {
		return errs.NewError(http.StatusNoContent, "未查询到数据")
	}
	return items.Get(0)
}

//sqlTradeOrderQuery 查询数据(订单记录)
const sqlTradeOrderQuery = `
select
t.order_id,
t.mer_no,
t.mer_order_no,
t.mer_product_id,
t.mer_shelf_id,
t.mer_product_no,
t.pl_id,
t.brand_no,
t.province_no,
t.city_no,
t.face,
t.num,
t.total_face,
t.account_name,
t.invoice_type,
t.sell_discount,
t.sell_amount,
t.mer_fee_discount,
t.trade_fee_discount,
t.payment_fee_discount,
t.can_split_order,
t.create_time,
t.order_timeout,
t.payment_timeout,
t.delivery_pause,
t.order_status,
t.payment_status,
t.delivery_status,
t.refund_status,
t.notify_status,
t.is_refund,
t.bind_face,
t.success_face,
t.success_sell_amount,
t.success_mer_fee,
t.success_mer_trade_fee,
t.success_mer_payment_fee,
t.success_cost_amount,
t.success_spp_fee,
t.success_spp_trade_fee,
t.success_spp_payment_fee,
t.profit 
from ots_trade_order t where 1 = 1
&t.order_id
&t.mer_no
&t.mer_order_no
&t.pl_id
&t.brand_no
&t.province_no`

//sqlTradeOrderCount 查询条数(订单记录)
const sqlTradeOrderCount = `
select count(1) from ots_trade_order t where 1 = 1
&t.order_id
&t.mer_no
&t.mer_order_no
&t.pl_id
&t.brand_no
&t.province_no`

//sqlSignleTradeOrder 查询单条数据(订单记录)
const sqlSignleTradeOrder = `
select
t.order_id,
t.mer_no,
t.mer_order_no,
t.mer_product_id,
t.mer_shelf_id,
t.mer_product_no,
t.pl_id,
t.brand_no,
t.province_no,
t.city_no,
t.face,
t.num,
t.total_face,
t.account_name,
t.invoice_type,
t.sell_discount,
t.sell_amount,
t.mer_fee_discount,
t.trade_fee_discount,
t.payment_fee_discount,
t.can_split_order,
t.create_time,
t.order_timeout,
t.payment_timeout,
t.delivery_pause,
t.order_status,
t.payment_status,
t.delivery_status,
t.refund_status,
t.notify_status,
t.is_refund,
t.bind_face,
t.success_face,
t.success_sell_amount,
t.success_mer_fee,
t.success_mer_trade_fee,
t.success_mer_payment_fee,
t.success_cost_amount,
t.success_spp_fee,
t.success_spp_trade_fee,
t.success_spp_payment_fee,
t.profit 
from ots_trade_order t where
t.order_id = @order_id `
