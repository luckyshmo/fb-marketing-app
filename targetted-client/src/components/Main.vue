<template>
    <div id="content">
        <h1 id="h1">Главная</h1>
        <h2 id="h2">Ваши рекламные кампании</h2>
        <p id="p1">Создайте рекламную кампанию, чтобы приводить новых клиентов в ваш бизнес</p>
        <router-link v-if="store.getters.GET_COMPANY_LIST.length < 3" :to="{path: '/company'}">
            <b-button variant="primary" class="main-button">Создать компанию</b-button>
        </router-link>
        <div v-if="store.getters.GET_COMPANY_LIST.length > 0">
            <h2 id="h2">Ваши рекламные кампании</h2>
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
    </div>
</template>
<script>
import store from '../../store/store'
export default {
    name: "main-page",
    data() {
        return{
            store,
        }
    },
    watch: {
        $route() {
            store.dispatch("getCompanyList")
        }
    },
    methods: {
        isMoney(company){
            return company.CurrentAmount > 0
        },
        isFb(company){
            return company.FbPageId.length > 0
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