package orders

import (
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/db"
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
func QueryByOrderID(orderID string) (t types.XMap, err error) {
	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectTradeOrderByOrderID, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}

//GetOrderByDB 查询订单
func GetOrderByDB(db db.IDBExecuter, orderID string) (t types.XMap, err error) {
	data, err := db.Query(sql.SelectTradeOrderByOrderID, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}

//GetOrderForBind 查询订单
func GetOrderForBind(orderID string) (t types.XMap, err error) {

	//查询订单是否需要绑定处理
	orders, err := hydra.C.DB().GetRegularDB().Query(sql.SelectTradeOrderByOrderIDForBind, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		return nil, err
	}
	return orders.Get(0), nil
}
