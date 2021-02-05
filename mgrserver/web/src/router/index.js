
import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);
export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'menus',
      component: () => import('../pages/system/menus.vue'),
      children:[
				{
					path: 'merchant/product',
					name: 'MerchantProduct',
					component: () => import('../pages/merchant/merchant.product.list.vue')
				},
				{
					path: 'merchant/product/detail',
					name: 'MerchantProductDetail',
					component: () => import('../pages/merchant/merchant.product.detail.vue')
				},
				{
					path: 'product/line',
					name: 'ProductLine',
					component: () => import('../pages/product/product.line.list.vue')
				},
				{
					path: 'product/line/detail',
					name: 'ProductLineDetail',
					component: () => import('../pages/product/product.line.detail.vue')
				},
				{
					path: 'trade/order',
					name: 'TradeOrder',
					component: () => import('../pages/trade/trade.order.list.vue')
				},
				{
					path: 'trade/order/detail',
					name: 'TradeOrderDetail',
					component: () => import('../pages/trade/trade.order.detail.vue')
				},
				{
					path: 'audit/info',
					name: 'AuditInfo',
					component: () => import('../pages/audit/audit.info.list.vue')
				},
				{
					path: 'audit/info/detail',
					name: 'AuditInfoDetail',
					component: () => import('../pages/audit/audit.info.detail.vue')
				},
				{
					path: 'product/flow',
					name: 'ProductFlow',
					component: () => import('../pages/product/product.flow.list.vue')
				},
				{
					path: 'product/flow/detail',
					name: 'ProductFlowDetail',
					component: () => import('../pages/product/product.flow.detail.vue')
				},
				{
					path: 'refund/apply',
					name: 'RefundApply',
					component: () => import('../pages/refund/refund.apply.list.vue')
				},
				{
					path: 'refund/apply/detail',
					name: 'RefundApplyDetail',
					component: () => import('../pages/refund/refund.apply.detail.vue')
				},
				{
					path: 'supplier/ecode',
					name: 'SupplierEcode',
					component: () => import('../pages/supplier/supplier.ecode.list.vue')
				},
				{
					path: 'supplier/ecode/detail',
					name: 'SupplierEcodeDetail',
					component: () => import('../pages/supplier/supplier.ecode.detail.vue')
				},
				{
					path: 'supplier/info',
					name: 'SupplierInfo',
					component: () => import('../pages/supplier/supplier.info.list.vue')
				},
				{
					path: 'supplier/info/detail',
					name: 'SupplierInfoDetail',
					component: () => import('../pages/supplier/supplier.info.detail.vue')
				},
				{
					path: 'supplier/product',
					name: 'SupplierProduct',
					component: () => import('../pages/supplier/supplier.product.list.vue')
				},
				{
					path: 'supplier/product/detail',
					name: 'SupplierProductDetail',
					component: () => import('../pages/supplier/supplier.product.detail.vue')
				},
				{
					path: 'supplier/shelf',
					name: 'SupplierShelf',
					component: () => import('../pages/supplier/supplier.shelf.list.vue')
				},
				{
					path: 'supplier/shelf/detail',
					name: 'SupplierShelfDetail',
					component: () => import('../pages/supplier/supplier.shelf.detail.vue')
				},
				{
					path: 'trade/delivery',
					name: 'TradeDelivery',
					component: () => import('../pages/trade/trade.delivery.list.vue')
				},
				{
					path: 'trade/delivery/detail',
					name: 'TradeDeliveryDetail',
					component: () => import('../pages/trade/trade.delivery.detail.vue')
				},
				{
					path: 'dictionary/info',
					name: 'DictionaryInfo',
					component: () => import('../pages/dictionary/dictionary.info.list.vue')
				},
				{
					path: 'dictionary/info/detail',
					name: 'DictionaryInfoDetail',
					component: () => import('../pages/dictionary/dictionary.info.detail.vue')
				},
				{
					path: 'merchant/shelf',
					name: 'MerchantShelf',
					component: () => import('../pages/merchant/merchant.shelf.list.vue')
				},
				{
					path: 'merchant/shelf/detail',
					name: 'MerchantShelfDetail',
					component: () => import('../pages/merchant/merchant.shelf.detail.vue')
				},
				{
					path: 'notify/info',
					name: 'NotifyInfo',
					component: () => import('../pages/notify/notify.info.list.vue')
				},
				{
					path: 'notify/info/detail',
					name: 'NotifyInfoDetail',
					component: () => import('../pages/notify/notify.info.detail.vue')
				},
				{
					path: 'system/task',
					name: 'SystemTask',
					component: () => import('../pages/system/system.task.list.vue')
				},
				{
					path: 'system/task/detail',
					name: 'SystemTaskDetail',
					component: () => import('../pages/system/system.task.detail.vue')
				},
				{
					path: 'account/info',
					name: 'AccountInfo',
					component: () => import('../pages/account/account.info.list.vue')
				},
				{
					path: 'account/info/detail',
					name: 'AccountInfoDetail',
					component: () => import('../pages/account/account.info.detail.vue')
				},
				{
					path: 'account/record',
					name: 'AccountRecord',
					component: () => import('../pages/account/account.record.list.vue')
				},
				{
					path: 'account/record/detail',
					name: 'AccountRecordDetail',
					component: () => import('../pages/account/account.record.detail.vue')
				},
				{
					path: 'merchant/info',
					name: 'MerchantInfo',
					component: () => import('../pages/merchant/merchant.info.list.vue')
				},
				{
					path: 'merchant/info/detail',
					name: 'MerchantInfoDetail',
					component: () => import('../pages/merchant/merchant.info.detail.vue')
				},
      ]
    }
  ]
})

