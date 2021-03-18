
package scheme

import (
	"github.com/micro-plat/hydra"
	_ "github.com/go-sql-driver/mysql"
)
		
func init() {
	//注册服务包
	hydra.OnReadying(func() error {
		hydra.Installer.DB.AddSQL(
		ots_merchant_info,
		ots_merchant_shelf,
		ots_merchant_product,
		ots_supplier_info,
		ots_supplier_shelf,
		ots_supplier_product,
		ots_supplier_ecode,
		ots_trade_order,
		ots_trade_delivery,
		ots_refund_apply,
		ots_notify_info,
		ots_audit_info,
		ots_product_line,
		ots_product_flow,
		SEQ_IDS,
		)
		return nil
	}) 
}
