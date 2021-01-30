package delivery

import (
	"net/http"
	"time"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/emshop/ots/otsserver/modules/spp"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//SaveStart 保存开始发货的结果消息
func SaveStart(deliveryID string, discount types.Decimal, resultCode string, returnMsg string) error {

	//根据发货信息查询错误码
	var result enums.FlowStatus
	var msg = ""
	delivery, err := Get(deliveryID)
	if err != nil {
		result = enums.Unknown
		msg = "无法获取发货记录"
	} else {
		result, msg = spp.GetDealCode(
			delivery.GetString(sql.FieldSppNo),
			enums.ResultFromRequest,
			delivery.GetInt(sql.FieldPlID),
			resultCode)
	}

	//根据发货结果修改发货记录状态
	returnMsg = types.GetString(returnMsg, msg)
	var row int64
	switch result {
	case enums.Failed:
		row, err = hydra.C.DB().GetRegularDB().Execute(sql.UpdateTradeDeliveryForDeliveryingFailed, map[string]interface{}{
			sql.FieldDeliveryID: deliveryID,
			sql.FieldResultCode: resultCode,
			sql.FieldReturnMsg:  returnMsg,
		})
	default: //非失败都纳入成功处理，等待后续通知、查询或人工审核
		row, err = hydra.C.DB().GetRegularDB().Execute(sql.UpdateTradeDeliveryForDeliveryingSuccess, map[string]interface{}{
			sql.FieldDeliveryID:   deliveryID,
			sql.FieldResultCode:   resultCode,
			sql.FieldReturnMsg:    returnMsg,
			sql.FieldCostDiscount: discount.String(),
		})
	}
	if err != nil {
		return err
	}
	if row == 0 {
		return errs.NewError(http.StatusNoContent, "无须发货处理(save.start)")
	}
	return nil
}

//StartNow 开始发货
func StartNow(deliveryID string) (*TradeDelivery, error) {

	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}

	//修改发货记录为正在发货
	row, err := db.Execute(sql.UpdateTradeDeliveryForStart, map[string]interface{}{
		sql.FieldDeliveryID: deliveryID,
	})
	if err != nil {
		db.Rollback()
		return nil, err
	}
	if row == 0 {
		db.Rollback()
		return nil, errs.NewError(http.StatusNoContent, "无须发货处理(start.now)")
	}

	//查询发货记录
	data, err := db.Query(sql.SelectTradeDeliveryForStart, map[string]interface{}{
		sql.FieldDeliveryID: deliveryID,
	})
	if err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	delviery := &TradeDelivery{}
	err = data.Get(0).ToAnyStruct(delviery)
	return delviery, err
}

//TradeDelivery 订单发货表
type TradeDelivery struct {

	//DeliveryID 发货编号
	deliveryID string `json:"delivery_id"`

	//OrderID 订单编号
	orderID string `json:"order_id"`

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
