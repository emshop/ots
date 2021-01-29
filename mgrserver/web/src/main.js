

import "jquery"
import "bootstrap"
 
import Vue from 'vue'
import App from './App'
import router from './router'

import VueCookies from 'vue-cookies'
Vue.use(VueCookies);

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);

import utility from './utility'
Vue.use(utility,true);

Vue.config.productionTip = false;

router.beforeEach((to, from, next) => {
    /* 路由发生变化修改页面title */   
    next()
})


  /* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    components: {
        App
    },
    template: '<App/>'
});

