package mers

import (
	"time"

	"github.com/micro-plat/lib4go/types"
)

//MerchantProduct 商户商品
type MerchantProduct struct {

	//MerProductID 商品编号
	MerProductID int `json:"mer_product_id"`

	//MerShelfID 货架编号
	MerShelfID int `json:"mer_shelf_id"`

	//MerNo 商户编号
	MerNo string `json:"mer_no"`

	//PlID 产品线
	PlID int `json:"pl_id"`

	//BrandNo 品牌
	BrandNo string `json:"brand_no"`

	//ProvinceNo 省份
	ProvinceNo string `json:"province_no"`

	//CityNo 城市
	CityNo string `json:"city_no"`

	//Face 面值
	Face types.Decimal `json:"face"`

	//MerProductNo 商户商品编号
	MerProductNo string `json:"mer_product_no"`

	//Discount 销售折扣（以面值算）
	Discount types.Decimal `json:"discount"`

	//Status 状态(0.是,1.否)
	Status int `json:"status"`

	//CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
}
