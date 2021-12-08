import Router from 'vue-router'
import Vue from 'vue'
import mainPage from '@src/pages/Main.vue'
import login from '@src/pages/login/Login.vue'
import questions from '@src/pages/Questions.vue'
import register from '@src/pages/login/Register.vue'
import createCampaign from '@src/pages/CreateCampaign.vue'
import imagesPreview from '@src/pages/CreativesPreview.vue'
import paymentPage from '@src/pages/Payment.vue'
import CampaignBalance from '@src/pages/CampaignBalance.vue'
import adminka from '@src/pages/Adminka.vue'
import UnattachedAccount from "@src/pages/advertising-campaign/UnattachedAccount";

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
      component: login,
      meta: {
        hideHeader: true
      }
    },
    {
      path: '/questions',
      name: 'questions',
      component: questions,
      meta: {
        hideHeader: true
      }
    },
    {
      path: '/register',
      name: 'register',
      component: register,
      meta: {
        hideHeader: true
      }
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
      path: '/campaign',
      name: 'create-campaign',
      component: createCampaign,
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
      path: '/campaign/:id',
      name: 'edit-campaign',
      component: createCampaign,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/campaign-balance/:id',
      name: 'campaign-balance',
      component: CampaignBalance
    },
    {
      path: '/unattached-account',
      name: 'unattached-account',
      component: UnattachedAccount
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
