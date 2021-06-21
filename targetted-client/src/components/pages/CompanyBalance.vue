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
                <p id="balance">{{company.CurrentAmount}}₽</p>
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
              @closePopup="closeInfoPopup"
            >
                <h1 style="margin-top: 60px; margin-bottom: 20px">Пополнение</h1>
                <b-form-input
                class="form-input"
                v-model="paymentAmount"
                :state="validateState1('paymentAmount')"
                placeholder="Введите сумму"
                ></b-form-input>
                <b-form-invalid-feedback 
                class="error-message">
                    Минимальная сумма пополнения 2500₽
                </b-form-invalid-feedback>
                <button 
                class="main-button-big"
                @click="render"
                >Оплатить</button>
                <div id="payment-form"></div>
            </popup>

            <b-form @submit.prevent="startCompany()">
                <h2 id="h2">Рекламный бюджет</h2>
                <b-form-group
                        label="Бюджет на сутки"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group-main"
                        label-for="input-horizontal"
                >
                    <b-form-input
                    class="form-input"
                    v-model="company.DailyAmount"
                    :state="validateState('DailyAmount')"
                    placeholder="Введите сумму"
                    ></b-form-input>
                    <b-form-invalid-feedback 
                    class="error-message">
                        Минимальная сумма 200₽ в сутки
                    </b-form-invalid-feedback>
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
                    :state="validateState('Days')"
                    v-model="company.Days"
                    placeholder="Введите колчиество дней"
                    ></b-form-input>
                    <b-form-invalid-feedback 
                    class="error-message">
                        Минимально один день
                    </b-form-invalid-feedback>
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
                @click="updateCompany()"
            >
                {{isEdit ? "Обновить":"Запустить"}} кампанию
            </b-button>
        </div>
    </div>
</template>
<script>
import store from '../../../store/store'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL;
import router from '../../../router/router'
import axios from 'axios'
import popup from '../popup.vue'
import { validationMixin } from "vuelidate";
import { required, minValue} from "vuelidate/lib/validators";
export default {
    mixins: [validationMixin],
    components: {
      popup
    },
    validations: {
        paymentAmount:{
            required,
            minValue: minValue(1)
        },
        company: {
            DailyAmount: {
                required,
                minValue: minValue(200)
            },
            Days: {
                required,
                maxValue: minValue(1),
            }
        }
    },
    data(){
        return {
            confirmationToken: '',
            paymentAmount: 0,
            paymentID: '',
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
        validateState(name) {
            const { $dirty, $error } = this.$v.company[name];
            return $dirty ? !$error : null;
        },
        validateState1(name) {
            const { $dirty, $error } = this.$v[name];
            return $dirty ? !$error : null;
        },
        render(){
            this.$v.paymentAmount.$touch();
            if (this.$v.paymentAmount.$anyError) {
                return;
            }
            const data = `{
                "amount": "${this.paymentAmount}.00"
            }`
            axios({url: `${VUE_APP_API_URL}/api/company/paymentToken`, data: data, method: 'POST' })
            .then(resp => {
                console.log(resp)
                this.paymentID = resp.data.id
                // setInterval(async () => {
                    axios({url: `${VUE_APP_API_URL}/api/company/paymentStatus/${this.paymentID}`, data: data, method: 'POST' })
                    .then(resp => {
                        console.log("status resp: ", resp.data)
                        if (resp.data.status === "succeeded"){
                            this.company.DailyAmount = this.company.DailyAmount + resp.data.amount.value
                        }
                    })
                    .catch(err => {
                        console.log(err)
                    })
                // }, 3000)
                const checkout = new window.YooMoneyCheckoutWidget({
                    confirmation_token: resp.data.confirmation.confirmation_token, //Token that must be obtained from YooMoney before the payment process
                    return_url: window.location.origin + router.history.current.fullPath, //URL to the payment completion page
                    error_callback: function(error) {
                        console.log(error)
                        //Processing of initialization errors
                    }
                })
                checkout.render('payment-form')
            })
            .catch(err => {
                console.log(err)
            })
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
        updateCompany(){
            this.$v.company.$touch();
            if (this.$v.company.$anyError) {
                return;
            }
            const companyData = new FormData();
            companyData.append("FbPageId", this.company.FbPageId)
            if (this.isEdit) {
                companyData.append("Id", this.company.Id)
            }
            companyData.append("CompanyName", this.company.CompanyName)
            companyData.append("CompnayPurpose", this.company.CompnayPurpose)
            companyData.append("CompanyField", this.company.CompanyField)
            companyData.append("BusinessAddress", this.company.BusinessAddress)
            companyData.append("ImagesDescription", this.company.ImagesDescription)
            companyData.append("ImagesSmallDescription", this.company.ImagesSmallDescription)
            companyData.append("CreativeStatus", this.company.CreativeStatus)
            companyData.append("PostDescription", this.company.PostDescription)
            companyData.append("DailyAmount",this.company.DailyAmount)
            companyData.append("Days",this.company.Days)
            axios({url: `${VUE_APP_API_URL}/api/company/${this.company.Id}`, data: companyData, method: 'PUT' })
            .then(resp => {
                console.log(resp)
                router.push({path: '/main'})
            })
            .catch(err => {
                if (err.response.data.message === 'pq: duplicate key value violates unique constraint "ad_company_name_user_id_key"') {
                    this.isCompanyExist = true;
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