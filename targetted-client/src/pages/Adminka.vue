<template>
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
          v-model="campaignFilter"
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
              v-for="campaign in user.campaigns"
              :key="campaign.Id"
              style="padding: 30px; background: rgb(255, 255, 255); margin: 10px; border-radius: 30px"
            >
              <h4>{{ campaign.CampaignName }}</h4>
              <p>
                facebook page ID:
                <b>{{ campaign.FbPageId }}</b>
              </p>
              <div v-if="campaign.FbPageId.length > 0">
                <a
                  target="_blank"
                  :href="getFBRedirect(campaign.FbPageId)"
                > link </a>
                <button
                  class="main-button"
                  style=""
                  @click="checkIfPageIsOwned(campaign.FbPageId)"
                >
                  Проверить
                </button>
              </div>
              <p>ID: <b>{{ campaign.Id }}</b></p>
              <p>Название: <b>{{ campaign.CampaignName }}</b></p>
              <p>Цель: <b>{{ campaign.CampaignObjective }}</b></p>
              <p>Сфера деятельности: <b>{{ campaign.CampaignField }}</b></p>
              <p>Бизнесс адрес: <b>{{ campaign.BusinessAddress }}</b></p>
              <p>Статус креативов: <b>{{ campaign.CreativeStatus }}</b></p>
              <p>Описание под постом: <b>{{ campaign.PostDescription }}</b></p>
              <p>Дневной бюджет <b>{{ campaign.DailyAmount }}</b></p>
              <p>Количество дней: <b>{{ campaign.Days }}</b></p>
              <p>Дата создания: <b>{{ campaign.CreationDate }}</b></p>
              <p>Дата запуска: <b>{{ campaign.StartDate }}</b></p>
              <!-- <button class="main-button">Выгрузить креативы</button> -->
              <div>
                <h5>Истории</h5>
                <div id="image-block">
                  <div
                    v-for="(image) in filteredImageNames(campaign.Id, true)"
                    :key="image.name"
                  >
                    <div>
                      <img
                        id="preview"
                        :src="getImageByName(image.name, campaign.UserId, campaign.Id)"
                      >
                    </div>
                  </div>
                </div>
                <h5>Посты</h5>
                <div id="image-block">
                  <div
                    v-for="(image) in filteredImageNames(campaign.Id, false)"
                    :key="image.name"
                  >
                    <div>
                      <img
                        id="preview"
                        :src="getImageByName(image.name, campaign.UserId, campaign.Id)"
                      >
                    </div>
                  </div>
                </div>
              </div>

              <button
                v-if="campaign.IsStarted"
                class="main-button"
                style="margin-top: 25px"
                @click="stopCampaign(campaign.Id)"
              >
                Остановить кампанию
              </button>
              <button
                v-if="!campaign.IsStarted"
                class="main-button"
                style="margin-top: 25px"
                @click="startCampaign(campaign.Id)"
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
import { instance as axios } from '@src/_services/index'
import store from '@src/store/store'
import loading from '@src/components/Loading.vue'
import { APP_UI_URL } from '../constants'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL

export default {
  components: {
    loading
  },
  data () {
    return {
      isLoading: false,
      userFilter: '',
      campaignFilter: '',
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
      // this.isLoading = true

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
          return (user.name.toLowerCase().match(this.userFilter.trim().toLowerCase()) ||
          user.email.toLowerCase().match(this.userFilter.trim().toLowerCase()) ||
          user.phoneNumber.toLowerCase().match(this.userFilter.trim().toLowerCase()) ||
          user.id.toLowerCase().match(this.userFilter.trim().toLowerCase()) ||
          user.TimeCreated.toLowerCase().match(this.userFilter.trim().toLowerCase()))
        })
      }

      if (this.campaignFilter.length > 1) {
        this.currentPage = 1
        uArr = uArr.filter((user) => {
          return user.campaigns.filter((com) => {
            return com.Id.toLowerCase().match(this.campaignFilter.trim().toLowerCase()) ||
            com.CampaignName.toLowerCase().match(this.campaignFilter.trim().toLowerCase()) ||
            com.PostDescription.toLowerCase().match(this.campaignFilter.trim().toLowerCase())
          })?.length > 0
        })
      }

      // this.isLoading = false

      return uArr
    },
    getPUsers (pageNumber, perPage) {
      return this.filteredUsers().slice(pageNumber * perPage - perPage, pageNumber * perPage)
    },
    getFBRedirect (id) {
      return `https://facebook.com/${id}`
    },
    changeBalance (id, amount) {
      axios.post({ url: `${VUE_APP_API_URL}/api/user/${id}/update-balance/${amount}` })
        .then(resp => {
          this.$alert(resp.statusText)
        })
        .catch(err => {
          this.$alert(err.response.data.message)
        })
    },
    startCampaign (id) {
      axios.post({ url: `${VUE_APP_API_URL}/api/campaign/start/${id}` })
        .then(resp => {
          this.$alert(resp.statusText)
        })
        .catch(err => {
          this.$alert(err.response.data.message)
        })
    },
    stopCampaign (id) {
      axios.post({ url: `${VUE_APP_API_URL}/api/campaign/stop/${id}` })
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
      axios.get({ url: `${VUE_APP_API_URL}/api/facebook/pending` })
        .then(resp => {
          this.pendingPages = resp.data
        })
        .catch(err => {
          alert(err.response)
          console.log(err.response)
        })
    },
    checkIfPageIsOwned (id) {
      axios.get({ url: `${VUE_APP_API_URL}/api/facebook/owned/${id}` })
        .then(resp => {
          console.log(resp)
          this.$alert('Кампания добавлена в БМ')
        })
        .catch(err => {
          console.log(err.response)
          this.$alert('Кампания отсутствует в БМ')
        })
    },
    makeClaimRequest () {
      axios.post({ url: `${VUE_APP_API_URL}/api/facebook/claim/${this.fbIDforClaim}` })
        .then(resp => {
          console.log(resp)
        })
        .catch(err => {
          console.log(err.response)
        })
    },
    removePageFromPendingList (id) {
      const url = `${VUE_APP_API_URL}/api/facebook/pending/${id}`
      axios.delete({ url })
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
      const url = `${VUE_APP_API_URL}/api/user/`
      axios.get({ url })
        .then(resp => {
          this.users = resp.data
          for (let i = 0; i < this.users.length; i++) {
            this.getAddCampaigns(this.users[i].id, i)
          }
        })
        .catch(err => {
          console.log(err)
        })
    },
    getAddCampaigns (id, i) { // вообще во VUE 3 должно работать...
      const url = `${VUE_APP_API_URL}/api/user/${id}/campaign/`

      axios.get({ url })
        .then(resp => {
          this.users[i].campaigns = []
          if (resp.data != null) {
            this.users[i].campaigns.push(...resp.data)
            for (let i = 0; i < resp.data.length; i++) {
              this.getCampaignImagesNames(resp.data[i].Id, resp.data[i].UserId)
            }
            this.isAmount = true
            this.isAmount = false //! пздц костыль
          }
        })
        .catch(err => {
          console.log(err)
          this.users[i].campaigns = []
        })
    },
    getCampaignImagesNames (cID, uID) {
      axios.get({
        url: `${VUE_APP_API_URL}/api/user/${uID}/campaign/${cID}/images/`
      }).then(resp => {
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
      return `${APP_UI_URL}/images/${uID}/${cID}${name}`
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
