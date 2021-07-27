<template lang="">
  <div
    id="content"
  >
    <loading v-if="isLoading"/>
    <div v-if="showInfo">
      <div>
        <b-form-input
          v-model="fbIDforClaim"
          class="form-input"
          disabled-field
          style="margin-bottom: 20px"
          placeholder="facebook id"
        />
        <button
          class="main-button"
          style=""
          @click="makeClaimRequest()"
        >
          Выслать запрос
        </button>

        <button
          class="main-button"
          style=""
          @click="getPendingPages()"
        >
          Pending pages
        </button>
        <div
          v-for="page in pendingPages"
          :key="page"
        >
          <a
            target="_blank"
            :href="getFBRedirect(page)"
          > {{ getFBRedirect(page) }} </a>
          <b-icon
            icon="trash-fill"
            aria-hidden="true"
            @click="removePageFromPendingList(page)"
          />
        </div>
      </div>

      <div
        class="overflow-auto"
        style="margin-top: 50px;"
      >
        <b-form-input
          v-model="userFilter"
          class="form-input"
          disabled-field
          style="margin-bottom: 20px"
          placeholder="Поиск по людям"
        />

        <b-form-input
          v-model="companyFilter"
          class="form-input"
          disabled-field
          style="margin-bottom: 20px"
          placeholder="Поиск по кампаниям"
        />

        <b-form-checkbox
          v-model="isAmount"
          value="true"
          unchecked-value="false"
        >
          Only users with amount
        </b-form-checkbox>

        <b-pagination
          v-model="currentPage"
          :total-rows="rows"
          :per-page="perPage"
          aria-controls="my-table"
        />

        <div
          v-for="user in getPUsers(currentPage, perPage)"
          :key="user.email"
        >
          <div class="user-content">
            <h3>Данные пользователя</h3>
            <p>ID: <b>{{ user.id }}</b></p>
            <p>Имя: <b>{{ user.name }}</b></p>
            <p>email: <b>{{ user.email }}</b></p>
            <p>телефон: <b>{{ user.phoneNumber }}</b></p>
            <p>баланс: <b>{{ user.amount }}</b></p>
            <p>дата регистрации: <b>{{ user.TimeCreated }}</b></p>
            <b-form-input
              v-model="user.amount"
              class="form-input"
              disabled-field
              style="margin-bottom: 20px"
              placeholder="Введите сумму"
            />
            <button
              class="main-button"
              style="margin-bottom: 25px"
              @click="changeBalance(user.id, user.amount)"
            >
              Поменять баланс
            </button>
            <h3>Данные кампаний</h3>
            <div
              v-for="company in user.companies"
              :key="company.Id"
              style="padding: 30px; background: rgb(255, 255, 255); margin: 10px; border-radius: 30px"
            >
              <h4>{{ company.CompanyName }}</h4>
              <p>
                facebook page ID:
                <b>{{ company.FbPageId }}</b>
              </p>
              <div v-if="company.FbPageId.length > 0">
                <a
                  target="_blank"
                  :href="getFBRedirect(company.FbPageId)"
                > link </a>
                <button
                  class="main-button"
                  style=""
                  @click="checkIfPageIsOwned(company.FbPageId)"
                >
                  Проверить
                </button>
              </div>
              <p>ID: <b>{{ company.Id }}</b></p>
              <p>Название: <b>{{ company.CompanyName }}</b></p>
              <p>Цель: <b>{{ company.CompnayPurpose }}</b></p>
              <p>Сфера деятельности: <b>{{ company.CompanyField }}</b></p>
              <p>Бизнесс адрес: <b>{{ company.BusinessAddress }}</b></p>
              <p>Статус креативов: <b>{{ company.CreativeStatus }}</b></p>
              <p>Надписи сториз: <b>{{ company.ImagesDescription }}</b></p>
              <p>Написи постов: <b>{{ company.ImagesSmallDescription }}</b></p>
              <p>Описание под постом: <b>{{ company.PostDescription }}</b></p>
              <p>Дневной бюджет <b>{{ company.DailyAmount }}</b></p>
              <p>Количество дней: <b>{{ company.Days }}</b></p>
              <p>Дата создания: <b>{{ company.CreationDate }}</b></p>
              <p>Дата запуска: <b>{{ company.StartDate }}</b></p>
              <!-- <button class="main-button">Выгрузить креативы</button> -->
              <div>
                <h5>Истории</h5>
                <div id="image-block">
                  <div
                    v-for="(image) in filteredImageNames(company.Id, true)"
                    :key="image.name"
                  >
                    <div>
                      <img
                        id="preview"
                        :src="getImageByName(image.name, company.UserId, company.Id)"
                      >
                    </div>
                  </div>
                </div>
                <h5>Посты</h5>
                <div id="image-block">
                  <div
                    v-for="(image) in filteredImageNames(company.Id, false)"
                    :key="image.name"
                  >
                    <div>
                      <img
                        id="preview"
                        :src="getImageByName(image.name, company.UserId, company.Id)"
                      >
                    </div>
                  </div>
                </div>
              </div>

              <button
                v-if="company.IsStarted"
                class="main-button"
                style="margin-top: 25px"
                @click="stopCompany(company.Id)"
              >
                Остановить кампанию
              </button>
              <button
                v-if="!company.IsStarted"
                class="main-button"
                style="margin-top: 25px"
                @click="startCompany(company.Id)"
              >
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
import loading from '../Loading.vue'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL
export default {
  components: {
    loading
  },
  data () {
    return {
      isLoading: false,
      userFilter: '',
      companyFilter: '',
      isAmount: false,
      currentPage: 1,
      perPage: 15,
      showInfo: false,
      store,
      users: [],
      imageNames: [],
      fbIDforClaim: '',
      pendingPages: []
    }
  },
  $route () {
    this.showInfo = store.getters.GET_EMAIL === 'admin@admin.com'
    this.getUsers()
  },
  computed: {
    rows () {
      return this.filteredUsers().length
    }
  },
  beforeMount () {
    this.showInfo = store.getters.GET_EMAIL === 'admin@admin.com'
  },
  mounted () {
    this.showInfo = store.getters.GET_EMAIL === 'admin@admin.com'
    this.getUsers()
  },
  methods: {
    filteredUsers () {
      this.isLoading = true

      let uArr = this.users
      if (this.isAmount === 'true') {
        this.currentPage = 1
        uArr = this.users.filter((user) => {
          return user.amount > 0
        })
      }
      if (this.userFilter.length > 1) {
        this.currentPage = 1
        uArr = uArr.filter((user) => {
          const isFiltered = (user.name.toLowerCase().match(this.userFilter.trim().toLowerCase()) ||
          user.email.toLowerCase().match(this.userFilter.trim().toLowerCase()) ||
          user.phoneNumber.toLowerCase().match(this.userFilter.trim().toLowerCase()) ||
          user.id.toLowerCase().match(this.userFilter.trim().toLowerCase()) ||
          user.TimeCreated.toLowerCase().match(this.userFilter.trim().toLowerCase()))
          return isFiltered
        })
      }
      if (this.companyFilter.length > 1) {
        this.currentPage = 1
        uArr = uArr.filter((user) => {
          return user.companies.filter((com) => {
            const isF = com.Id.toLowerCase().match(this.companyFilter.trim().toLowerCase()) ||
            com.CompanyName.toLowerCase().match(this.companyFilter.trim().toLowerCase()) ||
            com.PostDescription.toLowerCase().match(this.companyFilter.trim().toLowerCase())
            return isF
          }).length > 0
        })
      }

      this.isLoading = false

      return uArr
    },
    getPUsers (pageNumber, perPage) {
      return this.filteredUsers().slice(pageNumber * perPage - perPage, pageNumber * perPage)
    },
    getFBRedirect (id) {
      return `https://facebook.com/${id}`
    },
    changeBalance (id, amount) {
      axios({ url: `${VUE_APP_API_URL}/api/user/${id}/update-balance/${amount}`, method: 'POST' })
        .then(resp => {
          this.$alert(resp.statusText)
        })
        .catch(err => {
          this.$alert(err.response.data.message)
        })
    },
    startCompany (id) {
      axios({ url: `${VUE_APP_API_URL}/api/company/start/${id}`, method: 'POST' })
        .then(resp => {
          this.$alert(resp.statusText)
        })
        .catch(err => {
          this.$alert(err.response.data.message)
        })
    },
    stopCompany (id) {
      axios({ url: `${VUE_APP_API_URL}/api/company/stop/${id}`, method: 'POST' })
        .then(resp => {
          this.$alert(resp.statusText)
        })
        .catch(err => {
          this.$alert(err.response.data.message)
        })
    },
    isStoriesImage (name) {
      return name.includes('stories')
    },
    filteredImageNames (id, isStories) {
      return this.imageNames.filter(image => image.id === id && this.isStoriesImage(image.name) === isStories)
    },
    getPendingPages () {
      axios({ url: `${VUE_APP_API_URL}/api/facebook/pending`, method: 'GET' })
        .then(resp => {
          this.pendingPages = resp.data
        })
        .catch(err => {
          alert(err.response)
          console.log(err.response)
        })
    },
    checkIfPageIsOwned (id) {
      axios({ url: `${VUE_APP_API_URL}/api/facebook/owned/${id}`, method: 'GET' })
        .then(resp => {
          console.log(resp)
          this.$alert('Кампания доавблена в БМ')
        })
        .catch(err => {
          console.log(err.response)
          this.$alert('Кампания отсутствует в БМ')
        })
    },
    makeClaimRequest () {
      axios({ url: `${VUE_APP_API_URL}/api/facebook/claim/${this.fbIDforClaim}`, method: 'POST' })
        .then(resp => {
          console.log(resp)
        })
        .catch(err => {
          console.log(err.response)
        })
    },
    removePageFromPendingList (id) {
      axios({ url: `${VUE_APP_API_URL}/api/facebook/pending/${id}`, method: 'DELETE' })
        .then(resp => {
          this.pendingPages = this.pendingPages.filter(ID => ID !== id)
          console.log(resp)
        })
        .catch(err => {
          this.$alert(err.response)
          console.log(err.response)
        })
    },
    getUsers () {
      axios({ url: `${VUE_APP_API_URL}/api/user/`, method: 'GET' })
        .then(resp => {
          this.users = resp.data
          for (let i = 0; i < this.users.length; i++) {
            this.getAddCompanies(this.users[i].id, i)
          }
        })
        .catch(err => {
          console.log(err)
        })
    },
    getAddCompanies (id, i) { // вообще во VUE 3 должно работать...
      axios({ url: `${VUE_APP_API_URL}/api/user/${id}/company/`, method: 'GET' })
        .then(resp => {
          this.users[i].companies = []
          if (resp.data != null) {
            this.users[i].companies.push(...resp.data)
            for (let i = 0; i < resp.data.length; i++) {
              this.getCompanyImagesNames(resp.data[i].Id, resp.data[i].UserId)
            }
            this.isAmount = true
            this.isAmount = false //! пздц костыль
          }
        })
        .catch(err => {
          console.log(err)
          this.users[i].companies = []
        })
    },
    getCompanyImagesNames (cID, uID) {
      axios({ url: `${VUE_APP_API_URL}/api/user/${uID}/company/${cID}/images/`, method: 'GET' })
        .then(resp => {
          if (resp.data != null) {
            for (let i = 0; i < resp.data.length; i++) {
              this.imageNames.push({
                id: cID,
                name: resp.data[i]
              })
            }
          }
        })
    },
    getImageByName (name, uID, cID) {
      return `https://client.targetted.online/images/${uID}/${cID}${name}`
    }
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
