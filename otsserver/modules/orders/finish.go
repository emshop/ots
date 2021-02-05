package orders

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

//Finish 完结订单状态
func Finish(orderID string) error {

	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectTradeOrderByOrderID, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		return err
	}

	var stat enums.OrderStatus = enums.OrderStatus(data.Get(0).GetInt(sql.FieldOrderStatus))
	if stat == enums.OrderSuccess || stat == enums.OrderFailed {
		return errs.NewError(http.StatusNoContent, "订单无须处理")
	}

	row, err := hydra.C.DB().GetRegularDB().Execute(sql.UpdateTradeOrderForClose, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		return err
	}
	if row == 0 {
		return errs.NewError(http.StatusAccepted, "订单暂时无法处理")
	}
	return nil

}
