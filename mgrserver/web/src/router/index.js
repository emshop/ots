
import Vue from 'vue';
import Router from 'vue-router';
import menus from '@/pages/system/menus';
import order from '@/pages/trade/order';

Vue.use(Router);
export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: '/',
      component: menus,
      children:[{
        path: 'trade/order',
        name: 'order',
        component: order,
        titile:"首页"
      }]
    }
  ]
})

