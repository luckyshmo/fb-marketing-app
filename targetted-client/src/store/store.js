import Vue from 'vue'
import Vuex from 'vuex'

import actions from './actions'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: localStorage.getItem('token') || '',
    // tmp default state, for demo
    adCompanyList: localStorage.getItem('user_company') || [{
      Id: 'test-url',
      CompanyName: 'Test url',
      FbPageId: '',
      Date: '10 Сентября'
    },
    {
      Id: 'test-3456789',
      CompanyName: 'https://www.figma.com/',
      FbPageId: '2345678',
      IsStarted: true,
      Date: '1 Мая'
    }],
    user: localStorage.getItem('user') || '',
    email: localStorage.getItem('email') || '',
    account: localStorage.getItem('account') || {},
    status: '',
    adCompany: {},
    adCompanyImages: []
  },
  mutations: {
    set_adCompany (state, adCompany) {
      state.adCompany = adCompany
    },
    set_adCompanyImages (state, images) {
      state.adCompanyImages = images
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
    set_user_company (state, companies) {
      state.adCompanyList = companies
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
    GET_COMPANY_DATA (state) {
      if (typeof state.adCompany.FbPageId === 'undefined') {
        const data = {
          c: {
            FbPageId: '',
            CompanyName: '',
            CompanyPurpose: '',
            CompanyField: '',
            BusinessAdress: '',
            Images: [],
            ImagesDescription: [],
            ImagesSmall: [],
            ImagesSmallDescription: [],
            CreativeStatus: '',
            PostDescription: ''
          },
          i: state.adCompanyImages
        }
        return data
      }
      const data = {
        c: state.adCompany,
        i: state.adCompanyImages
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
    GET_COMPANY_LIST (state) {
      if (Array.isArray(state.adCompanyList)) {
        return state.adCompanyList
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
