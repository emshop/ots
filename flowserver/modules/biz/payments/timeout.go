package payments

import (
	"errors"

	"github.com/emshop/ots/flowserver/modules/const/fields"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//timeout 处理订单支付超时
func timeout(orderID string) (bool, error) {

	//构建事务
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return false, err
	}
	//执行SQL语句
	_, err = db.ExecuteBatch(dealTimeoutOrder, types.XMap{
		fields.FieldOrderID: orderID,
	},
	)
	if errors.Is(err, errs.ErrNotExist) {
		db.Rollback()
		return false, nil
	}
	if err != nil {
		db.Rollback()
		return false, err
	}
	db.Commit()
	return true, nil

}
