package deliverys

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/emshop/ots/flowserver/modules/const/xerr"
	"github.com/emshop/ots/flowserver/modules/pkgs"
	"github.com/emshop/ots/mgrserver/api/modules/const/field"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//Bind 绑定订单
func Bind(orderID string) (string, error) {

	//----------------处理订单超时-------------------------
	xdb, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return "", err
	}
	input := types.XMap{fields.FieldOrderID: orderID}
	_, err = xdb.ExecuteBatch(dealOrderTimeout, input.ToMap())
	if err == nil {
		xdb.Commit()
		return "", errs.NewErrorf(204, "订单(%s)已超时%w", orderID, xerr.ErrOrderTimeout)
	}
	xdb.Rollback()
	if err != nil && !errors.Is(err, errs.ErrNotExist) {
		return "", err
	}

	//获取绑定订单
	order, err := getBindOrder(input)
	if err != nil {
		return "", err
	}

	//绑定商品
	return bindByProducts(order)

}

func getBindOrder(input types.IXMap) (types.IXMap, error) {
	//----------------处理订单正常绑定------------------------
	db := hydra.C.DB().GetRegularDB()
	orders, err := db.Query(checkOrderForBind, input.ToMap())
	if err != nil {
		return nil, err
	}
	if orders.Len() == 0 {
		return nil, errs.NewError(204, fmt.Errorf("订单(%s)无须进行绑定(单品),%w", input.GetString(fields.FieldOrderID), errs.ErrNotExist))
	}

	//检查是否是复合商品
	if orders.Get(0).GetInt(field.FieldPlType) != 0 {
		orders, err = db.Query(checkOrderForBindPackage, input.ToMap())
		if err != nil {
			return nil, err
		}
		if orders.Len() == 0 {
			return nil, errs.NewError(204, fmt.Errorf("订单(%s)无须进行绑定(复品),%w", input.GetString(fields.FieldOrderID), errs.ErrNotExist))
		}
	}
	return orders.Get(0), nil

}

func bindByProducts(order types.IXMap) (string, error) {

	//2. 查询订单支持的上游产品
	db := hydra.C.DB().GetRegularDB()
	products, err := db.ExecuteBatch(querySppProducts, order.ToMap())
	if errors.Is(err, errs.ErrNotExist) || products.Len() == 0 {
		return "", errs.NewErrorf(http.StatusAccepted, "暂无上游产品可绑定%s", order.GetString(fields.FieldOrderID))
	}
	if err != nil {
		return "", err
	}

	//3. 根据商品创建发货记录
	deliveryID, err := pkgs.GetDeliveryID()
	if err != nil {
		return "", err
	}
	for i, product := range products.Maps() {

		//处理输入参数
		product.SetValue(fields.FieldDeliveryID, deliveryID)
		product.Merge(order)

		//创建发货记录
		xdb, err := hydra.C.DB().GetRegularDB().Begin()
		if err != nil {
			return "", err
		}
		_, err = xdb.ExecuteBatch(binds, product.ToMap())
		if err != nil {
			xdb.Rollback()
			if i == products.Len()-1 {
				return "", err
			}
			continue
		}
		xdb.Commit()
		return deliveryID, nil
	}
	return "", errs.NewErrorf(http.StatusAccepted, "无可用的上游可绑定%s", order.GetString(fields.FieldOrderID))
}
