<template>
    <div id="app">
      <Header/>
      <div class="main">
        <Main/>
      </div>
      <Footer/>
    </div>
</template>

<script>
import Main from './components/Main.vue'
import Footer from './components/Footer.vue'
import Header from './components/Header.vue'

export default {
  name: 'App',
  components: {
    Main,
    Footer,
    Header
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;600;900&display=swap');
@import url('http://fonts.cdnfonts.com/css/monument-extended');

.main {
  margin: 0px 50px 0px 50px !important; 
  background-color: whitesmoke !important;
  border-radius: 25px !important;
}
#app {
  /* -webkit-font-smoothing: antialiased !important;
  -moz-osx-font-smoothing: grayscale !important; */
  /* text-align: center; */
  color: aquamarine !important;
  background-color: black !important;
}
html { 
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

#content {
    padding: 40px;
}

#h1 {
  margin-top: 0px;
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  font-size: 66px;
  line-height: 100%;

  color: #000000;
}
#p1 {
  font-family: Montserrat;
  font-style: normal;
  font-weight: normal;
  font-size: 20px;
  line-height: 32px;

  color: #000000;

}

#h2 {
    margin-top: 50px;
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  font-size: 36px;
  line-height: 44px;

  color: #000000;
}

</style>
