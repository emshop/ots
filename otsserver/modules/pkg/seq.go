package pkg

import (
	"fmt"
	"time"

	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
)

//GetOrderID 获取订单号
func GetOrderID() (string, error) {
	id, err := getSEQID("", "order_id")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("RO%s%s", time.Now().Format("20060102"), id), nil
}

//GetDeliveryID 获取发货编号
func GetDeliveryID() (string, error) {
	id, err := getSEQID("", "delivery_id")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", time.Now().Format("060102"), id), nil
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
