<template>
    <div>
      <loading v-if="isLoading" />
      <div id="content-wrapper">
        <div id="content">
          <b-row class="app-new-steps-header">
            <b-col cols="8">
              <router-link id="app-link-back" v-if="currentStep === 1" :to="{name: 'mainPage'}">
                ← Назад
              </router-link>
              <p class="text-muted clickable" v-else @click="goStepBack" style="float:left;margin-right: 4px;line-height: 20px;">← Назад</p>
         
                <div class="text-muted d-block d-md-none d-lg-none d-xl-none">&nbsp;∙&nbsp;Шаг {{currentStep}} из {{totalSteps}}</div>
            </b-col>
          </b-row>
        
            <b-row>
              <b-col sm="12" md="8" lg="8" xl="8" class="app-new-main-content"> 

              <Step1 :label_cols="label_cols"
                      :content_cols="content_cols"
                      :isEdit="isEdit"
                      @logout="logout"
                      v-if="currentStep ===  1"
                      @next="saveAndNext">

                      <template v-slot:header>
                        <h1 id="h2">О бизнесе</h1>
                      </template>
              </Step1>

              <Step2 :label_cols="label_cols"
                      :content_cols="content_cols"
                      :isEdit="isEdit"
                      v-if="currentStep ===  2"
                      :company="company"
                      @next="saveAndNext">
                      <template v-slot:header>
                        <h1 id="h2">Изображения </h1>
                      </template>
              </Step2>

              <Step3 :label_cols="label_cols"
                      :content_cols="content_cols"
                      :isEdit="isEdit"
                      v-if="currentStep ===  3"
                      :company="company"
                      @next="saveAndNext">
                      <template v-slot:header>
                        <h1 id="h2">Аудитория</h1>
                      </template>
              </Step3>
<!-- 
              <Step4 :label_cols="label_cols"
                      :content_cols="content_cols"
                      :isEdit="isEdit"
                      v-if="currentStep ===  4"
                      :company="company"
                      @next="saveAndNext">
                      <template v-slot:header>
                        <h1 id="h2">Привязка страницы</h1>
                      </template>
              </Step4> -->

              <Step5 :label_cols="label_cols"
                      :content_cols="content_cols"
                      :isEdit="isEdit"
                      v-if="currentStep ===  4"
                      :company="company"
                      @next="saveAndNext">
                      <template v-slot:header>
                        <h1 id="h2">Охват рекламы</h1>
                      </template>
              </Step5>

          </b-col>
          <b-col cols="3" class=" d-none d-md-block d-lg-block d-xl-block">
          <aside>
            <h3 class="app-new-header">
              Заполнено {{(displayNumber)*100}}%
              </h3>
              <ul class="app-new-progress-text">
                <li :class="{active:(currentStep>=1)}">
                  О бизнесе
                </li>
                <li :class="{active:(currentStep>=2)}">
                  Изображения
                </li>
                <!-- <li :class="{active:(currentStep>=3)}">
                  Привязка аккаунта
                </li> -->
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
        </div>
      </div>
    </div>

