import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL
const timeout = 10000
import router from '../router/router'

Vue.use(Vuex);

let store = new Vuex.Store({
    state: {
        token: localStorage.getItem('token') || '',
        adCompanyList: localStorage.getItem('user_company') || [],
        user : localStorage.getItem('user') || {},  
        account : localStorage.getItem('account') || {},
        status: '',
        adCompany: {},
        adCompanyImages: [],
    },
    mutations: {
        set_adCompany(state, adCompany) {
            state.adCompany = adCompany
        },
        set_adCompanyImages(state, images){
            state.adCompanyImages = images
        },
        auth_request(state){
            state.status = 'loading'
        },
        auth_success(state, token){
            state.status = 'success'
            state.token = token
        },
        set_user(state, user){
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
                commit('save_request') //TOdo
                axios({url: `${VUE_APP_API_URL}/api/company/`, data: companyData, method: 'POST', timeout: 100000 })
                .then(resp => {
                    // resolve(resp)
                    console.log("only r", resp)
                    resolve(resp)
                    // axios({url: `${VUE_APP_API_URL}/api/company/`, method: 'GET', timeout: timeout })
                    // .then(resp => {
                    //     console.log("user companies resp", resp)
                    //     localStorage.setItem('user_company', resp.data)
                    //     commit('set_user_company', resp.data)
                    //     resolve(resp)
                    // })
                    // .catch(err => {
                    //     console.log(err)
                    //     reject(err)
                    // })
                    // // localStorage.setItem('user_company', resp.data) //?
                    // // commit('set_user_company', resp.data) //?
                    // resolve(resp)
                })
                .catch(err => {
                    console.log(err)
                    reject(err)
                })
            })
        },
        updateCompany({commit}, companyData) {
            console.log(companyData)
            let id = ""
            return new Promise((resolve, reject) => {
                commit('save_request') //TOdo
                axios({url: `${VUE_APP_API_URL}/api/company/${id}`, data: companyData, method: 'PUT', timeout: timeout })
                .then(resp => {
                    axios({url: `${VUE_APP_API_URL}/api/company/${id}`, method: 'GET', timeout: timeout })
                    .then(resp => {
                        localStorage.setItem('user_company', resp.data)
                        commit('set_user_company', resp.data)
                        console.log(resp.data)
                        resolve(resp)
                        router.push({path: '/company-balance/'+ resp.data.Id, query: { isEdit: false }})
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
                axios({url: `${VUE_APP_API_URL}/api/company/`, method: 'GET', timeout: timeout })
                .then(resp => {
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
        getCompanyByID({commit}, id){
            function getCompany(id) {
                return axios({url: `${VUE_APP_API_URL}/api/company/${id}`, method: 'GET', timeout: timeout })
                .then(resp => {
                    return resp.data
                })
                .catch(err => {
                    console.log(err)
                })
            }
            function getCompanyImages(id) {
                return axios({url: `${VUE_APP_API_URL}/api/company/${id}/images/`, method: 'GET', timeout: timeout })
                .then(resp => {
                    return resp.data
                })
                .catch(err => {
                    console.log(err)
                })
            }
            Promise.all([getCompany(id), getCompanyImages(id)])
            .then(function (results) {
                const companyData = results[0]
                const imagesNames = results[1]
                console.log(companyData)
                console.log("im", imagesNames)
                commit('set_adCompany', companyData)
                commit('set_adCompanyImages', imagesNames) //TODO
            });
        },
        login({commit}, user){
            return new Promise((resolve, reject) => {
                commit('auth_request')
                axios({url: `${VUE_APP_API_URL}/auth/sign-in`, data: user, method: 'POST', timeout: timeout })
                .then(resp => {
                    console.log("usr + resp: ", resp.data)
                    const token = resp.data.token
                    const user = resp.data.user
                    localStorage.setItem('token', token)
                    localStorage.setItem('user', user)
                    axios.defaults.headers.common['Authorization'] = token
                    axios.defaults.headers.post['Authorization'] = token
                    commit('auth_success', token)
                    commit('set_user', user)
                    resolve(resp)
                })
                .catch(err => {
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
            //     'page_id': this.form.FbPageId,
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
        GET_COMPANY_DATA(state){  
            if (typeof state.adCompany.FbPageId === 'undefined') {
                let data = {
                    c: {
                        FbPageId: '',
                        CompanyName: '',
                        CompnayPurpose: '',
                        CompanyField: '',
                        BusinessAdress: '',
                        Images: [],
                        ImagesDescription: [],
                        ImagesSmall: [],
                        ImagesSmallDescription: [],
                        CreativeStatus: '',
                        PostDescription: '',
                    },
                    i: state.adCompanyImages,
                }
                return data
            }         
            let data = {
                c: state.adCompany,
                i: state.adCompanyImages,
            }
            return data
        },
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
            if (Array.isArray(state.adCompanyList)){
                return state.adCompanyList
            }
            return []
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