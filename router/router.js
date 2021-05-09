import Router from 'vue-router'
import Vue from 'vue'
import mainPage from '../src/components/Main.vue'
import login from '../src/components/login/Login.vue'
import register from '../src/components/login/Register.vue'
import createCompany from '../src/components/pages/CreateCompany.vue'

Vue.use(Router);

let router = new Router({
    routes: [
      {
        path: '/login',
        name: 'login',
        component: login
      },
      {
        path: '/register',
        name: 'register',
        component: register
      },
      {
        path: '/main',
        name: 'mainPage',
        component: mainPage
      },
      {
        path: '/createCompany',
        name: 'createCompany',
        component: createCompany
      },
    ]
  })
  
  export default router;