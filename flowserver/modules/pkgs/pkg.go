package pkgs

import (
	"fmt"
	"strings"
	"time"

	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
)

//GetOrderID 获取订单号
func GetOrderID() (string, error) {
	id, err := getSEQID("", "order_id")
	if err != nil {
		return "", err
	}
	if len(id) > 8 {
		return fmt.Sprintf("R%s%s", time.Now().Format("060102"), id), nil
	}
	return fmt.Sprintf("R%s%s%s", time.Now().Format("060102"), strings.Repeat("0", 8-len(id)), id), nil
}

//GetDeliveryID 获取发货编号
func GetDeliveryID() (string, error) {
	id, err := getSEQID("", "delivery_id")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", time.Now().Format("060102"), id), nil
}

//GetRepBatchID 获取后补批次号
func GetRepBatchID() (string, error) {
	return getSEQID("", "replenish")
}

//getSEQID 获取流水号
func getSEQID(prefix string, tpName string) (string, error) {
	id, _, err := hydra.C.DB().GetRegularDB().Executes(InsertSEQID, map[string]interface{}{
		fields.FieldName: tpName,
	})
	if err != nil || id == 0 {
		return "", fmt.Errorf("获取流水号失败:%w %d", err, id)
	}
	return fmt.Sprintf("%s%d", prefix, id), nil
}

//InsertSEQID 添加seqID
const InsertSEQID = `insert into seq_ids(name)values(@name)`
