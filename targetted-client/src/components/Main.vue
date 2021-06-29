<template>
    <div id="content">
        <h1 id="h1">Главная</h1>
        <div v-if="store.getters.GET_COMPANY_LIST.length > 0">
            <h2 id="h2">Общий баланс</h2>

            <b-form-group
                label="Всего на счету"
                :label-cols="3"
                :content-cols="7"
                id="input-group-main"
                label-for="input-horizontal"
            >
                <p id="balance">{{user.amount}} ₽</p>
            </b-form-group>

            <b-button 
                variant="primary"
                class="main-button-grey"
                @click="push"
            >
                Пополнить баланс
            </b-button>
        </div>
        <h2 id="h2">Ваши рекламные кампании</h2>
        <p 
        v-if="!(store.getters.GET_COMPANY_LIST.length > 0)" 
        id="p1">Создайте рекламную кампанию, чтобы приводить новых<br>клиентов в ваш бизнес</p>
        <div v-if="store.getters.GET_COMPANY_LIST.length > 0">
            <div>
                <div
                v-for="company in store.getters.GET_COMPANY_LIST" 
                :key="company.Id"
                >
                    <router-link :to="{path: '/company-balance/'+ company.Id, query: { isEdit: true }}">
                        <div class="c-div">
                            <div class='l'>
                                <p class="c-name">
                                    {{company.CompanyName}}
                                </p>
                            </div>
                            <div class='r'>
                                <div class="c-status">
                                    <div class="elipse" id="white" v-if="!isFb(company)"></div>
                                    <div class="elipse" id="yellow" v-if="isFb(company) && !isMoney(company)"></div>
                                    <div class="elipse" id="green" v-if="isFb(company) && isMoney(company)"></div>
                                    <p class="c-status-text">{{getStatus(company)}}</p>
                                </div>
                            </div>
                        </div>
                    </router-link>
                </div>                
            </div>
        </div>
        <router-link v-if="store.getters.GET_COMPANY_LIST.length < 3" :to="{path: '/company'}">
            <b-button style="margin-top: 32px" variant="primary" class="main-button">Создать кампанию</b-button>
        </router-link>
    </div>
</template>
<script>
import store from '../../store/store'
import axios from 'axios'
import router from '../../router/router'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL //TODO GLOABL APP CONST
export default {
    name: "main-page",
    data() {
        return{
            user: {},
            store,
        }
    },
    beforeMount(){
        axios({url: `${VUE_APP_API_URL}/api/user/0`, method: 'GET' })
        .then(resp => {
            this.user = resp.data
        })
    },
    watch: {
        $route(to) {
            store.dispatch("getCompanyList")
            console.log("TO", to)
            if (to.name === "mainPage"){
                axios({url: `${VUE_APP_API_URL}/api/user/0`, method: 'GET' })
                .then(resp => {
                    this.user = resp.data
                })
            }
        }
    },
    methods: {
        push(){
            router.push({path: '/company-balance/'+ store.getters.GET_COMPANY_LIST[0].Id, query: { isEdit: true }})
        },
        isMoney(company){
            return company.CurrentAmount > 0
        },
        isFb(company){
            return company.FbPageId.length > 0
        },
        getStatus(company){
            // if (company.FbPageId.length === 0) {
            //     return "facebook не подключен"
            // }
            if (company.CurrentAmount === 0){
                return "закончился рекламный бюджет"
            }
            if (!company.IsStarted){
                return "настраивается"
            }
            return "работает"
        },
        getAdCompanyList(){
            store.dispatch("getCompanyList")
        },
    },
    mounted(){
        this.getAdCompanyList()
    }
}
</script>
<style>
.c-div{
    cursor: pointer;
    margin-top: 40px;
    background: #F3F3F3;
    border-radius: 20px;
    height: 112px;
    width: 100%;
    overflow: hidden;
    position: relative;
}
.l {
  position: absolute;
  text-align: center;
  margin-top: 46px;
  margin-left: 35px;
  height: 25%;
}
a {
    text-decoration: none !important;
}
.r {
  position: absolute;
  right: 0px;
  top: 0px;
  margin-right: 30px;
  margin-top: 29px;
  height: 25%;
}
.c-status-text{
    color: white;  
    margin-left: 30px;
    margin-right: 20px;
    margin-top: 2px;
    font-family: Montserrat;
    font-style: normal;
    padding: 0px !important;
}
#white{
    background: white;
}
#yellow{
    background: yellow;
}
#green{
    background: lightgreen;
}
.elipse{
    position: absolute;
    width: 10px;
    margin-top: 9px;
    margin-left: 9px;
    height: 10px;
    border-radius: 5px;
}
.c-status{
    text-align: center;
    padding-top: 13px;
    padding-left: 10px;
    height: 55px;
    padding-bottom: 1px; 
    background: #000000;
    border-radius: 16px;
}
.c-name{
    font-family: Montserrat;
    font-style: normal;
    font-weight: bold;
    font-size: 1.3em;
    color: #000000;
    text-align: left;
    width: 600px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}
@media (max-width: 600px){
    .c-name{
        max-width: 260px; text-overflow: ellipsis; overflow: hidden; white-space: nowrap;
    }
    .c-status{
       height: auto;
       padding-bottom: 4px; 
    }
    .c-status-text{
        margin-left: 15px;
        margin-right: 10px;
        margin-bottom: 10px;
    }
    .c-div{
        margin-top: 15px;
        height: auto;
    }
    .l {
        position: relative;
        text-align: left;
        margin-top: 20px;
        margin-left: 20px;
        height: auto;
    }
    .r {
        position: relative;
        right: auto;
        top: auto;
        margin-right: auto;
        margin-top: auto;
        height: auto;
        margin: 20px;
        
    }
}
</style>