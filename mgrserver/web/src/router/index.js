
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
        path: 'index',
        name: 'index',
        component: () => import('../pages/trade/trade.order.list.vue'),
        meta: { title: "交易订单" }
        },
        {
          path:'trade/order.detail',
          name: 'trade/order.detail',
          component: () => import('../pages/trade/trade.order.detail.vue'),
          meta: { title: "订单详情" }
        } , {
          path:'trade/delivery',
          name: 'trade/delivery',
          component: () => import('../pages/trade/trade.delivery.list.vue'),
          meta: { title: "发货列表" }
        },
        {
          path:'trade/delivery.detail',
          name: 'trade/delivery.detail',
          component: () => import('../pages/trade/trade.delivery.detail.vue'),
          meta: { title: "发货详情" }
        } ,
        {
          path:'merchant/info',
          name: 'merchant/merchant.info',
          component: () => import('../pages/merchant/merchant.info.list.vue'),
          meta: { title: "发货详情" }
        } ,{
          path:'merchant/info.detail',
          name: 'merchant/info.detail',
          component: () => import('../pages/merchant/merchant.info.detail.vue'),
          meta: { title: "发货详情" }
        } ,

        {
          path:'merchant/shelf',
          name: 'merchant/merchant.shelf',
          component: () => import('../pages/merchant/merchant.shelf.list.vue'),
          meta: { title: "发货详情" }
        } ,{
          path:'merchant/shelf.detail',
          name: 'merchant/shelf.detail',
          component: () => import('../pages/merchant/merchant.shelf.detail.vue'),
          meta: { title: "发货详情" }
        } ,
        {
          path:'merchant/product',
          name: 'merchant/merchant.product',
          component: () => import('../pages/merchant/merchant.product.list.vue'),
          meta: { title: "发货详情" }
        } ,{
          path:'merchant/product.detail',
          name: 'merchant/product.detail',
          component: () => import('../pages/merchant/merchant.product.detail.vue'),
          meta: { title: "发货详情" }
        } ,
      ]
    }
  ]
})

