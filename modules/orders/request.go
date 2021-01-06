package orders

import (
	"github.com/emshop/ots/modules/const/db/biz"
	"github.com/emshop/ots/modules/pkg"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//Query 查询商家订单
func Query(merNo string, merOrderNo string) (t types.XMap, err error) {
	data, err := hydra.C.DB().GetRegularDB().Query(biz.SelectTradeOrderByMerchant, map[string]interface{}{
		biz.FieldTradeOrderMerNo:      merNo,
		biz.FieldTradeOrderMerOrderNo: merOrderNo,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}

//Create 创建交易订单
func Create(merNo string, merOrderNo string, shelf types.XMap, product types.XMap, num int) (types.XMap, bool, error) {

	//查询订单记录
	order, err := Query(merNo, merOrderNo)
	if err != nil || order.Len() > 0 {
		return order, false, err
	}
	orderID, err := pkg.GetOrderID()
	if err != nil {
		return nil, false, err
	}
	//构建输入数据

	input := types.NewXMapByMap(shelf)
	input.Merge(product)
	input.SetValue(biz.FieldTradeOrderMerNo, merNo)
	input.SetValue(biz.FieldTradeOrderMerOrderNo, merOrderNo)
	input.SetValue(biz.FieldTradeOrderOrderID, orderID)
	input.SetValue(biz.FieldTradeOrderNum, num)

	//添加数据
	_, err = hydra.C.DB().GetRegularDB().Execute(biz.InsertTradeOrder, input)

	// 查询订单
	order, err1 := Query(merNo, merOrderNo)
	if err1 != nil || order.Len() == 0 {
		return nil, false, err // 未查询到订单，返回创建错误
	}

	//检查是否是本次创建的订单
	if input.GetString(biz.FieldTradeOrderOrderID) == order.GetString(biz.FieldTradeOrderOrderID) {
		return order, true, nil
	}
	return order, false, nil

}
