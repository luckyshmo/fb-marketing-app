import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

let store = new Vuex.Store({
    state: {
        authenticated: false
    },
    mutations: {
        SET_AUTH: (state, auth) => state.authenticated = auth
    },
    actions: {
        setAuth: ({ commit }, auth) => commit('SET_AUTH', auth)
    },
    getters: {
        IS_AUTH(state) {
            return state.authenticated;
        },
    },
    modules: {}
});

export default store;