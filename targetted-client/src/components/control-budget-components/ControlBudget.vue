<template>
<div class="control__budget">
  <router-link :to="{name: 'mainPage'}">
    <p
      id="navigation-text"
      style="margin:52px 0 0 0; display: flex; align-items: center"
    >
      ← Назад
    </p>
  </router-link>
  <h2 class="control__budget_title">Бюджет <br/>и длительность</h2>
  <div>
    <div class="control__budget_block_left">
      <div class="control__budget_range">
        <div>
          <h4>Бюджет в день <span>:</span> <span>{{ controlBudgetData.budgetPerDay }} ₽</span></h4>
          <vue-range-slider
            v-model="controlBudgetData.budgetPerDay"
            :tooltip-style="tooltip"
            :min="minMoney"
            :max="maxMoney"
            :step="moneyStep"
            height="8"
            :process-style="processStyle"
            :slider-style="bgStyle"
            :bg-style="processStyle"
          ></vue-range-slider>
        </div>
        <div>
          <h4>Длительность <span>:</span>  <span>ещё {{ controlBudgetData.duration }} дня</span></h4>
          <vue-range-slider
            v-model="controlBudgetData.duration"
            height="8"
            :process-style="processStyle"
            :slider-style="bgStyle"
            :bg-style="processStyle"
            :tooltip-style="tooltip"
            :min="minDay"
            :max="maxDay"
            :step="dayStep"
          ></vue-range-slider>
        </div>
      </div>
      <div class="control__budget_text">
        <div>
          <h2>{{ controlBudgetData.willSeeMin }}  – {{ controlBudgetData.willSeeMax }} человек</h2>
          <p>Увидят вашу рекламу</p>
        </div>
        <div>
          <h2>{{ controlBudgetData.willGoMin }}  – {{ controlBudgetData.willGoMax }} человек</h2>
          <p>Перейдут по вашей рекламе</p>
        </div>
        <div>
          <h2>{{ controlBudgetData.totalCostsMoney }} ₽ за {{ controlBudgetData.totalCostsDay }} дней</h2>
          <p>Общие затраты</p>
        </div>
      </div>
    </div>
    <div class="control__budget_block_right">
      <p>На балансе кампании</p>
      <h2>{{ controlBudgetData.campaignBalance }} ₽</h2>
      <div v-if="controlBudgetData.buttonBlock">
        <div></div>
        <p>{{ controlBudgetData.buttonBlockText }}</p>
        <button>{{ controlBudgetData.buttonValue }}</button>
      </div>
    </div>
</div>
</div>
</template>

<script>
import 'vue-range-component/dist/vue-range-slider.css'
import VueRangeSlider from 'vue-range-component'

export default {
  name: "ControlBudget",
  props: {
    controlBudgetData: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      processStyle: {
        background: '#F3F3F3'
      },
      bgStyle: {
        background: '#000000',
        width: '32px',
        height: '32px',
        top: '-12px'
      },
      minDay: 1,
      maxDay: 100,
      dayStep: 1,
      minMoney: 100,
      maxMoney: 5000,
      moneyStep: 50,
      tooltip: {display: 'none'}
    }
  },
  components : {
    VueRangeSlider
  }
}
</script>

<style lang="scss">
@import '@src/assets/styles/vars.scss';
.control__budget_title{
  font-size: 56px !important;
  line-height: 60px !important;
  margin-top: 12px;
}
.control__budget{
  width: 100%;
}
.control__budget > div{
  display: flex;
  justify-content: space-between;
}
.control__budget_block_left{
  max-width: 583px;
  width: 100%;
}
.control__budget_range{
  width: 100%;
  display: flex;
  justify-content: space-between;
}
.control__budget_range > div{
  max-width: 276px;
  width: 100%;
}
.control__budget_range > div > h4{
  margin-top: 40px;
  font-size: 16px;
  line-height: 24px;
  font-weight: bold;
}
.control__budget_text{
  width: 100%;
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}
.control__budget_text > div{
  max-width: 276px;
  width: 100%;
}
.control__budget_text > div >h2{
  margin-top: 40px;
}
.control__budget_text > div >p{
  color: $gray;
  margin: 0;
}
.control__budget_text > div:nth-child(3) > h2{
  margin-top: 20px;
}
.control__budget_block_right{
  max-width: 300px;
  width: 100%;
  background: $light;
  border-radius: 16px;
  height: fit-content;
}
.control__budget_block_right > p{
    padding: 24px 0 0 24px;
    margin-bottom: 0;
    color: $gray;
}
.control__budget_block_right > h2{
  font-size: 32px !important;
  line-height: 38px !important;
  padding-left: 24px;
  margin-bottom: 20px;
}
.control__budget_block_right > div > div{
  width: 100%;
  height: 1px;
  background-color: #DDDDDD;
}
.control__budget_block_right > div > p{
  padding: 20px  24px 16px;
  color: $gray;
}
.control__budget_block_right > div > button{
  font-weight: 500;
  font-size: 16px;
  line-height: 24px;
  padding: 9px 43px 11px;
  border: 0;
  background: #000;
  color: #fff;
  border-radius: 6px;
  margin: 0 0 40px 24px;
}
@media (max-width: 600px) {
  .control__budget_title{
    font-size: 32px !important;
    line-height: 38px !important;
  }
  .control__budget > div {
    flex-wrap: wrap;
  }
  .control__budget_range {
    flex-wrap: wrap;
  }
  .control__budget_range > div {
    max-width: 600px;
    width: 100%;
  }
  .control__budget_range > div > h4 {
    width: 100%;
    display: flex;
    justify-content: space-between;
    margin-bottom: 12px;
  }
  .control__budget_range > div > h4 > span:nth-child(1){
    display: none;
  }
  .control__budget_range > div > h4 > span:nth-child(2){
   text-transform: capitalize;
  }
  .control__budget_text{
    flex-wrap: wrap;
  }
  .control__budget_text > div {
    max-width: 600px;
    width: 100%;
  }
  .control__budget_text > div:nth-child(2) > h2 {
    margin-top: 20px;
  }
  .control__budget_block_right{
    margin: 40px 0 80px;
  }
}
</style>
