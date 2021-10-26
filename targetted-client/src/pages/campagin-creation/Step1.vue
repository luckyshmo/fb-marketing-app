<template>
  <section class="step1">
    <slot name="header"></slot>
    <p>Заполните информацию о своём бизнесе.</p>
      <br class="d-none d-sm-block">
      <br class="d-none d-sm-block">
    <b-form>
      <!-- <div v-if="!(store.getters.GET_FB_PAGES.length > 0) && !isRequestSent && !pageSubmitted">
        <p>Раздайте доступ к вашей Facebook и Instagram странице<br>для запуска и управления рекламой от имени ваших
          страниц. </p>

      </div> -->
      <!-- <div v-if="store.getters.GET_FB_PAGES.length > 0 && !isRequestSent && !pageSubmitted">
        <div>
          <p>Выберите страницу которую хотите привязать</p>
          <b-form-group label="Выберите страницу" :label-cols="label_cols" :content-cols="content_cols"
            id="input-group-main" label-for="input-horizontal">
            <b-form-radio-group v-model="company.FbPageId" :options="store.getters.GET_FB_PAGES"></b-form-radio-group>
          </b-form-group>
        </div>
        <b-button class="main-button" @click="sendFbRequest()">
          Привязать
        </b-button>
      </div> -->
      <!-- <div v-if="isRequestSent && !pageSubmitted">
        <p>Зайди в аккаунт на Facebook и подтверди привязку страницы в сообщениях</p>
        <b-button class="main-button" target="_blank" rel="noopener noreferrer" :href=getFBRedirect()>
          Перейти в facebook
        </b-button>
        <b-button class="main-button" id="primary-under" @click="checkPageSubmitted()">
          Я подтвердил в сообщениях
        </b-button>
      </div> -->
      <!-- <div v-if="pageSubmitted">
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
      </b-button> -->
      <b-form-group
        label="Что рекламируете"
        :label-cols="label_cols"
        :content-cols="content_cols"
        id="input-group-main"
        label-for="input-horizontal">
        <b-form-radio-group
          :disabled="isEdit"
          v-model="company.CompanyTarget"
          :options="['Аккаунт в Instagram или Facebook', 'Сайт']"
          class="app-new-radio" />
