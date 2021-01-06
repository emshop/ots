package mers

import (
	"github.com/emshop/ots/modules/const/db/biz"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

//GetShelfByID 获取货架信息
func GetShelfByID(merShelfID int) (types.XMap, error) {
	data, err := hydra.C.DB().GetRegularDB().Query(biz.SelectMerchantShelf, map[string]interface{}{
		biz.FieldTradeOrderMerShelfID: merShelfID,
	})
	if err != nil {
		return nil, err
	}
	return data.Get(0), nil
}
