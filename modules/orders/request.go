package orders

import (
	"fmt"

	"github.com/emshop/ots/modules/const/enums"
	"github.com/emshop/ots/modules/const/sql"
	"github.com/emshop/ots/modules/pkg"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Create 创建交易订单
func Create(merNo string, merOrderNo string, accountName string, num int, notifyURL string, shelf types.XMap, product types.XMap) (types.XMap, bool, error) {

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
	input.SetValue(sql.FieldMerNo, merNo)
	input.SetValue(sql.FieldMerOrderNo, merOrderNo)
	input.SetValue(sql.FieldAccountName, accountName)
	input.SetValue(sql.FieldOrderID, orderID)
	input.SetValue(sql.FieldNum, num)
	input.SetValue(sql.FieldNotifyUrl, notifyURL)

	//开启事务
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return nil, false, err
	}

	//添加通知记录、订单记录
	if notifyURL != "" {
		_, err = db.Execute(sql.InsertNotifyInfo, input)
	}
	if err == nil {
		//添加订单
		_, err = db.Execute(sql.InsertTradeOrder, input)
	}

	if err != nil {
		db.Rollback()
		return nil, false, errs.NewError(int(enums.CodeUnknownErr), fmt.Errorf("无法生成订单%w", err))
	}
	db.Commit()

	// 查询订单
	order, _ = Query(merNo, merOrderNo)
	if order.Len() == 0 {
		return nil, false, errs.NewError(int(enums.CodeUnknownErr), "订单未创建成功")
	}

	//检查是否是本次创建的订单
	if input.GetString(sql.FieldOrderID) == order.GetString(sql.FieldOrderID) {
		return order, true, nil
	}
	return order, false, nil

}
