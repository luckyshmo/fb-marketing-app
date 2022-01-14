<template>
<div class="campagin__component">
  <b-row class="app-new-steps-header">
    <b-col style="padding: 0; display: flex; align-items: center">
      <router-link id="app-link-back" :to="{path: '/'}" style="display: flex; align-items: center">

        <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
             x="0px" y="0px"
             viewBox="0 0 512 512" style="enable-background:new 0 0 512 512;" xml:space="preserve">
<path d="M3.919,243.077c-0.223,0.33-0.414,0.675-0.618,1.015c-0.186,0.31-0.382,0.614-0.552,0.936
	c-0.186,0.349-0.346,0.709-0.514,1.066c-0.157,0.332-0.321,0.658-0.464,0.998c-0.144,0.349-0.261,0.706-0.388,1.06
	c-0.129,0.362-0.268,0.718-0.38,1.089c-0.107,0.358-0.188,0.721-0.279,1.085c-0.093,0.374-0.199,0.743-0.275,1.125
	c-0.084,0.422-0.133,0.849-0.194,1.275c-0.047,0.326-0.109,0.645-0.143,0.976c-0.15,1.53-0.15,3.07,0,4.6
	c0.034,0.33,0.096,0.65,0.143,0.976c0.061,0.425,0.11,0.853,0.194,1.275c0.076,0.382,0.18,0.749,0.275,1.125
	c0.092,0.362,0.171,0.726,0.279,1.085c0.112,0.369,0.251,0.726,0.38,1.089c0.127,0.354,0.244,0.711,0.388,1.06
	c0.143,0.34,0.307,0.666,0.464,0.998c0.168,0.355,0.327,0.715,0.514,1.064c0.171,0.321,0.366,0.625,0.552,0.936
	c0.203,0.34,0.394,0.684,0.618,1.015c0.234,0.351,0.493,0.68,0.745,1.015c0.205,0.272,0.393,0.549,0.608,0.813
	c0.489,0.596,1.002,1.168,1.548,1.711l116.36,116.36c4.544,4.544,10.501,6.817,16.455,6.817c5.956,0,11.913-2.271,16.455-6.817
	c9.089-9.089,9.089-23.824,0-32.912l-76.636-76.636h409.272c12.853,0,23.273-10.42,23.273-23.273
	c0-12.853-10.42-23.273-23.273-23.273H79.456l76.636-76.636c9.089-9.089,9.089-23.824,0-32.912c-9.087-9.089-23.824-9.089-32.912,0
	L6.82,239.538c-0.546,0.543-1.06,1.116-1.548,1.711c-0.216,0.264-0.403,0.541-0.608,0.813
	C4.412,242.397,4.153,242.726,3.919,243.077z"/>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
          <g>
</g>
</svg>

        На главную
      </router-link>
    </b-col>
  </b-row>
  <div class="unattached-account__block">
    <div class="unattached-account__img">
      <img src="@src/assets/img/image_2.png" alt="unattached-account__img">
    </div>
    <div class="unattached-account__block_right">
      <h2>
        Рекламная кампания от 7 сентября
      </h2>
      <div>
        <div :style="{backgroundColor : headerActiveColor}"></div>
        <p>{{ headerText }}</p>
      </div>
    </div>
  </div>
  <div class="page_anchor" v-if="section">
    <h2>{{sectionTitle}}</h2>
    <p v-html="sectionText"></p>
    <button class="app-new-submit-button">{{sectionButton}}</button>
  </div>
  <div class="statistic" v-if="statistics">
    <h2>Статистика</h2>
    <p>
      Количество показов<span>:</span><span>{{statisticsImpressions}}</span>
    </p>
    <p>
      Количество переходов<span>:</span><span>{{statisticsTransitions}}</span>
    </p>
    <p>
      Количество написавших<span>:</span><span>{{statisticsWrote}}</span>
    </p>
    <button class="app-new-submit-button statistic__button">Вся статистика</button>
  </div>
  <div class="budget__duration" v-if="budget">
    <h2>Бюджет и длительность</h2>
    <p>
      Баланс<span>:</span><span>{{balance}}</span>
    </p>
    <p>
      Длительность<span>:</span><span>{{duration}}</span>
    </p>
    <p>
      Бюджет в день<span>:</span><span>{{budgetDay}}</span>
    </p>
    <button class="app-new-submit-button budget__button">{{budgetButton}}</button>
  </div>
  <div class="campaign__management">
    <h2>Управление кампанией</h2>
    <div v-for="val in management"
         :key="val">
      <p :class="{'red__color' : val.colorRed, 'purple__color' : val.colorPurple }">
        {{val.text}}
        <svg width="7" height="12" viewBox="0 0 7 12" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M1 1L5.5 6L1 11" stroke="black" stroke-width="2" stroke-linecap="round"/>
        </svg>
      </p>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: "CampaginComponent",
  props:[
    'headerActiveColor',
    'headerText',
    'statistics',
    'statisticsImpressions',
    'statisticsTransitions',
    'statisticsWrote',
    'section',
    'sectionTitle',
    'sectionText',
    'sectionButton',
    'budget',
    'balance',
    'duration',
    'budgetDay',
    'budgetButton',
    'management'
  ],
  data(){
    return{

    }
  }
}
</script>

