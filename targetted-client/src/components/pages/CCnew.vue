<template>
    <div>
      <loading v-if="isLoading" />
      <div id="content-wrapper">
        <div id="content">
          <router-link :to="{name: 'mainPage'}">
            <p id="navigation-text" style="margin:0;">← К списку кампаний</p>
          </router-link>
          <!-- <b-card class="mt-3" header="Form Data Result">
                <pre class="m-0">{{ company }}</pre>
            </b-card> -->
          <!-- <h1>{{isEdit ? "Редактирование":"Создание"}} кампании</h1> -->
          
          <div> 


            <Step1 :label_cols="label_cols"
                    :content_cols="content_cols"
                    :isEdit="isEdit"
                    @logout="logout"
                    v-if="currentStep ===  1"
                    @next="saveAndNext">

                     <template v-slot:header>
                       <h1 id="h2">Шаг 1.<br>О бизнесе</h1>
                     </template>
            </Step1>

            <Step2 :label_cols="label_cols"
                    :content_cols="content_cols"
                    :isEdit="isEdit"
                    v-if="currentStep ===  2"
                    :company="company"
                    @next="saveAndNext">
                    <template v-slot:header>
                      <h1 id="h2">Шаг 2.<br>Изображения </h1>
                     </template>
            </Step2>


            <Step3 :label_cols="label_cols"
                    :content_cols="content_cols"
                    :isEdit="isEdit"
                    v-if="currentStep ===  3"
                    :company="company"
                    @next="saveAndNext">
                    <template v-slot:header>
                      <h1 id="h2">Шаг 3.<br>Аудитория</h1>
                     </template>
            </Step3>

            <Step4 :label_cols="label_cols"
                    :content_cols="content_cols"
                    :isEdit="isEdit"
                    v-if="currentStep ===  4"
                    :company="company"
                    @next="saveAndNext">
                    <template v-slot:header>
                      <h1 id="h2">Шаг 4.<br>Бизнес-аккаунт</h1>
                     </template>
            </Step4>


            <Step5 :label_cols="label_cols"
                    :content_cols="content_cols"
                    :isEdit="isEdit"
                    v-if="currentStep ===  5"
                    :company="company"
                    @next="saveAndNext">
                    <template v-slot:header>
                      <h1 id="h2">Шаг 5.<br>Охват</h1>
                     </template>
            </Step5>

          </div>
        </div>
      </div>
    </div>

</template>
<script>
import accountService from '../../_services/account.service'
import store from '../../../store/store'
import router from '../../../router/router'
import {instance as axios} from '../../_services/index';
import popup from '../BigPopup.vue'
import { validationMixin } from 'vuelidate'
import { required, maxLength, minLength } from 'vuelidate/lib/validators'

import Step1 from './campagin-creation/Step1.vue';
import Step2 from './campagin-creation/Step2.vue';
import Step3 from './campagin-creation/Step3.vue';
import Step4 from './campagin-creation/Step4.vue';
import Step5 from './campagin-creation/Step5.vue';

