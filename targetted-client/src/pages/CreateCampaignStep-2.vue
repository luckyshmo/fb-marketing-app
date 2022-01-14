<template>
  <!--  :adaptive="true"-->
  <section class="steps">
    <loading v-if="isLoading"/>
    <modal name="popup-show"
           :width="windowInnerWidth > 680 ? 676 : '100%' "
           height="auto"
           :scrollable="true"
           classes="popup-block"
           style="background-color: rgba(0, 0, 0, 0.5) !important; padding: 100px 0 !important;"
    >
      <div class="modal__button">
        <div>
          <b-icon
            class="x-button"
            icon="x"
            @click="$modal.hide('popup-show')"
          />
        </div>
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
                Сделайте потенциальным клиентам выгодное предложение с призывом к конкретному действию. Чтобы
                подтолкнуть к совершению покупки прямо сейчас, можно использовать эффект дефицита или предложить
                приятный бонус. </p>
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
    <modal name="stories-modal"
           :width="windowInnerWidth > 680 ? 676 : '100%'"
           height="auto"
           :max-width="676"
           :min-width="300"
           :scrollable="true"
           style="background-color: rgba(0, 0, 0, 0.5) !important;"
           classes="stories-modal"
    >
      <div class="modal__button">
        <div>
          <b-icon
            class="x-button"
            icon="x"
            @click="$modal.hide('stories-modal')"
          />
        </div>
      </div>
      <h2 class="stories-modal__title">Реклама в сториз</h2>
      <carousel :items="1"
                :dots="false"
                :autoWidth="windowInnerWidth > 680 ? true : false"
                class="stories-modal-carousel"
                :navText="['', `<svg aria-hidden='true' focusable='false' data-prefix='fas' data-icon='arrow-right' class='svg-inline--fa fa-arrow-right fa-w-14'
      role='img' xmlns='http://www.w3.org/2000/svg' viewBox='0 0 448 512'>
      <path fill='currentColor'
            d='M190.5 66.9l22.2-22.2c9.4-9.4 24.6-9.4 33.9 0L441 239c9.4 9.4 9.4 24.6 0 33.9L246.6 467.3c-9.4 9.4-24.6 9.4-33.9 0l-22.2-22.2c-9.5-9.5-9.3-25 .4-34.3L311.4 296H24c-13.3 0-24-10.7-24-24v-32c0-13.3 10.7-24 24-24h287.4L190.9 101.2c-9.8-9.3-10-24.8-.4-34.3z'>
</path>
      </svg>`]"
      >
        <div class="stories-modal__owl" v-for="(ImageBig, index) in campaign.Images">
          <h3 class="stories-modal__owl_title">{{ advertisingOption[index] }} вариант рекламы</h3>
          <div class="stories-modal__owl_block">
            <div class="stories-modal__owl_block__img">
              <div class="stories-modal__owl_block__img_absolut">
                <div>
                  <img :src="campaign.ImagesSrc[index]" alt="aaa"/>
                </div>
                <div>
                  <p>Название компании</p>
                  <p>Реклама</p>
                </div>
              </div>
              <img id="preview" :src="campaign.ImagesSrc[index]" alt="aaa"/>
            </div>
          </div>
        </div>
      </carousel>
      <h2 class="stories-modal__title_two">Реклама в ленте</h2>
      <carousel :items="1"
                :dots="false"
                :autoWidth="windowInnerWidth > 680 ? true : false"
                class="stories-modal-carousel"
                :navText="['', `<svg aria-hidden='true' focusable='false' data-prefix='fas' data-icon='arrow-right' class='svg-inline--fa fa-arrow-right fa-w-14'
      role='img' xmlns='http://www.w3.org/2000/svg' viewBox='0 0 448 512'>
      <path fill='currentColor'
            d='M190.5 66.9l22.2-22.2c9.4-9.4 24.6-9.4 33.9 0L441 239c9.4 9.4 9.4 24.6 0 33.9L246.6 467.3c-9.4 9.4-24.6 9.4-33.9 0l-22.2-22.2c-9.5-9.5-9.3-25 .4-34.3L311.4 296H24c-13.3 0-24-10.7-24-24v-32c0-13.3 10.7-24 24-24h287.4L190.9 101.2c-9.8-9.3-10-24.8-.4-34.3z'>
