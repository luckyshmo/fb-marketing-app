<template>
<section>
   <popup
        v-if="isInfoPopupVisible"
        @closePopup="closeInfoPopup"
      >
 
       <Questions/>
      </popup>

 <b-row id="content-header">
        <b-col cols="3" class="app-new-nav-left">
      <a
      href="https://targetted.ru/"
      class="logo">
        <img
          id="l-im"
          src="@/assets/logo-new.png"
        >
        <div class="logo-bg"></div>
    </a>

        </b-col>
       <b-col cols="9" class="app-new-nav-right">

     
     <button
        variant="primary"
        class="app-new-button-sm"
        @click="showPopupInfo"
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
        @click="$router.push('/login')"
      >
        Войти
      </button>

       </b-col>
 </b-row>
  
  </section>
  <!-- <div class="container">
    
    <div
      id="left"
      class="nav"
    >
     
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
     
    </div>
  </div> -->
</template>
<script>
import popup from '@/components/popup.vue'
import store from '@/store/store'
import Questions from '@/pages/Questions.vue'
import Login from '@/pages/login/Login.vue'

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
    showPopupComponent(name){
      this.isInfoPopupVisible = true
      this.popupComponent = name
    },
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

#content-header {
  max-width: 1060px;
  padding: 24px 24px 24px 0;
  margin: 0% auto 0% auto !important;
  border-radius: 25px !important;
}

@media (max-width: 600px) {
    #content-header{
      padding: 22px 24px;
    }
}

#content-header {
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
    width: 120px;
    height: 20px;
    z-index: -1;
    top: 28px;
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
@media (min-width: 465px){
  #right {
        justify-content: space-between;
  }
}

</style>
