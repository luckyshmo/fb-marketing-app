<template>
    <div id="content-wrapper">
        <div id="content">
            <router-link :to="{name: 'mainPage'}">
                <p id="navigation-text" style="margin:0;">← К списку кампаний</p>
            </router-link>
            <div class="company-status">
                <div class='left'>
                    <h1 id="h1">Название кампании</h1>
                </div>
                <div class='right'>
                    <div class="c-status">
                        <div class="elipse" id="red" v-if="!isFb(company)"></div>
                        <div class="elipse" id="green" v-if="isFb(company)"></div>
                        <p class="c-status-text">{{getStatus(company)}}</p>
                    </div>
                </div>
            </div>
    
            <h2 id="h2">Общий баланс</h2>
            <b-button 
                v-if="store.getters.GET_FB_PAGES.length > 0"
                variant="primary"
                id="main-button"
                style="margin-top: 30px; background: #F3F3F3; color: black"
                @click="logout"
            >
                Привязать другой аккаунт
            </b-button>

            <p>{{company}}</p>

            <b-form @submit.prevent="startCompany()">
                <h2 id="h2">Рекламный бюджет</h2>
                <b-form-group
                        label="Бюджет на сутки"
                        description="Минимальная сумма 200₽ в сутки"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                >
                    <b-form-input
                    class="form-input"
                    v-model="balanceForm.dailyAmount"
                    placeholder="Введите сумму"
                    ></b-form-input>
                </b-form-group>
                <b-form-group
                        label="Количество дней"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                >
                    <b-form-input
                    class="form-input"
                    v-model="balanceForm.days"
                    placeholder="Введите кодчиество дней"
                    ></b-form-input>
                </b-form-group>
            </b-form>
            <h2 id="h2">Настройки кампании</h2>
            <b-button 
                        v-if="store.getters.GET_FB_PAGES.length > 0"
                        variant="primary"
                        id="main-button"
                        style="margin-top: 30px; background: #F3F3F3; color: black"
                        @click="logout"
            >
                Привязать другой аккаунт
            </b-button>
            <b-button
                id="submit-button"
                type="submit"
            >
                Запустить кампанию
            </b-button>
        </div>
    </div>
</template>
<script>
import store from '../../../store/store'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL;
import axios from 'axios'
export default {
    data(){
        return {
            store,
            label_cols: this.getWidth().label,
            content_cols: this.getWidth().content,
            company: {
                FbPageId: 'dEFAULT',
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
            },
            balanceForm: {
                currentAmount: 4500,
                dailyAmount: 1000,
                days: 3,
            }
        }
    },
    watch: {
        $route(to, from) {
            console.log("from", from)
            if (typeof to.params.id !== undefined){
                axios({url: `${VUE_APP_API_URL}/api/company/${to.params.id}`, method: 'GET' })
                    .then(resp => {
                        console.log("setAdd company", resp.data)
                        this.company = resp.data
                    })
            }
            
        }
    },
    methods: {
        startCompany(){
            console.log(this.balanceForm)
        },
        getStatus(company){
            if (company.FbPageId.length === 0) {
                return "FB не подключен"
            }
            return "Запущена"
        },
        isFb(company){
            console.log("try status", company.CompanyName)
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
.company-status{
    height: 70px;
    width:100%;
    overflow: hidden;
    position: relative;
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
</style>