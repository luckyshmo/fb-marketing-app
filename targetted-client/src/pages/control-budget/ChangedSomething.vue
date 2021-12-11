<template>
<div class="changed__something">
  <ControlBudget
    :controlBudgetData="controlBudget"
  />
</div>
</template>

<script>
import ControlBudget from "@src/components/control-budget-components/ControlBudget";
export default {
  name: "ChangedSomething",
  data(){
    return{
      controlBudget: {
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
        buttonBlockText: 'После завершения кампании\n' +
          'на балансе останется 1 000 ₽',
        buttonValue: 'Сохранить изменения'
      }
    }
  },
  components: {ControlBudget}
}
</script>

<style scoped>

</style>
