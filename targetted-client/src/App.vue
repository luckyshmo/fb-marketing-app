<template>
  <div id="app">
    <Header />
    <div class="main">
      <keep-alive>
        <router-view />
      </keep-alive>
    </div>
    <Footer />
  </div>
</template>

<script>
import Footer from './components/Footer.vue'
import Header from './components/Header.vue'
import store from '../store/store'
import router from '../router/router'
export default {
  name: 'App',
  components: {
    Footer,
    Header
  },
  created: function () {
    this.$http.interceptors.response.use(undefined, function (err) {
      return new Promise(function () {
        if (err.status === 401 && err.config && !err.config.__isRetryRequest) {
          store.dispatch('logout')
        }
        throw err
      })
    })
  },
  methods: {
    logout: function () {
      store.dispatch('logout')
        .then(() => {
          router.push('/login')
        })
    }
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;600;900&display=swap');
@import url('https://fonts.googleapis.com/css2?family=IBM+Plex+Sans:wght@400;700&display=swap');
/* @import url('') format("opentype"); */
@font-face {
    font-family: Monument Extended;
    font-weight: bold;
    src: url("./assets/MonumentExtended-Regular.otf") format("opentype");
}

#app {
   /* -webkit-font-smoothing: antialiased !important;
  -moz-osx-font-smoothing: grayscale !important; */
  /* text-align: center; */
  overflow: hidden;

}
html {
  overflow-y: auto;
  height: 100% !important;
  width: 100% !important;
}
body {
  background: #fff;
}
/* Global styles */

.main-button-grey{
    background: #F3F3F3 !important;
    color: black !important;
    outline: none !important;
    border-radius: 8px !important;
    padding: 12px 28px !important;
    border:none !important;
    font-family: 'IBM Plex Sans', sans-serif !important;
    font-style: normal !important;
    font-weight: 400 !important;
    line-height: 24px;
    height: 48px;
}

.main-button-big{
  color:white !important;
  outline: none !important;
  height: 56px;
  background-color: #6C1BD2 !important;
  border-radius: 8px !important;
  padding: 12px 28px !important;
  border:none !important;

  line-height: 32px;
  font-family: Montserrat !important;
  font-style: normal !important;
  font-weight: 600 !important;
  font-size: 20px;
}

.main-button {
  color:white !important;
  outline: none !important;
  height: 48px;
  background-color: #6C1BD2 !important;
  border-radius: 8px !important;
  padding: 12px 28px !important;
  border:none !important;

  line-height: 24px;
  font-family: Montserrat !important;
  font-style: normal !important;
  font-weight: 600 !important;
  font-size: 16px;
}
.main-button-grey:hover{
  background-color: #EEEEEE !important;
}
.main-button:hover {
  background-color: #5101B5 !important;
}
.main-button-big:hover {
  background-color: #5101B5 !important;
}
@media (max-width: 600px) {
    .main-button {
      height: 48px;
      font-size: 16px;
    }
}

.submit-button{
  outline: none !important;
  position: absolute !important;
  margin-top: 138px; /* //TODO dependce on content-wrapper */
  color:white !important;
  background-color: #FF62B7 !important;
  border-radius: 8px !important;
  padding: 12px 28px !important;
  height: 56px;
  border:none !important;
  font-family: Montserrat !important;
  font-style: normal !important;
  font-weight: 600 !important;
  line-height: 32px;
  font-size: 20px;
}
.submit-button:hover{
  background-color:#F248A4 !important;
}

#content {
  padding: 80px;
  width: 1220px;
  margin: 0% auto 0% auto !important;
  background-color: white !important;
  border-radius: 25px !important;
}
#content-wrapper {
  width: 1220px;
  margin: 0% auto 0% auto !important;
  border-radius: 25px !important;
  padding-bottom: 45px;
}
@media (max-width: 600px) {
  #content{
    width: 375px;
  }
  #content-wrapper{
    width: 375px;
  }
}
#content-login {
  padding: 80px 40px 80px 40px;
  margin: 0px 50px 0px 50px !important;
  background-color: white !important;
  border-radius: 25px !important;
  max-width:640px;
  margin:0 auto !important;
}

#navigation-text {
    font-family: Montserrat;
    font-style: normal;
    font-weight: normal;
    font-size: 1em;
    max-width: 250px;
    margin: 32px auto 0px auto;
    color: #767676;
}

h1 {
  margin: 0px;
  font-family: IBM Plex Sans !important;
  font-weight: 900 !important;
  line-height: 72px !important;
  font-size: 72px !important;
  color: black;
}
#h1-centered {
  font-family: IBM Plex Sans;
  font-style: normal;
  font-weight: 900;
  font-size: 72px;
  line-height: 72px;
  text-align: center;
}
p {
  font-family: Montserrat;
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  line-height: 24px;
  color: black;
}

#h2-n {
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  /* font-size: 2.3em; */
  font-size: 36px;

  color: #000000;
}
#h2 {
  margin-top: 0px;
  margin-bottom: 0px;
  font-family: IBM Plex Sans;
  font-style: normal;
  font-weight: 900;
  /* font-size: 2.3em; */
  font-size: 36px;

  color: #000000;
}

#form-centred{
    text-align: center !important;
}
#input-group-main{
  margin-top: 32px;
  color: black;
  display: flex;
  font-family: Montserrat;
  font-style: normal;
  font-weight: normal;
  font-weight: 400;
  font-size: 1em;
}
.x-button{
  height: 40px;
  width: 40px;
  color: black;
  background: #f3f3f3;
  border-radius: 20px;
  cursor: pointer;
}
@media (max-width: 600px) {
    #input-group-main{
        display:grid;
    }
    #h1-centered {
      font-size: 36px;
      line-height: 44px;
    }
    h1 {
      font-size: 32px !important;
      line-height: 44px !important;
    }
    #h2 {
      font-size: 24px;
      line-height: 32px;
      margin-top: 0px;
      margin-bottom: 0px;
    }
}
.form-input{
    height: 48px;
    border-radius: 8px !important;
    border: none !important;
    background-color: #F3F3F3 !important;
    color: black !important;
}

::-webkit-scrollbar {
  width: 0px;
}

@media (max-width: 600px) {
    .main-button {
      padding: 12px 10px !important;
    }
    .submit-button{
      margin-top: 75px;
      padding: 12px 10px !important;
    }
    #content{
      padding: 40px;
    }
}
</style>
