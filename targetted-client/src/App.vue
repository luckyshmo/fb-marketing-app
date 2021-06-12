<template>
    <div id="app">
      <Header/>
      <div class="main">
        <keep-alive>
            <router-view></router-view>
        </keep-alive>
      </div>
      <Footer/>
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
  methods: {
    logout: function () {
      store.dispatch('logout')
      .then(() => {
        router.push('/login')
      })
    }
  },
  created: function () {
      this.$http.interceptors.response.use(undefined, function (err) {
        return new Promise(function () {
          if (err.status === 401 && err.config && !err.config.__isRetryRequest) {
            store.dispatch('logout')
          }
          throw err;
        });
      });
    }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;600;900&display=swap');
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
  color: red;
  background-color: black !important;
}
html { 
  overflow-y: auto;
  height: 100% !important;
  width: 100% !important;
  background-color: black !important;
}

/* Global styles */
#main-button {
    color:white;
     outline: none !important;
    background-color: #6C1BD2;
    border-radius: 8px;
    padding: 12px 28px;
    border:none;
    font-family: Montserrat;
    font-style: normal;
    font-weight: 600;
}

#submit-button{
   outline: none !important;
    position: absolute;
    margin-top: 115px; /* //TODO dependce on content-wrapper */
    color:white;
    background-color: #FF62B7;
    border-radius: 8px;
    padding: 12px 28px;
    border:none;
    font-family: Montserrat;
    font-style: normal;
    font-weight: 600;
}

#content {
  padding: 80px;
  max-width: 1220px;
  margin: 0% auto 0% auto !important; 
  background-color: white !important;
  border-radius: 25px !important;
}
#content-wrapper {
  max-width: 1220px;
  margin: 0% auto 0% auto !important; 
  background-color: #6C1BD2 !important;
  border-radius: 25px !important;
  padding-bottom: 120px;
}
#content-login {
  padding: 80px 40px 80px 40px;
  margin: 0px 50px 0px 50px !important; 
  background-color: white !important;
  border-radius: 25px !important;
  max-width:600px;
  margin:0 auto !important;
}

#navigation-text {
    font-family: Montserrat;
    font-style: normal;
    font-weight: normal;
    font-size: 1em;
    /* line-height: 24px; */
    max-width: 250px;
    margin: 32px auto 0px auto;
    color: #767676;
}

#h1 {
  margin: 0px;
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  /* font-size: 3.5em; */
  font-size: 72px;
  color: #000000;
}
#h1-centered {
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  font-size: 72px;
  /* font-size: 3.5em; */
  line-height: 100%;
  text-align: center;
  color: #000000;
}
#p1 {
  font-family: Montserrat;
  font-style: normal;
  font-weight: normal;
  /* font-size: 1.3em; */
  font-size: 16px;

  color: #000000;

}

#h2 {
  margin-top: 60px;
  margin-bottom: 20px;
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  /* font-size: 2.3em; */
  font-size: 36px;

  color: #000000;
}

#form-centred{
    text-align: center !important;
}
#input-group{
  margin: 32px 20px 32px 20px !important;
}
.form-input{
  max-width: 540px !important;
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
  height: 35px;
  width: 35px;
  color: black;
  background: #e4e4e4;
  border-radius: 17.5px;
  cursor: pointer;
}
@media (max-width: 600px) {
    #input-group-main{
        display:grid;
    }
    #h1-centered {
      font-size: 2.5em;
    }
    #h1 {
      font-size: 2.5em;
    }
}
.form-input{
    height: 48px;
    border-radius: 8px !important;
    border: none !important;
    background-color: #F3F3F3 !important;
    color: #767676 !important;
}

::-webkit-scrollbar {
  width: 0px;
}

</style>
