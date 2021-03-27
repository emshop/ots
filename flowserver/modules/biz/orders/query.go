package orders

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//QueryDetail
func QueryDetail(merNo string, merOrderNo string) (t types.IXMap, err error) {
	data, err := hydra.C.DB().GetRegularDB().Query(selectTradeOrderByMerchantForDetail, map[string]interface{}{
		fields.FieldMerNo:      merNo,
		fields.FieldMerOrderNo: merOrderNo,
	})
	if err != nil {
		return nil, err
	}
	if data.Len() == 0 {
		return nil, errs.NewError(http.StatusNoContent, "未查询到订单")
	}
	order := data.Get(0)
	if order.GetString(fields.FieldOrderStatus) == strings.ToUpper(enums.LbSuccess) &&
		order.GetInt(fields.FieldHasFeedback) == int(enums.Success) {
		deliverys, err := hydra.C.DB().GetRegularDB().Query(selectDeliveryReuqestParams, order.ToMap())
		if err != nil {
			return nil, err
		}
		list := make([]string, 0, deliverys.Len())
		for i := 0; i < deliverys.Len(); i++ {
			list = append(list, deliverys.Get(i).GetString(fields.FieldRequestParams))
		}
		order.SetValue("products", fmt.Sprintf("[%s]", strings.Join(list, ",")))
	}
	return order, nil
}

//Query 查询商家订单
func Query(merNo string, merOrderNo string) (t types.IXMap, err error) {
	data, err := hydra.C.DB().GetRegularDB().Query(selectTradeOrderByMerchant, map[string]interface{}{
		fields.FieldMerNo:      merNo,
		fields.FieldMerOrderNo: merOrderNo,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}

//QueryByOrderID 查询订单
func QueryByOrderID(orderID string) (t types.IXMap, err error) {
	data, err := hydra.C.DB().GetRegularDB().Query(selectTradeOrderByOrderID, map[string]interface{}{
		fields.FieldOrderID: orderID,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}
