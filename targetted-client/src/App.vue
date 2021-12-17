<template>
  <div id="app">
    <section class="content">
      <Header />
      <keep-alive>
        <router-view />
      </keep-alive>
    </section>
    <Footer />
  </div>
</template>

<script>
import Footer from '@src/components/Footer.vue'
import Header from '@src/components/Header.vue'
import store from '@src/store/store'
import router from '@src/router/router'
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

@font-face {
    font-family: Monument Extended;
    font-weight: bold;
    src: url("./assets/MonumentExtended-Regular.otf") format("opentype");
}

/* Global styles */
html, body, #app {
  height: 100%;
  width: 100%;
}

#app {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  width: 1008px; /* visibleWidth(960) + 2 * px(24) */
  padding: 0 24px;
  margin: 0 auto;
}

@media (max-width: 1008px) {
  #app{
    position: relative;
  }
}

.btn.main-button-grey {
    background: #F3F3F3 !important;
    color: black !important;
    outline: none !important;
    border-radius: 6px;
    padding: 9px 24px 11px;
    border:none !important;
    font-family: 'IBM Plex Sans', sans-serif !important;
    font-style: normal !important;
    font-weight: 400 !important;
    line-height: 24px;
}

.main-button-big {
  color:white !important;
  outline: none !important;
  height: 56px;
  background-color: #6C1BD2 !important;
  border-radius: 8px !important;
  padding: 12px 28px !important;
  border:none !important;
  line-height: 32px;
  font-family: Montserrat, sans-serif !important;
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
  font-family: Montserrat , sans-serif !important;
  font-style: normal !important;
  font-weight: 600 !important;
  font-size: 16px;
}

.main-button-grey:hover {
  background-color: #EAEAEA !important;
}

.main-button:hover {
  background-color: #5101B5 !important;
}

.main-button-big:hover {
  background-color: #5101B5 !important;
}
@media (max-width: 600px) {
    #app {
      max-width: 600px;
    }

    .main-button {
      height: 48px;
      font-size: 16px;
    }
}

.submit-button {
  outline: none !important;
  position: absolute !important;
  margin-top: 138px; /* TODO dependce on content-wrapper */
  color: white !important;
  background-color: #FF62B7 !important;
  border-radius: 8px !important;
  padding: 12px 28px !important;
  height: 56px;
  border:none !important;
  font-family: Montserrat , sans-serif !important;
  font-style: normal !important;
  font-weight: 600 !important;
  line-height: 32px;
  font-size: 20px;
}

.submit-button:hover {
  background-color:#F248A4 !important;
}

#content-login {
  padding: 80px 40px 80px 40px;
  background-color: white !important;
  border-radius: 25px !important;
  max-width:640px;
  margin:0 auto;
}

#navigation-text {
    font-style: normal;
    font-weight: normal;
    font-size: 1em;
    max-width: 250px;
    margin: 32px auto 0 auto;
    color: #767676;
}

h1 {
  margin: 0;
  font-weight: bold !important;
  line-height: 56px !important;
  font-size: 60px !important;
  color: black;
}

@media (max-width: 600px) {

}

#h1-centered {
  font-family: IBM Plex Sans , sans-serif;
  font-style: normal;
  font-weight: 900;
  font-size: 72px;
  line-height: 72px;
  text-align: center;
}

p {
  font-style: normal;
  font-weight: normal;
  font-size: 16px;
  line-height: 24px;
  color: black;
  margin-bottom: 20px;
}

#h2-n {
  font-style: normal;
  font-weight: 900;
  /* font-size: 2.3em; */
  font-size: 36px;

  color: #000000;
}

#h2 {
  margin-top: 0;
  margin-bottom: 0;
  font-family: IBM Plex Sans , sans-serif;
  font-style: normal;
  font-weight: 900;
  /* font-size: 2.3em; */
  font-size: 36px;
  color: #000000;
}

.md-button{
  height: 24px;
  width: 24px;
  color: black;
  background: #f3f3f3;
  border-radius: 20px;
  cursor: pointer;
}

.x-button{
  height: 40px;
  width: 40px;
  color: black;
  background: #f3f3f3;
  border-radius: 20px;
  cursor: pointer;
}

.x-button:hover {
  background: #EAEAEA;
}

.form-input,
.form-select {
    height: 44px;
    border-radius: 6px !important;
    border: none !important;
    background-color: #F3F3F3 !important;
    color: black !important;
}

.form-select:focus,
.form-select:hover,
.form-input:focus,
.form-input:hover {
  background-color: #EAEAEA !important;
}

.form-input.width-1-2 {
  max-width: 45%;
}

@media (max-width: 450px) {
  .form-input.width-1-2 {
    max-width: 99%;
  }
}

::-webkit-scrollbar {
  width: 0;
}

@media (max-width: 600px) {
    .main-button {
      padding: 12px 10px !important;
    }

    .submit-button{
      margin-top: 75px;
      padding: 12px 10px !important;
    }
}
</style>
