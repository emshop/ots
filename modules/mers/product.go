package mers

import (
	"github.com/emshop/ots/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//GetProductByID 获取产品信息
func GetProductByID(merProductID int) (types.XMap, error) {
	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectMerchantProduct, map[string]interface{}{
		sql.FieldMerProductID: merProductID,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}
