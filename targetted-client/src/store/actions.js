import axios from 'axios'

//TODO: replace with axios instance
const VUE_APP_API_URL = process.env.VUE_APP_API_URL
const timeout = 10000
//!!

function buildFormData(formData, data, parentKey) {
  if (data && typeof data === 'object' && !(data instanceof Date) && !(data instanceof File)) {
    Object.keys(data).forEach(key => {
      buildFormData(formData, data[key], parentKey ? `${parentKey}[${key}]` : key);
    });
  } else {
    const value = data == null ? '' : data;

    formData.append(parentKey, value);
  }
}

const jsonToFormData = (data, isEdit) => {
    console.log(data, isEdit)
    //todo, from SO
    const companyData = new FormData()
    buildFormData(companyData, data);
  
    // companyData.append('FbPageId', data.FbPageId)
    // if (isEdit) {
    //   companyData.append('Id', data.Id)
    // }
    // companyData.append('CompanyName', data.CompanyName)
    // companyData.append('CompnayPurpose', data.CompnayPurpose)
    // companyData.append('CompanyField', data.CompanyField)
    // companyData.append('BusinessAddress', data.BusinessAddress)
    // companyData.append('ImagesDescription', data.ImagesDescription)
    // companyData.append('ImagesSmallDescription', data.ImagesSmallDescription)
    // companyData.append('CreativeStatus', data.CreativeStatus)
    // companyData.append('PostDescription', data.PostDescription)
    // companyData.append('DailyAmount', data.DailyAmount)
    // companyData.append('Days', data.Days)
    // Array.from(data.ImagesSmall).forEach(Image => {
    //   companyData.append('ImageSmall', Image)
    // })
    // Array.from(data.Images).forEach(Image => {
    //   companyData.append('Image', Image)
    // })
  
    return companyData;
  }

  
const actions = {
    saveCompany ({ commit }, companyData) {

      const data = jsonToFormData(companyData);
      const url = `${VUE_APP_API_URL}/api/company/`;

      return new Promise((resolve, reject) => {
        commit('save_request') // TOdo
        axios.post(url, data)
          .then(resp => {
            axios.get(`${VUE_APP_API_URL}/api/company/`)
              .then(resp => {
                console.log('user companies resp', resp)
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
            //TODO remove
            resolve('test')


            console.log(err)
            reject(err)
          })
      })
    },
    updateCompany ({ commit }, companyData) {
      
      const data = jsonToFormData(companyData, true);

      const id = ''
      return new Promise((resolve, reject) => {
        commit('save_request') // TOdo
        debugger
        const url = `${VUE_APP_API_URL}/api/company/${id}`
        axios.put(url, {data })
          .then(resp => {
            resolve(resp)
          })
          .catch(err => {
            console.log(err)
            reject(err)
          })
      })
    },
    getCompanyList ({ commit }) {
      return new Promise((resolve, reject) => {
        const url = `${VUE_APP_API_URL}/api/company/`;
        axios.get(url)
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
    getCompanyByID ({ commit }, id) {
      function getCompany (id) {
        const url = `${VUE_APP_API_URL}/api/company/${id}`;
        return axios.get(url)
          .then(resp => {
            return resp.data
          })
          .catch(err => {
            console.log(err)
          })
      }
      function getCompanyImages (id) {
        return axios.get({ url: `${VUE_APP_API_URL}/api/company/${id}/images/`, timeout })
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
          console.log('im', imagesNames)
          commit('set_adCompany', companyData)
          commit('set_adCompanyImages', imagesNames) // TODO
        })
    },
    login ({ commit }, user) {
      return new Promise((resolve, reject) => {
        commit('auth_request')
        const url = `${VUE_APP_API_URL}/auth/sign-in`;
        axios.post(url, { data: user, timeout })
          .then(resp => {
            console.log('usr + resp: ', resp.data)
            const token = resp.data.token
            const user = resp.data.user
            localStorage.setItem('token', token)
            console.log('set User email', user.email)
            localStorage.setItem('email', user.email)
            axios.defaults.headers.common.Authorization = token
            axios.defaults.headers.post.Authorization = token
            commit('auth_success', token)
            commit('set_email', user.email)
            resolve(resp)
          })
          .catch(err => {
            localStorage.removeItem('token')
            reject(err)
          })
      })
    },
    logout ({ commit }) {
      return new Promise((resolve, reject) => {
        console.log(reject) // ?
        commit('logout')
        localStorage.removeItem('token')
        localStorage.removeItem('email')
        localStorage.removeItem('user')
        delete axios.defaults.headers.common.Authorization
        resolve()
      })
    },
    setAccount ({ commit }, token) {
      if (token === '') {
        localStorage.setItem('account', {})
        commit('set_account', {})
        console.log('unloginFB')
        return
      }

      function getUserAccount (token) {
        return axios.get(`https://graph.facebook.com/v10.0/me?access_token=${token}`)
      }
      function getUserPages (token) {
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

          const account = {
            pages: pages,
            facebookId: data.id,
            name: data.name,
            token: token
          }

          localStorage.setItem('account', account)
          commit('set_account', account)
        })
    }
  }

export default actions;