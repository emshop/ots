package notices

import (
	"errors"
	"net/http"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/dbs"
	"github.com/emshop/ots/otsserver/modules/const/sql"
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
		sql.FieldOrderID:   orderID,
		sql.FieldNotifyMsg: types.GetString(msg, types.DecodeString(status, enums.Success, "充值成功", "订单无须处理")),
	}

	//执行通知结果保存
	switch status {
	case enums.Success:
		order, err = dbs.Executes(db, input, updateNotifyForSuccess...)

	default:
		order, err = dbs.Executes(db, input, updateNotifyInfoForUnknown)
	}
	if errors.Is(err, xerr.ErrNOTEXISTS) {
		db.Rollback()
		return nil, errs.NewStopf(http.StatusNoContent, "订单(%s)无须处理通知", orderID)
	}
	if err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return order, nil
}
