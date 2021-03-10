package deliverys

import (
	"fmt"

	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
)

//GetDealCode 根据错误码获取供货商业务处理结果
func GetDealCode(deliveryID string, source enums.ResultSource, code string) (enums.FlowStatus, string) {
	data, err := hydra.C.DB().GetRegularDB().Query(queryDealCode, map[string]interface{}{
		fields.FieldDeliveryID: deliveryID,
		fields.FieldCategory:   source,
		fields.FieldErrorCode:  code,
	})
	if err != nil || data.Len() == 0 {
		return enums.Unknown, fmt.Sprintf("错误码(delivery_id:%s,source:%s,code:%s)未配置", deliveryID, source, code)
	}
	dealCode := data.Get(0).GetInt(fields.FieldDealCode, -1)
	switch dealCode {
	case int(enums.Success):
		return enums.Success, data.Get(0).GetString(fields.FieldErrorDesc)
	case int(enums.Failed):
		return enums.Failed, data.Get(0).GetString(fields.FieldErrorDesc)
	default:
		return enums.Unknown, data.Get(0).GetString(fields.FieldErrorDesc)
	}
}
