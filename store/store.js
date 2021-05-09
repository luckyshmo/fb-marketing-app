import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

let store = new Vuex.Store({
    state: {
        authenticated: false,
        FBtoken: ''
    },
    mutations: {
        SET_AUTH: (state, auth) => state.authenticated = auth,
        SET_FBTOKEN: (state, token) => state.FBtoken = token
    },
    actions: {
        setAuth: ({ commit }, auth) => commit('SET_AUTH', auth),
        setFBToken: ({ commit }, token) => commit('SET_FBTOKEN', token)
    },
    getters: {
        IS_AUTH(state) {
            return state.authenticated;
        },
        GET_FB_TOKEN(state) {
            return state.FBtoken;
        }
    },
    modules: {}
});

export default store;