package delivery

import (
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//Get 获取发货信息
func Get(deliveryID string) (types.XMap, error) {
	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectTradeDelivery, map[string]interface{}{
		sql.FieldDeliveryID: deliveryID,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}
