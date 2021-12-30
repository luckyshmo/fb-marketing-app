<template>
<div class="page-anchor">
  <modal name="page-anchor-modal" classes="page-anchor-modal"
         :width="770"
         :height="406"
         :max-width="770"
         :adaptive="true"
         :min-width="300"
         :min-height="208"
         :max-height="406"
  >
    <div class="page-anchor-modal__block">
      <iframe
        type="text/html"
        src="https://www.youtube.com/embed/EgpUx9_4ZIQ?autoplay=1&origin=https://youtu.be/EgpUx9_4ZIQ"/>
      <div
        id="ytcloser">
        <b-icon
          class="x-button"
          icon="x"
          @click="closePageAnchorVideo()"/>
      </div>
    </div>
  </modal>
  <b-row class="app-new-steps-header">
    <b-col style="padding: 0; display: flex; align-items: center">
      <router-link id="app-link-back" :to="{path: '/'}" style="display: flex; align-items: center">

        <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
             x="0px" y="0px"
             viewBox="0 0 512 512" style="enable-background:new 0 0 512 512;" xml:space="preserve">
<path d="M3.919,243.077c-0.223,0.33-0.414,0.675-0.618,1.015c-0.186,0.31-0.382,0.614-0.552,0.936
	c-0.186,0.349-0.346,0.709-0.514,1.066c-0.157,0.332-0.321,0.658-0.464,0.998c-0.144,0.349-0.261,0.706-0.388,1.06
	c-0.129,0.362-0.268,0.718-0.38,1.089c-0.107,0.358-0.188,0.721-0.279,1.085c-0.093,0.374-0.199,0.743-0.275,1.125
	c-0.084,0.422-0.133,0.849-0.194,1.275c-0.047,0.326-0.109,0.645-0.143,0.976c-0.15,1.53-0.15,3.07,0,4.6
	c0.034,0.33,0.096,0.65,0.143,0.976c0.061,0.425,0.11,0.853,0.194,1.275c0.076,0.382,0.18,0.749,0.275,1.125
	c0.092,0.362,0.171,0.726,0.279,1.085c0.112,0.369,0.251,0.726,0.38,1.089c0.127,0.354,0.244,0.711,0.388,1.06
	c0.143,0.34,0.307,0.666,0.464,0.998c0.168,0.355,0.327,0.715,0.514,1.064c0.171,0.321,0.366,0.625,0.552,0.936
	c0.203,0.34,0.394,0.684,0.618,1.015c0.234,0.351,0.493,0.68,0.745,1.015c0.205,0.272,0.393,0.549,0.608,0.813
	c0.489,0.596,1.002,1.168,1.548,1.711l116.36,116.36c4.544,4.544,10.501,6.817,16.455,6.817c5.956,0,11.913-2.271,16.455-6.817
	c9.089-9.089,9.089-23.824,0-32.912l-76.636-76.636h409.272c12.853,0,23.273-10.42,23.273-23.273
	c0-12.853-10.42-23.273-23.273-23.273H79.456l76.636-76.636c9.089-9.089,9.089-23.824,0-32.912c-9.087-9.089-23.824-9.089-32.912,0
	L6.82,239.538c-0.546,0.543-1.06,1.116-1.548,1.711c-0.216,0.264-0.403,0.541-0.608,0.813
	C4.412,242.397,4.153,242.726,3.919,243.077z"/>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
</svg>

        На главную
      </router-link>
    </b-col>
  </b-row>
  <h2>Привязка страницы</h2>
  <p>Привяжите страницу Facebook, чтобы мы могли запустить рекламу от имени вашей компании. Это безопасно — мы не
    получим доступ к вашему логину, паролю, личному аккаунту и сообщениям.</p>
    <button class="app-new-submit-button" v-if="pageAnchorButton" @click="SelectPageShow()">Начать привязку</button>
    <div class="loggedIn" v-if="selectPage">
      <h2>Вы авторизовались</h2>
      <router-link to="/page-anchor">Изменить аккаунт</router-link>
    </div>
    <div class="select__page" v-if="selectPage">
      <h2>Выбор страницы</h2>
      <div v-if="selectingPageShow" class="select__page_block_one">
        <p>Выбрана страница {{selectPageValue}}</p>
        <button type="button" @click="ShowSelectingPage()">Выбрать другую страницу</button>
      </div>
      <div v-if="!selectingPageShow" class="select__page_block_two">
        <p>Выберите страницу, от имени которой будет показываться реклама:</p>
        <b-form-group
          id="input-group-main"
          style="margin-bottom: 8px !important;"
        >
          <b-form-radio-group
            :options="['Hostted.ru', 'Ногти Рукти' , 'Prof4nails' , 'Wresks']"
            v-model="selectPageValue"
            class="app-new-radio" />
        </b-form-group>
        <button class="app-new-submit-button" @click="ShowSelectingPage()">Выбрать</button>
      </div>
    </div>
    <div class="instagram__block" v-if="selectingPageShow">
      <h2>Привяжите Instagram к странице Facebook</h2>
      <p>Чтобы продолжить, выберите связанную страницу или привяжите Instagram-аккаунт к этой странице: </p>
      <p><span>1.</span>Зайдите в приложение Instagram.</p>
      <p><span>2.</span>Откройте свой профиль и коснитесь значка меню в верхнем правом углу.</p>
      <p><span>3.</span>Нажмите Настройки.</p>
      <p><span>4.</span>Нажмите Аккаунт.</p>
      <p><span>5.</span>Нажмите Преобразовать в бизнес-аккаунт.</p>
      <p><span>6.</span>Подключите бизнес-аккаунт к Странице Facebook, которая связана с вашей компанией. С вашим аккаунтом Instagram может быть связана только одна Страница Facebook.</p>
      <p><span>7.</span>Добавьте категорию компании и контактную информацию.</p>
      <p><span>8.</span>Нажмите Готово.</p>
      <button type="button" class="app-new-submit-button" @click="GetAccountlinked()">Аккаунт привязан</button>
      <p class="error-message" v-if="errorMess">Страница не привязана, попробуйте ещё раз</p>
    </div>
    <div class="video__block" v-if="selectPage">
      <h2>Возникли сложности?</h2>
      <p>Посмотрите видео-инструкцию о том, как создать и привязать бизнес-аккаунт.</p>
      <b-button
        variant="primary"
        class="main-button-grey"
        @click="showPageAnchorVideo()">
        <b-icon font-scale=".98" class="app-mid-size" icon="play-circle-fill"/>
        Посмотреть видео
      </b-button>
    </div>
