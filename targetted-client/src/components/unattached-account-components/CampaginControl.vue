<template>
  <div class="campaign__management">
    <h2>{{campaginControlData.title}}</h2>
    <div v-for="val in campaginControlData.campaginControlValue"
         :key="val">
      <p :class="{'red__color' : val.colorRed, 'purple__color' : val.colorPurple }">
        {{val.title}}
        <svg viewBox="0 0 7 12" fill="none" xmlns="http://www.w3.org/2000/svg" v-if="val.svg">
          <path d="M1 1L5.5 6L1 11" stroke="black" stroke-width="2" stroke-linecap="round"/>
        </svg>
      </p>
      </div>
  </div>
</template>

<script>
export default {
  name: "CampaginControl",
  props: {
    campaginControlData: {
      type: Object,
      required: true
    }
  }
}
</script>

<style lang="scss">
@import '@src/assets/styles/vars.scss';
.campaign__management{
  margin: 40px 0 200px;
}
.campaign__management > div{
  padding: 0 16px;
  background-color: $light;
}
.campaign__management > div:nth-child(2){
  margin-top: 16px;
  border-top-left-radius: 16px;
  border-top-right-radius: 16px;
}
.campaign__management > div:last-child{
  border-bottom-left-radius: 16px;
  border-bottom-right-radius: 16px;
}
.campaign__management > div:last-child::after{
  display: none;
}
.campaign__management > div::after{
  content: " ";
  display: block;
  width: 100%;
  height: 1px;
  background-color: $rectangle;
}
.campaign__management > div > p{
  padding: 16px 0 ;
  margin: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.campaign__management > div > p > svg{
  width: 8px;
  height: 8px;
  fill: $black;
  cursor: pointer;
}
.red__color{
  color : #FF2F12
}
.purple__color{
  color : #6C1BD2
}
</style>
