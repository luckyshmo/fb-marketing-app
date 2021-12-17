<template>
<div class="no__money">
  <ControlBudget
    :controlBudgetData="noMoneyData"
  />
</div>
</template>

<script>
import ControlBudget from "@src/components/control-budget-components/ControlBudget";
export default {
  name: "NoMoney",
  components:{
    ControlBudget
  },
  data(){
    return{
      noMoneyData: {
        budgetPerDay: 500,
        duration: 3,
        willSeeMin: 3000,
        willSeeMax: 5000,
        willGoMin: 1500,
        willGoMax: 2000,
        totalCostsMoney: 1500,
        totalCostsDay: 3,
        campaignBalance: 2500,
        buttonBlock: true,
        buttonBlockText: 'Недостаточно средств для изменения бюджета и длительности',
        buttonValue: 'Пополнить на 1 000 ₽'
      }
    }
  }
}
</script>

<style scoped>

</style>
