<template>
    <div>
        <loading v-if="isLoading"/>
        <div id="content-wrapper">
            <div id="content">
            <router-link :to="{name: 'mainPage'}">
                <p id="navigation-text" style="margin:0;">← К списку кампаний</p>
            </router-link>
            <!-- <b-card class="mt-3" header="Form Data Result">
                <pre class="m-0">{{ company }}</pre>
            </b-card> -->
            <h1>{{isEdit ? "Редактирование":"Создание"}} кампании</h1>
            <div>
                <b-form @submit.prevent="createCompany()">
                    <div>
                        <h2 id="h2">Доступ к странице Facebook и Instagram</h2>
                        <div v-if="!(store.getters.GET_FB_PAGES.length > 0) && !isRequestSent && !pageSubmitted">
                            <p>Раздайте доступ к вашей Facebook и Instagram странице<br>для запуска и управления рекламой от имени ваших страниц. </p>
                            <b-button
                                v-if="!(store.getters.GET_FB_PAGES.length > 0)"

                                class="main-button"
                                @click="loginFB"
                            >
                                У меня есть бизнес-аккаунт
                            </b-button>
                            <!-- //TODO ситуация с осутсвием страниц -->
                            <popup
                            v-if="isInfoPopupVisible"
                            popupTitle="Инструкция по созданию бизнесс-аккаунта"
                            @closePopup="closeInfoPopup"
                            >
                                <div style="text-align: left; padding: 30px">
                                    <h1 style="margin=0px !important;">Как создать<br>бизнес-аккаунт</h1>
                                    <p>Профессиональный аккаунт в Instagram необходим для того,<br>чтобы мы могли запустить и управлять рекламными кампаниями <br>от имени вашего бизнеса. </p>
                                    <div class="num-wrapper">
                                        <div class="num">
                                            <div class="num-num">
                                                1
                                            </div>
                                        </div>
                                        <p class="num-text">Откройте мобильное приложение Инстаграм.</p>
                                    </div>
                                    <div class="num-wrapper">
                                        <div class="num">
                                            <div class="num-num">
                                                2
                                            </div>
                                        </div>
                                        <p class="num-text">Зайдите в настройки профиля. </p>
                                    </div>
                                    <div class="num-wrapper">
                                        <div class="num">
                                            <div class="num-num">
                                                3
                                            </div>
                                        </div>
                                        <p class="num-text">Нажмите на вкладку “Переключиться на профессиональный аккаунт”.</p>
                                    </div>
                                    <div class="num-wrapper">
                                        <div class="num">
                                            <div class="num-num">
                                                4
                                            </div>
                                        </div>
                                        <p class="num-text">Дайте разрешение приложению Инстаграма управлять вашими Фейсбук-страницами. Не вашим личным аккаунтом, а именно страницами в Фейсбуке, к которым вы имеете доступ. Вы можете выбрать или существующую страницу в Facebook, к которой у вас есть доступ, или создать новую. </p>
                                    </div>
                                    <div class="num-wrapper">
                                        <div class="num">
                                            <div class="num-num">
                                                5
                                            </div>
                                        </div>
                                        <p class="num-text">Настройка контактной информации для профиля Инстаграма. Укажите в соответствующие поля рабочий email компании, номер телефона и физический адрес (если есть).</p>
                                    </div>
                                    <div class="num-wrapper">
                                        <div class="num">
                                            <div class="num-num">
                                                6
                                            </div>
                                        </div>
                                        <p class="num-text">Когда вы прошли все шаги и включили профессиональный аккаунт, то, пожалуйста, вернитесь на сайт client.targetted.online и нажмите кнопку “у меня есть профессиональный аккаунт” и следуйте дальнейшим инструкциям на сайте.</p>
                                    </div>
                                </div>
                            </popup>
                            <b-button
                                v-if="!(store.getters.GET_FB_PAGES.length > 0)"
                                class="main-button"
                                id="primary-under"
                                @click="showPopupInfo"
                            >
                                Нет бизнес-аккаунта
                            </b-button>
                        </div>
                        <div v-if="store.getters.GET_FB_PAGES.length > 0 && !isRequestSent && !pageSubmitted">
                            <div>
                                <p>Выберите страницу которую хотите привязать</p>
                                <!-- //TODO SELECT -->
                                <b-form-group
                                label="Выберите страницу"
                                :label-cols="label_cols"
                                :content-cols="content_cols"
                                id="input-group-main"
                                label-for="input-horizontal"
                                >
                                    <b-form-radio-group
                                        v-model="company.FbPageId"
                                        :options="store.getters.GET_FB_PAGES"
                                    ></b-form-radio-group>
                                </b-form-group>
                            </div>
                            <b-button
                                class="main-button"
                                @click="sendFbRequest()"
                            >
                                Привязать
                            </b-button>
                        </div>
                        <div v-if="isRequestSent && !pageSubmitted">
                            <p>Зайди в аккаунт на Facebook и подтверди привязку страницы в сообщениях</p>
                            <b-button
                                class="main-button"
                                target="_blank"
                                rel="noopener noreferrer"
                                :href=getFBRedirect()
                            >
                            Перейти в facebook
                            </b-button>
                            <b-button

                                class="main-button"
                                id="primary-under"
                                @click="checkPageSubmitted()"
                            >
                                Я подтвердил в сообщениях
                            </b-button>
                        </div>
                        <div v-if="pageSubmitted">
                            <div class="c-status" style="margin-top: 30px; max-width: 800px">
                                <div class="elipse" id="green"></div>
                                <p class="c-status-text" id="c-status-text-u" style="text-align: left;"
                                >Страница {{company.FbPageId}} привязана к targetted</p>
                            </div>
                        </div>
                        <b-button
                            v-if="store.getters.GET_FB_PAGES.length > 0 ||
                            (isRequestSent && pageSubmitted) ||
                            (store.getters.GET_FB_PAGES.length == 0 || isRequestSent)"

                            class="main-button-grey"
                            style="margin-top: 30px; background: #F3F3F3; color: black"
                            @click="logout"
                        >
                            Привязать другой аккаунт
                        </b-button>
                    </div>

                    <h2 id="h2">Основное</h2>

                    <b-form-group
                        label="Название кампании"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group-main"
                        label-for="input-horizontal"
                    >
                        <b-form-input
                        class="form-input"
                        id="c-name"
                        v-model="company.CompanyName"
                        :disabled="isEdit"
                        :state="validateState('CompanyName')"
                        placeholder="Введите название"
                        @click="resetNameErr()"
                        ></b-form-input>
                        <b-form-invalid-feedback
                        class="error-message"
                        id="c-name">
                            Название должно быть между 3 и 30 символами
                        </b-form-invalid-feedback>
                        <small v-if="isCompanyExist" class="error-message">
                            Кампания с таким иминем уже создана
                        </small>
                    </b-form-group>

                    <b-form-group
                        label="Цель кампании"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group-main"
                        label-for="input-horizontal"
                    >
                        <b-form-radio-group
                            v-model="company.CompnayPurpose"
                            :disabled="isEdit"
                            :options="[
                                'Сообщения в директ',
                                'Подписки в instagram',
                                'Заявки через форму обратной связи',
                                'Целевое действие на сайте'
                            ]"
                        ></b-form-radio-group>
                    </b-form-group>

                    <b-form-group
                        label="Сфера деятельности"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group-main"
                        label-for="input-horizontal"
                    >
                        <b-form-input
                        class="form-input"
                        :disabled="isEdit"
                        v-model="company.CompanyField"
                        :state="validateState('CompanyField')"
                        placeholder="Вставьте ссылку"
                        ></b-form-input>
                        <b-form-invalid-feedback
                        class="error-message">
                            Обязательное поле
                        </b-form-invalid-feedback>
                    </b-form-group>

                    <b-form-group
                        label="Адрес бизнеса"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group-main"
                        label-for="input-horizontal"
                        description="Нужен только для офлайн бизнеса"
                    >
                        <b-form-input
                        class="form-input"
                        :disabled="isEdit"
                        v-model="company.BusinessAddress"
                        placeholder="Точный адрес"
                        ></b-form-input>
                    </b-form-group>
                <h2 id="h2">Креативы</h2>
                <div v-if="isEdit" >
                    <b-form-group
                    v-if="imageNames.length > 0"
                    label="Уже имеющиеся креативы"
                    :label-cols="label_cols"
                    :content-cols="content_cols"
                    id="input-group-main"
                    label-for="input-horizontal"
                    >
                        <div id="image-block">
                            <div
                            v-for="(name) in imageNames"
                            :key="name">
                            <div>
                                <img id="preview" :src="getImageByName(name)"/>
                            </div>
                            </div>
                        </div>
                    </b-form-group>
                </div>

                <div v-if="!isEdit">
                    <b-form-group
                            label="Наличие креативов"
                            :label-cols="label_cols"
                            :content-cols="content_cols"
                            id="input-group-main"
                            label-for="input-horizontal"
                    >
                        <b-form-radio-group
                            v-model="company.CreativeStatus"
                            :options="[
                                'Есть рекламные креативы',
                                'Создать рекламные креативы'
                            ]"
                        ></b-form-radio-group>
                    </b-form-group>

                    <div class="creative-message">
                        <div
                        v-if="isCreative()">
                            Для создания рекламных креативов загрузите картинки и напишите текст, который будет на них отображаться. <a href="https://docs.google.com/document/d/1gqOkpxDJ1wNt-AYlt5Q1Et1kF8NRLRYdG-dXK7WdT1k/edit?usp=sharing" style="color: #6C1BD2" target="_blank">Посмотрите пример,</a> как это будет выглядеть и подробную инструкцию.
                        </div>
                        <div
                        v-if="!isCreative()">
                            <a href="https://docs.google.com/document/d/1DGJQw2lw_Y6KzyPeqA6To8u9UuiYi2Kokx_yjvTxgKE/edit?usp=sharing" style="color: #6C1BD2" target="_blank">Воспользуйтесь советами</a> при самостоятельном создании креативов, чтобы увеличить эффективность рекламной кампании.
                        </div>
                    </div>

                    <b-form-group
                            :label="getStoriesLabel()"
                            :label-cols="label_cols"
                            :state="validateImages()"
                            id="input-group-main"
                            label-for="input-horizontal"
                            description="До 5 слайдов в сториз"
                    >
                        <div id="image-block">
                            <div
                            v-for="(Image, key) in Images"
                            :key="key">
                                <div>
                                    <div
                                    id="icon-div-image">
                                        <b-icon
                                        @click.stop="removeImage(Image)"
                                        class="x-button"
                                        icon="x"></b-icon>
                                    </div>
                                    <img id="preview" :ref="'Image'" />
                                </div>
                            </div>

                            <div v-if="Images.length < 5">
                                <input
                                style="display: none"
                                type="file"
                                multiple
                                accept="Image/gif, Image/jpeg, Image/png, Image/jpg"
                                @change="onFileSelected"
                                ref="fileInput">
                                <div
                                @click="$refs.fileInput.click()"
                                id="load-frame">
                                <!-- Обводка у кнопок -->
                                <!-- ПОехал размер -->
                                    <p id="load-file">Загрузить<br>файл</p>
                                    <p id="file-size-big">Размер<br>1920х1080рх</p>
                                    <!-- TODO FILL -->
                                </div>
                            </div>
                        </div>
                        <b-form-invalid-feedback
                        class="error-message">
                            Необходим минимум один файл
                        </b-form-invalid-feedback>
                    </b-form-group>
                    <div v-if="isCreative()">
                        <div
                        v-for="(Image, index) in Images"
                        :key="Image.name">
                            <b-form-group
                                :label="textOnSlide(index)"
                                :label-cols="label_cols"
                                :content-cols="content_cols"
                                id="input-group-main"
                                label-for="input-horizontal"
                            >
                                <b-form-input
                                class="form-input"
                                v-model="company.ImagesDescription[index]"
                                placeholder="Введите текст"
                                ></b-form-input>
                            </b-form-group>
                        </div>
                    </div>
                    <b-form-group
                            :label="getPostLabel()"
                            :label-cols="label_cols"
                            :state="validateImagesSmall()"
                            id="input-group-main"
                            label-for="input-horizontal"
                            description="До 5 слайдов в посте"
                    >

                        <div id="image-block">
                            <div
                            v-for="(Image, key) in ImagesSmall"
                            :key="key" style="width: 160px; height: 160px;">
                                <div id="icon-div-image">
                                    <b-icon
                                    @click.stop="removeImageSmall(Image)"
                                    class="x-button"
                                    icon="x"></b-icon>
                                </div>
                                <img id="preview-small" :ref="'ImageSmall'" />
                            </div>
                            <div v-if="ImagesSmall.length < 5">
                                <input
                                style="display: none"
                                type="file"
                                multiple
                                accept="Image/gif, Image/jpeg, Image/png, Image/jpg"
                                @change="onSmallFileSelected"
                                ref="smallFileInput">
                                <div
                                @click="$refs.smallFileInput.click()"
                                id="load-frame-small">
                                    <p id="load-file">Загрузить<br>файл</p>
                                    <p id="file-size">Размер<br>1080х1080рх</p>
                                </div>
                            </div>
                        </div>
                        <b-form-invalid-feedback
                        class="error-message">
                            Необходим минимум один файл
                        </b-form-invalid-feedback>
                    </b-form-group>
                    <div v-if="isCreative()">
                        <div
                        v-for="(Image, index) in ImagesSmall"
                        :key="Image.name">
                            <b-form-group
                                :label="textOnImage(index)"
                                :label-cols="label_cols"
                                :content-cols="content_cols"
                                id="input-group-main"
                                label-for="input-horizontal"
                            >
                                <b-form-input
                                class="form-input"
                                v-model="company.ImagesSmallDescription[index]"
                                placeholder="Введите текст"
                                ></b-form-input>
                            </b-form-group>
                        </div>
                    </div>
                </div>
                <b-form-group
                v-if="ImagesSmall.length > 0"
                        label="Описание под постом в ленте"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group-main"
                        label-for="input-horizontal"
                    >
                        <b-form-textarea
                        class="form-input"
                        style="height: 100px"
                        v-model="company.PostDescription"
                        placeholder="Введите текст"
                        ></b-form-textarea>
                    </b-form-group>
                    <b-button
                        class="submit-button"
                        type="submit"
                    >
                        {{isEdit ? "Назад":"Продолжить"}}
                    </b-button>
                </b-form>
            </div>
        </div>
        </div>
    </div>