</div>
</template>

<script>
export default {
  name: "PageAnchor",
  data() {
    return {
      pageAnchorButton : true,
      selectPage: false,
      pageAnchorVideo: false,
      selectPageValue : '',
      selectingPageShow : false,
      errorMess : false
    }
  },
  methods: {
    SelectPageShow() {
      this.pageAnchorButton = false
      this.selectPage = true
    },
    showPageAnchorVideo(){
      this.pageAnchorVideo = true
      this.$modal.show('page-anchor-modal')
    },
    closePageAnchorVideo(){
      this.pageAnchorVideo = false
      this.$modal.hide('page-anchor-modal')
    },
    ShowSelectingPage(){
      if (this.selectPageValue){
        this.selectingPageShow = !this.selectingPageShow
      }
    },
    GetAccountlinked(){
      this.errorMess = true
    }
  },
}
</script>

<style lang="scss">
@import '@src/assets/styles/vars.scss';
.page-anchor{
  max-width: 580px;
  width: 100%;
}
.page-anchor > h2{
  font-size: 56px !important;
  line-height: 60px !important;
  margin-top: -2px !important;
  margin-bottom: 12px !important;
}
.page-anchor > p{
  margin-bottom: 0 !important;
}
.page-anchor > button{
  margin-top: 20px;
}
.loggedIn{
  margin-top: 40px;
  margin-bottom: 0 !important;
}
.loggedIn > a{
  margin-top: 8px;
  font-size: 16px;
  line-height: 24px;
  color: $violet;
  &:hover{
    color: $violet;
  }
}
.select__page{
  margin-top: 40px;
  width: 100%;
  & > .select__page_block_two {
    & > p{
      margin-bottom: 12px;
    }
    & > button{
      margin-top: -7px;
    }
  }
}
.select__page_block_one{
  & > button{
    border: 0;
    background-color: transparent;
    color: $violet;
    font-style: normal;
    font-weight: normal;
    font-size: 16px;
    line-height: 24px;
    padding: 0;
  }
  & > p{
    margin: 0;
  }
}
.video__block{
  margin-top: 120px;
  margin-bottom: 195px;
  & > p{
    max-width: 385px;
  }
  & > button{
    margin-top: 4px;
    & > svg {
      margin-right: 8px;
    }
  }
}
.page-anchor-modal{
  background-color: transparent !important;
  box-shadow: none !important;
}
.page-anchor-modal__block{
  width: 100%;
  display: flex;
  justify-content: space-between;
  & > iframe{
    max-width: 729px;
    width: 100%;
    height: 406px;
    border-radius: 12px;
  }
}
.instagram__block{
  margin-top: 40px;
  & > p{
    margin-bottom: 12px;
    & > span{
      margin-right: 6px;
      height: 100%;
    }
  }
  & > button{
    margin-top: 8px ;
  }
  & > .error-message{
    margin-top: 8px;
  }
}

@media (max-width: 600px) {
  .page-anchor{
    .vm--container{
      background-color: $black;
    }
    #ytcloser{
      position: fixed;
      top: 24px;
      right: 24px;
    }
    .page-anchor-modal{
      height: 208px !important;
    }
    iframe{
      height: 208px;
    }
  }
  .page-anchor > h2{
    font-size: 32px !important;
    line-height: 38px !important;
  }
  .video__block{
    margin: 80px 0 75px;
  }
}
</style>
