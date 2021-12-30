<template>
<!--  :adaptive="true"-->
  <section class="steps">
    <loading v-if="isLoading"/>
    <modal name="popup-show"
           :width="windowInnerWidth > 680 ? 676 : '95%' "
           height="auto"
           :scrollable="true"
           classes="popup-block"
           style="background-color: rgba(0, 0, 0, 0.5) !important; padding: 100px 0 !important;"
    >
      <div class="modal__button">
        <b-icon
          class="x-button"
          icon="x"
          @click="$modal.hide('popup-show')"
        />
      </div>

      <div class="popup__block">
        <div class="popup__modal">
          <div class="popup__top_block">
            <h2>Советы по созданию креативов</h2>
            <p>Креативы должны быть в 2 размерах: 1920×1080px и 1080×1080px.</p>
            <p>Лучше загрузить несколько вариантов креативов. Например, разные фотографии или видео, разные акции и
              специальные предложения. Это увеличит эффективность рекламной кампании</p>
          </div>
          <div class="popup__advice">
            <h2>Советы</h2>
            <div>
              <p>1.</p>
              <p>Заинтересуйте заманчивым предложением на креативе. Например, дайте скидку на первое посещение;
                предложите бесплатную консультацию; расскажите клиентам, что вы давно работаете на рынке и у вас много
                отличных отзывов.</p>
            </div>
            <div>
              <p>2.</p>
              <p>Сделайте описание, которое расскажет пользователю подробнее о вашем бизнесе и предложении. Например, вы
                можете перечислить ваши преимущества для клиента, рассказать подробнее об акции, указать как можно
                купить ваш товар или услугу. </p>
            </div>
            <div>
              <p>3.</p>
              <p>Используйте изображения из наиболее популярных постов и сторизов, если они подходят для рекламируемой
                услуги, товара или профиля.</p>
            </div>
          </div>
          <div class="popup__descriptions">
            <h2>Шаблон описания</h2>
            <div>
              <p>Заголовок</p>
              <p>Сделайте его цепляющим, интересным. Кпримеру, укажите главное преимущество товара/услуги, задайте
                вопрос или укажите цифры.</p>
            </div>
            <div>
              <p>Особенности продукта, услуги</p>
              <p>
                В этом блоке сделайте описание свойств товара/услуги.
              </p>
            </div>
            <div>
              <p>Список с преимуществами</p>
              <p>
                Озвучьте достоинства продукта/услуги в виде списка. </p>
            </div>
            <div>
              <p> Призыв к действию</p>
              <p>
                Сделайте потенциальным клиентам выгодное предложение с призывом к конкретному действию. Чтобы подтолкнуть к совершению покупки прямо сейчас, можно использовать эффект дефицита или предложить приятный бонус.  </p>
            </div>
            <div>
              <p>Пример описания</p>
              <p>
                НАЧНЕМ ГОТОВИТЬСЯ К ЛЕТУ!<br/>🌹Весна уже пришла это значит, что скоро мы снимем куртки!</p>
            </div>
            <button>Показать всё</button>
          </div>
          <div class="popup__service">
            <div>
              <h2>Сервис для креативов</h2>
              <p>
                <a href="https://www.google.com">Canva.com</a>
                — есть готовые шаблоны для создания рекламных постов и сторизов, бесплатные стоковые фото,
                30-дневный бесплатный период.
                <router-link to="/">Инструкция</router-link>
                по созданию креативов для рекламы в Facebook и Instagram.
              </p>
            </div>
            <div>
              <h2>Примеры</h2>
              <p>Бесплатное первое посещение.</p>
              <p>Цель — привлечь клиента на запись, а потом сделать из него постоянного клиента. </p>
              <div>
                <img src="../assets/img/image 5.png" alt="img">
              </div>
            </div>
            <div>
              <p>Скидка на услугу.</p>
              <p>
                Цель — увеличить заинтересованность клиента в вашем предложении. </p>
              <div>
                <img src="../assets/img/image 7.png" alt="img">
              </div>
            </div>
          </div>
          <div class="popup__bottom">
            <h2>Пример для карусели</h2>
            <div>
              <div>
                <div>
                  <img src="../assets/img/image 8.png" alt="img">
                </div>
                <p>Стоимость услуги<br/> со скидкой и описанием. </p>
              </div>
              <div>
                <div>
                  <img src="../assets/img/image 9.png" alt="img">
                </div>
                <p>Подробное описание<br/> услуг + подарок. </p>
              </div>
              <div>
                <div>
                  <img src="../assets/img/image 10.png" alt="img">
                </div>
                <p>Отзывы помогают<br/> вызвать больше доверия. </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </modal>
    <b-row class="app-new-steps-header">
      <b-col style="padding: 0; display: flex; align-items: center">
        <router-link id="app-link-back" :to="{path: '/campaign-step-1'}" style="display: flex; align-items: center">

          <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
               x="0px" y="0px"
               viewBox="0 0 512 512" style="enable-background:new 0 0 512 512;" xml:space="preserve">
