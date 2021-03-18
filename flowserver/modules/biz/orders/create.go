package orders

import (
	"errors"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/dbs"
	"github.com/emshop/ots/flowserver/modules/pkgs"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Create 创建交易订单
func Create(merNo string, merOrderNo string, merProductID int, accountName string, num int, notifyURL string) (types.IXMap, bool, error) {

	//1. 查询订单记录
	order, err := Query(merNo, merOrderNo)
	if err != nil || order.Len() > 0 {
		return order, false, err
	}

	//2. 创建订单
	orderID, err := pkgs.GetOrderID()
	if err != nil {
		return nil, false, err
	}
	input := types.NewXMap()
	input.SetValue(fields.FieldMerNo, merNo)
	input.SetValue(fields.FieldMerOrderNo, merOrderNo)
	input.SetValue(fields.FieldAccountName, accountName)
	input.SetValue(fields.FieldOrderID, orderID)
	input.SetValue(fields.FieldNum, num)
	input.SetValue(fields.FieldNotifyUrl, notifyURL)
	input.SetValue(fields.FieldMerProductID, merProductID)

	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return nil, false, err
	}
	_, err = dbs.Executes(db, input, createOrder...)
	if errors.Is(err, xerr.ErrNOTEXISTS) {
		db.Rollback()
		return nil, false, errs.NewError(int(enums.CodeParamNoSetting), "商品不存在或未启用")
	}
	if err != nil {
		db.Rollback()
		return nil, false, err
	}
	db.Commit()

	// 3. 查询订单
	order, _ = Query(merNo, merOrderNo)
	if order.Len() == 0 {
		return nil, false, errs.NewError(int(enums.CodeUnknownErr), "订单未创建成功")
	}

	//4. 检查是否是本次创建的订单
	if input.GetString(fields.FieldOrderID) == order.GetString(fields.FieldOrderID) {
		return order, true, nil
	}
	return order, false, nil
}
