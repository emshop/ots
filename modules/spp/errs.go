package spp

import (
	"github.com/emshop/ots/modules/const/enums"
	"github.com/emshop/ots/modules/const/sql"
	"github.com/micro-plat/hydra"
)

//GetDealCode 根据错误码获取供货商业务处理结果
func GetDealCode(sppNO string, source enums.ResultSource, plID int, code string) enums.FlowStatus {
	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectSupplierErrorCode, map[string]interface{}{
		sql.FieldSppNo:     sppNO,
		sql.FieldCategory:  source,
		sql.FieldErrorCode: code,
		sql.FieldPlID:      plID,
	})
	if err != nil || data.Len() == 0 {
		return enums.Unknown
	}
	dealCode := data.Get(0).GetInt(sql.FieldDealCode, -1)
	switch dealCode {
	case int(enums.Success):
		return enums.Success
	case int(enums.Failed):
		return enums.Failed
	default:
		return enums.Unknown
	}
}
