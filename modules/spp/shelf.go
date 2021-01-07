package spp

import (
	"github.com/emshop/ots/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//GetShelfByID 获取spp货架信息
func GetShelfByID(sppShelfID int) (types.XMap, error) {
	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectSupplierShelf, map[string]interface{}{
		sql.FieldSppShelfID: sppShelfID,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}
