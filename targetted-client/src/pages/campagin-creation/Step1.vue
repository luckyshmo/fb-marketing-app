<template>
  <div>
    <slot name="header"></slot>
    <br>
    <b-form>
      <div v-if="!(store.getters.GET_FB_PAGES.length > 0) && !isRequestSent && !pageSubmitted">
        <p>Раздайте доступ к вашей Facebook и Instagram странице<br>для запуска и управления рекламой от имени ваших
          страниц. </p>
   
      </div>
      <div v-if="store.getters.GET_FB_PAGES.length > 0 && !isRequestSent && !pageSubmitted">
        <div>
          <p>Выберите страницу которую хотите привязать</p>
          <!-- //TODO SELECT -->
          <b-form-group label="Выберите страницу" :label-cols="label_cols" :content-cols="content_cols"
            id="input-group-main" label-for="input-horizontal">
            <b-form-radio-group v-model="company.FbPageId" :options="store.getters.GET_FB_PAGES"></b-form-radio-group>
          </b-form-group>
        </div>
        <b-button class="main-button" @click="sendFbRequest()">
          Привязать
        </b-button>
      </div>
      <div v-if="isRequestSent && !pageSubmitted">
        <p>Зайди в аккаунт на Facebook и подтверди привязку страницы в сообщениях</p>
        <b-button class="main-button" target="_blank" rel="noopener noreferrer" :href=getFBRedirect()>
          Перейти в facebook
        </b-button>
        <b-button class="main-button" id="primary-under" @click="checkPageSubmitted()">
          Я подтвердил в сообщениях
        </b-button>
      </div>
      <div v-if="pageSubmitted">
        <div class="c-status" style="margin-top: 30px; max-width: 800px">
          <div class="elipse" id="green"></div>
          <p class="c-status-text" id="c-status-text-u" style="text-align: left;">Страница {{company.FbPageId}} привязана
            к targetted</p>
        </div>
      </div>
      <b-button v-if="store.getters.GET_FB_PAGES.length > 0 ||
              (isRequestSent && pageSubmitted) ||
              (store.getters.GET_FB_PAGES.length == 0 || isRequestSent)" class="main-button-grey"
        style="margin-top: 30px; background: #F3F3F3; color: black" @click="this.$emit('logout')">
        Привязать другой аккаунт
      </b-button>



      <h2 id="h2">Основное</h2>

      <b-form-group label="Название кампании" :label-cols="label_cols" :content-cols="content_cols" id="input-group-main"
      label-for="input-horizontal">
      <b-form-input class="form-input" v-model="$v.company.CompanyName.$model" :disabled="isEdit"
          placeholder="Введите название"
          @click="resetNameErr()"></b-form-input>
          <small
              v-if="$v.company.CompanyName.$dirty && !$v.company.CompanyName.required"
              class="error-message"
            >
               Название должно быть между 3 и 30 символами
            </small>
      <small v-if="isCompanyExist" class="error-message">
          Кампания с таким иминем уже создана
      </small>
      </b-form-group>

      <b-form-group label="Цель кампании" :label-cols="label_cols" :content-cols="content_cols" id="input-group-main"
      label-for="input-horizontal">
      <b-form-radio-group v-model="company.CompnayPurpose" :disabled="isEdit" :options="[
                      'Сообщения в директ',
                      'Подписки в instagram',
                      'Заявки через форму обратной связи',
                      'Целевое действие на сайте'
                  ]"></b-form-radio-group>
      </b-form-group>

      <b-form-group label="Сфера деятельности" :label-cols="label_cols" :content-cols="content_cols" id="input-group-main"
      label-for="input-horizontal">
      <b-form-input class="form-input" :disabled="isEdit" v-model="$v.company.CompanyField.$model"
          placeholder="Вставьте ссылку"></b-form-input>
           <small
              v-if="$v.company.CompanyField.$dirty && !$v.company.CompanyField.required"
              class="error-message"
            >
              Пустое поле 
            </small>
      </b-form-group>

          <b-form-group label="Адрес бизнеса" :label-cols="label_cols" :content-cols="content_cols" id="input-group-main"
          label-for="input-horizontal" description="Нужен только для офлайн бизнеса">
          <b-form-input class="form-input" :disabled="isEdit" v-model="$v.company.BusinessAddress.$model" placeholder="Точный адрес">
          </b-form-input>
             <small
              v-if="$v.company.BusinessAddress.$dirty && !$v.company.BusinessAddress.required"
              class="error-message"
            >
              Укажите адрес
            </small>
      </b-form-group> 
    
      <b-form-group class="mt-5">

          <b-row>
                <b-col cols="6" sm="8" md="3" lg="3" xl="3">
              <b-button type="button"
                        class="app-new-submit-button"
                        @click="sendData">
                  {{isEdit ? "Назад":"Продолжить"}}
              </b-button>
                </b-col>
                <b-col cols="6" sm="8" md="6" lg="6" xl="6" class="app-new-small-text-fit">
                  <p class="text-muted">Заполните все данные, чтобы продолжить</p>
                </b-col>
          </b-row>
      </b-form-group>
            
      </b-form>
  </div>
