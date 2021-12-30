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
import AdvertisingCampaign from "@src/pages/AdvertisingCampaign";
import ControlBudget from "@src/pages/ControlBudget";
import CreateCampaignStep1 from "@src/pages/CreateCampaignStep-1";
import CreateCampaignStep2 from "@src/pages/CreateCampaignStep-2";
import CreateCampaignStep3 from "@src/pages/CreateCampaignStep-3";
import CreateCampaignStep5 from "@src/pages/CreateCampaignStep-5";
import PageAnchor from "@src/pages/PageAnchor";
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
      path: '/advertising-campaign',
      name: 'AdvertisingCampaign',
      component: AdvertisingCampaign
    },
    {
      path: '/control-budget',
      name: 'ControlBudget',
      component: ControlBudget
    },
    {
      path: '/campaign-step-1',
      name: 'CreateCampaignStep1',
      component: CreateCampaignStep1
    },
    {
      path: '/campaign-step-2',
      name: 'CreateCampaignStep2',
      component: CreateCampaignStep2
    },
    {
      path: '/campaign-step-3',
      name: 'CreateCampaignStep3',
      component: CreateCampaignStep3
    },
    {
      path: '/campaign-step-5',
      name: 'CreateCampaignStep5',
      component: CreateCampaignStep5
    },
    {
      path: '/page-anchor',
      name: 'PageAnchor',
      component: PageAnchor
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
