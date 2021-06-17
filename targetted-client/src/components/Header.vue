<template>
    <div id="container">
        <div class="nav" id="left">
          <popup
              v-if="isInfoPopupVisible"
              popupTitle="Напишите нам"
              @closePopup="closeInfoPopup"
          >
            <div>
              <p id="p1">Свяжитесь с нами, если у вас возникли вопросы или предложения :)</p>
              <div class="popup-footer">
                <p>Whatsup</p>
                <p>Telegram</p>
              </div>
            </div>
          </popup>
        <div v-if="isLoggedIn" @click="showPopupInfo" id="write-us-text">Написать нам</div>
        </div>
        <div class="logo" id="middle">
          <img id="l-im" src="logo.svg">
        </div>
        <div class="nav" id="right">
            <div id="r-text" @click="logout" v-if="isLoggedIn">Выход</div>
        </div>
    </div>
</template>
<script>
import popup from './popup.vue'
import store from '../../store/store'
  export default {
    components: {
      popup
    },
    data() {
      return {
        isInfoPopupVisible: false
      }
    },
    computed : {
      isLoggedIn : function(){ return store.getters.isLoggedIn}
    },
    methods: {
      closeInfoPopup() {
        this.isInfoPopupVisible = false;
      },
      showPopupInfo() {
        this.isInfoPopupVisible = true;
      },
      logout: function () {
        store.dispatch('logout')
        .then(() => {
          this.$router.push('/login')
        })
      }
    },
  }
</script>
<style>
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