</template>
<script>
import accountService from '../../_services/account.service'
import store from '../../../store/store'
import router from '../../../router/router'
import axios from 'axios'
import popup from '../BigPopup.vue'
import { validationMixin } from 'vuelidate'
import { required, maxLength, minLength } from 'vuelidate/lib/validators'
import loading from '../Loading.vue'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL
export default {
  name: 'CreateCompany',
  mixins: [validationMixin],
  components: {
    popup,
    loading
  },
  data () {
    return {
      store,
      isLoading: false,
      isCompanyExist: false,
      isInfoPopupVisible: false,
      label_cols: this.getWidth().label,
      content_cols: this.getWidth().content,
      isRequestSent: false,
      pageSubmitted: false,
      imageNames: [],
      ImagesSmall: [],
      Images: [],
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
            if (err.response.status === 401) {
              store.dispatch('logout')
                .then(() => {
                  this.$router.push('/login')
                })
            }
          })
      } else {
        this.reset()
      }
    }
  },
  computed: {
    isEdit () {
      return this.$route.query.isEdit === 'true'
    }
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
          if (err.response.status === 401) {
            router.push({ path: '/login' })
          }
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
    reset () {
      this.isCompanyExist = false
      this.Images = []
      this.ImagesSmall = []
      this.isLoading = false
      this.company = {
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
      this.isInfoPopupVisible = false
      this.isRequestSent = false
      this.pageSubmitted = false
      this.label_cols = this.getWidth().label
      this.content_cols = this.getWidth().content
    },
    getFBRedirect () {
      return `https://facebook.com/${this.company.FbPageId}/settings/?tab=admin_roles`
    },
    resetNameErr () {
      this.isCompanyExist = false
    },
    validateState (name) {
      const { $dirty, $error } = this.$v.company[name]
      return $dirty ? !$error : null
    },
    validateImages () {
      return this.Images.length === 0
    },
    validateImagesSmall () {
      return this.ImagesSmall.length === 0
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
          if (err.response.status === 401) {
            store.dispatch('logout')
              .then(() => {
                this.$router.push('/login')
              })
          }
        })
    },
    getImageByName (name) {
      const uID = this.company.UserId
      const cID = this.company.Id
      return `https://client.targetted.online/images/${uID}/${cID}${name}`
    },
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
    closeInfoPopup () {
      this.isInfoPopupVisible = false
    },
    showPopupInfo () {
      this.isInfoPopupVisible = true
    },
    removeImageSmall (Image) {
      this.remove(this.ImagesSmall, Image)
      for (let i = 0; i < this.ImagesSmall.length; i++) {
        const reader = new FileReader()
        reader.onload = () => {
          this.$refs.ImageSmall[i].src = reader.result
        }

        reader.readAsDataURL(this.ImagesSmall[i])
      }
    },
    removeImage (Image) {
      this.remove(this.Images, Image)
      for (let i = 0; i < this.Images.length; i++) {
        const reader = new FileReader()
        reader.onload = () => {
          this.$refs.Image[i].src = reader.result
        }

        reader.readAsDataURL(this.Images[i])
      }
    },
    remove (arr, img) {
      for (let i = 0; i < arr.length; i++) {
        if (arr[i].name === img.name) {
          arr.splice(i, 1)
        }
      }
    },
    getStoriesLabel () {
      if (!this.isCreative()) {
        return 'Креативы для Сториз'
      }
      return 'Картинки для Сториз'
    },
    getPostLabel () {
      if (!this.isCreative()) {
        return 'Креативы для поста в ленте'
      }
      return 'Картинки для поста в ленте'
    },
    isCreative () {
      return this.company.CreativeStatus === 'Создать рекламные креативы'
    },
    textOnSlide (index) {
      return `Текст на слайде ${index + 1}`
    },
    textOnImage (index) {
      return `Текст на картинке ${index + 1}`
    },
    loginFB () {
      accountService.login()
    },
    createCompany () {
      if (!this.isEdit) { // TODO isEdit = isInView for some time
        this.$v.$touch()
        if (this.$v.$anyError) {
          window.scrollTo(0, 100)
          return
        }
        console.log('begin gen comp')
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
              console.log(err.response)
              console.log(err.response.data)
              console.log(err.response.data.message)
              if (err.response.data.message === 'pq: duplicate key value violates unique constraint "ad_company_name_user_id_key"') {
                this.isCompanyExist = true
              }
              if (err.response.status === 401) {
                store.dispatch('logout')
                  .then(() => {
                    this.$router.push('/login')
                  })
              }
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
              if (err.response?.status === 401) {
                store.dispatch('logout')
                  .then(() => {
                    this.$router.push('/login')
                  })
              }
            })
        }
      } else {
        router.push({ path: '/company-balance/' + this.company.Id, query: { isEdit: false } })
      }
    },
    sendFbRequest () {
      this.isLoading = true
      axios({ url: `${VUE_APP_API_URL}/api/facebook/claim/${this.company.FbPageId}`, method: 'POST' })
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
      axios({ url: `${VUE_APP_API_URL}/api/facebook/owned/${this.company.FbPageId}`, method: 'GET' })
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
    logout () {
      this.pageSubmitted = false
      this.isRequestSent = false
      accountService.logout()
    },
    onFileSelected (e) {
      if (e != null) {
        const selectedFiles = e.target.files
        const len = Math.min(selectedFiles.length, 5 - this.Images.length)
        for (let i = 0; i < len; i++) {
          this.Images.push(selectedFiles[i])
        }

        for (let i = 0; i < this.Images.length; i++) {
          const reader = new FileReader()
          reader.onload = () => {
            this.$refs.Image[i].src = reader.result
          }

          reader.readAsDataURL(this.Images[i])
        }
      }
    },
    onSmallFileSelected (e) {
      if (e != null) {
        const selectedFiles = e.target.files
        const len = Math.min(selectedFiles.length, 5 - this.ImagesSmall.length)
        for (let i = 0; i < len; i++) {
          this.ImagesSmall.push(selectedFiles[i])
        }

        for (let i = 0; i < this.ImagesSmall.length; i++) {
          const reader = new FileReader()
          reader.onload = () => {
            this.$refs.ImageSmall[i].src = reader.result
          }

          reader.readAsDataURL(this.ImagesSmall[i])
        }
      }
    }
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
