<template lang="">
    <div 
    id="content">
    <div v-if="showInfo">
        <div>
            <b-form-input
            class="form-input"
            disabled-field
            style="margin-bottom: 20px"
            v-model="fbIDforClaim"
            placeholder="facebook id"
            ></b-form-input>
            <button
            @click="makeClaimRequest()" 
            class="main-button"
            style="">
                Выслать запрос
            </button>

            <button
            @click="getPendingPages()" 
            class="main-button"
            style="">
                Pending pages
            </button>
            <div
            v-for="page in pendingPages"
            :key=page>
                <a target="_blank" :href="getFBRedirect(page)" > {{getFBRedirect(page)}} </a>
                <b-icon
                icon="trash-fill"
                aria-hidden="true"
                @click="removePageFromPendingList(page)"></b-icon>
            </div>
        </div>

        <div class="overflow-auto" style="margin-top: 50px;">

            <b-pagination
            v-model="currentPage"
            :total-rows="rows"
            :per-page="perPage"
            aria-controls="my-table"
            ></b-pagination>

            <b-list
            id="my-table"
            :items="users"
            :per-page="perPage"
            :current-page="currentPage"
            small
            >
            </b-list>

            <div
        v-for="user in getPUsers(currentPage, perPage)"
        :key=user.email
        >
            <div class="user-content">
                <!-- {{user}} -->
                <h3>Данные пользователя</h3>
                <p>ID: <b>{{user.id}}</b></p>
                <p>Имя: <b>{{user.name}}</b></p>
                <p>email: <b>{{user.email}}</b></p>
                <p>телефон: <b>{{user.phoneNumber}}</b></p>
                <p>баланс: <b>{{user.amount}}</b></p>
                <p>дата регистрации: <b>{{user.TimeCreated}}</b></p>
                <b-form-input
                class="form-input"
                disabled-field
                style="margin-bottom: 20px"
                v-model="user.amount"
                placeholder="Введите сумму"
                ></b-form-input>
                <button
                @click="changeBalance(user.id, user.amount)" 
                class="main-button"
                style="margin-bottom: 25px">
                    Поменять баланс
                </button>  
                <h3>Данные кампаний</h3>
                <div
                v-for="company in filteredCompanies(user.id)"
                :key="company.Id"
                style="padding: 30px; background: rgb(255, 255, 255); margin: 10px; border-radius: 30px"
                >
                <!-- {{company}} -->
                    <h4>{{company.CompanyName}}</h4>
                    <p>facebook page ID: 
                        <b>{{company.FbPageId}}</b>
                    </p>
                    <div v-if="company.FbPageId.length > 0">
                        <a target="_blank" :href="getFBRedirect(company.FbPageId)" > link </a>
                        <button
                        @click="checkIfPageIsOwned(company.FbPageId)" 
                        class="main-button"
                        style="">
                            Проверить
                        </button>
                    </div>
                    <p>ID: <b>{{company.Id}}</b></p>
                    <p>Название: <b>{{company.CompanyName}}</b></p>
                    <p>Цель: <b>{{company.CompnayPurpose}}</b></p>
                    <p>Сфера деятельности: <b>{{company.CompanyField}}</b></p>
                    <p>Бизнесс адрес: <b>{{company.BusinessAddress}}</b></p>
                    <p>Статус креативов: <b>{{company.CreativeStatus}}</b></p>
                    <p>Надписи сториз: <b>{{company.ImagesDescription}}</b></p>
                    <p>Написи постов: <b>{{company.ImagesSmallDescription}}</b></p>
                    <p>Описание под постом: <b>{{company.PostDescription}}</b></p>
                    <p>Дневной бюджет <b>{{company.DailyAmount}}</b></p>
                    <p>Количество дней: <b>{{company.Days}}</b></p>
                    <p>Дата создания: <b>{{company.CreationDate}}</b></p>
                    <p>Дата запуска: <b>{{company.StartDate}}</b></p>
                    <!-- <button class="main-button">Выгрузить креативы</button> -->
                    <div>
                        <h5>Истории</h5>
                        <div id="image-block">
                            <div 
                            v-for="(image) in filteredImageNames(company.Id, true)"
                            :key="image.name">
                                <div>
                                    <img 
                                    id="preview"
                                    :src="getImageByName(image.name, company.UserId, company.Id)"/>
                                </div>
                            </div>
                        </div>
                        <h5>Посты</h5>
                        <div id="image-block">
                            <div 
                            v-for="(image) in filteredImageNames(company.Id, false)"
                            :key="image.name">
                                <div>
                                    <img 
                                    id="preview"
                                    :src="getImageByName(image.name, company.UserId, company.Id)"/>
                                </div>
                            </div>
                        </div>
                    </div>
                    
                    <button 
                    class="main-button"
                    style="margin-top: 25px"
                    @click="stopCompany(company.Id)"
                    v-if="company.IsStarted">
                        Остановить кампанию
                    </button>
                    <button 
                    class="main-button"
                    style="margin-top: 25px"
                    @click="startCompany(company.Id)"
                    v-if="!company.IsStarted">
                        Запустить кампанию
                    </button>  
                </div>
            </div>
        </div>
        </div>
        
    </div>       
    </div>
