import Router from 'vue-router'
import Vue from 'vue'
import mainPage from '../src/components/Main.vue'
import login from '../src/components/login/Login.vue'
import register from '../src/components/login/Register.vue'
import createCompany from '../src/components/pages/CCnew.vue'
//import createCompany from '../src/components/pages/CreateCompany.vue'
import store from '../store/store'
import CompanyBalance from '../src/components/pages/CompanyBalance.vue'
import adminka from '../src/components/pages/Adminka.vue'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '*',
      redirect: '/main' // TODO NOT FOUND?
    },
    {
      path: '/',
      name: 'default',
      redirect: '/main'
    },
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
      component: mainPage,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/adminka',
      name: 'adminka',
      component: adminka,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/company',
      name: 'create-company',
      component: createCompany,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/company/:id',
      name: 'edit-company',
      component: createCompany,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/company-balance/:id',
      name: 'company-balance',
      component: CompanyBalance
    }
  ]
})

router.beforeEach((to, from, next) => {
  
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // this route requires auth, check if logged in
    // if not, redirect to login page.
    if (!store.getters.isLoggedIn) {
      next({ name: 'login' })
    } else {
      next() // go to wherever I'm going
    }
  } else {
    next() // does not require auth, make sure to always call next()!
  }

  next()
})

export default router