</template>

<script>
import store from '@/store/store'
import router from '@/router/router'
import popup from '@/components/BigPopup.vue'

import { validationMixin } from 'vuelidate'
import { required, maxLength, minLength } from 'vuelidate/lib/validators'

const VUE_APP_API_URL = process.env.VUE_APP_API_URL

export default {
    name: 'Step1',
      props: ['label_cols', 'content_cols', 'isEdit'],
      mixins: [validationMixin],
      components: {
        popup
      },
      validations: {
        company: {
          CompanyName: {
            required,
            maxLength: maxLength(30),
            minLength: minLength(3)
          },
          BusinessAddress: {
            required
          },
          CompanyField: {
            required
          }
        }
      },
      data: function () {
        return {
          store, //fixme
          isLoading: false,
          isRequestSent: false,
          pageSubmitted: false,
          isCompanyExist: false,
          company: {
            FbPageId: '',
            UserId: '',
            Id: '',
            CompanyName: '',
            CompnayPurpose: 'Сообщения в директ',
            CompanyField: '',
            BusinessAddress: '',
            ImagesDescription: [],
            ImagesSmallDescription: [],
            CreativeStatus: 'Есть рекламные креативы',
            PostDescription: '',
            DailyAmount: 0,
            Days: 0
          }
        }
      },
      methods: {
        sendData() {
          this.$v.company.$touch()
          if (this.$v.company.$anyError) {
            return
          }

          this.$emit('next', this.company)
        },
  
        // validateState(name) {
        //   const {
        //     $dirty,
        //     $error
        //   } = this.$v.company[name]
        //   return $dirty ? !$error : null
        // },
        resetNameErr() {
          this.isCompanyExist = false
        },
        getFBRedirect () {
          return `https://facebook.com/${this.company.FbPageId}/settings/?tab=admin_roles`
       },
        sendFbRequest () {
          this.isLoading = true
          axios.post({ url: `${VUE_APP_API_URL}/api/facebook/claim/${this.company.FbPageId}`})
            .then(resp => {
              this.isLoading = false
              console.log(resp)
              this.isRequestSent = true
            })
            .catch(err => {
              this.isLoading = false
              this.$alert(err.response.data.message)
            })
        }, 
        checkPageSubmitted () {
          this.isLoading = true
          axios.get({ url: `${VUE_APP_API_URL}/api/facebook/owned/${this.company.FbPageId}`})
            .then(resp => {
              this.isLoading = false
              console.log(resp)
              this.pageSubmitted = true
            })
            .catch(err => {
              this.isLoading = false
              console.log(err.response.data)
              this.$alert(err.response.data.message)
            })
        },
      },
      
      beforeMount() {

        if (!(typeof this.$router.history.current.params.id === 'undefined')) {
          axios.get({
              url: `${VUE_APP_API_URL}/api/company/${this.$router.history.current.params.id}`
            })
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
              if (err.response.status === 401) {
                router.push({
                  path: '/login'
                })
              }
            })
        }
      }
}
</script>

<style>

</style>