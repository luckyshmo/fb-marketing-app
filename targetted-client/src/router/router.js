import Router from 'vue-router'
import Vue from 'vue'
import mainPage from '@/pages/Main.vue'
import login from '@/pages/login/Login.vue'
import questions from '@/pages/Questions.vue'
import register from '@/pages/login/Register.vue'
import createCompany from '@/pages/CCnew.vue'
import imagesPreview from '@/pages/CreativesPreview.vue'
import paymentPage from '@/pages/Payment.vue'
//import createCompany from '../src/components/pages/CreateCompany.vue'
import store from '../store/store'
import CompanyBalance from '@/pages/CompanyBalance.vue'
import adminka from '@/pages/Adminka.vue'

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
      path: '/questions',
      name: 'questions',
      component: questions
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
      path: '/questions',
      name: 'questions',
      component: questions,
      meta: {
        requiresAuth: false
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
      path: '/payment',
      name: 'campagin-payment',
      component: paymentPage,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/preview',
      name: 'images-preview',
      component: imagesPreview,
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
  
  // if (to.matched.some(record => record.meta.requiresAuth)) {
  //   // this route requires auth, check if logged in
  //   // if not, redirect to login page.
  //   if (!store.getters.isLoggedIn) {
  //     next({ name: 'login' })
  //   } else {
  //     next() // go to wherever I'm going
  //   }
  // } else {
  //   next() // does not require auth, make sure to always call next()!
  // }

  next()
})

export default router
