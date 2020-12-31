package orders

//TradeOrder 订单记录
type TradeOrder struct {
	//OrderID 订单编号
	OrderID int64 `json:"order_id"  valid:"required"`

	//MerNo 商户编号
	MerNo string `json:"mer_no"  valid:"required"`

	//MerOrderNo 商户订单编号
	MerOrderNo string `json:"mer_order_no"  valid:"required"`

	//MerProductID 商品编号
	MerProductID int `json:"mer_product_id"  valid:"required"`

	//Face 商品面值
	Face float32 `json:"face"  valid:"required"`

	//Num 商品数量
	Num int `json:"num"  valid:"required"`

	//TotalFace 商品总面值
	TotalFace float32 `json:"total_face"  valid:"required"`
}

//Query 查询商家订单
func Query(merNo string, merOrderNo string) (*TradeOrder, error) {
	return nil, nil
}

//Create 创建订单
func Create(merNo string, merOrderNo string, shelf *MerchantShelf, product *MerchantProduct, num int) (*TradeOrder, error) {
	return nil, nil
}
