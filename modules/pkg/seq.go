package pkg

import (
	"fmt"

	"github.com/emshop/ots/modules/const/db/biz"
	"github.com/micro-plat/hydra"
)

//GetOrderID 获取订单号
func GetOrderID() (string, error) {
	return getSEQID("", "ORDER_ID")
}

//getSEQID 获取流水号
func getSEQID(prefix string, tpName string) (string, error) {
	id, _, err := hydra.C.DB().GetRegularDB().Executes(biz.InsertSEQID, map[string]interface{}{
		biz.FieldName: tpName,
	})
	if err != nil || id == 0 {
		return "", fmt.Errorf("获取流水号失败:%w %d", err, id)
	}
	return fmt.Sprintf("%s%d", prefix, id), nil
}
