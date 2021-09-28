<template>
  <div id="content">
    <div v-if="store.getters.GET_COMPANY_LIST.length > 0">
      <h1 id="h1">
        Главная
      </h1>
      <div>
        <!-- <h2 id="h2">
          Общий баланс
        </h2> -->
        <!-- <b-form-group
          id="input-group-main"
          label="Всего на счету"
          :label-cols="3"
          :content-cols="7"
          label-for="input-horizontal"
        >
          <p id="balance">
            {{ user.amount }} ₽
          </p>
        </b-form-group> -->

        <b-button
          variant="primary"
          class="main-button-grey mt-2"
          @click="push"
        >
          Пополнить баланс
        </b-button>
      </div>
      <h2 id="h2" class="mt-4">
        Ваши рекламные кампании
      </h2>

         <p>Здесь вы можете ознакомиться со своими рекламными кампаниями.</p>

      <div>
        <div>
          <div
            v-for="company in store.getters.GET_COMPANY_LIST"
            :key="company.Id"
          >
            <router-link :to="{path: '/company-balance/'+ company.Id, query: { isEdit: true }}">
              <div class="c-div">
                <div class="l">
                  <p class="c-date">{{company.Date}}</p>
                  <p class="c-name">
                    {{ company.CompanyName }}
                  </p>
                </div>
                <div class="r">
                  <div class="c-status">
                    <div
                      v-if="!isFb(company)"
                      id="white"
                      class="elipse"
                    />
                    <div
                      v-if="isFb(company) && !isMoney()"
                      id="yellow"
                      class="elipse"
                    />
                    <div
                      v-if="isFb(company) && isMoney()"
                      id="green"
                      class="elipse"
                    />
                    <p class="c-status-text">
                      {{ getStatus(company) }}
                    </p>
                  </div>
                </div>
              </div>
            </router-link>
          </div>
        </div>
      </div>
    </div>

    <div v-if="!(store.getters.GET_COMPANY_LIST.length > 0)">
      <h1 id="h1">
        Добро<br>пожаловать!
      </h1>
      <p id="p1">
        Для привлечения новых клиентов осталось<br>совсем немного. Создайте свою рекламную кампанию<br>или посмотрите видео-инструкцию, если возникли вопросы :)
      </p>
    </div>

    <div class="mt-5">
      <router-link
        v-if="store.getters.GET_COMPANY_LIST.length < 3"
        :to="{path: '/company'}"
      >
        <b-button
          class="app-new-submit-button"
        >
          Запустить рекламу
        </b-button>
      </router-link>
      <b-button
        variant="primary"
        class="main-button-grey ml-0 ml-sm-2 ml-lg-2 ml-md-2 mt-lg-0 mt-md-0 mt-sm-0 mt-2"
        @click="showVideo"
      >
        <b-icon icon="play-circle-fill"/>
         Инструкция
      </b-button>
    </div>

    <div v-if="isVideo">
      <iframe
        id="ytplayer"
        type="text/html"
        src="https://www.youtube.com/embed/EgpUx9_4ZIQ?autoplay=1&origin=https://youtu.be/EgpUx9_4ZIQ"
        frameborder="0"
      />
      <div
        id="ytcloser"
      >
        <b-icon
          class="x-button"
          icon="x"
          @click="closeVideo"
        />
      </div>
    </div>
  </div>
</template>
<script>
import store from '@/store/store'
import {instance as axios} from '@/_services/index';
import router from '@/router/router'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL // TODO GLOABL APP CONST
export default {
  name: 'MainPage',
  data () {
    return {
      user: {},
      isVideo: false,
      store
    }
  },
  watch: {
    $route (to) {
      this.isVideo = false
      store.dispatch('getCompanyList')
      console.log('TO', to)
      if (to.name === 'mainPage') {
        axios.get({ url: `${VUE_APP_API_URL}/api/user/0`})
          .then(resp => {
            this.user = resp.data
          })
      }
    }
  },
  beforeMount () {
    axios.get({ url: `${VUE_APP_API_URL}/api/user/0`})
      .then(resp => {
        this.user = resp.data
      })
  },
  mounted () {
    this.getAdCompanyList()
  },
  methods: {
    showVideo () {
      this.isVideo = true
    },
    closeVideo () {
      this.isVideo = false
    },
    push () {
      router.push({ path: '/company-balance/' + store.getters.GET_COMPANY_LIST[0].Id, query: { isEdit: true } })
    },
    isMoney () {
      return this.user.Amount > 0
    },
    isFb (company) {
      return company.FbPageId.length > 0
    },
    getStatus (company) {
      // if (company.FbPageId.length === 0) {
      //     return "facebook не подключен"
      // }
      if (this.user.Amount === 0) {
        return 'закончился рекламный бюджет'
      }
      if (!company.IsStarted) {
        return 'настраивается'
      }
      return 'работает'
    },
    getAdCompanyList () {
      store.dispatch('getCompanyList')
    }
  }
}
</script>
<style lang="scss">
  @import '@/assets/styles/vars.scss';

.campagin-item {
  & .date {
    font-size: 14px;
    line-height: 18px;
    color: $gray;
  }
  & .link {
    font-size: $baseFont;
    line-height: $baseLH;
  }

  & .status {
    font-size: 14px;
    line-height: 18px;
    color: $white;
    border-radius: 6px;
  }
}
#ytplayer{
    margin-top: 30px;
    width: 640px;
    height: 360px;
}
#ytcloser{
    position: absolute;
    margin-left: 620px;
    margin-top: -385px;
}
.c-date {
    font-size: 14px;
    line-height: 18px;
    color: $gray;
}
.c-div{
    cursor: pointer;
    margin-top: 40px;
    background: $light;
    border-radius: 20px;
    height: 112px;
    width: 100%;
    overflow: hidden;
    position: relative;
}
.l {
  position: absolute;
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
    #ytcloser{
        position: absolute;
        margin-left: 270px;
        margin-top: -195px;
    }
    #ytplayer{
        width: 290px;
        height: 170px;
    }
    .c-name{
        max-width: 260px; text-overflow: ellipsis; overflow: hidden; white-space: nowrap;
    }
    .c-status{
       height: auto;
       padding-bottom: 4px;
       display: inline-block;
    padding: 0;
    }
    .c-status-text{
           margin-left: 25px;
    margin-right: 10px;
    margin-bottom: 3px;
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
