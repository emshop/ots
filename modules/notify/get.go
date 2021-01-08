package notify

import (
	"fmt"
	"net/http"

	"github.com/emshop/ots/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

func Get(orderID int64) (*NotifyInfo, error) {
	//修改支付状态
	row, err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateNotifyInfoForStart, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errs.NewError(http.StatusNoContent, "通知记录无须处理")
	}
	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectNotifyInfoForStart, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		return nil, err
	}
	if data.Len() == 0 {
		return nil, fmt.Errorf("未查询到订单通知记录")
	}
	n := &NotifyInfo{}
	err = data.Get(0).ToAnyStruct(n)
	return n, err
}

//NotifyInfo 订单通知表
type NotifyInfo struct {

	//OrderID 订单编号
	OrderID int64 `json:"order_id"`

	//MerNo 商户编号
	MerNo string `json:"mer_no"`

	//MerOrderNo 商户订单编号
	MerOrderNo string `json:"mer_order_no"`

	//AccountName 用户账户信息
	AccountName string `json:"account_name"`

	//InvoiceType 开票方式（1.不开发票）
	InvoiceType int `json:"invoice_type"`

	//SellDiscount 销售折扣
	SellDiscount types.Decimal `json:"sell_discount"`

	//SellAmount 总销售金额
	SellAmount types.Decimal `json:"sell_amount"`

	//NotifyURL 通知地址
	NotifyURL string `json:"notify_url"`
}
