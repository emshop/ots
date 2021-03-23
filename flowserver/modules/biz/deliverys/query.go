package deliverys

import (
	"errors"
	"net/http"

	"github.com/emshop/ots/flowserver/modules/const/fields"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//GetQuery 查询发货
func GetQuery(deliveryID string) (t types.IXMap, err error) {
	//启动事务获取数据
	xdb, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}
	tx, err := xdb.ExecuteBatch(deliveryQuery, map[string]interface{}{fields.FieldDeliveryID: deliveryID})
	if err == nil {
		xdb.Commit()
		return tx.Get(0), nil
	}
	xdb.Rollback()
	if errors.Is(err, errs.ErrNotExist) {
		return nil, errs.NewErrorf(http.StatusNoContent, "发货记录无须进行查询处理%w", errs.ErrNotExist)
	}
	return nil, err
}
