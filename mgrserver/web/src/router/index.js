
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
					path: 'merchant/info',
					name: 'MerchantInfo',
					component: () => import('../pages/merchant/merchant.info.list.vue')
				},
				{
					path: 'merchant/info/detail',
					name: 'MerchantInfoDetail',
					component: () => import('../pages/merchant/merchant.info.detail.vue')
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
					path: 'supplier/error/code',
					name: 'SupplierErrorCode',
					component: () => import('../pages/supplier/error/supplier.error.code.list.vue')
				},
				{
					path: 'supplier/error/code/detail',
					name: 'SupplierErrorCodeDetail',
					component: () => import('../pages/supplier/error/supplier.error.code.detail.vue')
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
					path: 'merchant/shelf',
					name: 'MerchantShelf',
					component: () => import('../pages/merchant/merchant.shelf.list.vue')
				},
				{
					path: 'merchant/shelf/detail',
					name: 'MerchantShelfDetail',
					component: () => import('../pages/merchant/merchant.shelf.detail.vue')
				},
      ]
    }
  ]
})

