<template>
    <div id="content-wrapper">
        <div id="content">
            <router-link :to="{name: 'mainPage'}">
                <p id="navigation-text" style="margin:0;">← К списку кампаний</p>
            </router-link>
            <div id="company-status-wrapper">
                <h1 id="h1" style="max-width: 700px; text-overflow: ellipsis; overflow: hidden; white-space: nowrap;">
                    {{company.CompanyName}}
                </h1>
                <div class="c-status" id="c-status">
                    <div class="elipse" id="white" v-if="!isFb(company)"></div>
                    <div class="elipse" id="yellow" v-if="isFb(company) && !isMoney(company)"></div>
                    <div class="elipse" id="green" v-if="isFb(company) && isMoney(company)"></div>
                    <p class="c-status-text">{{getStatus(company)}}</p>
                </div>
            </div>
    
            <h2 id="h2">Общий баланс</h2>

            <b-form-group
                label="Всего на счету"
                :label-cols="label_cols"
                :content-cols="content_cols"
                id="input-group-main"
                label-for="input-horizontal"
            >
                <p id="balance">{{company.CurrentAmount}}000₽</p>
            </b-form-group>

            <b-button 
                variant="primary"
                class="main-button-grey"
                @click="showPopupInfo"
            >
                Пополнить баланс
            </b-button>

            <popup
              v-if="isInfoPopupVisible"
              style="margin-top: -250px;"
              popupTitle="Пополнение счета"
              @closePopup="closeInfoPopup"
              @renderAction="render"
            >
                 <div id="payment-form"></div>
            </popup>

            <b-form @submit.prevent="startCompany()">
                <h2 id="h2">Рекламный бюджет</h2>
                <!-- //TODO descrip отступ сверху, цвет, размер -->
                <b-form-group
                        label="Бюджет на сутки"
                        description="Минимальная сумма 200₽ в сутки"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group-main"
                        label-for="input-horizontal"
                >
                    <b-form-input
                    class="form-input"
                    v-model="company.dailyAmount"
                    placeholder="Введите сумму"
                    ></b-form-input>
                </b-form-group>
                <b-form-group
                        label="Количество дней"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group-main"
                        label-for="input-horizontal"
                >
                    <b-form-input
                    class="form-input"
                    v-model="company.days"
                    placeholder="Введите кодчиество дней"
                    ></b-form-input>
                </b-form-group>
            </b-form>
            <div>
                <h2 id="h2">Настройки кампании</h2>
                <router-link :to="{path: '/company/'+ company.Id, query: { isEdit: true }}">
                <b-button 
                    variant="primary"
                    class="main-button-grey"
                >
                    Настройки компании
                </b-button>
                </router-link>
            </div>
            <b-button
                class="submit-button"
                type="submit"
            >
                {{isEdit ? "Обновить":"Запустить"}} кампанию
            </b-button>
        </div>
    </div>
</template>
<script>
import store from '../../../store/store'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL;
import axios from 'axios'
import popup from '../popup.vue'
export default {
    components: {
      popup
    },
    data(){
        return {
            store,
            isInfoPopupVisible: false,
            label_cols: this.getWidth().label,
            content_cols: this.getWidth().content,
            company: {
                FbPageId: '',
                CompanyName: '',
                CompnayPurpose: '',
                CompanyField: '',
                BusinessAdress: '',
                Images: [],
                ImagesDescription: [],
                ImagesSmall: [],
                ImagesSmallDescription: [],
                CreativeStatus: '',
                PostDescription: '',
                CurrentAmount: 0,
                DailyAmount: 0,
                Days: 0,
            },
        }
    },
    mounted() {
        
    },
    watch: {
        $route(to) {
            if (!(typeof to.params.id === 'undefined')){
                axios({url: `${VUE_APP_API_URL}/api/company/${to.params.id}`, method: 'GET' })
                .then(resp => {
                    this.company = resp.data
                })
            }
        }
    },
    computed: {
        isEdit() {
            return this.$route.query.isEdit
        }
    },
    beforeMount(){
        let mScript = document.createElement('script')
        mScript.setAttribute('src', 'https://yookassa.ru/checkout-widget/v1/checkout-widget.js')
        document.head.appendChild(mScript)

        axios({url: `${VUE_APP_API_URL}/api/company/${this.$router.history.current.params.id}`, method: 'GET' })
        .then(resp => {
            this.company = resp.data
        }) 
    },
    methods: {
        render(){
            const checkout = new window.YooMoneyCheckoutWidget({
                confirmation_token: 'confirmation-token', //Token that must be obtained from YooMoney before the payment process
                return_url: 'https://merchant.site', //URL to the payment completion page
                error_callback: function(error) {
                    console.log(error)
                    //Processing of initialization errors
                }
            })
            checkout.render('payment-form')
            //После отображения платежной формы метод render возвращает Promise (можно не использовать).
            .then(() => {
                //Код, который нужно выполнить после отображения платежной формы.
            });
        },
        closeInfoPopup() {
            this.isInfoPopupVisible = false;
        },
        showPopupInfo() {
            this.isInfoPopupVisible = true;
        },
        startCompany(){
            console.log(this.balanceForm)
        },
        isMoney(company){
            return company.CurrentAmount > 0
        },
        getStatus(company){
            if (company.FbPageId.length === 0) {
                return "facebook не подключен"
            }
            if (company.CurrentAmount === 0){
                return "закончился рекламный бюджет"
            }
            return "в работе"
        },
        isFb(company){
            return company.FbPageId.length > 0
        },
        getWidth() {
            let width = Math.max(
                document.body.scrollWidth,
                document.documentElement.scrollWidth,
                document.body.offsetWidth,
                document.documentElement.offsetWidth,
                document.documentElement.clientWidth
            );
            if (width < 570){
                return {
                    label:12,
                    content:12
                };
            }
            return {
                    label:3,
                    content:9
            };
        },

    }
}
</script>
<style>
#c-status{
    margin: 12px;
}
#company-status-wrapper{
    display: flex;
    align-items: stretch;
}
#balance{
    font-family: Montserrat;
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
  right: 0px;
  top: 0px;
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
    .company-status{
        height: auto;
        position:relative;
    }
}
</style>