const companyDefault = {
        FbPageId: '',
        Id: '',
        CompanyName: '',
        CompnayPurpose: 'Сообщения в директ',
        CompanyField: '',
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

import loading from '../Loading.vue'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL
export default {
  name: 'CreateCompany',
  mixins: [validationMixin],
  components: {
    popup,
    loading,
    Step1,
    Step2,
    Step3,
    Step4,
    Step5,
  },
  data () {
    return {
      currentStep: 1,
      totalSteps: 5,
      store,
      isLoading: false,
      isCompanyExist: false,
      isInfoPopupVisible: false,
      isRequestSent: false,
      pageSubmitted: false,
      imageNames: [],
      ImagesSmall: [],
      Images: [],
      company: companyDefault
    }
  },
  watch: {
    $route (to) {
      console.log('route ', store.getters.GET_EMAIL)
      window.scrollTo(0, 100)
      this.isLoading = false
      if (!(typeof to.params.id === 'undefined')) {
        axios({ url: `${VUE_APP_API_URL}/api/company/${to.params.id}`, method: 'GET' })
          .then(resp => {
            this.company = resp.data

            this.getImages()
            if (this.company.FbPageId !== '') {
              this.isRequestSent = true
              this.pageSubmitted = true
            }
          })
          .catch(err => {
            console.log(err)
            //TODO: remove, replaced by axios interceptor
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
  computed: {
    isEdit () {
      return this.$route.query.isEdit === 'true'
    },
    label_cols() { return this.getWidth().label},
    content_cols() { return this.getWidth().content },
  },
  beforeMount () {

    console.log('BM ', store.getters.GET_EMAIL)
    if (!(typeof this.$router.history.current.params.id === 'undefined')) {
      axios({ url: `${VUE_APP_API_URL}/api/company/${this.$router.history.current.params.id}`, method: 'GET' })
        .then(resp => {
          this.company = resp.data

          this.getImages()
          if (this.company.FbPageId !== '') {
            this.isRequestSent = true
            this.pageSubmitted = true
          }
        })
        .catch(err => {
          console.log(err)
          //TODO: remove, replaced by axios interceptor
          // if (err.response.status === 401) {
          //   router.push({ path: '/login' })
          // }
        })
    }
  },
  validations: {
    company: {
      CompanyName: {
        required,
        maxLength: maxLength(30),
        minLength: minLength(3)
      },
      CompanyField: {
        required
      }
    }
  },
  methods: {
    saveAndNext(data){
        // console.log(data.company)

        if(data && data.company){
          this.company = {
            ...this.company,
            ...data.company
          }
        }

        this.currentStep = this.currentStep + 1;
        if(this.currentStep > this.totalSteps) {
            this.createCompany();
        }
    },
    reset () {
      this.isCompanyExist = false
      this.Images = []
      this.ImagesSmall = []
      this.isLoading = false
      this.company = companyDefault;
      this.isInfoPopupVisible = false
      this.isRequestSent = false
      this.pageSubmitted = false
    },
    getImages () {
      axios({ url: `${VUE_APP_API_URL}/api/company/${this.company.Id}/images/`, method: 'GET' })
        .then(resp => {
          if (resp.data == null) {
            this.imageNames = []
          } else {
            this.imageNames = resp.data
          }
        })
        .catch(err => {
          console.log(err)
          //TODO: remove, replaced by axios interceptor
          // if (err.response.status === 401) {
          //   store.dispatch('logout')
          //     .then(() => {
          //       this.$router.push('/login')
          //     })
          // }
        })
    },


    //TODO: extract
    getWidth () {
      const width = Math.max(
        document.body.scrollWidth,
        document.documentElement.scrollWidth,
        document.body.offsetWidth,
        document.documentElement.offsetWidth,
        document.documentElement.clientWidth
      )
      if (width < 570) {
        return {
          label: 12,
          content: 12
        }
      }
      return {
        label: 3,
        content: 9
      }
    },

    createCompany () {
      if (!this.isEdit) { // TODO isEdit = isInView for some time
        this.$v.$touch()
        if (this.$v.$anyError) {
          window.scrollTo(0, 100)
          return
        }
        
        this.isLoading = true
        const companyData = new FormData()
        companyData.append('FbPageId', this.company.FbPageId)
        if (this.isEdit) {
          companyData.append('Id', this.company.Id)
        }
        companyData.append('CompanyName', this.company.CompanyName)
        companyData.append('CompnayPurpose', this.company.CompnayPurpose)
        companyData.append('CompanyField', this.company.CompanyField)
        companyData.append('BusinessAddress', this.company.BusinessAddress)
        companyData.append('ImagesDescription', this.company.ImagesDescription)
        companyData.append('ImagesSmallDescription', this.company.ImagesSmallDescription)
        companyData.append('CreativeStatus', this.company.CreativeStatus)
        companyData.append('PostDescription', this.company.PostDescription)
        companyData.append('DailyAmount', this.company.DailyAmount)
        companyData.append('Days', this.company.Days)
        Array.from(this.ImagesSmall).forEach(Image => {
          companyData.append('ImageSmall', Image)
        })
        Array.from(this.Images).forEach(Image => {
          companyData.append('Image', Image)
        })
        if (!this.isEdit) {
          console.log('dispatch')
          store.dispatch('saveCompany', companyData)
            .then((resp) => {
              this.reset()
              router.push({ path: '/company-balance/' + resp.data, query: {} }) // TODO QUERY
            })
            .catch(err => {
              this.isLoading = false
              console.log(err)
              // console.log(err.response)
              // console.log(err.response.data)
              // console.log(err.response.data.message)
              if (err.response.data.message === 'pq: duplicate key value violates unique constraint "ad_company_name_user_id_key"') {
                this.isCompanyExist = true
              }
              
              //TODO: remove, replaced by axios interceptor
              // if (err.response.status === 401) {
              //   store.dispatch('logout')
              //     .then(() => {
              //       this.$router.push('/login')
              //     })
              // }
            })
        } else {
          axios({ url: `${VUE_APP_API_URL}/api/company/${this.company.Id}`, data: companyData, method: 'PUT' })
            .then(resp => {
              console.log(resp.status)
              router.push({ path: '/company-balance/' + this.company.Id, query: { isEdit: false } })
            })
            .catch(err => {
              this.isLoading = false
              console.log(err)
              if (err.response?.data.message === 'pq: duplicate key value violates unique constraint "ad_company_name_user_id_key"') {
                this.isCompanyExist = true
              }
              
              //TODO: remove, replaced by axios interceptor
              // if (err.response?.status === 401) {
              //   store.dispatch('logout')
              //     .then(() => {
              //       this.$router.push('/login')
              //     })
              // }
            })
        }
      } else {
        router.push({ path: '/company-balance/' + this.company.Id, query: { isEdit: false } })
      }
    },

    logout () {
      this.pageSubmitted = false
      this.isRequestSent = false
      accountService.logout()
    },

  }
}
</script>
<style>
.num-wrapper{
    display: flex;
    align-items: stretch;
    margin-bottom: 16px;
}
.num{
    background: #F3F3F3;
    width: 48px;
    height: 48px;
    border-radius: 24px;
}
.num-num{
    padding-top: 11px;
    padding-left: 19px;
}
.num-text{
    margin-left: 16px;
    max-width: 600px;
}
.creative-message{
    margin: 28px 0px 40px 0px;
    padding: 40px;
    width: 800px;
    height: 128px;
    background: #F3F3F3;
    border-radius: 20px;
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
        margin-left: 0px;
        margin-top: 20px;
    }
    .creative-message{
        width: auto!important;
        height: auto!important;
        padding: 25px!important;
    }
}
.custom-radio{
    margin: 10px;
}
.custom-control-input{
    margin-right: 3px;
}
#preview{
    width: 160px;
    height: 280px;
    object-fit: cover;
    border-radius: 20px;
}
#preview-small{
    width: 160px;
    height: 160px;
    object-fit: cover;
    border-radius: 20px;
}
#image-block {
    max-width: 530px;
    display: grid;
    grid-template-columns: repeat(auto-fill,minmax(160px, 1fr));
    justify-content: space-between;
    align-items: center;
    grid-gap: 24px;
}
@media (max-width: 790px) {
    #image-block {
    max-width: 350px;
    }
}
@media (max-width: 450px) {
    #image-block {
        max-width: 200px;
    }
}
#load-frame {
    border: 2px dashed #CCCCCC;
    border-radius: 20px;
    width: 160px;
    height: 280px;
}
#load-frame-small {
    border: 2px dashed #CCCCCC;
    border-radius: 20px;
    width: 160px;
    height: 160px;
}
#load-frame:hover {
    background: #F3F3F3;
}
#load-frame-small:hover {
    background: #F3F3F3;
}
#load-file {
    font-family: Montserrat;
    font-style: normal;
    font-weight: 600;
    font-size: 16px;
    line-height: 24px;
    margin: 20px auto;
    text-align: center;
    color: #000000;
}
#file-size-big {
    font-family: Montserrat;
    font-style: normal;
    font-weight: normal;
    font-size: 12px;
    line-height: 20px;
    /* bottom: 0;//TODO почему то не работает */
    margin-top: 140px;
    text-align: center;
    color: #767676;
}
#file-size {
    font-family: Montserrat;
    font-style: normal;
    font-weight: normal;
    font-size: 12px;
    line-height: 20px;
    bottom: 0;
    text-align: center;
    color: #767676;
}
#list {
  font-family: Montserrat;
  font-style: normal;
  font-weight: normal;
  font-size: 20px;
  line-height: 32px;

  color: #000000;

}
#list-item {
    border-radius: 8px;
    border: none;
    background-color: #e4e4e4;
    margin: 5px;

}
</style>
