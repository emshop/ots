package mers

//MerchantProduct 商户商品
type MerchantProduct struct {
	//MerProductID 商品编号
	MerProductID int `json:"mer_product_id"  valid:"required"`

	//MerShelfID 货架编号
	MerShelfID int `json:"mer_shelf_id"  valid:"required"`

	//MerNo 商户编号
	MerNo string `json:"mer_no"  valid:"required"`

	//PlID 产品线
	PlID int `json:"pl_id"  valid:"required"`

	//BrandNo 品牌
	BrandNo string `json:"brand_no"  valid:"required"`

	//ProvinceNo 省份
	ProvinceNo string `json:"province_no"  valid:"required"`

	//CityNo 城市
	CityNo string `json:"city_no"  valid:"required"`

	//Face 面值
	Face float32 `json:"face"  valid:"required"`

	//MerProductNo 商户商品编号
	MerProductNo string `json:"mer_product_no"  valid:"required"`

	//Discount 销售折扣（以面值算）
	Discount float32 `json:"discount"  valid:"required"`
}
