
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
        component: () => import('../pages/trade/order.list.vue'),
        titile:"首页"
      },{
        path: 'ots/trade/order.detail/:no',
        name: 'index',
        component: () => import('../pages/trade/order.detail.vue'),
        titile:"首页"
      }]
    }
  ]
})

