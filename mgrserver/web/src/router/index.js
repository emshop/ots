
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
					path: 'trade/delivery',
					name: 'TradeDelivery',
					component: () => import('../pages/trade/trade.delivery.list.vue')
				},
				{
					path: 'trade/delivery/detail',
					name: 'TradeDeliveryDetail',
					component: () => import('../pages/trade/trade.delivery.detail.vue')
				},
      ]
    }
  ]
})

