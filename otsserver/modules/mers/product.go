package mers

import (
	"fmt"

	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//GetProductByID 获取产品信息
func GetProductByID(merNo string, merProductID int) (types.XMap, error) {
	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectMerchantProduct, map[string]interface{}{
		sql.FieldMerProductID: merProductID,
		sql.FieldMerNo:        merNo,
	})
	if err != nil {
		return nil, err
	}
	if data.Get(0).GetInt(sql.FieldPlID) == 0 {
		return nil, fmt.Errorf("商户产品配置错误")
	}
	return data.Get(0), nil
}