</template>
<script>
import accountService from '@/_services/account.service'
import store from '@/store/store'
import router from '@/router/router'
import {instance as axios} from '@/_services/index';
import popup from '@/components/BigPopup.vue'
import { validationMixin } from 'vuelidate'
import { required, maxLength, minLength } from 'vuelidate/lib/validators'
import loading from '@/components/Loading.vue'

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
  };
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
      currentStep: parseInt(localStorage.getItem('campagin_step')) || 1,
      totalSteps: 4,
      store,
      isLoading: false,
      isCompanyExist: false,
      isInfoPopupVisible: false,
      isRequestSent: false,
      pageSubmitted: false,
      imageNames: [],
      ImagesSmall: [],
      Images: [],
      company: companyDefault,
      interval:false,
      displayNumber:0,
    }
  },
  computed: {
    number() {
      return this.currentStep/this.totalSteps;
    },
    isEdit () {
      return this.$route.query.isEdit === 'true'
    },
    label_cols() { return this.getWidth().label},
    content_cols() { return this.getWidth().content },
  },
  watch: {
    	number: function(){
    
    	clearInterval(this.interval);
    
    		if(this.number == this.displayNumber){
          return;
        }
    
      	this.interval = window.setInterval(function(){
        
        	if(this.displayNumber != this.number){
        
        		var change = (this.number - this.displayNumber) / 10;
            
            change = change >= 0 ? Math.ceil(change) : Math.floor(change);
          
          	this.displayNumber = this.displayNumber + change;
          
          }
        
        }.bind(this), 20);
      
    },
    $route (to) {
      console.log('route ', store.getters.GET_EMAIL)
      window.scrollTo(0, 100)
      this.isLoading = false
      if (!(typeof to.params.id === 'undefined')) {
        axios.get({ url: `${VUE_APP_API_URL}/api/company/${to.params.id}`})
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
  beforeMount () {

    console.log('BM ', store.getters.GET_EMAIL, 'campId:', this.$router.history.current?.params?.id )
    if (!this.$router.history.current?.params?.id === 'undefined') {
      axios.get({ url: `${VUE_APP_API_URL}/api/company/${this.$router.history.current.params.id}`})
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
  created() {
      window.addEventListener('resize', this.getWidth);
      this.getWidth();
  },
  destroyed() {
      window.removeEventListener('resize', this.getWidth);
  },

  methods: {
    goStepBack(){
      this.currentStep = this.currentStep - 1;
      localStorage.setItem('campagin_step', this.currentStep)
    },
    saveAndNext(data){

        if(data){
          this.company = {
            ...this.company,
            ...data
          }
        }

        const nextStepNumber = this.currentStep + 1;
        if(this.totalSteps < nextStepNumber) {
          this.updateCompany();
          localStorage.setItem('campagin_step', 1)

          //CHECK MONEY ENOUGH
          this.$router.push('/payment')
          //OR 
          //this.$router.push('/main')
          return
        }

        //todo: mv to service
        localStorage.setItem('campagin_step', nextStepNumber)
        localStorage.setItem('campagin_data', JSON.stringify(this.company));

        if(!this.isEdit && this.currentStep === 2) {
            this.createCompany();
        } else {
            this.updateCompany();
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
      axios.get({ url: `${VUE_APP_API_URL}/api/company/${this.company.Id}/images/`})
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
        return {
          label: 12,
          content: 12
        }
    },

    updateCompany(){
      store.dispatch('saveCompany', this.company)
    },

    createCompany () {
      if (!this.isEdit) { // TODO isEdit = isInView for some time
        // this.$v.$touch()
        // if (this.$v.$anyError) {
        //   window.scrollTo(0, 100)
        //   return
        // }
        
        this.isLoading = true
        
        if (!this.isEdit) {
          console.log('dispatch')
          store.dispatch('saveCompany', this.company)
            .then((resp) => {
              this.reset()
              router.push({ path: `/company-balance/${resp.data}`, query: {} }) // TODO QUERY
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
          axios.get({ url: `${VUE_APP_API_URL}/api/company/${this.company.Id}`, data: companyData})
            .then(() => {
              router.push({ path: `/company-balance/${this.company.Id}`, query: { isEdit: false } })
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
        router.push({ path: `/company-balance/${this.company.Id}`, query: { isEdit: false } })
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
<style lang="scss">
  @import '@/assets/styles/vars.scss';

#app-link-back {
  // float: left;
  margin-right: 5px;
  float:left;
  line-height:20px;
  color: #6c757d ;

  &:hover {
    color:$black;
  }
}

.app-new-progress-text  {
  list-style: none;
  margin-left: -20px;
  padding-left: 0px;
}
.app-new-progress-text li {
  color: $gray;
  font-size: $baseFont;
  line-height: 46px;

  &::before {
    margin-right: 10px;
    color: $gray;
    content: "•";
  }
  
  &.active {
    color: $black;
    
    &::before {
      color: $lime;
      content: "•";
    }
  }
}

.app-new-steps-header {
  font-size: 14px;
  margin-bottom: -10px;
  & p {
    font-size: 14px;
  }
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
}

aside {
  margin-top: 35px;
}

.custom-control-input{
    margin-right: 3px;
}

@media (max-width: 790px) {
    #image-block {
    max-width: 350px;
    }
}
</style>
