package mers

import (
	"github.com/emshop/ots/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//GetShelfByID 获取货架信息
func GetShelfByID(merShelfID int) (types.XMap, error) {
	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectMerchantShelf, map[string]interface{}{
		sql.FieldMerShelfID: merShelfID,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}
