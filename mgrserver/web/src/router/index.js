
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
      ]
    }
  ]
})