<path d="M3.919,243.077c-0.223,0.33-0.414,0.675-0.618,1.015c-0.186,0.31-0.382,0.614-0.552,0.936
	c-0.186,0.349-0.346,0.709-0.514,1.066c-0.157,0.332-0.321,0.658-0.464,0.998c-0.144,0.349-0.261,0.706-0.388,1.06
	c-0.129,0.362-0.268,0.718-0.38,1.089c-0.107,0.358-0.188,0.721-0.279,1.085c-0.093,0.374-0.199,0.743-0.275,1.125
	c-0.084,0.422-0.133,0.849-0.194,1.275c-0.047,0.326-0.109,0.645-0.143,0.976c-0.15,1.53-0.15,3.07,0,4.6
	c0.034,0.33,0.096,0.65,0.143,0.976c0.061,0.425,0.11,0.853,0.194,1.275c0.076,0.382,0.18,0.749,0.275,1.125
	c0.092,0.362,0.171,0.726,0.279,1.085c0.112,0.369,0.251,0.726,0.38,1.089c0.127,0.354,0.244,0.711,0.388,1.06
	c0.143,0.34,0.307,0.666,0.464,0.998c0.168,0.355,0.327,0.715,0.514,1.064c0.171,0.321,0.366,0.625,0.552,0.936
	c0.203,0.34,0.394,0.684,0.618,1.015c0.234,0.351,0.493,0.68,0.745,1.015c0.205,0.272,0.393,0.549,0.608,0.813
	c0.489,0.596,1.002,1.168,1.548,1.711l116.36,116.36c4.544,4.544,10.501,6.817,16.455,6.817c5.956,0,11.913-2.271,16.455-6.817
	c9.089-9.089,9.089-23.824,0-32.912l-76.636-76.636h409.272c12.853,0,23.273-10.42,23.273-23.273
	c0-12.853-10.42-23.273-23.273-23.273H79.456l76.636-76.636c9.089-9.089,9.089-23.824,0-32.912c-9.087-9.089-23.824-9.089-32.912,0
	L6.82,239.538c-0.546,0.543-1.06,1.116-1.548,1.711c-0.216,0.264-0.403,0.541-0.608,0.813
	C4.412,242.397,4.153,242.726,3.919,243.077z"/>
            <g>
</g>
            <g>
</g>
            <g>
</g>
            <g>
</g>
            <g>
</g>
            <g>
</g>
            <g>
</g>
            <g>
</g>
            <g>
</g>
            <g>
</g>
            <g>
</g>
            <g>
</g>
            <g>
</g>
            <g>
</g>
            <g>
