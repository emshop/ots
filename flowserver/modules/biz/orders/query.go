package orders

import (
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//Query 查询商家订单
func Query(merNo string, merOrderNo string) (t types.IXMap, err error) {
	data, err := hydra.C.DB().GetRegularDB().Query(SelectTradeOrderByMerchant, map[string]interface{}{
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
	data, err := hydra.C.DB().GetRegularDB().Query(SelectTradeOrderByOrderID, map[string]interface{}{
		fields.FieldOrderID: orderID,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}
