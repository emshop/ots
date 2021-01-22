package mers

import (
	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//GetShelfByID 获取货架信息
func GetShelfByID(merNo string, merShelfID int, num int) (types.XMap, error) {
	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectMerchantShelf, map[string]interface{}{
		sql.FieldMerShelfID: merShelfID,
		sql.FieldMerNo:      merNo,
	})
	if err != nil || data.Len() == 0 {
		return nil, errs.NewErrorf(int(enums.CodeParamNoSetting), "商户或货架未配置 %s", err)
	}
	if data.Get(0).GetInt(sql.FieldLimitCount) < num {
		return nil, errs.NewError(int(enums.CodeOverLimit), "购买数量超过限制")
	}
	return data.Get(0), nil
}