<!--          @change="setFieldValue('CompanyTarget', $event)" />-->
<!--        <small-->
<!--          v-if="customErrors.company.CompanyTarget"-->
<!--          class="error-message">-->
<!--          Не указана цель рекламы-->
<!--        </small>-->
      </b-form-group>

      <div v-if="this.company.CompanyTarget !== undefined">

        <b-form-group
          label="Ссылка на аккаунт"
          :label-cols="label_cols"
          :content-cols="content_cols"
          id="input-group-main"
          label-for="input-horizontal">
          <b-form-input class="form-input"
                        :disabled="isEdit"
                        :state="$v.company.AccountURL.$dirty ? !$v.company.AccountURL.$error : null"
                        v-model="$v.company.AccountURL.$model"
                        placeholder="Введите ссылку" />
          <small
            v-if="$v.company.AccountURL.$dirty && !$v.company.AccountURL.required"
            class="error-message">
            Укажите URL
          </small>
        </b-form-group>

        <b-form-group
          label="Категория деятельности"
          :label-cols="label_cols"
          :content-cols="content_cols"
          id="input-group-main"
          label-for="input-horizontal">
          <b-form-select
            class="form-select"
            v-model="company.CompanyFieldArea"
            :options="['Выберите из спиcка', 'Интернет-реклама', 'Сайт']"
            @change="setFieldValue('CompanyFieldArea', $event)" />
          <p
            v-if="company.CompanyFieldArea === 'Интернет-реклама'"
            class="app-new-info__bottom">
            Стоимость перехода по рекламе ~ 10–30 ₽
          </p>
          <small
            v-if="customErrors.company.CompanyFieldArea"
            class="error-message">
            Не указана cфера деятельности
          </small>
        </b-form-group>

        <b-form-group
          label="Особенности бизнеса"
          :label-cols="label_cols"
          :content-cols="content_cols"
          id="input-group-main"
          label-for="input-horizontal">
          <p class="app-new-info">
            Укажите услуги, которые оказываете
          </p>
          <b-form-textarea
            class="form-input"
            :disabled="isEdit"
            :state="$v.company.CompanyField.$dirty ? !$v.company.CompanyField.$error : null"
            v-model="$v.company.CompanyField.$model"
            rows="6"
            max-rows="9"
            placeholder="Введите описание">
          </b-form-textarea>
          <small
            v-if="$v.company.CompanyField.$dirty && !$v.company.CompanyField.required"
            class="error-message">
            Пустое поле
          </small>
        </b-form-group>

        <!--    <b-form-group-->
        <!--      label="Название кампании"-->
        <!--      :label-cols="label_cols"-->
        <!--      :content-cols="content_cols"-->
        <!--      id="input-group-main"-->
        <!--      label-for="input-horizontal">-->
        <!--      <b-form-input-->
        <!--        class="form-input"-->
        <!--        v-model="$v.company.CompanyName.$model"-->
        <!--        :state="$v.company.CompanyName.$dirty ? !$v.company.CompanyName.$error : null"-->
        <!--        :disabled="isEdit"-->
        <!--        placeholder="Введите название"-->
        <!--        @click="resetNameErr()">-->
        <!--      </b-form-input>-->
        <!--      <small-->
        <!--        v-if="$v.company.CompanyName.$dirty && !$v.company.CompanyName.required"-->
        <!--        class="error-message">-->
        <!--        Название должно быть между 3 и 30 символами-->
        <!--      </small>-->
        <!--      <small v-if="isCompanyExist" class="error-message">-->
        <!--        Кампания с таким именем уже создана-->
        <!--      </small>-->
        <!--    </b-form-group>-->

        <b-form-group
          label="Где оказываете услуги"
          :label-cols="label_cols"
          :content-cols="content_cols"
          id="input-group-main"
          label-for="input-horizontal">
          <p class="app-new-info">
            Если у вас офлайн бизнес, то укажите точный адрес оказания услуг
          </p>
          <b-form-input
            class="form-input"
            :disabled="isEdit"
            :state="$v.company.BusinessAddress ? !$v.company.BusinessAddress.$error : null"
            v-model="$v.company.BusinessAddress.$model"
            placeholder="Город или улица">
          </b-form-input>
          <small
            v-if="$v.company.BusinessAddress.$dirty && !$v.company.BusinessAddress.required"
            class="error-message">
            Укажите адрес
          </small>
        </b-form-group>

        <b-form-group
          label="Цель кампании"
          :label-cols="label_cols"
          :content-cols="content_cols"
          id="input-group-main"
          label-for="input-horizontal">
          <b-form-radio-group
            v-model="company.CompanyPurpose"
            :disabled="isEdit"
            class="app-new-radio"
            :options="[
            'Заявки в директ',
            'Новые подписчики',
            'Заявки через лид-форму'
          ]" />
        </b-form-group>

        <b-form-group
          label="Номер телефона"
          id="input-group-main">
          <p class="app-new-info">
            Мы не покажем его в рекламе
          </p>
          <b-form-input
            :state="$v.userData.phone.$dirty ? !$v.userData.phone.$error : null"
            v-model="$v.userData.phone.$model"
            class="form-input width-1-2"
            placeholder="Введите телефон" />
          <small
            v-if="$v.userData.phone.$dirty && !$v.userData.phone.required"
            class="error-message">
            Телефон на указан
          </small>
          <small
            v-if="$v.userData.phone.$dirty && !$v.userData.phone.numeric"
            class="error-message">
            Только цифры в телефоне
          </small>
        </b-form-group>

        <b-form-group class="mt-5">
          <b-row>
            <b-col cols="6" sm="8" md="3" lg="3" xl="3">
              <b-button
                type="button"
                :class="{'disabled-look': !$v.$anyDirty || $v.$anyError}"
                class="app-new-submit-button"
                @click="sendData">
                {{isEdit ? "Назад":"Продолжить"}}
              </b-button>
            </b-col>
            <b-col cols="6" sm="8" md="6" lg="6" xl="6" class="app-new-small-text-fit">
              <p class="text-muted text-pt-desktop">
                Заполните все данные, чтобы продолжить
              </p>
            </b-col>
          </b-row>
        </b-form-group>

      </div>

    </b-form>
  </section>
