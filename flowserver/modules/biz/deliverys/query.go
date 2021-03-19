package deliverys

import (
	"errors"
	"net/http"

	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/dbs"
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
	t, err = dbs.Executes(xdb, types.XMap{fields.FieldDeliveryID: deliveryID}, deliveryQuery...)
	if err == nil {
		xdb.Commit()
		return t, nil
	}
	xdb.Rollback()
	if errors.Is(err, xerr.ErrNOTEXISTS) {
		return nil, errs.NewErrorf(http.StatusNoContent, "发货记录无须进行查询处理%w", xerr.ErrNOTEXISTS)
	}
	return nil, err
}
