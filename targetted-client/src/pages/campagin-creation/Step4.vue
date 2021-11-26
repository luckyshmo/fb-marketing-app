<template>
   <div>
        <slot name="header"></slot>
        <b-form>
        <p>
            Привяжите страницу Facebook к&nbsp;targetted, чтобы настроить и&nbsp;запустить рекламную кампанию от имени вашего бизнеса.
        </p>

        <b-button type="button"
                 v-if="!isConnectStarted"
                class="app-new-submit-button"
                @click="startConnecting">
        Начать привязку
      </b-button>
       <b-button type="button"
                class="main-button-grey ml-0 ml-sm-2 ml-lg-2 ml-md-2 mt-lg-0 mt-md-0 mt-sm-0 mt-2"
                @click="sendData">
        {{isEdit ? "Назад":"Привязать позже"}}
      </b-button>

        <b-row align-h="between" no-gutters v-if="isConnectStarted">
          <h2>Наличие бизнес-аккаунта Instagram</h2>
          <p>Укажите, есть ли у вас бизнес-аккаунт в Instagram.</p>
            <b-col cols="6">
                <b-button v-if="!(store.getters.GET_FB_PAGES.length > 0)"    class="app-new-submit-button less-padding"
                  @click="loginFB">
                Есть аккаунт
              </b-button>
            </b-col>
            <b-col cols="6">
                <b-button v-if="!(store.getters.GET_FB_PAGES.length > 0)"    class="app-new-submit-button less-padding"
                @click="showPopupInfo">
                Нет аккаунта
              </b-button>
            </b-col>
        </b-row>

         <popup v-if="isInfoPopupVisible" popupTitle="Инструкция по созданию бизнесс-аккаунта"
          @closePopup="closeInfoPopup">
          <div style="text-align: left; padding: 30px">
            <h1 style="margin=0px !important;">Как создать<br>бизнес-аккаунт</h1>
            <p>Профессиональный аккаунт в Instagram необходим для того,<br>чтобы мы могли запустить и управлять рекламными
              кампаниями <br>от имени вашего бизнеса. </p>
            <div class="num-wrapper">
              <div class="num">
                <div class="num-num">
                  1
                </div>
              </div>
              <p class="num-text">Откройте мобильное приложение Инстаграм.</p>
            </div>
            <div class="num-wrapper">
              <div class="num">
                <div class="num-num">
                  2
                </div>
              </div>
              <p class="num-text">Зайдите в настройки профиля. </p>
            </div>
            <div class="num-wrapper">
              <div class="num">
                <div class="num-num">
                  3
                </div>
              </div>
              <p class="num-text">Нажмите на вкладку “Переключиться на профессиональный аккаунт”.</p>
            </div>
            <div class="num-wrapper">
              <div class="num">
                <div class="num-num">
                  4
                </div>
              </div>
              <p class="num-text">Дайте разрешение приложению Инстаграма управлять вашими Фейсбук-страницами. Не вашим
                личным аккаунтом, а именно страницами в Фейсбуке, к которым вы имеете доступ. Вы можете выбрать или
                существующую страницу в Facebook, к которой у вас есть доступ, или создать новую. </p>
            </div>
            <div class="num-wrapper">
              <div class="num">
                <div class="num-num">
                  5
                </div>
              </div>
              <p class="num-text">Настройка контактной информации для профиля Инстаграма. Укажите в соответствующие поля
                рабочий email компании, номер телефона и физический адрес (если есть).</p>
            </div>
            <div class="num-wrapper">
              <div class="num">
                <div class="num-num">
                  6
                </div>
              </div>
              <p class="num-text">Когда вы прошли все шаги и включили профессиональный аккаунт, то, пожалуйста, вернитесь
                на сайт client.targetted.online и нажмите кнопку “у меня есть профессиональный аккаунт” и следуйте
                дальнейшим инструкциям на сайте.</p>
            </div>
          </div>
        </popup>

        <div class="mt-5 mb-5" v-if="isConnectStarted">
          <h2>Возникли сложности?</h2>
          <p>Посмотрите видео-инструкцию или пропустите этот шаг и мы поможем вам привязать аккаунт позже.</p>
           <b-button
        variant="primary"
        class="main-button-grey"
        @click="showHelpVideo"
      >
        <b-icon icon="play-circle-fill"/>
         Инструкция
      </b-button>
        </div>

    </b-form>
  </div>
</template>

<script>
import store from '@src/store/store'
import accountService from '@src/_services/account.service'
import popup from '@src/components/BigPopup.vue'

export default {
  name: 'Step4',
  props: ['label_cols', 'content_cols', 'campaign', 'isEdit'],
  components: {
    popup
  },
  data: function () {
    return {
      store, // fixme
      isInfoPopupVisible: false,
      isConnectStarted: false
    }
  },
  methods: {
    startConnecting () {
      this.isConnectStarted = true
    },
    showHelpVideo () {

    },
    loginFB () {
      accountService.login()
    },
    sendData () {
      // todo
      // вопрос, какие тут данные отправляем?
      this.$emit('next')
    },
    showPopupInfo () {
      this.isInfoPopupVisible = true
    },
    closeInfoPopup () {
      this.isInfoPopupVisible = false
    }
  }
}
</script>

<style lang="scss">
  @import '@src/assets/styles/vars.scss';

  .num-wrapper{
      display: flex;
      align-items: stretch;
      margin-bottom: 16px;
  }
  .num{
      background: $light;
      width: 48px;
      height: 48px;
      border-radius: 24px;
  }

  .num-num{
      padding-top: 11px;
      padding-left: 19px;
  }

  .num-text{
      margin-left: 16px;
      max-width: 600px;
  }

</style>
