import Vue from 'vue'
import App from './App.vue'
import router from '../router/router'
import Axios from 'axios'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
// Import Bootstrap an BootstrapVue CSS files (order is important)
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue)
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin)
Vue.config.productionTip = false

//this.$http для использования axios
Vue.prototype.$http = Axios;

const token = localStorage.getItem('user-token')
if (token) {
  //добавляем токен во все запросы
  Vue.prototype.$http.defaults.headers.common['Authorization'] = token
}

Vue.config.productionTip = false

import initFacebookSdk from '../src/_helpers/init-facebook-sdk';
// fakeBackend();
initFacebookSdk();
new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
