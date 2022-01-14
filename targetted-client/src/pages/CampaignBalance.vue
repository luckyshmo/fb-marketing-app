<template>
  <div id="content-wrapper">
    <loading v-if="isLoading" />
    <div id="content">
      <router-link :to="{name: 'mainPage'}">
        <p
          id="navigation-text"
          style="margin:0;"
        >
          ← Назад
        </p>
      </router-link>

    <h1>Реклама от 9 сентября</h1>

      <div
        v-if="getWidthPx() > 600"
        id="campaign-status-wrapper"
      >

        <!-- <h1
          id="h1"
          style="max-width: 700px; text-overflow: ellipsis; overflow: hidden; white-space: nowrap;"
        >
          Реклама от 9 сентября
        </h1> -->
        <div
          id="c-status"
          class="c-status"
        >
          <div
            v-if="!isFb(campaign)"
            id="white"
            class="elipse"
          />
          <div
            v-if="isFb(campaign) && !isMoney(campaign)"
            id="yellow"
            class="elipse"
          />
          <div
            v-if="isFb(campaign) && isMoney(campaign)"
            id="green"
            class="elipse"
          />
          <p class="c-status-text">
            {{ getStatus(campaign) }}
          </p>
        </div>
      </div>

      <div v-if="!(getWidthPx() > 600)">
        <h1
          id="h1"
          style="max-width: 300px; text-overflow: ellipsis; overflow: hidden; white-space: nowrap;"
        >
          {{ campaign.CampaignName }}
        </h1>
        <div
          id="c-status"
          class="c-status"
        >
          <div
            v-if="!isFb(campaign)"
            id="white"
            class="elipse"
          />
          <div
            v-if="isFb(campaign) && !isMoney(campaign)"
            id="yellow"
            class="elipse"
          />
          <div
            v-if="isFb(campaign) && isMoney(campaign)"
            id="green"
            class="elipse"
          />
          <p class="c-status-text">
            {{ getStatus(campaign) }}
          </p>
        </div>
      </div>

  <br>

   <section class="app-new-stats">
            <div class="app-new-stats-type">instagram.com/targetted</div>
            <div class="app-new-stats-details">Аккаунт в Instagram или Facebook</div>

            <div class="app-new-stats-type">Заявки в директ</div>
            <div class="app-new-stats-details">Цель рекламной кампании</div>

            <div class="app-new-stats-type">1 498 ₽ за 7 дней</div>
            <div class="app-new-stats-details">Общие затраты</div>
      </section>

      <!-- <b-form-group
        id="input-group-main"
        label="Всего на счету"
        :label-cols="label_cols"
        :content-cols="content_cols"
        label-for="input-horizontal"
      >
        <p id="balance">
          {{ campaign.CurrentAmount }} ₽
        </p>
      </b-form-group>

      <b-button
        variant="primary"
        class="main-button-grey"
        style="margint-top: 20px"
        @click="showPopupInfo"
      >
        Пополнить баланс
      </b-button>

      <popup
        v-if="isInfoPopupVisible"
        style="margin-top: -250px;"
        @closePopup="closeInfoPopup"
      >
        <h1 style="margin-top: 60px; margin-bottom: 20px">
          Пополнение
        </h1>
        <p>Минимальная сумма пополнения 500 ₽.<br>Сумма пополнения идет на рекламный<br>бюджет вашей кампании.</p>

        <div v-if="getWidthPx() > 600">
          <div
            v-if="!isFormRendered"
            id="campaign-status-wrapper"
            style="margin: 0 auto; width: 430px;"
          >
            <b-form-input
              v-model="paymentAmount"
              style="margin: 28px 12px 0px 0px"
              class="form-input"
              :state="validateState1('paymentAmount')"
              placeholder="Введите сумму"
            />
            <b-form-invalid-feedback
              class="error-message"
            >
              Минимальная сумма пополнения 500₽
            </b-form-invalid-feedback>
            <button
              style="margin-top: 28px;"
              class="main-button"
              @click="render"
            >
              Пополнить
            </button>
          </div>
        </div>
        <div v-if="!(getWidthPx() > 600)">
          <div
            v-if="!isFormRendered"
            style="margin: 0 auto;"
          >
            <b-form-input
              v-model="paymentAmount"
              class="form-input"
              style="margin: 0 auto;"
              :state="validateState1('paymentAmount')"
              placeholder="Введите сумму"
            />
            <b-form-invalid-feedback
              class="error-message"
            >
              Минимальная сумма пополнения 500₽
            </b-form-invalid-feedback>
            <button
              style="margin-top: 28px;"
              class="main-button"
              @click="render"
            >
              Пополнить
            </button>
          </div>
        </div>
        <div id="payment-form" />
      </popup>

      <b-form @submit.prevent="startCampaign()">
        <h2 id="h2">
          Рекламный бюджет
        </h2>
        <b-form-group
          id="input-group-main"
          label="Бюджет на сутки"
          :label-cols="label_cols"
          :content-cols="content_cols"
          label-for="input-horizontal"
        >
          <b-form-input
            v-model="campaign.DailyAmount"
            class="form-input"
            :state="validateState('DailyAmount')"
            placeholder="Введите сумму"
          />
          <b-form-invalid-feedback
            class="error-message"
          >
            Минимальная сумма 200₽ в сутки
          </b-form-invalid-feedback>
        </b-form-group>
        <b-form-group
          id="input-group-main"
          label="Количество дней"
          :label-cols="label_cols"
          :content-cols="content_cols"
          label-for="input-horizontal"
        >
          <b-form-input
            v-model="campaign.Days"
            class="form-input"
            :state="validateState('Days')"
            placeholder="Введите колчиество дней"
          />
          <b-form-invalid-feedback
            class="error-message"
          >
            Минимально один день
          </b-form-invalid-feedback>
        </b-form-group>
      </b-form> -->
      <!-- <div>
        <h2 id="h2">
          Настройки кампании
        </h2>
        <p>Пока вы не можете изменить настройки кампании. Если хотите внести<br>какие-то изменения, то напишите нам или создайте новую кампанию.</p>
        <router-link :to="{path: '/campaign/'+ campaign.Id, query: { isEdit: true }}">
          <b-button
            variant="primary"
            class="main-button-grey"
            style="margint-top: 20px"
          >
            К настройкам кампании
          </b-button>
        </router-link>
      </div> -->
      <!-- <b-button
        class="submit-button"
        type="submit"
        @click="updateCampaign()"
      >
        {{ isEdit ? "Обновить":"Запустить" }} кампанию
      </b-button> -->
    </div>
  </div>
