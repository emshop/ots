package orders

import (
	"github.com/emshop/ots/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//Query 查询商家订单
func Query(merNo string, merOrderNo string) (t types.XMap, err error) {
	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectTradeOrderByMerchant, map[string]interface{}{
		sql.FieldMerNo:      merNo,
		sql.FieldMerOrderNo: merOrderNo,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}

//QueryByOrderID 查询订单
func QueryByOrderID(orderID int64) (t types.XMap, err error) {
	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectTradeOrderByMerchant, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}
