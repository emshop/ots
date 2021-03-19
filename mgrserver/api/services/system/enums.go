package system

func init() {
	enumsMap["account_info"] = `select 'account_info' type , t.account_id value ,CONCAT(t.account_name," (",t.groupx,")") name  from beanpay_account_info t order by t.account_name asc,t.groupx asc`

}
