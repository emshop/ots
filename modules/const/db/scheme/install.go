
package scheme

import (
	"github.com/micro-plat/hydra"
)
		
func init() {
	//注册服务包
	hydra.DBCli.OnStarting(func(c hydra.ICli) error {
		hydra.Installer.DB.AddSQL(
		ots_merchant_info,
		ots_merchant_shelf,
		ots_merchant_product,
		ots_supplier_info,
		ots_supplier_shelf,
		ots_supplier_product,
		ots_supplier_error_code,
		ots_trade_order,
		ots_trade_delivery,
		ots_trade_lifetime,
		ots_refund_apply,
		ots_notify_info,
		ots_audit_info,
		ots_product_line,
		ots_product_flow,
		dds_dictionary_info,
		ots_canton_info,
		)
		return nil
	})

}
