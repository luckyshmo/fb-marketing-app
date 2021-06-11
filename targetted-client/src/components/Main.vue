<template>
    <div id="content">
        <h1 id="h1">Главная</h1>
        <h2 id="h2">Ваши рекламные кампании</h2>
        <p id="p1">Создайте рекламную кампанию, чтобы приводить новых клиентов в ваш бизнес</p>
        <router-link :to="{path: '/company'}">
            <b-button variant="primary" id="main-button">Создать компанию</b-button>
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
                                    <div class="elipse" id="red" v-if="!isFb(company)"></div>
                                    <div class="elipse" id="green" v-if="isFb(company)"></div>
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
    name: "main",
    data() {
        return{
            store,
        }
    },
    methods: {
        isFb(company){
            return company.FbPageId.length > 0
        },
        getStatus(company){
            if (company.FbPageId.length === 0) {
                return "FB не подключен"
            }
            return "Запущена"
        },
        getImageByName(name){
            console.log("store USER: ", store.getters.GET_COMPANY_DATA.c.UserId)
            console.log("store COMPANY: ", store.getters.GET_COMPANY_DATA.c.Id)
            let uID = store.getters.GET_COMPANY_DATA.c.UserId
            let cID = store.getters.GET_COMPANY_DATA.c.Id
            return `https://client.targetted.online/images/${uID}/${cID}/${name}`
        },
        getAdCompanyList(){
            store.dispatch("getCompanyList")
        },
        printComData(id){
            store.dispatch("getCompanyByID", id)
        }
    },
    mounted(){
        this.getAdCompanyList()
    }
}
</script>
<style>
.c-div{
    cursor: pointer;
    margin-top: 20px;
    background: lightgray;
    border-radius: 20px;
    height: 100px;
    width:100%;
    overflow: hidden;
    position: relative;
}
.l {
  position: absolute;
  text-align: center;
  margin-top: 35px;
  margin-left: 30px;
  height: 25%;
}
.r {
  position: absolute;
  right: 0px;
  top: 0px;
  margin-right: 30px;
  margin-top: 25px;
  height: 25%;
}
.c-status-text{
    color: white;  
    margin-left: 30px;
    margin-right: 20px;
    font-family: Montserrat;
    font-style: normal;
    font-size: 1em;
}
#red{
    background: red;
}
#green{
    background: lightgreen;
}
.elipse{
    position: absolute;
    width: 15px;
    margin-top: 5px;
    margin-left: 5px;
    height: 15px;
    border-radius: 7.5px;
}
.c-status{
    text-align: center;
    padding-top: 13px;
    padding-left: 10px;
    /* width: 150px; */
    height: 50px;
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
    width: 450px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}
</style>