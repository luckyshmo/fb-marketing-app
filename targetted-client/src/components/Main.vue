<template>
    <div id="content">
        <h1 id="h1">Главная</h1>
        <h2 id="h2">Ваши рекламные кампании</h2>
        <p id="p1">Создайте рекламную кампанию, чтобы приводить новых клиентов в ваш бизнес</p>
        <router-link :to="{name: 'createCompany'}">
            <b-button variant="primary" id="main-button">Создать компанию</b-button>
        </router-link>

        <div v-if="store.getters.GET_COMPANY_LIST.length > 0">
            <h2 id="h2">Ваши рекламные компании</h2>
            <div>
                <div
                v-for="company in store.getters.GET_COMPANY_LIST" 
                :key="company.Id">
                    <div class="c-div">
                        <div id='l'>
                            <p id="c-name">
                                {{company.CompanyName}}
                            </p>
                        </div>
                        <div id='r'>
                            <div id="c-status">
                                <div id="elipse"></div>
                                <p id="c-status-text">запущена</p>
                            </div>
                        </div>
                    </div>
                </div>                
            </div>
        </div>
    </div>
</template>
<script>
import store from '../../store/store'
export default {
    name: 'CreateCompany',
    data() {
        return{
            store,
        }
    },
    methods: {
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
    margin-top: 20px;
    background: lightgray;
    border-radius: 20px;
    height: 100px;
    width:100%;
    overflow: hidden;
  position: relative;
}
#l {
  position: absolute;
  text-align: center;
  margin-top: 35px;
  margin-left: 30px;
  height: 25%;
}
#r {
  position: absolute;
  right: 0px;
  top: 0px;
  margin-right: 30px;
  margin-top: 25px;
  height: 25%;
}
#c-status-text{
    color: white;  
    font-family: Montserrat;
    font-style: normal;
    font-size: 1em;
}
#elipse{
    background: lightgreen;
    position: absolute;
    width: 15px;
    margin-top: 5px;
    margin-left: 5px;
    height: 15px;
    border-radius: 7.5px;
}
#c-status{
    text-align: center;
    padding-top: 13px;
    padding-left: 10px;
    width: 150px;
    height: 50px;
    background: #000000;
    border-radius: 16px;
}
#c-name{
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