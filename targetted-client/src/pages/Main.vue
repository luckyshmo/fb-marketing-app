<template>
  <section class="main">
    <template v-if="isVideo" class="app-modal-white">
      <div class="video-pp-wrapper">
        <div>
          <iframe
            id="ytplayer"
            type="text/html"
            src="https://www.youtube.com/embed/EgpUx9_4ZIQ?autoplay=1&origin=https://youtu.be/EgpUx9_4ZIQ"/>
          <div
            id="ytcloser">
            <b-icon
              class="x-button"
              icon="x"
              @click="closeVideo"/>
          </div>
        </div>
      </div>
    </template>
    <template v-if="store.getters.GET_CAMPAIGN_LIST.length > 0">
      <h1>
        Александр, добро пожаловать!
      </h1>
      <p>
        Здесь вы можете ознакомиться со&nbsp;своими рекламными кампаниями.
      </p>
      <div>
        <div
          v-for="campaign in store.getters.GET_CAMPAIGN_LIST"
          :key="campaign.Id"
          class="campagin-item">
          <router-link :to="{path: '/campaign-balance/'+ campaign.Id, query: { isEdit: true }}">
            <b-row>
              <b-col cols="10">
                <p class="c-date">Создана {{campaign.Date}}</p>
                <p class="c-name">
                  {{ campaign.CampaignName }}
                </p>
                <div class="c-status">
                  <div
                    v-if="!isFb(campaign)"
                    id="white"
                    class="elipse" />
                  <div
                    v-if="isFb(campaign) && !isMoney()"
                    id="yellow"
                    class="elipse" />
                  <div
                    v-if="isFb(campaign) && isMoney()"
                    id="green"
                    class="elipse" />
                  <p class="c-status-text">
                    {{ getStatus(campaign) }}
                  </p>
                </div>
              </b-col>
              <b-col cols="2" class="right-pointer">
                <b-icon icon="chevron-right" @click="openCampaignDetails"/>
              </b-col>
            </b-row>
          </router-link>
        </div>
      </div>
    </template>
    <template v-if="!(store.getters.GET_CAMPAIGN_LIST.length > 0)">
      <h1>
        Добро<br>пожаловать!
      </h1>
      <p>
        Для привлечения новых клиентов осталось<br>совсем немного. Создайте свою рекламную кампанию<br>или посмотрите видео-инструкцию, если возникли вопросы :)
      </p>
    </template>
    <div class="mt-4">
      <router-link
        v-if="store.getters.GET_CAMPAIGN_LIST.length < 3"
        :to="{path: '/campaign-step-1'}">
        <b-button
          class="app-new-submit-button">
          Создать новую
        </b-button>
      </router-link>
      <br class="d-block d-sm-none d-md-none d-lg-none">
      <b-button
        variant="primary"
        class="main-button-grey ml-0 ml-sm-2 ml-lg-2 ml-md-2 mt-lg-0 mt-md-0 mt-sm-0 mt-2"
        @click="showVideo">
        <b-icon font-scale=".98" class="app-mid-size" icon="play-circle-fill"/>
        Инструкция
      </b-button>
    </div>
  </section>
</template>
<script>
import store from '@src/store/store'
import { instance as axios } from '@src/_services/index'
import router from '@src/router/router'

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
      store.dispatch('getCampaignList')
      console.log('TO', to)
      if (to.name === 'mainPage') {
        axios.get({ url: `${VUE_APP_API_URL}/api/user/0` })
          .then(resp => {
            this.user = resp.data
          })
      }
    }
  },
  beforeMount () {
    axios.get({ url: `${VUE_APP_API_URL}/api/user/0` })
      .then(resp => {
        this.user = resp.data
      })
  },
  mounted () {
    this.getAdCampaignList()
  },
  methods: {
    openCampaignDetails () {

    },
    showVideo () {
      this.isVideo = true
    },
    closeVideo () {
      this.isVideo = false
    },
    push () {
      router.push({ path: '/campaign-balance/' + store.getters.GET_CAMPAIGN_LIST[0].Id, query: { isEdit: true } })
    },
    isMoney () {
      return true
      // return this.user.Amount > 0
    },
    isFb (campaign) {
      return campaign.FbPageId.length > 0
    },
    getStatus (campaign) {
      // if (campaign.FbPageId.length === 0) {
      //     return "facebook не подключен"
      // }
      if (this.user.Amount === 0) {
        return 'ожидает оплаты'
      }
      if (!campaign.IsStarted) {
        return 'черновик'
      }
      return 'в работе'
    },
    getAdCampaignList () {
      store.dispatch('getCampaignList')
    }
  }
}
</script>
<style lang="scss">
  @import '@src/assets/styles/vars.scss';

  .main {
    margin: 82px 0 45px;
  }

  .video-pp-wrapper {
    background: rgba(0, 0, 0, 0.5);
    z-index: 1000;
    position: absolute;
    width: 100%;
    left: 0;
    top: 0;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .video-pp-wrapper > div{
    display: flex;
  }
  .right-pointer {
    justify-content: flex-end;
    display: flex;
    align-items: center;
    color: #000;
  }

  .campagin-item {
    background: $light;
    border-radius: 16px;
    padding: 16px;
    margin-bottom: 22px;

   &:hover {
     background: #EAEAEA;
   }
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
    width: 640px;
    height: 360px;
    border-radius: 12px;
  }

  #ytcloser{
    margin-left: 16px;
    & > svg{
      background-color: transparent;
      width: 24px;
      height: 24px;
      fill: $white;
    }
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
    right: 0;
    top: 0;
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
    padding: 0 !important;
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
    margin-top: 12px;
    margin-left: 12px;
    border-radius: 5px;
  }

  .c-status{
    text-align: center;
    // padding-top: 13px;
    // padding-left: 10px;
    height: 32px;
    padding-bottom: 1px;
    background: $black;
    border-radius: 6px;
    display: inline-block;
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

  @media (min-width: 600px) {
    .campagin-item {
      max-width: 50%;
    }
  }

  @media (max-width: 600px) {
    .video-pp-wrapper{
      width: 100%;
      background: $black;
      & > div{
        width: 100%;
      }
    }

    #ytcloser {
      position: absolute;
      top: 24px;
      right: 24px;
    }

    #ytplayer {
      width: 100%;
      border-radius: 0;
    }

    .c-name {
      max-width: 260px;
      text-overflow: ellipsis;
      overflow: hidden;
      white-space: nowrap;
      margin-top: 3px;
      margin-bottom: 7px;
    }

    .c-status{
      display: inline-block;
      padding: 0;
    }

    .c-status-text{
      margin-left: 27px;
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
        height: auto;
        margin: 16px;
    }
  }
</style>
