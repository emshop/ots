package deliverys

import (
	"errors"
	"net/http"

	"github.com/emshop/ots/flowserver/modules/const/fields"

	"github.com/emshop/ots/flowserver/modules/pkgs"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//QueryReplenish 查询后补数据
func QueryReplenish() (t types.XMaps, err error) {
	//获取批次号
	batchID, err := pkgs.GetRepBatchID()
	if err != nil {
		return nil, err
	}

	//启动事务获取数据
	xdb, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}
	t, err = xdb.ExecuteBatch(reps, types.XMap{fields.FieldBatchID: batchID})
	if err == nil {
		xdb.Commit()
		return t, nil
	}
	xdb.Rollback()
	if errors.Is(err, errs.ErrNotExist) {
		return nil, errs.NewError(http.StatusNoContent, "没有需要后补的发货数据")
	}
	return nil, err
}
