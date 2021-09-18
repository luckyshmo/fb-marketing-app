<template>
   <div>
        <slot name="header"></slot>
        <b-form>
        <p>
            Привяжите страницу Facebook к targetted, чтобы настроить и запустить рекламную кампанию от имени вашего бизнеса.
        </p>

        <b-row align-h="between" no-gutters>
            <b-col cols="6">
                <b-button v-if="!(store.getters.GET_FB_PAGES.length > 0)"    class="app-new-submit-button"
                  @click="loginFB">
                Есть аккаунт
              </b-button>
            </b-col>
            <b-col cols="6">
                <b-button v-if="!(store.getters.GET_FB_PAGES.length > 0)"    class="app-new-submit-button"
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

        <b-button type="button"
                class="app-new-submit-button"
                @click="sendData">
        {{isEdit ? "Назад":"Продолжить"}}
      </b-button>
    </b-form>
  </div>
</template>

<script>
import store from '@/../store/store'

export default {
  name: 'Step4',
    props: ['label_cols', 'content_cols', 'company', 'isEdit'],
    data: function () {
        return {
          store, //fixme
          isInfoPopupVisible: false,
        }
      },
    methods: {
        sendData(){
            //todo
            this.$emit('next');
        },
        showPopupInfo() {
          this.isInfoPopupVisible = true
        },
        closeInfoPopup() {
          this.isInfoPopupVisible = false
        },
    }
}
</script>

<style lang="scss">
  @import '@/../../../assets/vars.scss';
.num-wrapper{
    display: flex;
    align-items: stretch;
    margin-bottom: 16px;
}
.num{
    background: #F3F3F3;
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