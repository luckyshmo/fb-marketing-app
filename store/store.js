import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex);

let store = new Vuex.Store({
    state: {
        FBtoken: '',
        token: localStorage.getItem('token') || '',
        user : {},  
    },
    mutations: {
        SET_FBTOKEN: (state, token) => state.FBtoken = token,
        auth_request(state){
            state.status = 'loading'
        },
        auth_success(state, token, user){
            state.status = 'success'
            state.token = token
            state.user = user
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
        setFBToken: ({ commit }, token) => commit('SET_FBTOKEN', token),
        login({commit}, user){
            return new Promise((resolve, reject) => {
                commit('auth_request')
                console.log(user)
                axios({url: 'http://localhost:3000/auth/sign-in', data: user, method: 'POST' })
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
                axios({url: 'http://localhost:3000/auth/sign-up', data: user, method: 'POST', headers: {
                    // remove headers
                  } })
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
        }
    },
    getters: {
        IS_AUTH(state) {
            return state.authenticated;
        },
        GET_FB_TOKEN(state) {
            return state.FBtoken;
        },

        isLoggedIn: state => !!state.token,
        authStatus: state => state.status,
    },
    modules: {}
});

export default store;