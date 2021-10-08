<template>
  <div id="content">

    <template v-if="isVideo" class="app-modal-white">
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
    </template>

    <template v-if="store.getters.GET_COMPANY_LIST.length > 0">
      <h1 id="h1">
        Александр, добро пожаловать!
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

        <!-- <b-button
          variant="primary"
          class="main-button-grey mt-2"
          @click="push"
        >
          Пополнить баланс
        </b-button> -->
      </div>

         <p>Здесь вы можете ознакомиться со&nbsp;своими рекламными кампаниями.</p>

      <div>
        <div>

<!-- 
 <b-row class="mt-5">
        <b-col cols="6" sm="8" md="3" lg="3" xl="3"> -->


          <div
            v-for="company in store.getters.GET_COMPANY_LIST"
            :key="company.Id"
          >
            <router-link :to="{path: '/company-balance/'+ company.Id, query: { isEdit: true }}">
              <div class="c-div">
                <div class="l">
                  <p class="c-date">Создана {{company.Date}}</p>
                  <p class="c-name">
                    {{ company.CompanyName }}
                  </p>
                    <b-icon
                      class="md-button app-block-icon"
                      icon="chevron-right"
                      @click="openCompanyDetails"
                    />
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
    </template>

    <template v-if="!(store.getters.GET_COMPANY_LIST.length > 0)">
      <h1 id="h1">
        Добро<br>пожаловать!
      </h1>
      <p id="p1">
        Для привлечения новых клиентов осталось<br>совсем немного. Создайте свою рекламную кампанию<br>или посмотрите видео-инструкцию, если возникли вопросы :)
      </p>
    </template>

    <div class="mt-5">
      <router-link
        v-if="store.getters.GET_COMPANY_LIST.length < 3"
        :to="{path: '/company'}"
      >
        <b-button
          class="app-new-submit-button"
        >
          Создать новую
        </b-button>
      </router-link>
      <br class="d-block d-sm-none d-md-none d-lg-none">
      <b-button
        variant="primary"
        class="main-button-grey ml-0 ml-sm-2 ml-lg-2 ml-md-2 mt-lg-0 mt-md-0 mt-sm-0 mt-2"
        @click="showVideo"
      >
        <b-icon font-scale=".98" class="app-mid-size" icon="play-circle-fill"/>
         Инструкция
      </b-button>
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
    openCompanyDetails(){

    },
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
      return true
      //return this.user.Amount > 0
    },
    isFb (company) {
      return company.FbPageId.length > 0
    },
    getStatus (company) {
      // if (company.FbPageId.length === 0) {
      //     return "facebook не подключен"
      // }
      if (this.user.Amount === 0) {
        return 'ожидает оплаты'
      }
      if (!company.IsStarted) {
        return 'черновик'
      }
      return 'в работе'
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
    margin-bottom: 0;
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
    font-style: normal;
    font-size: 14px;
    padding: 0px !important;
}
#white{
    background: white;
}
#yellow{
    background: $lime;
}
#green{
    background:#E2FF12;
}
.elipse{
    position: absolute;
    width: 8px;
    height: 8px;
    margin-top: 9px;
    margin-left: 9px;
    border-radius: 5px;
}
.c-status{
    text-align: center;
    padding-top: 13px;
    padding-left: 10px;
    height: 55px;
    padding-bottom: 1px;
    background: $black;
    border-radius: 6px;
}
.c-name{
    font-style: normal;
    font-weight: bold;
    font-size: 16px;
    line-height: 24px;
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
        margin-top: 14px;
        margin-left: 16px;
        height: auto;
    }
    .r {
        position: relative;
        right: auto;
        top: auto;
        margin-right: auto;
        margin-top: auto;
        height: auto;
        margin: 16px;

    }
}
</style>
