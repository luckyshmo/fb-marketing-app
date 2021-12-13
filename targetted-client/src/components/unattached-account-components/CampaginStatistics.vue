<template>
  <div class="statistic">
    <h2>{{campaginStatisticsData.title}}</h2>
    <p v-for="val in campaginStatisticsData.campaginStatisticsValue" :key="val.title">
      {{val.title}}<span>:</span><span>{{val.value}}</span>
    </p>
    <button class="app-new-submit-button statistic__button">{{campaginStatisticsData.buttonValue}}</button>
  </div>
</template>

<script>
export default {
  name: "CampaginStatistics",
  props: {
    campaginStatisticsData: {
      type: Object,
      required: true
    }
  }
}
</script>

<style lang="scss">
@import '@src/assets/styles/vars.scss';
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
  .statistic > p{
    display: flex;
    justify-content: space-between;
  }
  .statistic > p > span:nth-child(1){
    display: none;
  }
}
</style>
