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
.main {
  margin: 0px 50px 0px 50px !important; 
  background-color: whitesmoke !important;
  border-radius: 25px;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: aquamarine;
}
html { 
  height: 100% !important;
  width: 100% !important;
  background-color: black;
}
h1 {
  margin-top: 10px;
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  font-size: 66px;
  line-height: 100%;

  color: #000000;
}
p1 {
  font-family: Montserrat;
  font-style: normal;
  font-weight: normal;
  font-size: 20px;
  line-height: 32px;

  color: #000000;

}

h2 {
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  font-size: 36px;
  line-height: 44px;

  color: #000000;
}

.button-main {

  border-style: solid;
  border-color: #6C1BD2;

  /* Main */
  background-color: #6C1BD2 ;
  /* linear-gradient(0deg, #6C1BD2, #6C1BD2), linear-gradient(91.15deg, #0704FD -5.97%, #F498AD 49.57%, #BB44CB 116.59%); */
  border-radius: 8px;

  font-family: Montserrat;
  font-style: normal;
  font-weight: 600;
  font-size: 20px;
  line-height: 32px;
  text-align: center;
  color: #FFFFFF;
}

</style>
