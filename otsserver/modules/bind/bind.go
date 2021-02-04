package bind

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/delivery"
	"github.com/emshop/ots/otsserver/modules/spp"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

//Bind 订单绑定上游渠道
func Bind(orderID string) (string, error) {

	//查询订单是否需要绑定处理
	orders, err := hydra.C.DB().GetRegularDB().Query(sql.SelectTradeOrderByOrderIDForBind, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		return "", err
	}

	//订单无需绑定
	if orders.Len() != 1 {
		return "", errs.NewError(http.StatusNoContent, "订单无须绑定")
	}

	//查询上游产品
	products, err := spp.GetProducts(
		orders.Get(0).GetInt(sql.FieldPlID),
		orders.Get(0).GetString(sql.FieldBrandNo),
		orders.Get(0).GetString(sql.FieldProvinceNo),
		orders.Get(0).GetString(sql.FieldCityNo),
		orders.Get(0).GetInt(sql.FieldCanSplitOrder, 1),
		orders.Get(0).GetInt(sql.FieldTotalFace),
		orders.Get(0).GetDecimal(sql.FieldSellDiscount))
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
		deliveryID, ok, err := delivery.Create(orders.Get(0), shelf, product)
		if err != nil || ok {
			return deliveryID, err
		}
	}

	return "", errs.NewErrorf(http.StatusAccepted, "无可用的上游可绑定%d", orderID)
}
