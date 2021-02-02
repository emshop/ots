package api

import (
	"github.com/micro-plat/hydra"
	"github.com/emshop/ots/mgrserver/api/services/audit"
	"github.com/emshop/ots/mgrserver/api/services/merchant"
	"github.com/emshop/ots/mgrserver/api/services/product"
	"github.com/emshop/ots/mgrserver/api/services/refund"
	"github.com/emshop/ots/mgrserver/api/services/supplier"
	"github.com/emshop/ots/mgrserver/api/services/system"
	"github.com/emshop/ots/mgrserver/api/services/trade"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	hydra.OnReady(func() {
		hydra.S.Web("/supplier/shelf", supplier.NewSupplierShelfHandler())
		hydra.S.Web("/trade/delivery", trade.NewTradeDeliveryHandler())
		hydra.S.Web("/trade/order", trade.NewTradeOrderHandler())
		hydra.S.Web("/merchant/product", merchant.NewMerchantProductHandler())
		hydra.S.Web("/audit/info", audit.NewAuditInfoHandler())
		hydra.S.Web("/product/flow", product.NewProductFlowHandler())
		hydra.S.Web("/product/line", product.NewProductLineHandler())
		hydra.S.Web("/refund/apply", refund.NewRefundApplyHandler())
		hydra.S.Web("/supplier/product", supplier.NewSupplierProductHandler())
		hydra.S.Web("/system/enums", system.NewSystemEnumsHandler())
		hydra.S.Web("/merchant/info", merchant.NewMerchantInfoHandler())
		hydra.S.Web("/merchant/shelf", merchant.NewMerchantShelfHandler())
		hydra.S.Web("/supplier/ecode", supplier.NewSupplierEcodeHandler())
		hydra.S.Web("/supplier/info", supplier.NewSupplierInfoHandler())
	})
}

