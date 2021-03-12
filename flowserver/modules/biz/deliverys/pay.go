package deliverys

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/dbs"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Pay 处理订单支付
func Pay(deliveryID string) error {
	input := types.XMap{
		fields.FieldDeliveryID: deliveryID,
	}

	//启动事务进行支付处理
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}

	//修改支付状态
	order, err := dbs.Executes(db, input, updateDeliveryForPaying...)
	if errors.Is(err, xerr.ErrNOTEXISTS) {
		db.Rollback()
		return errs.NewErrorf(http.StatusNoContent, "发货编号(%s)无须进行发货支付", deliveryID)
	}
	if err != nil {
		db.Rollback()
		return fmt.Errorf("发货编号(%s)发货支付出错%w", deliveryID, err)
	}

	//获取账户信息
	account := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountSupplier))

	//处理交易扣款
	rs, err := account.DeductAmount(db,
		order.GetString(sql.FieldSppNo),
		order.GetString(sql.FieldDeliveryID),
		beanpay.AccountTradeType,
		order.GetFloat64(sql.FieldCostAmount),
		"供货商扣款")
	if err != nil {
		db.Rollback()
		return fmt.Errorf("发货支付失败(%s)%v", order.GetString(sql.FieldSppNo), err)
	}
	if rs.GetCode() != beanpay.Success {
		db.Rollback()
		return errs.NewErrorf(rs.GetCode(), "%s扣款失败%s", deliveryID, rs.GetCode())
	}
	db.Commit()
	return nil
}
