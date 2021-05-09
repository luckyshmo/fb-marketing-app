<template>
    <div id="content">
        <router-link :to="{name: 'mainPage'}">
            <p id="navigation-text">← К списку кампаний</p>
        </router-link>
        <h1 id="h1">Создание компании</h1>
        <h2 id="h2">Доступ к кабинету Facebook</h2>
        <p id="p1">Привяжите свой аккаунт Facebook к targetted, чтобы натсроить и запустить рекламунюу компанию </p>
        <b-button 
            variant="primary"
            id="main-button"
            @click="loginFB"
        >
            Настроить доступ
        </b-button>
        <b-button 
            variant="primary"
            id="main-button"
            @click="logout"
        >
            Разлогин
        </b-button>
        <!-- <div v-if="store.getters.GET_FB_TOKEN.len"> -->
            <p>{{store.getters.GET_FB_TOKEN}}</p>
        <!-- </div> -->
        <h2 id="h2">Основное</h2>
        <p id="p1">Для сроклла</p>
        <p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p><p id="p1">Для сроклла</p>
        <h2 id="h2">Креативы</h2>
    </div>
</template>
<script>
import accountService from '../../_services/account.service';
import store from '../../../store/store'
export default {
    name: 'CreateCompany',
    data() {
        return{
            store,
        }
    },
    methods: {
        loginFB() {
            accountService.login()
        },
        logout(){
            accountService.logout()
        }
    }
}
</script>
<style>
</style>