package notices

import (
	"errors"
	"net/http"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Save 保存通知结果
func Save(orderID string, status enums.FlowStatus, msg string) (order types.IXMap, err error) {
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}

	input := types.XMap{
		fields.FieldOrderID:   orderID,
		fields.FieldNotifyMsg: types.GetString(msg, types.DecodeString(status, enums.Success, "通知成功", "结果未知")),
	}

	//执行通知结果保存
	var orders types.XMaps
	switch status {
	case enums.Success:
		orders, err = db.ExecuteBatch(updateNotifyForSuccess, input)
	default:
		orders, err = db.ExecuteBatch(updateNotifyInfoForUnknown, input)
	}
	if err == nil {
		db.Commit()
		return orders.Get(0), nil
	}

	db.Rollback()
	if errors.Is(err, errs.ErrNotExist) {
		return nil, errs.NewErrorf(http.StatusNoContent, "订单(%s)无须处理通知%w", orderID, errs.ErrNotExist)
	}
	return nil, err
}
