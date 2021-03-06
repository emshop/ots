package spp

import (
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//GetProducts 根据产品线，省，市获取上游产品信息
func GetProducts(plid int, brandNO string, provinceNO string, cityNO string, canSplitOrder int, face int, discount types.Decimal) (types.XMaps, error) {
	input := map[string]interface{}{
		sql.FieldPlID:          plid,
		sql.FieldBrandNo:       brandNO,
		sql.FieldProvinceNo:    provinceNO,
		sql.FieldCityNo:        cityNO,
		sql.FieldFace:          face,
		sql.FieldCostDiscount:  discount,
		sql.FieldCanSplitOrder: canSplitOrder,
	}
	rows, err := hydra.C.DB().GetRegularDB().Query(sql.SelectSupplierProduct, input)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
