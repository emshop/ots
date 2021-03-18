package notices

import (
	"errors"
	"net/http"

	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/dbs"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Start 开始启动通知
func Start(orderID string) (*NotifyInfo, error) {

	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}

	//修改无须通知记录
	data, err := dbs.Executes(db, types.XMap{fields.FieldOrderID: orderID}, updateNoNeedNotices...)
	if err == nil {
		db.Commit()
		return nil, errs.NewErrorf(http.StatusNoContent, "订单(%s)无须通知%w", orderID, xerr.ErrNOTEXISTS)
	}
	db.Rollback()
	if err != nil && !errors.Is(err, xerr.ErrNOTEXISTS) {
		return nil, err
	}

	//处理正常待通知
	db, err = hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}
	//需要通知则按正常流程修改状态
	data, err = dbs.Executes(db, types.XMap{fields.FieldOrderID: orderID}, startNotices...)
	if errors.Is(err, xerr.ErrNOTEXISTS) {
		db.Rollback()
		return nil, errs.NewErrorf(http.StatusNoContent, "订单(%s)已通知%w", orderID, xerr.ErrNOTEXISTS)
	}
	if err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	n := &NotifyInfo{}
	err = data.ToAnyStruct(n)
	return n, err
}

//NotifyInfo 订单通知表
type NotifyInfo struct {

	//OrderID 订单编号
	OrderID string `json:"order_id"`

	//MerNo 商户编号
	MerNo string `json:"mer_no"`

	//MerOrderNo 商户订单编号
	MerOrderNo string `json:"mer_order_no"`

	//AccountName 用户账户信息
	AccountName string `json:"account_name"`

	//SellDiscount 销售折扣
	SellDiscount types.Decimal `json:"sell_discount"`

	//NotifyURL 通知地址
	NotifyURL string `json:"notify_url"`

	//订单状态
	Status int `json:"status"`
}
