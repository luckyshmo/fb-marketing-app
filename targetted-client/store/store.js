import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL;

Vue.use(Vuex);

let store = new Vuex.Store({
    state: {
        token: localStorage.getItem('token') || '',
        adCompanyList: localStorage.getItem('user_company') || [],
        user : {},  
        account : localStorage.getItem('account') || {},
        status: '',
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
        set_user_company(state, companies){
            state.adCompanyList = companies
        },
        save_request(state){
            state.status = 'loading'
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
        saveCompany({commit}, companyData) {
            return new Promise((resolve, reject) => {
                commit('save_request')
                axios({url: `${VUE_APP_API_URL}/api/company/`, data: companyData, method: 'POST' })
                .then(resp => {
                    console.log(resp)
                    axios({url: `${VUE_APP_API_URL}/api/company/`, method: 'GET' })
                    .then(resp => {
                        console.log(resp)
                        localStorage.setItem('user_company', resp.data)
                        commit('set_user_company', resp.data)
                        resolve(resp)
                    })
                    .catch(err => {
                        console.log(err)
                        reject(err)
                    })
                    // localStorage.setItem('user_company', resp.data) //?
                    // commit('set_user_company', resp.data) //?
                    resolve(resp)
                })
                .catch(err => {
                    console.log(err)
                    reject(err)
                })
            })
        },
        getCompanyList({commit}){
            return new Promise((resolve, reject) => {
                axios({url: `${VUE_APP_API_URL}/api/company/`, method: 'GET' })
                .then(resp => {
                    console.log(resp)
                    localStorage.setItem('user_company', resp.data)
                    commit('set_user_company', resp.data)
                    resolve(resp)
                })
                .catch(err => {
                    console.log(err)
                    reject(err)
                })
            })
        },
        login({commit}, user){
            return new Promise((resolve, reject) => {
                commit('auth_request')
                console.log(user)                
                // const baseUrl = `${process.env.VUE_APP_API_URL}/accounts`;
                axios({url: `${VUE_APP_API_URL}/auth/sign-in`, data: user, method: 'POST' })
                .then(resp => {
                    const token = resp.data.token
                    const user = resp.data.user
                    console.log("Token to AUTH:", token)
                    localStorage.setItem('token', token)
                    // Add the following line:
                    axios.defaults.headers.common['Authorization'] = token
                    axios.defaults.headers.post['Authorization'] = token
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
                    axios.defaults.headers.post['Authorization'] = token
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
        // setAdCompanyList({commit}){
            // let appToken = 'EAAEDuTXOcAgBAEbAJLLg00LDOJH4LyOekYZCWtJhjul3xbrUpQZCWt0LEDTlpQrsxhwWUZBSjZAA5OyRMgZB0g83zIIKXNQRys82ZAajuUGAmZAmQGy5kH242uZAZABoMjgebiuGQkcjKJ5Kd8xyWXThFQytJP1ATmHNNQvPZA0I1RROQAbmWUJS8HgyFMtWkETMecbEPUNLC4zgZDZD'
            // let companyId = 856950044859235
            // let files = {
            //     'page_id': this.form.fbPageId,
            //     'permitted_tasks': '[\'MANAGE\', \'CREATE_CONTENT\', \'MODERATE\', \'ADVERTISE\', \'ANALYZE\']',
            //     'access_token': appToken,
            // }
            // axios.get(`https://graph.facebook.com/v10.0/${companyId}/client_pages`, files)
            // .then(resp) {

            // }
    //         	curl -G \
	// -d "access_token=EAAEDuTXOcAgBAEbAJLLg00LDOJH4LyOekYZCWtJhjul3xbrUpQZCWt0LEDTlpQrsxhwWUZBSjZAA5OyRMgZB0g83zIIKXNQRys82ZAajuUGAmZAmQGy5kH242uZAZABoMjgebiuGQkcjKJ5Kd8xyWXThFQytJP1ATmHNNQvPZA0I1RROQAbmWUJS8HgyFMtWkETMecbEPUNLC4zgZDZD" \
	// "https://graph.facebook.com/v10.0/856950044859235/client_pages"
        // },
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
        GET_FB_PAGES(state){
            if (typeof state.account.pages === 'undefined'){
                return []
            }
            let pages = []
            for (let i = 0; i < state.account.pages.length; i++){
                pages[i] = {
                    page: state.account.pages[i],
                    name: state.account.pages[i].name,
                    value: state.account.pages[i].id,
                    text: state.account.pages[i].name
                }
            }
            return pages
        },
        GET_COMPANY_LIST(state){
            return state.adCompanyList
        },
        GET_USER(state){
            return state.user
        },
        GET_TOKEN(state){
            return state.token
        },
        isLoggedIn: state => !!state.token,
        authStatus: state => state.status,
    },
    modules: {}
});

export default store;