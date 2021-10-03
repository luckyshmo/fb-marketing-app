<template>
  <div class="container">
    <a
      href="https://targetted.ru/"
      target="_blank"
      class="logo"
    >
        <img
          id="l-im"
          src="@/assets/logo-new.png"
        >
        <div class="logo-bg"></div>
    </a>
    <div
      id="left"
      class="nav"
    >
      <popup
        v-if="isInfoPopupVisible"
        @closePopup="closeInfoPopup"
      >
        <div class="p-wrapper">
          <h1>Напишите нам</h1>
          <p
            id="p1"
            style="margin-top: 20px"
          >
            Свяжитесь с нами, если у вас возникли вопросы или проблемы
          </p>
          <div>
            <a
              id="social-link"
              href="https://wa.me/79952335503"
              target="_blank"
              rel="noopener noreferrer"
            >
              Whatsapp
            </a>
            <a
              id="social-link"
              href="http://t.me/targetted"
              target="_blank"
              rel="noopener noreferrer"
            >
              Telegram
            </a>
          </div>
        </div>
      </popup>
      <div
        v-if="isLoggedIn"
        id="write-us-text"
        @click="showPopupInfo"
      >
        Написать нам
      </div>
    </div>

    <div
      id="right"
      class="nav"
    >
     <button
        variant="primary"
        class="app-new-button-sm"
        @click="$router.push('/questions')"
      >
        <img src="@/assets/q-icon.svg"/>
      </button>
      &nbsp;
      <button
         variant="primary"
        class="app-new-button-sm d-none d-md-block d-lg-block d-xl-block"
        @click="register"
      >
        Зарегистрироваться
      </button>
      &nbsp;
      <button
        variant="primary"
        class="app-new-button-sm"
        v-if="isLoggedIn"
        @click="logout"
      >
        Выход
      </button>
      &nbsp;
       <button
         variant="primary"
        class="app-new-button-sm"
        v-if="!isLoggedIn"
        @click="isLoggedIn = true"
      >
        Вход
      </button>
    </div>
  </div>
</template>
<script>
import popup from '@/components/popup.vue'
import store from '@/store/store'
export default {
  components: {
    popup
  },
  data () {
    return {
      isInfoPopupVisible: false
    }
  },
  computed: {
    isLoggedIn: function () { return store.getters.isLoggedIn }
  },
  methods: {
    closeInfoPopup () {
      this.isInfoPopupVisible = false
    },
    showPopupInfo () {
      this.isInfoPopupVisible = true
    },
    register(){
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
  @import '@/assets/styles/vars.scss';

#social-link{
  margin: 30px;
  color: #6C1BD2
}
.p-wrapper{
  padding: 80px 100px 80px 100px
}
.logo {
  margin-left: 6%;
  margin-top: 4px;
}
@media (min-width: 768px){
  .logo {
    margin-left: 0;
  }
}
.logo-bg {
    background: #E2FF12;
    filter: blur(9px);
    position: absolute;
    width: 100px;
    height: 20px;
    z-index: -1;
    top: 28px;
}
@media (max-width: 465px){
  #l-im{
    /* width: 153px; */
    padding-bottom: 15px;
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
  max-width: 1220px;
  display: flex;
	flex-direction: row;
	flex-wrap: nowrap;
	justify-content: space-between;
	align-items: stretch;
	align-content: stretch;
  margin: 0 auto;
  margin-top: 20px;
}
#left, #middle {display: inline-block;}
#left {width: 25%; padding-top: 40px;}
#middle {width: 50%;}
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

</style>
