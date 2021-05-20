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
// delete Axios.defaults.headers.common["Authorization"];

// Axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*'
// Axios.defaults.headers.common['Access-Control-Allow-Credentials'] = 'true'
// Axios.defaults.headers.common['Access-Control-Allow-Headers'] = 'Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With'
// Axios.defaults.headers.common['Access-Control-Allow-Methods'] = 'POST, OPTIONS, GET, PUT, DELETE'

Vue.config.productionTip = false

// setup fake backend
import fakeBackend from '../src/_helpers/fake-backend';
import initFacebookSdk from '../src/_helpers/init-facebook-sdk';
fakeBackend();
initFacebookSdk();
new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