</g>
</svg>

          Назад
        </router-link>
        <b-col>
          <div class="text-muted d-block">&nbsp;∙&nbsp;Шаг 2 из 4</div>
        </b-col>
      </b-col>
    </b-row>
    <b-row >
      <b-col class="step__col">

        <Step2
          :label_cols="12"
          :content_cols="12"
          :isEdit="false"
          :campaign="campaign"
          @next="saveAndNext">
          <template v-slot:header>
            <h1>Изображения </h1>
          </template>
        </Step2>
      </b-col>
      <b-col class="steps-process">
        <aside>
          <h2>
            Заполнено 50%
          </h2>
          <ul class="app-new-progress-text">
            <li>
              О бизнесе
            </li>
            <li class="active">
              Изображения
            </li>
            <li>
              Привязка аккаунта
            </li>
            <li>
              Охват рекламы
            </li>
            <li>
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
import {instance as axios} from '@src/_services/index'
import {validationMixin} from 'vuelidate'
import {required, maxLength, minLength} from 'vuelidate/lib/validators'
import loading from '@src/components/Loading.vue'
import Step2 from './campagin-creation/Step2.vue'

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
    loading,
    Step2,
  },
  data() {
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
      displayNumber: 0,
      windowInnerWidth : window.innerWidth,
    }
  },
  computed: {
    number() {
      return (this.currentStep / this.totalSteps) * 100
    },
    isEdit() {
      return this.$route.query.isEdit === 'true'
    },
    label_cols() {
      return this.getWidth().label
    },
    content_cols() {
      return this.getWidth().content
    }
  },
  watch: {
    windowWidth() {
      this.onResize()
    },
    $route(to) {
      console.log('route ', store.getters.GET_EMAIL)
      window.scrollTo(0, 100)
      this.isLoading = false
      if (!(typeof to.params.id === 'undefined')) {
        axios.get({url: `${VUE_APP_API_URL}/api/campaign/${to.params.id}`})
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
  mounted() {
    this.$nextTick(() => {
      window.addEventListener('resize', this.onResize);
    })
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
  beforeMount() {
    console.log('BM ', store.getters.GET_EMAIL, 'campId:', this.$router.history.current?.params?.id)
    if (!this.$router.history.current?.params?.id === 'undefined') {
      axios.get({url: `${VUE_APP_API_URL}/api/campaign/${this.$router.history.current.params.id}`})
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
    goStepBack() {
      this.currentStep = this.currentStep - 1
      localStorage.setItem('campagin_step', this.currentStep)
    },
    onResize() {
      this.windowInnerWidth = window.innerWidth
    },
    saveAndNext(data) {
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
    reset() {
      this.isCampaignExist = false
      this.Images = []
      this.ImagesSmall = []
      this.isLoading = false
      this.campaign = campaignDefault
      this.isInfoPopupVisible = false
      this.isRequestSent = false
      this.pageSubmitted = false
    },
    getImages() {
      axios.get({url: `${VUE_APP_API_URL}/api/campaign/${this.campaign.Id}/images/`})
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
    getWidth() {
      return {
        label: 12,
        content: 12
      }
    },

    updateCampaign() {
      store.dispatch('saveCampaign', this.campaign)
    },

    createCampaign() {
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
              router.push({path: `/campaign-balance/${resp.data}`, query: {}}) // TODO QUERY
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
          axios.get({url: `${VUE_APP_API_URL}/api/campaign/${this.campaign.Id}`, data: campaignData})
            .then(() => {
              router.push({path: `/campaign-balance/${this.campaign.Id}`, query: {isEdit: false}})
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
        router.push({path: `/campaign-balance/${this.campaign.Id}`, query: {isEdit: false}})
      }
    },

    logout() {
      this.pageSubmitted = false
      this.isRequestSent = false
      accountService.logout()
    }

  }
}
</script>
<style lang="scss">
@import '@src/assets/styles/vars.scss';
#app-link-back {
  // float: left;
  margin-right: 5px;
  float: left;
  line-height: 20px;
  color: #6c757d;

  &:hover {
    color: $black;
  }
  & > svg{
    width: 10px;
    fill: #6c757d;
    margin-right: 5px;
  }
  &:hover > svg{
    fill: $black;
  }
}
.popup-block{
  background: #FFFFFF;
  box-shadow: 0px 4px 24px rgba(0, 0, 0, 0.12), 0px 1px 3px rgba(0, 0, 0, 0.05) !important;
  border-radius: 20px !important;
}
.popup__block{
  padding: 40px 48px 56px;
}
.modal__button{
  position: relative;
  width: 100%;
  & > svg{
    position: absolute;
    top: 16px;
    right: 16px;
    padding: 4px;
  }
}
.popup__top_block > h2{
  font-size: 32px !important;
  font-style: normal;
  line-height: 38px !important;
}
.popup__top_block > p:nth-child(2){
  margin-top: 12px;
}
.popup__top_block > p:nth-child(3){
  margin-top: -3px;
}
.popup__advice{
  padding: 22px 0;
}
.popup__advice > div{
  width: 100%;
  display: flex;
}
.popup__advice > div > p:nth-child(1){
  max-width: 14px;
  width: 100%;
  margin-right: 8px;
}
.popup__advice > div > p{
  margin-bottom: 12px;
}
.popup__descriptions > h2{
  margin-top: 7px;
}
.popup__descriptions > div > p:nth-child(1){
  font-weight: bold;
  margin-bottom: 0;
}
.popup__descriptions > div > p:nth-child(2){
  margin-bottom: 12px;
}
.popup__descriptions > div:nth-child(6) > p:nth-child(2){
  margin-bottom: 0;
}
.popup__descriptions > button{
  border: none;
  background-color: transparent;
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  line-height: 24px;
  color: $violet;
  padding: 0;
}
.popup__service{
  padding-top: 40px;
}
.popup__service > div:nth-child(1) > p > a{
  color: $violet;
}
.popup__service > div:nth-child(2) > h2{
  margin-top: 40px;
}
.popup__service > div:nth-child(2) > p:nth-child(2){
  font-weight: bold;
  margin-bottom: 0;
}
.popup__service > div:nth-child(2) > p:nth-child(3){
  margin-bottom: 0;
}
.popup__service > div:nth-child(2) > div{
  max-width: 283px;
  width: 100%;
  height: 279px;
  margin-top: 12px;
  & > img{
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}
.popup__service > div:nth-child(3) > p:nth-child(1){
  font-weight: bold;
  margin-top: 28px;
  margin-bottom: 0;
}
.popup__service > div:nth-child(3) > p:nth-child(2){
  margin-bottom: 0;
}
.popup__service > div:nth-child(3) > div{
  max-width: 283px;
  width: 100%;
  height: 283px;
  margin-top: 12px;
  & > img{
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}
.popup__bottom{
  padding-top: 40px;
}
.popup__bottom > div{
  width: 100%;
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}
.popup__bottom > div > div{
  max-width: 185px;
  width: 100%;
  p{
    margin-top: 12px;
    font-style: normal;
    font-weight: normal;
    font-size: 14px !important;
    line-height: 18px !important;
    margin-bottom: 0;
  }
}
.popup__bottom > div > div > div{
  width: 100%;
  height: 328px;
  & > img{
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}
#icon-div-image {
  position: absolute;
  margin-left: 140px;
  margin-top: -15px;
}

#primary-under {
  margin-left: 20px;
}

@media (max-width: 722px) {
  #c-status-text-u {
    margin-left: 30px;
  }
  #primary-under {
    margin-left: 0;
    margin-top: 20px;
  }
}

@media (max-width: 600px) {
  .popup__block{
    padding: 20px 20px 30px;
    h2{
      font-size: 24px !important;
      line-height: 30px !important;
    }
  }
  .popup__bottom > h2{
    text-align: center;
  }
  .popup__bottom > div{
    flex-wrap: wrap;
    justify-content: center;
    margin-top: 10px;
  }
  .popup__bottom > div > div{
    margin-top: 15px;
  }
}


.custom-control-input {
  margin-right: 3px;
}
</style>
