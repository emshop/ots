package deliverys

import (
	"errors"
	"net/http"

	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/dbs"
	"github.com/emshop/ots/flowserver/modules/pkgs"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Bind 绑定订单
func Bind(orderID string) (string, error) {

	//检查订单是否需要绑定
	db := hydra.C.DB().GetRegularDB()
	input := types.XMap{fields.FieldOrderID: orderID}
	orders, err := db.Query(checkOrderForBind, input)
	if err != nil {
		return "", err
	}
	if orders.Len() == 0 {
		return "", errs.NewStop(204, "订单无须进行绑定")
	}

	//查询订单支持的上游产品
	products, err := dbs.Query(db, orders.Get(0), querySppProducts...)
	if errors.Is(err, xerr.ErrNOTEXISTS) || products.Len() == 0 {
		return "", errs.NewErrorf(http.StatusAccepted, "暂无上游产品可绑定%s%v", orderID, err)
	}
	if err != nil {
		return "", err
	}

	//根据商品创建发货记录
	deliveryID, err := pkgs.GetDeliveryID()
	if err != nil {
		return "", err
	}
	for i, product := range products {

		//处理输入参数
		product.SetValue(fields.FieldDeliveryID, deliveryID)
		product.Merge(orders.Get(0))

		//创建发货记录
		xdb, err := hydra.C.DB().GetRegularDB().Begin()
		if err != nil {
			return "", err
		}
		_, err = dbs.Executes(xdb, product, binds...)
		if err != nil {
			xdb.Rollback()
			if i == len(products)-1 {
				return "", err
			}
			continue
		}
		xdb.Commit()
		return deliveryID, nil
	}
	return "", errs.NewErrorf(http.StatusAccepted, "无可用的上游可绑定%s", orderID)
}