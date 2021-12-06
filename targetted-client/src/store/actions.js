import axios from 'axios'

// TODO: replace with axios instance
const VUE_APP_API_URL = process.env.VUE_APP_API_URL
const timeout = 10000
//! !

function buildFormData (formData, data, parentKey) {
  if (data && typeof data === 'object' && !(data instanceof Date) && !(data instanceof File)) {
    Object.keys(data).forEach(key => {
      buildFormData(formData, data[key], parentKey ? `${parentKey}[${key}]` : key)
    })
  } else {
    const value = data == null ? '' : data

    formData.append(parentKey, value)
  }
}

const jsonToFormData = (data, isEdit) => {
  console.log(data, isEdit)
  // todo, from SO
  const campaignData = new FormData()
  buildFormData(campaignData, data)

  // campaignData.append('FbPageId', data.FbPageId)
  // if (isEdit) {
  //   campaignData.append('Id', data.Id)
  // }
  // campaignData.append('CampaignName', data.CampaignName)
  // campaignData.append('CampaignObjective', data.CampaignObjective)
  // campaignData.append('CampaignField', data.CampaignField)
  // campaignData.append('BusinessAddress', data.BusinessAddress)
  // campaignData.append('ImagesDescription', data.ImagesDescription)
  // campaignData.append('ImagesSmallDescription', data.ImagesSmallDescription)
  // campaignData.append('CreativeStatus', data.CreativeStatus)
  // campaignData.append('PostDescription', data.PostDescription)
  // campaignData.append('DailyAmount', data.DailyAmount)
  // campaignData.append('Days', data.Days)
  // Array.from(data.ImagesSmall).forEach(Image => {
  //   campaignData.append('ImageSmall', Image)
  // })
  // Array.from(data.Images).forEach(Image => {
  //   campaignData.append('Image', Image)
  // })

  return campaignData
}

const actions = {
  saveCampaign ({ commit }, campaignData) {
    const data = jsonToFormData(campaignData)
    const url = `${VUE_APP_API_URL}/api/campaign/`

    return new Promise((resolve, reject) => {
      commit('save_request')
      axios.post(url, data)
        .then(resp => {
          axios.get(`${VUE_APP_API_URL}/api/campaign/`)
            .then(resp => {
              console.log('user campaigns resp', resp)
              localStorage.setItem('user_campaign', resp.data)
              commit('set_user_campaign', resp.data)
              resolve(resp)
            })
            .catch(err => {
              console.log(err)
              reject(err)
            })
            // localStorage.setItem('user_campaign', resp.data) //?
            // commit('set_user_campaign', resp.data) //?
          resolve(resp)
        })
        .catch(err => {
          // TODO remove
          resolve('test')

          console.log(err)
          //  reject(err)
        })
    })
  },
  updateCampaign ({ commit }, campaignData) {
    const data = jsonToFormData(campaignData, true)

    const id = ''
    return new Promise((resolve, reject) => {
      commit('save_request') // TOdo
      debugger
      const url = `${VUE_APP_API_URL}/api/campaign/${id}`
      axios.put(url, { data })
        .then(resp => {
          resolve(resp)
        })
        .catch(err => {
          console.log(err)
          reject(err)
        })
    })
  },
  getCampaignList ({ commit }) {
    return new Promise((resolve, reject) => {
      const url = `${VUE_APP_API_URL}/api/campaign/`
      axios.get(url)
        .then(resp => {
          console.log(resp.data)
          localStorage.setItem('user_campaign', resp.data)
          commit('set_user_campaign', resp.data)
          resolve(resp)
        })
        .catch(err => {
          console.log(err)
          reject(err)
        })
    })
  },
  getCampaignByID ({ commit }, id) {
    function getCampaign (id) {
      const url = `${VUE_APP_API_URL}/api/campaign/${id}`
      return axios.get(url)
        .then(resp => {
          return resp.data
        })
        .catch(err => {
          console.log(err)
        })
    }
    function getCampaignImages (id) {
      return axios.get({ url: `${VUE_APP_API_URL}/api/campaign/${id}/images/`, timeout })
        .then(resp => {
          return resp.data
        })
        .catch(err => {
          console.log(err)
        })
    }
    Promise.all([getCampaign(id), getCampaignImages(id)])
      .then(function (results) {
        const campaignData = results[0]
        const imagesNames = results[1]
        console.log(campaignData)
        console.log('im', imagesNames)
        commit('set_adCampaign', campaignData)
        commit('set_adCampaignImages', imagesNames) // TODO
      })
  },
  login ({ commit }, user) {
    return new Promise((resolve, reject) => {
      commit('auth_request')
      const url = `${VUE_APP_API_URL}/auth/sign-in`
      axios.post(url, user)
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

export default actions
