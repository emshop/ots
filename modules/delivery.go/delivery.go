package delivery

import (
	"net/http"

	"github.com/emshop/ots/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Create 构建发货信息
func Create(order types.XMap, shelf types.XMap, prod types.XMap) (bool, error) {
	input := types.NewXMapByMap(order)
	input.Merge(shelf)
	input.Merge(prod)

	//启动事务进行支付处理
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return false, err
	}

	//修改改订单的绑定金额
	row, err := db.Execute(sql.UpdateTradeOrderForBind, input)
	if err != nil || row == 0 {
		return false, errs.NewError(http.StatusAccepted, err)
	}

	//添加发货记录
	db.Execute(sql.InsertTradeDelivery, input)
	if err != nil || row == 0 {
		return false, errs.NewError(http.StatusAccepted, err)
	}
	db.Commit()
	return true, nil
}
