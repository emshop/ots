package bind

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/delivery"
	"github.com/emshop/ots/otsserver/modules/orders"
	"github.com/emshop/ots/otsserver/modules/spp"
	"github.com/emshop/ots/otsserver/modules/state"
	"github.com/micro-plat/lib4go/errs"
)

//Bind 订单绑定上游渠道
func Bind(orderID string) (string, error) {

	//查询订单，检查是否可以绑定
	order, err := orders.GetOrderForBind(orderID)
	if err != nil {
		return "", err
	}
	var binding state.Binding
	if !binding.Check(order) {
		return "", errs.NewStopf(http.StatusNoContent, "订单(%s)状态错误，无法绑定", orderID)
	}
	if binding.IsFinish(order) {
		return "", errs.NewStopf(http.StatusNoContent, "订单(%s)已全部绑定", orderID)
	}
	//查询上游产品
	products, err := spp.GetProducts(
		order.GetInt(sql.FieldPlID),
		order.GetString(sql.FieldBrandNo),
		order.GetString(sql.FieldProvinceNo),
		order.GetString(sql.FieldCityNo),
		order.GetInt(sql.FieldCanSplitOrder, 1),
		order.GetInt(sql.FieldFace),
		order.GetDecimal(sql.FieldSellDiscount))
	if err != nil {
		return "", err
	}

	//无可用产品
	if products.Len() == 0 {
		return "", errs.NewErrorf(http.StatusAccepted, "暂无上游产品可绑定%s", orderID)
	}

	//循环处理每次绑定
	for _, product := range products {

		//查询货架信息
		shelf, err := spp.GetShelfByID(product.GetInt(sql.FieldSppShelfID))
		if err != nil {
			return "", err
		}
		if shelf.Len() == 0 {
			continue
		}

		//创建发货记录
		deliveryID, ok, err := delivery.Create(order, shelf, product)
		if err != nil || ok {
			return deliveryID, err
		}
	}

	return "", errs.NewErrorf(http.StatusAccepted, "无可用的上游可绑定%d", orderID)
}