</path>
      </svg>`]"
      >
        <div class="stories-modal__owl_small" v-for="(ImageSmall, index) in campaign.ImagesSmall">
          <h3 class="stories-modal__owl_title_small">{{ advertisingOption[index] }} вариант рекламы</h3>
          <div class="stories-modal__owl_block_small">
            <div class="stories-modal__owl_block__img_absolut_small">
              <div>
                <img :src="campaign.ImagesSmallSrc[index]" alt="aaa"/>
              </div>
              <div>
                <p>Название компании</p>
                <p>Реклама</p>
              </div>
            </div>
            <div class="stories-modal__owl_block__img_small">
              <img id="preview" :src="campaign.ImagesSmallSrc[index]" alt="aaa"/>
            </div>
           <div class="stories-modal__owl_block__img_index">
             <p>Подробнее</p>
             <button type="button">
               <svg viewBox="0 0 7 12" fill="none" xmlns="http://www.w3.org/2000/svg">
                 <path d="M1 1L5.5 6L1 11" stroke="black" stroke-width="2" stroke-linecap="round"/>
               </svg>
             </button>
           </div>
            <div class="stories-modal__owl_block__icon">
              <div>
                <button type="button">
                  <svg id="like" viewBox="0 0 48 48">
                    <path d="M34.6 6.1c5.7 0 10.4 5.2 10.4 11.5 0 6.8-5.9 11-11.5 16S25 41.3 24 41.9c-1.1-.7-4.7-4-9.5-8.3-5.7-5-11.5-9.2-11.5-16C3 11.3 7.7 6.1 13.4 6.1c4.2 0 6.5 2 8.1 4.3 1.9 2.6 2.2 3.9 2.5 3.9.3 0 .6-1.3 2.5-3.9 1.6-2.3 3.9-4.3 8.1-4.3m0-3c-4.5 0-7.9 1.8-10.6 5.6-2.7-3.7-6.1-5.5-10.6-5.5C6 3.1 0 9.6 0 17.6c0 7.3 5.4 12 10.6 16.5.6.5 1.3 1.1 1.9 1.7l2.3 2c4.4 3.9 6.6 5.9 7.6 6.5.5.3 1.1.5 1.6.5.6 0 1.1-.2 1.6-.5 1-.6 2.8-2.2 7.8-6.8l2-1.8c.7-.6 1.3-1.2 2-1.7C42.7 29.6 48 25 48 17.6c0-8-6-14.5-13.4-14.5z"></path>
                  </svg>
                </button>
                <button type="button">
                  <svg id="comment" viewBox="0 0 48 48">
                    <path clip-rule="evenodd" d="M47.5 46.1l-2.8-11c1.8-3.3 2.8-7.1 2.8-11.1C47.5 11 37 .5 24 .5S.5 11 .5 24 11 47.5 24 47.5c4 0 7.8-1 11.1-2.8l11 2.8c.8.2 1.6-.6 1.4-1.4zm-3-22.1c0 4-1 7-2.6 10-.2.4-.3.9-.2 1.4l2.1 8.4-8.3-2.1c-.5-.1-1-.1-1.4.2-1.8 1-5.2 2.6-10 2.6-11.4 0-20.6-9.2-20.6-20.5S12.7 3.5 24 3.5 44.5 12.7 44.5 24z" fill-rule="evenodd"></path>
                  </svg>
                </button>
                <button type="button">
                  <svg id="send" viewBox="0 0 48 48">
                    <path d="M47.8 3.8c-.3-.5-.8-.8-1.3-.8h-45C.9 3.1.3 3.5.1 4S0 5.2.4 5.7l15.9 15.6 5.5 22.6c.1.6.6 1 1.2 1.1h.2c.5 0 1-.3 1.3-.7l23.2-39c.4-.4.4-1 .1-1.5zM5.2 6.1h35.5L18 18.7 5.2 6.1zm18.7 33.6l-4.4-18.4L42.4 8.6 23.9 39.7z"></path>
                  </svg>
                </button>
              </div>
              <div>
                <button type="button">
                  <svg id="save" viewBox="0 0 48 48">
                    <path d="M43.5 48c-.4 0-.8-.2-1.1-.4L24 29 5.6 47.6c-.4.4-1.1.6-1.6.3-.6-.2-1-.8-1-1.4v-45C3 .7 3.7 0 4.5 0h39c.8 0 1.5.7 1.5 1.5v45c0 .6-.4 1.2-.9 1.4-.2.1-.4.1-.6.1zM24 26c.8 0 1.6.3 2.2.9l15.8 16V3H6v39.9l15.8-16c.6-.6 1.4-.9 2.2-.9z"></path>
                  </svg>
                </button>
              </div>
            </div>
            <div class="stories-modal__owl_block__like">
              <p>Нравится: 100</p>
            </div>
            <div class="stories-modal__owl_block__footer">
              <p> <span>Запускайте рекламу</span> за пару кликов с помощью сервиса targetted.ru</p>
            </div>

          </div>
        </div>
      </carousel>
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
    <div>
      <b-col class="step__col">

        <Step2
          :label_cols="12"
          :content_cols="12"
          :isEdit="false"
          :campaign="campaign"
          :windowInnerWidth="windowInnerWidth"
          @next="saveAndNext">
          <template v-slot:header>
            <h1>Изображения </h1>
          </template>
        </Step2>
      </b-col>
      <b-col class="steps-process">
        <aside>
          <h2 @click="aa()">
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
    </div>
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
import carousel from 'vue-owl-carousel'

const campaignDefault = {
  FbPageId: '',
  Id: '',
  CampaignName: '',
  CampaignObjective: 'Сообщения в директ',
  CampaignField: '',
  BusinessAddress: '',
  Images: [],
  ImagesDescription: [],
  ImagesSrc: [],
  ImagesSmall: [],
  ImagesSmallDescription: [],
  ImagesSmallSrc: [],
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
    carousel
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
      windowInnerWidth: window.innerWidth,
      advertisingOption: ['Первый', 'Второй', 'Третий', 'Четвертый', 'Пятый']
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
    aa() {
      console.log(this.campaign)
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
  color: #767676;

  &:hover {
    color: $black;
  }

  & > svg {
    width: 10px;
    fill: #767676;
    margin-right: 5px;
  }

  &:hover > svg {
    fill: $black;
  }
}

.popup-block {
  background: #FFFFFF;
  box-shadow: 0px 4px 24px rgba(0, 0, 0, 0.12), 0px 1px 3px rgba(0, 0, 0, 0.05) !important;
  border-radius: 20px !important;
}

.popup__block {
  padding: 40px 48px 56px;
}

.modal__button {
  position: relative;
  width: 100%;

  & > div {
    width: 40px;
    height: 40px;
    background: #F3F3F3;
    position: absolute;
    top: 16px;
    right: 16px;
    border-radius: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }
}

.popup__top_block > h2 {
  font-size: 32px !important;
  font-style: normal;
  line-height: 38px !important;
}

.popup__top_block > p:nth-child(2) {
  margin-top: 12px;
}

.popup__top_block > p:nth-child(3) {
  margin-top: -3px;
}

.popup__advice {
  padding: 22px 0;
}

.popup__advice > div {
  width: 100%;
  display: flex;
}

.popup__advice > div > p:nth-child(1) {
  max-width: 14px;
  width: 100%;
  margin-right: 8px;
}

.popup__advice > div > p {
  margin-bottom: 12px;
}

.popup__descriptions > h2 {
  margin-top: 7px;
}

.popup__descriptions > div > p:nth-child(1) {
  font-weight: bold;
  margin-bottom: 0;
}

.popup__descriptions > div > p:nth-child(2) {
  margin-bottom: 12px;
}

.popup__descriptions > div:nth-child(6) > p:nth-child(2) {
  margin-bottom: 0;
}

.popup__descriptions > button {
  border: none;
  background-color: transparent;
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  line-height: 24px;
  color: $violet;
  padding: 0;
}

.popup__service {
  padding-top: 40px;
}

.popup__service > div:nth-child(1) > p > a {
  color: $violet;
}

.popup__service > div:nth-child(2) > h2 {
  margin-top: 40px;
}

.popup__service > div:nth-child(2) > p:nth-child(2) {
  font-weight: bold;
  margin-bottom: 0;
}

.popup__service > div:nth-child(2) > p:nth-child(3) {
  margin-bottom: 0;
}

.popup__service > div:nth-child(2) > div {
  max-width: 283px;
  width: 100%;
  height: 279px;
  margin-top: 12px;

  & > img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}

.popup__service > div:nth-child(3) > p:nth-child(1) {
  font-weight: bold;
  margin-top: 28px;
  margin-bottom: 0;
}

.popup__service > div:nth-child(3) > p:nth-child(2) {
  margin-bottom: 0;
}

.popup__service > div:nth-child(3) > div {
  max-width: 283px;
  width: 100%;
  height: 283px;
  margin-top: 12px;

  & > img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}

.popup__bottom {
  padding-top: 40px;
}

.popup__bottom > div {
  width: 100%;
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}

.popup__bottom > div > div {
  max-width: 185px;
  width: 100%;

  p {
    margin-top: 12px;
    font-style: normal;
    font-weight: normal;
    font-size: 14px !important;
    line-height: 18px !important;
    margin-bottom: 0;
  }
}

.popup__bottom > div > div > div {
  width: 100%;
  height: 328px;

  & > img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}

.stories-modal-carousel {
  padding: 0 48px;

  .owl-prev {
    display: none !important;
  }

  .owl-next {
    position: absolute;
    top: calc(50% - 40px);
    right: -25px;
    width: 40px;
    height: 40px;
    display: flex !important;
    justify-content: center;
    align-items: center;
    font-size: 14px !important;
    background: $white !important;
    color: $black !important;
    border-radius: 20px !important;

    & > svg {
      width: 14px;
    }
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

.stories-modal {
  border-radius: 20px !important;
}


.stories-modal__title {
  margin: 40px 0 0 48px;
  font-style: normal;
  font-weight: bold;
  font-size: 32px !important;
  line-height: 38px !important;

}

.stories-modal__owl {
  width: 345px;
  margin: 20px 20px 0 0;
}

.stories-modal__owl_title {
  font-style: normal;
  font-weight: bold;
  font-size: 24px;
  line-height: 30px;
  margin: 0;
}

.stories-modal__owl_block {
  width: 327px;
}

.stories-modal__owl_block__img {
  width: 100%;
  height: 572px;
  position: relative;
  margin-top: 16px;

  & > img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    border-radius: 12px;
  }

  & > .stories-modal__owl_block__img_absolut {
    position: absolute;
    left: 18px;
    top: 18px;
    display: flex;
    align-items: center;

    & > div:nth-child(1) {
      width: 40px;
      height: 40px;
      border-radius: 100%;

      & > img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        border-radius: 100%;
      }
    }

    & > div:nth-child(2) {
      margin-left: 6px;

      & > p:nth-child(1) {
        font-style: normal;
        font-weight: bold;
        font-size: 12px;
        line-height: 16px;
        color: $white;
        margin-bottom: 0;
      }

      & > p:nth-child(2) {
        font-style: normal;
        font-weight: normal;
        font-size: 11px;
        line-height: 14px;
        color: $white;
        margin-bottom: 0;
      }
    }
  }
}

.stories-modal__title_two {
  margin: 80px 0 0 48px;
  font-style: normal;
  font-weight: bold;
  font-size: 32px !important;
  line-height: 38px !important;
  color: $black;
}

.stories-modal__owl_small {
  margin-top: 20px;
  width: 345px;
}

.stories-modal__owl_title_small {
  font-style: normal;
  font-weight: bold;
  font-size: 24px;
  line-height: 30px;
  color: $black;
  margin-bottom: 0;
}

.stories-modal__owl_block_small {
  width: 327px;
  margin-top: 16px;
}

.stories-modal__owl_block__img_absolut_small {
  display: flex;

  & > div:nth-child(1) {
    width: 30px;
    height: 30px;
    border-radius: 100%;

    & > img {
      width: 100%;
      height: 100%;
      border-radius: 100%;
      object-fit: cover;
    }
  }

  & > div:nth-child(2) {
    margin-left: 8px;

    & > p:nth-child(1) {
      margin-bottom: 0;
      font-style: normal;
      font-weight: bold;
      font-size: 12px;
      line-height: 16px;
      color: $black;
    }

    & > p:nth-child(2) {
      font-style: normal;
      font-weight: normal;
      font-size: 11px;
      line-height: 14px;
      color: $black;
    }
  }
}

.stories-modal__owl_block__img_small {
  width: 100%;
  height: 300px;

  & > img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}
.stories-modal__owl_block__img_index{
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 0.5px solid rgba(0, 0, 0, 0.1);
  padding: 8px 5px 11px;
  & > p{
    font-style: normal;
    font-weight: normal;
    font-size: 13px;
    line-height: 17px;
    margin: 0;
  }
  & > button{
    width: 7px;
    cursor:pointer;
    border: 0;
    background-color: transparent;
    & > svg{
      width: 7px;
    }
  }
}
.stories-modal__owl_block__icon{
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 0 10px;
  & > div{
    & > button{
      width: auto;
      border: 0;
      background-color: transparent;
      cursor: pointer;
      padding: 0 18px 0 0;
      & > svg{
        width: 21px;
      }
    }
    & > button:nth-child(3){
      padding: 0;
    }
  }
  & > div:nth-child(2){
    & > button{
      padding: 0;
    }
  }
}
.stories-modal__owl_block__like{
  width: 100%;
  & > p{
    font-style: normal;
    font-weight: bold;
    font-size: 13px;
    line-height: 22px;
    margin: 0;
  }
}
.stories-modal__owl_block__footer{
  width: 100%;
  margin: 4px 0 56px;
  & > p{
    font-style: normal;
    font-weight: normal;
    font-size: 11.5px;
    line-height: 15px;
    margin: 0;
    & > span{
      font-weight: bold;
    }
  }
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
  .popup__block {
    padding: 20px 20px 30px;

    h2 {
      font-size: 24px !important;
      line-height: 30px !important;
    }
  }
  .popup__bottom > h2 {
    text-align: center;
  }
  .popup__bottom > div {
    flex-wrap: wrap;
    justify-content: center;
    margin-top: 10px;
  }
  .popup__bottom > div > div {
    margin-top: 15px;
  }
}
@media (max-width: 680px) {
  .stories-modal__owl {
    width: 100%;
    margin-top: 20px;
  }
  .stories-modal__owl_title{
    width: 100%;
    text-align: center;
  }
  .stories-modal__owl_block {
    max-width: 320px;
    margin: 0 auto;
  }
  .stories-modal-carousel .owl-next{
    right:5px
  }
  .stories-modal__owl_small {
    margin-top: 20px;
   width: 100%;
  }
  .stories-modal__owl_block_small {
    max-width: 320px;
    margin: 16px auto 0;
  }
  .stories-modal__owl_title_small{
    width: 100%;
    text-align: center;
  }
  .stories-modal-carousel{
    padding: 0;
  }
  .stories-modal__owl_block__img{
    height: 450px;
  }
  .stories-modal {
    border-radius: 0 !important;
  }
}
.custom-control-input {
  margin-right: 3px;
}
</style>