<style lang="scss">
@import '@src/assets/styles/vars.scss';
.unattached-account__block {
  width: 100%;
  display: flex;
  margin-top: 20px;
  border-radius: 16px;
  background-color: $light;
}
.unattached-account__img{
  width: 140px;
  height: 140px;
  border-radius: 12px;
  margin: 24px;
}
.unattached-account__img > img{
  width: 100%;
  height: 100%;
}
.unattached-account__block_right{
  max-width: 346px;
}
.unattached-account__block_right  > h2{
  margin-top: 38px !important;
  font-size: 32px !important;
  line-height: 38px !important;
  color: $black;
}
.unattached-account__block_right > div{
  width: 100%;
  display: flex;
  align-items: center;
  margin-top: 12px;
}
.unattached-account__block_right > div > p{
  font-size: 14px;
  line-height: 18px;
  color: $black;
  position: relative;
  margin: 0;
}
.unattached-account__block_right > div > div{
  width: 8px;
  height: 8px;
  //background-color: $gray;
  border-radius: 100%;
  margin-right: 7px;
}
.page_anchor{
  padding-top: 40px;
}
.page_anchor > p{
  margin: 8px 0 20px;
}
.budget__duration{
  margin-top: 40px;
}
.budget__duration > p{
  margin: 12px 0 0;
}
.budget__duration > p > span:nth-child(2){
  margin-left: 6px;
  font-weight: bold;
}
.budget__button{
  background: $light !important;
  margin-top: 20px;
}
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
  padding: 14px 0 16px;
  margin: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.campaign__management > div > p > svg{
  cursor: pointer;
}
.red__color{
  color : #FF2F12
}
.purple__color{
  color : #6C1BD2
}
.statistic{
  margin-top: 40px;
}
.statistic > p{
  margin: 12px 0 0;
}
.statistic > p > span:nth-child(2){
  margin-left: 6px;
  font-weight: bold;
}
.statistic__button{
  background: $light !important;
  margin-top: 20px;
}

@media (max-width: 600px) {
  .unattached-account__block{
    position: relative;
    justify-content: center;
    padding-bottom: 28px;
    margin-top: 70px;
  }
  .unattached-account__img{
    position: absolute;
    top: -50px;
    width: 100px;
    height: 100px;
    margin: 0;
  }
  .unattached-account__block_right{
    width: 100%;
    max-width: 254px;
    text-align: center;
  }
  .unattached-account__block_right > h2{
    margin-top: 66px !important;
    font-size: 24px !important;
    line-height: 30px !important;
  }
  .unattached-account__block_right > div{
    justify-content: center;
  }
  .campaign__management{
    margin: 40px 0 78px;
  }
  .budget__duration > p{
    display: flex;
    justify-content: space-between;
  }
  .budget__duration > p > span:nth-child(1){
    display: none;
  }
  .statistic > p{
    display: flex;
    justify-content: space-between;
  }
  .statistic > p > span:nth-child(1){
    display: none;
  }
}
</style>