</template>
<script>
import store from '@src/store/store'
import router from '@src/router/router'
import { instance as axios } from '@src/_services/index'
import { validationMixin } from 'vuelidate'
import { required, minValue } from 'vuelidate/lib/validators'
import loading from '@src/components/Loading.vue'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL
export default {
  components: {
    loading
  },
  mixins: [validationMixin],
  validations: {
    paymentAmount: {
      required,
      minValue: minValue(500)
    },
    campaign: {
      DailyAmount: {
        required,
        minValue: minValue(200)
      },
      Days: {
        required,
        maxValue: minValue(1)
      }
    }
  },
  data () {
    return {
      isLoading: false,
      confirmationToken: '',
      paymentAmount: null,
      paymentID: '',
      isFormRendered: false,
      store,
      isInfoPopupVisible: false,
      label_cols: this.getWidth().label,
      content_cols: this.getWidth().content,
      campaign: {
        FbPageId: '',
        CampaignName: '',
        CampaignObjective: '',
        CampaignField: '',
        BusinessAdress: '',
        Images: [],
        ImagesDescription: [],
        ImagesSmall: [],
        ImagesSmallDescription: [],
        CreativeStatus: '',
        PostDescription: '',
        CurrentAmount: 0,
        DailyAmount: 0,
        Days: 0
      }
    }
  },
  computed: {
    isEdit () {
      return this.$route.query.isEdit
    }
  },
  watch: {
    $route (to) {
      this.isLoading = false
      if (!(typeof to.params.id === 'undefined')) {
        axios({ url: `${VUE_APP_API_URL}/api/campaign/${to.params.id}`, method: 'GET' })
          .then(resp => {
            this.campaign = resp.data
          })
      }
    }
  },
  beforeMount () {
    const mScript = document.createElement('script')
    mScript.setAttribute('src', 'https://yookassa.ru/checkout-widget/v1/checkout-widget.js')
    document.head.appendChild(mScript)

    axios({ url: `${VUE_APP_API_URL}/api/campaign/${this.$router.history.current.params.id}`, method: 'GET' })
      .then(resp => {
        this.campaign = resp.data
      })
  },
  methods: {
    getWidthPx () {
      return Math.max(
        document.body.scrollWidth,
        document.documentElement.scrollWidth,
        document.body.offsetWidth,
        document.documentElement.offsetWidth,
        document.documentElement.clientWidth
      )
    },
    validateState (name) {
      const { $dirty, $error } = this.$v.campaign[name]
      return $dirty ? !$error : null
    },
    validateState1 (name) {
      const { $dirty, $error } = this.$v[name]
      return $dirty ? !$error : null
    },
    render () {
      this.$v.paymentAmount.$touch()
      if (this.$v.paymentAmount.$anyError) {
        return
      }
      const data = `{
                "amount": "${this.paymentAmount}.00"
            }`
      axios({
        url: `${VUE_APP_API_URL}/api/user/0/payment/token`,
        data: data,
        method: 'POST'
      })
        .then(resp => {
          this.paymentID = resp.data.id

          axios({
            url: `${VUE_APP_API_URL}/api/user/0/payment/status/${this.paymentID}`,
            data: data,
            method: 'POST'
          })
            .then(resp => {
              console.log('status resp: ', resp.data)
              if (resp.data.status === 'succeeded') {
                this.campaign.DailyAmount = this.campaign.DailyAmount + resp.data.amount.value
              }
            })
            .catch(err => {
              console.log(err)
            })
          const checkout = new window.YooMoneyCheckoutWidget({
            confirmation_token: resp.data.confirmation.confirmation_token, // Token that must be obtained from YooMoney before the payment process
            return_url: window.location.origin + router.history.current.fullPath, // URL to the payment completion page
            error_callback: function (error) {
              console.log(error)
              // Processing of initialization errors
            }
          })
          checkout.render('payment-form')
          this.isFormRendered = true
        })
        .catch(err => {
          console.log(err)
        })
    },
    closeInfoPopup () {
      this.isFormRendered = false
      this.isInfoPopupVisible = false
    },
    showPopupInfo () {
      this.isInfoPopupVisible = true
    },
    startCampaign () {
      console.log(this.balanceForm)
    },
    isMoney (campaign) {
      return campaign.CurrentAmount > 0
    },
    getStatus (campaign) {
      // if (campaign.FbPageId.length === 0) {
      //     return "facebook не подключен"
      // }
      if (campaign.CurrentAmount === 0) {
        return 'закончился рекламный бюджет'
      }
      if (!campaign.IsStarted) {
        return 'настраивается'
      }
      return 'в работе'
    },
    isFb (campaign) {
      return campaign.FbPageId.length > 0
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
    updateCampaign () {
      this.$v.campaign.$touch()
      if (this.$v.campaign.$anyError) {
        return
      }

      this.isLoading = true

      const campaignData = new FormData()
      campaignData.append('FbPageId', this.campaign.FbPageId)
      if (this.isEdit) {
        campaignData.append('Id', this.campaign.Id)
      }
      campaignData.append('CampaignName', this.campaign.CampaignName)
      campaignData.append('CampaignObjective', this.campaign.CampaignObjective)
      campaignData.append('CampaignField', this.campaign.CampaignField)
      campaignData.append('BusinessAddress', this.campaign.BusinessAddress)
      campaignData.append('ImagesDescription', this.campaign.ImagesDescription)
      campaignData.append('ImagesSmallDescription', this.campaign.ImagesSmallDescription)
      campaignData.append('CreativeStatus', this.campaign.CreativeStatus)
      campaignData.append('PostDescription', this.campaign.PostDescription)
      campaignData.append('DailyAmount', this.campaign.DailyAmount)
      campaignData.append('Days', this.campaign.Days)
      axios({ url: `${VUE_APP_API_URL}/api/campaign/${this.campaign.Id}`, data: campaignData, method: 'PUT' })
        .then(resp => {

          router.push({ path: '/main' })
        })
        .catch(err => {
          if (err.response.data.message === 'pq: duplicate key value violates unique constraint "ad_campaign_name_user_id_key"') {
            this.isCampaignExist = true
          }
          if (err.response.status === 401) {
            store.dispatch('logout')
              .then(() => {
                this.$router.push('/login')
              })
          }
        })
    }

  }
}
</script>
<style>
#campaign-status-wrapper{
    display: flex;
    align-items: stretch;
}
#balance{
    font-style: normal;
    font-weight: 600;
    font-size: 16px;
    line-height: 24px;
    padding-top: 7px;
}
.left {
    position: absolute;
    text-align: center;
}

.right {
  position: absolute;
  right: 0;
  top: 0;
  margin-top: 1%;
  margin-right: 30px;
}

@media (max-width: 600px) {
  .left{
       text-align: left;
      position:relative;
  }

  .right{
      position:relative;
  }

  .campaign-status{
      height: auto;
      position:relative;
  }
}
</style>
