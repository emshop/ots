package order

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/utility"
)

func Mock(ctx hydra.IContext) interface{} {
	ctx.Request().SetValue("mer_order_no", utility.GetGUID())
	return Request(ctx)
}
