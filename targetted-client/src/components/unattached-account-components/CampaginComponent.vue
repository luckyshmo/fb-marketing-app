<template>
<div class="campagin__component">
  <router-link :to="{name: 'mainPage'}">
    <p
      id="navigation-text"
      style="margin:52px 0 0 0; display: flex; align-items: center"
    >
      ← На главную
    </p>
  </router-link>
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
    <p>{{sectionText}}</p>
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
      Баланс<span>:</span><span>{{budgetDay}}</span>
    </p>
    <button class="app-new-submit-button budget__button">{{budgetButton}}</button>
  </div>
  <div class="campaign__management">
    <h2>Управление кампанией</h2>
    <div v-for="val in management"
         :key="val">
      <p :class="{'red__color' : val.colorRed, 'purple__color' : val.colorPurple }">
        {{val.text}}
        <svg viewBox="0 0 7 12" fill="none" xmlns="http://www.w3.org/2000/svg" v-if="val.svg">
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
  line-height: 38px;
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
  }
  .unattached-account__block_right > div{
    justify-content: center;
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