</template>

<script>
import store from '@src/store/store'
import router from '@src/router/router'
import popup from '@src/components/BigPopup.vue'
import { instance as axios } from '@src/_services/index'

import { validationMixin } from 'vuelidate'
import { required, maxLength, minLength, numeric } from 'vuelidate/lib/validators'

const VUE_APP_API_URL = process.env.VUE_APP_API_URL

export default {
  name: 'Step1',
  props: ['label_cols', 'content_cols', 'isEdit'],
  mixins: [validationMixin],
  components: {
    popup
  },
  validations: {
    userData: {
      phone: {
        required,
        numeric
        // minLength: minLength(7)
      }
    },
    company: {
      CompanyName: {
        required,
        maxLength: maxLength(30),
        minLength: minLength(3)
      },
      BusinessAddress: {
        required
      },
      AccountURL: {
        required
        // minLength: minLength(10)
      },
      CompanyField: {
        required
      },
      CompanyFieldArea: {
        required
      }
    }
  },
  data: function () {
    return {
      store, // fixme
      isLoading: false,
      isRequestSent: false,
      pageSubmitted: false,
      isCompanyExist: false,
      userData: {
        phone: ''
      },
      company: {
        FbPageId: '',
        CompanyTarget: '',
        UserId: '',
        Id: '',
        CompanyName: '',
        CompanyPurpose: 'Сообщения в директ',
        CompanyField: '',
        BusinessAddress: '',
        AccountURL: '',
        CompanyFieldArea: 'Выберите из спиcка',
        ImagesDescription: [],
        ImagesSmallDescription: [],
        CreativeStatus: 'Есть рекламные креативы',
        PostDescription: '',
        DailyAmount: 0,
        Days: 0
      },
      customErrors: {
        company: {
          CompanyFieldArea: false
        },
        anyError: false
      }
    }
  },
  methods: {
    sendData () {
      // валидация с помощью vuelidate
      this.$v.company.$touch()

      // кастомная валидация для checkbox/select
      Object.keys(this.customErrors.company).forEach(nameField => {
        if (!this.company[nameField]) {
          this.customErrors.company[nameField] = true
          this.customErrors.anyError = true
        }
      })

      if (this.$v.company.$anyError || this.customErrors.anyError) {
        window.scrollTo(0, 100)
        return
      }

      this.$emit('next', this.company)
    },
    setFieldValue (fieldName, value) {
      if (this.company[fieldName]) {
        this.company[fieldName] = value

        if (this.customErrors.company[fieldName]) {
          this.customErrors.company[fieldName] = false
        }
      }
    },
    // validateState(name) {
    //   const {
    //     $dirty,
    //     $error
    //   } = this.$v.company[name]
    //   return $dirty ? !$error : null
    // },
    resetNameErr () {
      this.isCompanyExist = false
    },
    getFBRedirect () {
      return `https://facebook.com/${this.company.FbPageId}/settings/?tab=admin_roles`
    },
    sendFbRequest () {
      this.isLoading = true
      axios.post({ url: `${VUE_APP_API_URL}/api/facebook/claim/${this.company.FbPageId}` })
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
      axios.get({ url: `${VUE_APP_API_URL}/api/facebook/owned/${this.company.FbPageId}` })
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
    }
  },

  mounted () {
    const cd = localStorage.getItem('campagin_data')
    try {
      const cdd = JSON.parse(cd)
      this.company = {
        ...cdd
      }
      // eslint-disable-next-line no-empty
    } catch {

    }
  },

  beforeMount () {
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
  select.form-input.custom-select {
      width: 640px;
      padding: 10px;
      margin-bottom: 12px;
  }
  @media(max-width: 550px) {
    select.form-input.custom-select {
      width: 99%;
    }
  }
</style>
