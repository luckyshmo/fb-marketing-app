import Vue from 'vue'
import Vuex from 'vuex'

import actions from './actions'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: localStorage.getItem('token') || '',
    adCampaignList: localStorage.getItem('user_campaign'),
    user: localStorage.getItem('user') || '',
    email: localStorage.getItem('email') || '',
    account: localStorage.getItem('account') || {},
    status: '',
    adCampaign: {},
    adCampaignImages: []
  },
  mutations: {
    set_adCampaign (state, adCampaign) {
      state.adCampaign = adCampaign
    },
    set_adCampaignImages (state, images) {
      state.adCampaignImages = images
    },
    auth_request (state) {
      state.status = 'loading'
    },
    auth_success (state, token) {
      state.status = 'success'
      state.token = token
    },
    set_user (state, user) {
      state.user = user
    },
    set_email (state, email) {
      state.email = email
    },
    set_user_campaign (state, campaigns) {
      state.adCampaignList = campaigns
    },
    save_request (state) {
      state.status = 'loading'
    },
    set_account (state, account) {
      state.account = account
    },
    auth_error (state) {
      state.status = 'error'
    },
    logout (state) {
      state = {
        status: '',
        user: '',
        email: '',
        token: ''
      }
    }
  },
  actions,
  getters: {
    GET_CAMPAIGN_DATA (state) {
      if (typeof state.adCampaign.FbPageId === 'undefined') {
        const data = {
          c: {
            FbPageId: '',
            CampaignName: '',
            CampaignObjective: '',
            CampaignField: '',
            BusinessAdress: '',
            Images: [],
            ImagesSmall: [],
            CreativeStatus: '',
            PostDescription: ''
          },
          i: state.adCampaignImages
        }
        return data
      }
      const data = {
        c: state.adCampaign,
        i: state.adCampaignImages
      }
      return data
    },
    GET_FB_ACCOUNT (state) {
      return state.account
    },
    GET_FB_PAGES (state) {
      if (typeof state.account.pages === 'undefined') {
        return []
      }
      const pages = []
      for (let i = 0; i < state.account.pages.length; i++) {
        pages[i] = {
          page: state.account.pages[i],
          name: state.account.pages[i].name,
          value: state.account.pages[i].id,
          text: state.account.pages[i].name
        }
      }
      return pages
    },
    GET_CAMPAIGN_LIST (state) {
      if (Array.isArray(state.adCampaignList)) {
        return state.adCampaignList
      }
      return []
    },
    GET_USER (state) {
      console.log(state.user)
      return state.user
    },
    GET_EMAIL (state) {
      console.log('get email: ', state.email)
      return state.email
    },
    GET_TOKEN (state) {
      return state.token
    },
    isLoggedIn: state => !!state.token,
    authStatus: state => state.status
  }
})

export default store
