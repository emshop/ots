package orders

import "github.com/emshop/ots/modules/mers"

//Query 查询商家订单
func Query(merNo string, merOrderNo string) (*TradeOrder, error) {
	return nil, nil
}

//Create 创建订单
func Create(merNo string, merOrderNo string, shelf *mers.MerchantShelf, product *mers.MerchantProduct, num int) (*TradeOrder, error) {
	return nil, nil
}
