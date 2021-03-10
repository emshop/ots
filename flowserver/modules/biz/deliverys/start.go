package deliverys

import (
	"errors"
	"net/http"
	"time"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/dbs"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//SaveStart 保存开始发货的结果消息
func SaveStart(deliveryID string, discount types.Decimal, resultCode string, returnMsg string) error {

	result, msg := GetDealCode(deliveryID, enums.ResultFromRequest, resultCode)

	//根据发货结果修改发货记录状态
	input := map[string]interface{}{
		fields.FieldDeliveryID:   deliveryID,
		fields.FieldResultCode:   resultCode,
		fields.FieldReturnMsg:    types.GetString(returnMsg, msg),
		fields.FieldCostDiscount: discount.String(),
	}
	var row int64
	var err error
	switch result {
	case enums.Failed:
		row, err = hydra.C.DB().GetRegularDB().Execute(UpdateTradeDeliveryForDeliveryingFailed, input)
	default: //非失败都纳入成功处理，等待后续通知、查询或人工审核
		row, err = hydra.C.DB().GetRegularDB().Execute(UpdateTradeDeliveryForDeliveryingSuccess, input)
	}
	if err != nil {
		return err
	}
	if row == 0 {
		return errs.NewError(http.StatusNoContent, "无须发货处理(save.start)")
	}
	return nil
}

//Start 开始发货
func Start(deliveryID string) (*TradeDelivery, error) {
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}

	//修改发货记录为正在发货
	data, err := dbs.Executes(db, types.XMap{fields.FieldDeliveryID: deliveryID}, deliveryStartNow...)
	if errors.Is(err, xerr.ErrNOTEXISTS) {
		return nil, errs.NewErrorf(http.StatusNoContent, "发货编号(%s)无须发货", deliveryID)
	}
	if err != nil {
		db.Rollback()
		return nil, err
	}

	db.Commit()
	delviery := &TradeDelivery{}
	err = data.ToAnyStruct(delviery)
	return delviery, err
}

//TradeDelivery 订单发货表
type TradeDelivery struct {

	//DeliveryID 发货编号
	DeliveryID string `json:"delivery_id"`

	//OrderID 订单编号
	OrderID string `json:"order_id"`

	//SppNo 供货商编号
	SppNo string `json:"spp_no"`

	//SppProductID 供货商商品编号
	SppProductID int `json:"spp_product_id"`

	//SppDeliveryNo 供货商发货编号
	SppDeliveryNo string `json:"spp_delivery_no"`

	//SppProductNo 供货商商品编号
	SppProductNo string `json:"spp_product_no"`

	//MerNo 商户编号
	MerNo string `json:"mer_no"`

	//MerProductID 商户商品编号
	MerProductID int `json:"mer_product_id"`

	//PlID 产品线
	PlID int `json:"pl_id"`

	//BrandNo 品牌
	BrandNo string `json:"brand_no"`

	//ProvinceNo 省份
	ProvinceNo string `json:"province_no"`

	//CityNo 城市
	CityNo string `json:"city_no"`

	//InvoiceType 开票方式（1.不开发票）
	InvoiceType int `json:"invoice_type"`

	//AccountName 用户账户信息
	AccountName string `json:"account_name"`

	//DeliveryStatus 发货状态
	DeliveryStatus int `json:"delivery_status"`

	//PaymentStatus 支付状态
	PaymentStatus int `json:"payment_status"`

	//CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`

	//Face 商品面值
	Face int `json:"face"`

	//Num 发货数量
	Num int `json:"num"`

	//TotalFace 发货总面值
	TotalFace int `json:"total_face"`

	//CostAmount 发货成本
	CostAmount types.Decimal `json:"cost_amount"`

	//SppFeeAmount 供货商佣金
	SppFeeAmount types.Decimal `json:"spp_fee_amount"`

	//TradeFeeAmount 发货服务费
	TradeFeeAmount types.Decimal `json:"trade_fee_amount"`

	//PaymentFeeAmount 支付服务费
	PaymentFeeAmount types.Decimal `json:"payment_fee_amount"`
}
