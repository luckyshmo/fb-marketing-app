import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL;

Vue.use(Vuex);

let store = new Vuex.Store({
    state: {
        token: localStorage.getItem('token') || '',
        user : {},  
        account : {},
    },
    mutations: {
        auth_request(state){
            state.status = 'loading'
        },
        auth_success(state, token, user){
            state.status = 'success'
            state.token = token
            state.user = user
        },
        set_account(state, account){
            state.account = account
        },
        auth_error(state){
            state.status = 'error'
        },
        logout(state){
            state.status = ''
            state.token = ''
        },
    },
    actions: {
        login({commit}, user){
            return new Promise((resolve, reject) => {
                commit('auth_request')
                console.log(user)                
                // const baseUrl = `${process.env.VUE_APP_API_URL}/accounts`;
                axios({url: `${VUE_APP_API_URL}/auth/sign-in`, data: user, method: 'POST' })
                .then(resp => {
                    console.log(resp.data)
                    const token = resp.data.token
                    const user = resp.data.user
                    localStorage.setItem('token', token)
                    // Add the following line:
                    axios.defaults.headers.common['Authorization'] = token
                    commit('auth_success', token, user)
                    resolve(resp)
                })
                .catch(err => {
                    commit('auth_error')
                    localStorage.removeItem('token')
                    reject(err)
                })
            })
        },
        register({commit}, user){
            return new Promise((resolve, reject) => {
                commit('auth_request')
                axios({url: `${VUE_APP_API_URL}/auth/sign-up`, data: user, method: 'POST'})
                .then(resp => {
                    console.log(resp.data)
                    const token = resp.data.token
                    const user = resp.data.user
                    localStorage.setItem('token', token)
                    // Add the following line:
                    axios.defaults.headers.common['Authorization'] = token
                    commit('auth_success', token, user)
                    resolve(resp)
                })
                .catch(err => {
                    commit('auth_error', err)
                    localStorage.removeItem('token')
                    reject(err)
                })
            })
        },
        logout({commit}){
            return new Promise((resolve, reject) => {
                console.log(reject) //?
                commit('logout')
                localStorage.removeItem('token')
                delete axios.defaults.headers.common['Authorization']
                resolve()
            })
        },
        setAccount({commit}, token){
            if (token === '') {
                localStorage.setItem('account', {})
                commit("set_account", {})
                console.log("unloginFB")
                return
            }

            function getUserAccount(token) {
                return axios.get(`https://graph.facebook.com/v10.0/me?access_token=${token}`)
            }
            function getUserPages(token) {
                return axios.get(`https://graph.facebook.com/v10.0/me/accounts?access_token=${token}`)
            }
            Promise.all([getUserAccount(token), getUserPages(token)])
            .then(function (results) {
                const uA = results[0]
                const uP = results[1]
                
                const data = uA.data
                const pages = uP.data.data
                console.log(pages)
                console.log(uA)

                let account = {
                    pages: pages,
                    facebookId: data.id,
                    name: data.name,
                    token: token,
                }

                localStorage.setItem('account', account)
                commit("set_account", account)
                

              });
        }
    },
    getters: {
        GET_FB_ACCOUNT(state){
            return state.account
        },
        isLoggedIn: state => !!state.token,
        authStatus: state => state.status,
    },
    modules: {}
});

export default store;