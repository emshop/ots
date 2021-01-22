package spp

import (
	"fmt"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
)

//GetDealCode 根据错误码获取供货商业务处理结果
func GetDealCode(sppNO string, source enums.ResultSource, plID int, code string) (enums.FlowStatus, string) {
	data, err := hydra.C.DB().GetRegularDB().Query(sql.SelectSupplierErrorCode, map[string]interface{}{
		sql.FieldSppNo:     sppNO,
		sql.FieldCategory:  source,
		sql.FieldErrorCode: code,
		sql.FieldPlID:      plID,
	})
	if err != nil || data.Len() == 0 {
		return enums.Unknown, fmt.Sprintf("错误码(sppNO:%s,category:%s,pl_id:%d,code:%s)未配置 %v", sppNO, source, plID, code, err)
	}
	dealCode := data.Get(0).GetInt(sql.FieldDealCode, -1)
	switch dealCode {
	case int(enums.Success):
		return enums.Success, data.Get(0).GetString(sql.FieldErrorDesc)
	case int(enums.Failed):
		return enums.Failed, data.Get(0).GetString(sql.FieldErrorDesc)
	default:
		return enums.Unknown, data.Get(0).GetString(sql.FieldErrorDesc)
	}
}
