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
    color:whitesmoke;
    background-color: #6C1BD2;
    border-radius: 8px;
    padding: 12px 28px;
    border:none;
    font-family: Montserrat;
    font-style: normal;
    font-weight: 600;
}

#submit-button{
    position: absolute;
    margin-top: 75px; /* //TODO dependce on content-wrapper */
    color:whitesmoke;
    /* margin-top: 15px; */
    background-color: #FF62B7;
    border-radius: 8px;
    padding: 12px 28px;
    border:none;
    font-family: Montserrat;
    font-style: normal;
    font-weight: 600;
}

#content {
  padding: 40px;
  max-width: 1200px;
  margin: 0% auto 0% auto !important; 
  background-color: whitesmoke !important;
  border-radius: 25px !important;
}
#content-wrapper {
  max-width: 1200px;
  margin: 0% auto 0% auto !important; 
  background-color: #6C1BD2 !important;
  border-radius: 25px !important;
  padding-bottom: 120px;
}
#content-login {
  padding: 40px;
  margin: 0px 50px 0px 50px !important; 
  background-color: whitesmoke !important;
  border-radius: 25px !important;
  width:600px;
  margin:0 auto !important;
}

#navigation-text {
    font-family: Montserrat;
    font-style: normal;
    font-weight: normal;
    margin: 10px;
    font-size: 1em;
    line-height: 24px;
    max-width: 190px;
    margin:0 auto;
    color: #767676;
}

#h1 {
  margin-top: 0px;
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  font-size: 3.5em;

  color: #000000;
}
#h1-centered {
  margin-top: 0px;
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  font-size: 3.5em;
  line-height: 100%;
  text-align: center;
  color: #000000;
}
#p1 {
  font-family: Montserrat;
  font-style: normal;
  font-weight: normal;
  font-size: 1.3em;

  color: #000000;

}

#h2 {
  margin-top: 50px;
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  font-size: 2.3em;

  color: #000000;
}

#form-centred{
    text-align: center !important;
}
#input-group{
  margin: 35px 20px 35px 20px;
}
#input-group1{
  margin-top: 10px;
  color: black;
  display: flex;
  font-family: Montserrat;
  font-style: normal;
  font-weight: normal;
  font-weight: 400;
  font-size: 1em;
}
#form-input{
    border-radius: 8px;
    border: none;
    background-color: #e4e4e4;
    color: #767676;
}

::-webkit-scrollbar {
  width: 0px;
}

/* Track */
::-webkit-scrollbar-track {
  background: #f1f1f1;
}

/* Handle */
::-webkit-scrollbar-thumb {
  background: #6C1BD2;
}

/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
  background: #FF62B7;
}

</style>
