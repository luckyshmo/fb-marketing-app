<template>
    <section class="steps">
      <loading v-if="isLoading" />
      <b-row class="app-new-steps-header">
        <b-col>
          <router-link id="app-link-back" v-if="currentStep === 1" :to="{name: 'mainPage'}">
            ← Назад
          </router-link>
          <p class="text-muted clickable" v-else @click="goStepBack" style="float:left;margin-right: 4px;line-height: 20px;">
            ← Назад
          </p>
          <b-col>
            <div class="text-muted d-block d-md-none d-lg-none d-xl-none">&nbsp;∙&nbsp;Шаг {{currentStep}} из {{totalSteps}}</div>
          </b-col>
        </b-col>
      </b-row>
      <b-row align-h="between">
        <b-col>
          <Step1
            :label_cols="label_cols"
            :content_cols="content_cols"
            :isEdit="isEdit"
            @logout="logout"
            v-if="currentStep ===  1"
            @next="saveAndNext">
            <template v-slot:header>
              <h1>О бизнесе</h1>
            </template>
          </Step1>

          <Step2
            :label_cols="label_cols"
            :content_cols="content_cols"
            :isEdit="isEdit"
            :campaign="campaign"
            v-if="currentStep ===  2"
            @next="saveAndNext">
            <template v-slot:header>
              <h1>Изображения </h1>
            </template>
          </Step2>

          <Step3
            :label_cols="label_cols"
            :content_cols="content_cols"
            :isEdit="isEdit"
            :campaign="campaign"
            v-if="currentStep ===  3"
            @next="saveAndNext">
            <template v-slot:header>
              <h1>Аудитория</h1>
            </template>
          </Step3>

          <Step4
            :label_cols="label_cols"
            :content_cols="content_cols"
            :isEdit="isEdit"
            :campaign="campaign"
            v-if="currentStep ===  4"
            @next="saveAndNext">
            <template v-slot:header>
              <h1>Привязка страницы</h1>
            </template>
          </Step4>

          <Step5
            :label_cols="label_cols"
            :content_cols="content_cols"
            :isEdit="isEdit"
            :campaign="campaign"
            v-if="currentStep ===  4"
            @next="saveAndNext">
            <template v-slot:header>
              <h1>Охват рекламы</h1>
            </template>
          </Step5>
        </b-col>
        <b-col class="steps-process">
          <aside>
            <h2>
              Заполнено {{displayNumber}}%
            </h2>
            <ul class="app-new-progress-text">
              <li :class="{active:(currentStep>=1)}">
                О бизнесе
              </li>
              <li :class="{active:(currentStep>=2)}">
                Изображения
              </li>
              <li :class="{active:(currentStep>=3)}">
                Привязка аккаунта
              </li>
              <li :class="{active:(currentStep>=3)}">
                Охват рекламы
              </li>
              <li :class="{active:(currentStep>=4)}">
                Аудитория
              </li>
            </ul>
          </aside>
        </b-col>
      </b-row>
    </section>
</template>
<script>

import accountService from '@src/_services/account.service'
import store from '@src/store/store'
import router from '@src/router/router'
import { instance as axios } from '@src/_services/index'
// import popup from '@src/components/BigPopup.vue'
import { validationMixin } from 'vuelidate'
import { required, maxLength, minLength } from 'vuelidate/lib/validators'
import loading from '@src/components/Loading.vue'

import Step1 from './campagin-creation/Step1.vue'
import Step2 from './campagin-creation/Step2.vue'
import Step3 from './campagin-creation/Step3.vue'
import Step4 from './campagin-creation/Step4.vue'
import Step5 from './campagin-creation/Step5.vue'

const campaignDefault = {
  FbPageId: '',
  Id: '',
  CampaignName: '',
  CampaignObjective: 'Сообщения в директ',
  CampaignField: '',
  BusinessAddress: '',
  Images: [],
  ImagesDescription: [],
  ImagesSmall: [],
  ImagesSmallDescription: [],
  CreativeStatus: 'Есть рекламные креативы',
  PostDescription: '',
  DailyAmount: 0,
  Days: 0
}
const VUE_APP_API_URL = process.env.VUE_APP_API_URL

