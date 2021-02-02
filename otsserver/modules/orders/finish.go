package orders

import (
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
)

//Finish 完结订单状态
func Finish(orderID string) error {

	row, err := hydra.C.DB().GetRegularDB().Scalar(sql.UpdateTradeOrderForClose, map[string]interface{}{
		sql.FieldOrderID: orderID,
	})
	if err != nil {
		return err
	}
	if row == 0 {
		return errs.NewError(http.StatusNoContent, "订单无须处理")
	}
	return nil

}
