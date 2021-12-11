<template>
  <div class="budget__duration">
    <h2>{{campaginBudgetData.title}}</h2>
    <p v-for="val in campaginBudgetData.campaginBudgetValue" :key="val.title">
      {{val.title}}<span>:</span><span>{{val.value}}</span>
    </p>
    <button class="app-new-submit-button budget__button">{{campaginBudgetData.buttonValue}}</button>
  </div>
</template>

<script>
export default {
  name: "CampaginBudget",
  props: {
    campaginBudgetData: {
      type: Object,
      required: true
    }
  }
}
</script>

<style lang="scss">
@import '@src/assets/styles/vars.scss';
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
@media (max-width: 600px) {
  .budget__duration > p{
    display: flex;
    justify-content: space-between;
  }
  .budget__duration > p > span:nth-child(1){
    display: none;
  }
}
</style>