export default {
  name: 'CreateCampaign',
  mixins: [validationMixin],
  components: {
    // popup,
    loading,
    Step1,
    Step2,
    Step3,
    Step4,
    Step5
  },
  data () {
    return {
      currentStep: parseInt(localStorage.getItem('campagin_step')) || 1,
      totalSteps: 4,
      store,
      isLoading: false,
      isCampaignExist: false,
      isInfoPopupVisible: false,
      isRequestSent: false,
      pageSubmitted: false,
      imageNames: [],
      ImagesSmall: [],
      Images: [],
      campaign: campaignDefault,
      interval: false,
      displayNumber: 0
    }
  },
  computed: {
    number () {
      return (this.currentStep / this.totalSteps) * 100
    },
    isEdit () {
      return this.$route.query.isEdit === 'true'
    },
    label_cols () { return this.getWidth().label },
    content_cols () { return this.getWidth().content }
  },
  watch: {
    $route (to) {
      console.log('route ', store.getters.GET_EMAIL)
      window.scrollTo(0, 100)
      this.isLoading = false
      if (!(typeof to.params.id === 'undefined')) {
        axios.get({ url: `${VUE_APP_API_URL}/api/campaign/${to.params.id}` })
          .then(resp => {
            this.campaign = resp.data

            this.getImages()
            if (this.campaign.FbPageId !== '') {
              this.isRequestSent = true
              this.pageSubmitted = true
            }
          })
          .catch(err => {
            console.log(err)
            // TODO: remove, replaced by axios interceptor
            // if (err.response.status === 401) {
            //   store.dispatch('logout')
            //     .then(() => {
            //       this.$router.push('/login')
            //     })
            // }
          })
      } else {
        this.reset()
      }
    }
  },
  mounted () {
    clearInterval(this.interval)

    if (this.number === this.displayNumber) {
      return
    }

    this.interval = window.setInterval(function () {
      if (this.displayNumber !== this.number) {
        var change = (this.number - this.displayNumber) / 10
        change = change >= 0 ? Math.ceil(change) : Math.floor(change)
        this.displayNumber = this.displayNumber + change
      }
    }.bind(this), 20)
  },
  beforeMount () {
    console.log('BM ', store.getters.GET_EMAIL, 'campId:', this.$router.history.current?.params?.id)
    if (!this.$router.history.current?.params?.id === 'undefined') {
      axios.get({ url: `${VUE_APP_API_URL}/api/campaign/${this.$router.history.current.params.id}` })
        .then(resp => {
          this.campaign = resp.data

          this.getImages()
          if (this.campaign.FbPageId !== '') {
            this.isRequestSent = true
            this.pageSubmitted = true
          }
        })
        .catch(err => {
          console.log(err)
          // TODO: remove, replaced by axios interceptor
          // if (err.response.status === 401) {
          //   router.push({ path: '/login' })
          // }
        })
    }
  },
  validations: {
    campaign: {
      CampaignName: {
        required,
        maxLength: maxLength(30),
        minLength: minLength(3)
      },
      CampaignField: {
        required
      }
    }
  },
  // created() {
  //     window.addEventListener('resize', this.getWidth);
  //     this.getWidth();

  // },
  // destroyed() {
  //     window.removeEventListener('resize', this.getWidth);
  // },

  methods: {
    goStepBack () {
      this.currentStep = this.currentStep - 1
      localStorage.setItem('campagin_step', this.currentStep)
    },
    saveAndNext (data) {
      if (data) {
        this.campaign = {
          ...this.campaign,
          ...data
        }
      }

      const nextStepNumber = this.currentStep + 1
      if (this.totalSteps < nextStepNumber) {
        this.updateCampaign()
        localStorage.setItem('campagin_step', 1)

        // CHECK MONEY ENOUGH
        this.$router.push('/payment')
        // OR
        // this.$router.push('/main')
        return
      }

      // todo: mv to service
      localStorage.setItem('campagin_step', nextStepNumber)
      localStorage.setItem('campagin_data', JSON.stringify(this.campaign))

      if (!this.isEdit && this.currentStep === 2) {
        this.createCampaign()
      } else {
        this.updateCampaign()
      }
      this.currentStep++
    },
    reset () {
      this.isCampaignExist = false
      this.Images = []
      this.ImagesSmall = []
      this.isLoading = false
      this.campaign = campaignDefault
      this.isInfoPopupVisible = false
      this.isRequestSent = false
      this.pageSubmitted = false
    },
    getImages () {
      axios.get({ url: `${VUE_APP_API_URL}/api/campaign/${this.campaign.Id}/images/` })
        .then(resp => {
          if (resp.data == null) {
            this.imageNames = []
          } else {
            this.imageNames = resp.data
          }
        })
        .catch(err => {
          console.log(err)
          // TODO: remove, replaced by axios interceptor
          // if (err.response.status === 401) {
          //   store.dispatch('logout')
          //     .then(() => {
          //       this.$router.push('/login')
          //     })
          // }
        })
    },

    // TODO: extract
    getWidth () {
      return {
        label: 12,
        content: 12
      }
    },

    updateCampaign () {
      store.dispatch('saveCampaign', this.campaign)
    },

    createCampaign () {
      if (!this.isEdit) { // TODO isEdit = isInView for some time
        // this.$v.$touch()
        // if (this.$v.$anyError) {
        //   window.scrollTo(0, 100)
        //   return
        // }

        this.isLoading = true

        if (!this.isEdit) {
          console.log('dispatch')
          store.dispatch('saveCampaign', this.campaign)
            .then((resp) => {
              this.reset()
              router.push({ path: `/campaign-balance/${resp.data}`, query: {} }) // TODO QUERY
            })
            .catch(err => {
              this.isLoading = false
              console.log(err)
              // console.log(err.response)
              // console.log(err.response.data)
              // console.log(err.response.data.message)
              if (err.response.data.message === 'pq: duplicate key value violates unique constraint "ad_campaign_name_user_id_key"') {
                this.isCampaignExist = true
              }

              // TODO: remove, replaced by axios interceptor
              // if (err.response.status === 401) {
              //   store.dispatch('logout')
              //     .then(() => {
              //       this.$router.push('/login')
              //     })
              // }
            })
        } else {
          axios.get({ url: `${VUE_APP_API_URL}/api/campaign/${this.campaign.Id}`, data: campaignData })
            .then(() => {
              router.push({ path: `/campaign-balance/${this.campaign.Id}`, query: { isEdit: false } })
            })
            .catch(err => {
              this.isLoading = false
              console.log(err)
              if (err.response?.data.message === 'pq: duplicate key value violates unique constraint "ad_campaign_name_user_id_key"') {
                this.isCampaignExist = true
              }

              // TODO: remove, replaced by axios interceptor
              // if (err.response?.status === 401) {
              //   store.dispatch('logout')
              //     .then(() => {
              //       this.$router.push('/login')
              //     })
              // }
            })
        }
      } else {
        router.push({ path: `/campaign-balance/${this.campaign.Id}`, query: { isEdit: false } })
      }
    },

    logout () {
      this.pageSubmitted = false
      this.isRequestSent = false
      accountService.logout()
    }

  }
}
</script>
<style lang="scss">
  @import '@src/assets/styles/vars.scss';

  //#app-link-back {
  //  // float: left;
  //  margin-right: 5px;
  //  float:left;
  //  line-height:20px;
  //  color: #6c757d ;
  //
  //  &:hover {
  //    color:$black;
  //  }
  //}

  .app-new-steps-header {
    font-size: 14px;
    line-height: 18px;
    margin: 52px 0 12px !important;
  }

  #icon-div-image{
      position: absolute;
      margin-left: 140px;
      margin-top: -15px;
  }

  #primary-under{
      margin-left: 20px;
  }

  @media (max-width: 722px) {
      #c-status-text-u{
          margin-left: 30px;
      }
      #primary-under{
          margin-left: 0;
          margin-top: 20px;
      }
  }
  .custom-control-input{
      margin-right: 3px;
  }
</style>
