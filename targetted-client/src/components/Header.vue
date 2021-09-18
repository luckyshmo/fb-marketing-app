<template>
  <div id="container">
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
    <a
      href="https://targetted.ru/"
      target="_blank"
    >
      <div
        id="middle"
        class="logo"
      >
        <img
          id="l-im"
          src="logo.svg"
        >
      </div>
    </a>
    <div
      id="right"
      class="nav"
    >
      <div
        v-if="isLoggedIn"
        id="r-text"
        @click="logout"
      >
        Выход
      </div>
    </div>
  </div>
</template>
<script>
import popup from '@/components/popup.vue'
import store from '@/../store/store'
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
    logout: function () {
      store.dispatch('logout')
        .then(() => {
          this.$router.push('/login')
        })
    }
  }
}
</script>
<style>
#social-link{
  margin: 30px;
  color: #6C1BD2
}
.p-wrapper{
  padding: 80px 100px 80px 100px
}
.logo {
    font-family: Monument Extended;
    font-style: normal;
    font-weight: normal;
    font-size: 1.3em;
    line-height: 50px;

    letter-spacing: 0.8em;
    text-transform: uppercase;

    background: -webkit-linear-gradient(left, #0704FD, #F498AD, #BB44CB);
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;

    width: 50%;
}
@media (max-width: 465px){
  #l-im{
    width: 183px;
    padding-bottom: 15px;
  }
  #r-text{
    position: absolute;
    right:20px;
    top: 55px;
  }
  .p-wrapper{
    padding: 20px 25px 20px 25px
  }
}
.nav {
    cursor: pointer;
    font-family: Montserrat;
    font-style: normal;
    font-weight: normal;
    font-size: 1em;
    text-align: center;

    color:white !important;

    width: 25%;
}

#container {
  height: 100px;
  max-width: 1220px;
  text-align: center;
  margin: 0 auto
}
#left, #middle, #right {display: inline-block;}
#left {width: 25%; padding-top: 40px;}
#middle {width: 50%;}
#right {width: 25%;}

</style>