</template>
<script>
import axios from 'axios'
import store from '../../../store/store'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL;
export default {
    data(){
        return {
            currentPage: 1,
            perPage: 5,
            showInfo: false,
            store,
            users: [],
            companies: [],
            imageNames: [],
            fbIDforClaim: "",
            pendingPages: [],
        }
    },
    $route() {
        console.log("route ", store.getters.GET_EMAIL)
        this.showInfo = store.getters.GET_EMAIL === 'admin@admin.com'
        this.getUsers()
    },
    beforeMount() {
        console.log("BM ", store.getters.GET_EMAIL)
        this.showInfo = store.getters.GET_EMAIL === 'admin@admin.com'
    },
    mounted() {
        console.log("M ", store.getters.GET_EMAIL)
        this.showInfo = store.getters.GET_EMAIL === 'admin@admin.com'
        this.getUsers()
    },
    computed: {
      rows() {
        return this.users.length
      }
    },
    methods: {
        getPUsers(pageNumber, perPage){
            return this.users.slice(pageNumber*perPage-perPage, pageNumber*perPage)
        },
        getFBRedirect(id){
            return `https://facebook.com/${id}`
        },
        changeBalance(id, amount){
            axios({url: `${VUE_APP_API_URL}/api/user/${id}/update-balance/${amount}`, method: 'POST' })
            .then(resp => {
                this.$alert(resp.statusText)
            })
            .catch(err => {
                this.$alert(err.response.data.message)
            })
        },
        startCompany(id){
            axios({url: `${VUE_APP_API_URL}/api/company/start/${id}`, method: 'POST' })
            .then(resp => {
                this.$alert(resp.statusText)
            })
            .catch(err => {
                this.$alert(err.response.data.message)
            })
        },
        stopCompany(id){
            axios({url: `${VUE_APP_API_URL}/api/company/stop/${id}`, method: 'POST' })
            .then(resp => {
                this.$alert(resp.statusText)
            })
            .catch(err => {
                this.$alert(err.response.data.message)
            })
        },
        isStoriesImage(name){
            return name.includes("stories")
        },
        filteredImageNames(id, isStories){
            return this.imageNames.filter(image => image.id === id && this.isStoriesImage(image.name) === isStories)
        },
        filteredCompanies(id){
            return this.companies.filter(c => c.UserId === id);
        },
        getPendingPages(){
            axios({url: `${VUE_APP_API_URL}/api/facebook/pending`, method: 'GET' })
            .then(resp => {
                console.log(resp)
                this.pendingPages = resp.data
            })
            .catch(err => {
                console.log(err.response)
            })
        },
        checkIfPageIsOwned(id){
            axios({url: `${VUE_APP_API_URL}/api/facebook/owned/${id}`, method: 'GET' })
            .then(resp => {
                console.log(resp)
                this.$alert("Кампания доавблена в БМ")
            })
            .catch(err => {
                console.log(err.response)
                this.$alert("Кампания отсутствует в БМ")
            })
        },
        makeClaimRequest(){
            axios({url: `${VUE_APP_API_URL}/api/facebook/claim/${this.fbIDforClaim}`, method: 'POST' })
            .then(resp => {
                console.log(resp)
            })
            .catch(err => {
                console.log(err.response)
            })
        },
        removePageFromPendingList(id){
            axios({url: `${VUE_APP_API_URL}/api/facebook/pending/${id}`, method: 'DELETE' })
            .then(resp => {
                this.pendingPages = this.pendingPages.filter(ID => ID != id);
                console.log(resp)
            })
            .catch(err => {
                this.$alert(err.response)
                console.log(err.response)
            })
        },
        getUsers(){
            axios({url: `${VUE_APP_API_URL}/api/user/`, method: 'GET' })
            .then(resp => {
                this.users = resp.data
                for (let i = 0; i < this.users.length; i++){
                    this.getAddCompanies(this.users[i].id)
                }
            })
            .catch(err => {
                console.log(err)
            })
        },
        getAddCompanies(id){
            axios({ url: `${VUE_APP_API_URL}/api/user/${id}/company/`, method: 'GET' })
            .then(resp => {
                if (resp.data != null){
                    this.companies.push(...resp.data)
                    for (let i = 0; i < resp.data.length; i++){
                        this.getCompanyImagesNames(resp.data[i].Id, resp.data[i].UserId)
                    }
                }
            });
        },
        getCompanyImagesNames(cID, uID){
            axios({url: `${VUE_APP_API_URL}/api/user/${uID}/company/${cID}/images/`, method: 'GET' })
            .then(resp => {
                if (resp.data != null){
                    for (let i = 0; i < resp.data.length; i++){
                        this.imageNames.push({
                            id: cID,
                            name: resp.data[i],
                        })
                    }
                }
            });
        },
        getImageByName(name, uID, cID){
            return `https://client.targetted.online/images/${uID}/${cID}${name}`
        },
    }
}
</script>
<style>
.user-content{
    background-color: rgb(235, 235, 235);
    border-radius: 30px;
    padding: 20px;
    margin-bottom: 30px;
}
.swal2-styled.swal2-confirm {
    background-color: #6C1BD2 !important;
}
</style>