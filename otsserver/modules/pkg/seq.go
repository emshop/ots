package pkg

import (
	"fmt"

	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
)

//GetOrderID 获取订单号
func GetOrderID() (string, error) {
	return getSEQID("", "order_id")
}

//GetDeliveryID 获取发货编号
func GetDeliveryID() (string, error) {
	return getSEQID("", "delivery_id")
}

//getSEQID 获取流水号
func getSEQID(prefix string, tpName string) (string, error) {
	id, _, err := hydra.C.DB().GetRegularDB().Executes(sql.InsertSEQID, map[string]interface{}{
		sql.FieldName: tpName,
	})
	if err != nil || id == 0 {
		return "", fmt.Errorf("获取流水号失败:%w %d", err, id)
	}
	return fmt.Sprintf("%s%d", prefix, id), nil
}
