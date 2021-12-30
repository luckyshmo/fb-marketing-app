<template>
  <section class="header">
<!--    <popup-->
<!--      v-if="isInfoPopupVisible"-->
<!--      @closePopup="closeInfoPopup">-->
<!--      <Questions/>-->
<!--    </popup>-->
    <b-row align-h="between">
      <b-col cols="auto">
        <a href="https://targetted.ru/" style="display: block;
          width: 116px;
          height: 31px;
          margin-left: -7px;
          margin-top: 2px;
          ">
          <img class="header-logo" src="@src/assets/logo.png" alt="logo" style="width: 100% ; height: 100%; object-fit: cover"/>
        </a>
      </b-col>
      <b-col cols="auto">
        <div class="header-buttons">
          <button
            variant="primary"
            class="app-new-button-sm"
            @click="showPopupInfo">
            <img src="@src/assets/q-icon.svg" alt=""/>
          </button>
          &nbsp;
          <button
            variant="primary"
            class="app-new-button-sm d-none d-md-block d-lg-block d-xl-block align-items-start"
            @click="register">
            Зарегистрироваться
          </button>
          &nbsp;
          <button
            variant="primary"
            class="app-new-button-sm"
            v-if="isLoggedIn"
            @click="logout">
            Выход
          </button>
          &nbsp;
          <button
            variant="primary"
            class="app-new-button-sm"
            v-if="!isLoggedIn"
            @click="$router.push('/login')">
            Войти
          </button>
        </div>
      </b-col>
    </b-row>
  </section>
  <!-- <div class="container">
    <div
      id="left"
      class="nav">
      <div
        v-if="isLoggedIn"
        id="write-us-text"
        @click="showPopupInfo">
        Написать нам
      </div>
    </div>
    <div
      id="right"
      class="nav">
    </div>
  </div> -->
</template>
<script>
import popup from '@src/components/popup.vue'
import store from '@src/store/store'
import Questions from '@src/pages/Questions.vue'
import Login from '@src/pages/login/Login.vue'

export default {
  components: {
    popup,
    Questions,
    Login
  },
  data () {
    return {
      isInfoPopupVisible: false,
      popupComponent: undefined
    }
  },
  computed: {
    isLoggedIn: function () { return store.getters.isLoggedIn }
  },
  methods: {
    showPopupComponent (name) {
      this.isInfoPopupVisible = true
      this.popupComponent = name
    },
    closeInfoPopup () {
      this.isInfoPopupVisible = false
    },
    showPopupInfo () {
      this.isInfoPopupVisible = true
    },
    register () {
      this.$router.push('/register')
    },
    logout: function () {
      store.dispatch('logout')
        .then(() => {
          this.$router.push('/login')
        })
    }
  }
}
</script>

<style lang="scss">
  @import '@src/assets/styles/vars.scss';

  .header {
    padding: 27px 0;
  }

  .header-logo {
    height: 32px;
  }

  .header-buttons {
    display: flex;
  }

  @media (max-width: 600px) {
      .header {
        padding: 24px 0;
      }
  }

  .header {
    & .app-new-nav-right {
      padding: 0;
    }

    & .app-new-nav-left  {
       padding: 0;
    }

    & .app-new-nav-right {
      display: flex;
      justify-content: flex-end;
      button {
          margin-left: 10px;
      }
    }
  }

  #social-link{
    margin: 30px;
    color: #6C1BD2
  }

  .p-wrapper{
    padding: 80px 100px 80px 100px
  }

  @media (max-width: 465px){
    #l-im{
      padding-bottom: 15px;
      width: 135px;
    }
    .p-wrapper{
      padding: 20px 25px 20px 25px
    }
  }

  .nav {
      cursor: pointer;
      font-style: normal;
      font-weight: normal;
      font-size: 1em;
      text-align: center;
      color: $white;
      width: 25%;
  }

  .container {
    height: 50px;
    max-width: 1060px;
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    justify-content: space-between;
    align-items: stretch;
    align-content: stretch;
    margin: 20px auto 0;
  }

  #left, #middle {
    display: inline-block;
  }

  #left {
    width: 25%;
    padding-top: 40px;
  }

  #middle {
    width: 50%;
  }

  #right {
    width: 25%;
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    justify-content: flex-end;
    align-items: stretch;
    align-content: stretch;

    button {
      margin-right: 10px;
    }
  }

  @media (min-width: 465px){
    #right {
      justify-content: space-between;
    }
  }

</style>
