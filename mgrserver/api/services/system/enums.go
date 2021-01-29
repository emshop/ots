
package system

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/dds/dds"
)

//EnumsHandler 枚举数据查询服务
type EnumsHandler struct {
}

//NewEnumsHandler 枚举数据查询服务
func NewEnumsHandler() *EnumsHandler {
	return &EnumsHandler{}
}

//QueryHandle 枚举数据查询服务
func (o *EnumsHandler) QueryHandle(ctx hydra.IContext) interface{} {

	//根据传入的枚举类型获取数据
	tp := ctx.Request().GetString("dic_type")
	if tp != "" {
		var items types.XMaps
		var err error
		if _, ok := enumsMap[tp]; !ok {
			items, err = dds.GetEnums(ctx, ctx.Request().GetMap())
		} else {
			items, err = hydra.C.DB().GetRegularDB().Query(enumsMap[tp], ctx.Request().GetMap())
		}
		if err != nil {
			return err
		}
		return items
	}

	//查询所有枚举数据
	list := types.XMaps{}
	for _, sql := range enumsMap {
		items, err := hydra.C.DB().GetRegularDB().Query(sql, ctx.Request().GetMap())
		if err != nil {
			return err
		}
		list = append(list, items...)
	}
	return list
}

var enumsMap = map[string]string{
"merchant_info":`select 'merchant_info' type , t.mer_no value , t.mer_name name  from ots_merchant_info t `,
"merchant_shelf":`select 'merchant_shelf' type , t.mer_shelf_id value , t.mer_shelf_name name  from ots_merchant_shelf t `,
"supplier_info":`select 'supplier_info' type , t.spp_no value , t.spp_name name  from ots_supplier_info t `,
"supplier_shelf":`select 'supplier_shelf' type , t.spp_shelf_id value , t.spp_shelf_name name  from ots_supplier_shelf t `,
"product_line":`select 'product_line' type , t.pl_id value , t.pl_name name  from ots_product_line t `,
"product_flow":`select 'product_flow' type , t.flow_id value , t.flow_name name , t.tag_name name  from ots_product_flow t `,
}