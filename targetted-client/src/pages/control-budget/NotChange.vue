<template>
<div class="not__change">
  <ControlBudget
    :controlBudgetData="notChangeData"
  />
</div>
</template>

<script>
import ControlBudget from "@src/components/control-budget-components/ControlBudget";
export default {
  name: "NotChange",
  data() {
    return {
      notChangeData: {
        budgetPerDay: 500,
        duration: 3,
        willSeeMin: 3000,
        willSeeMax: 5000,
        willGoMin: 1500,
        willGoMax: 2000,
        totalCostsMoney: 1500,
        totalCostsDay: 3,
        campaignBalance: 2500,
        buttonBlock: false,
        buttonBlockText: 'После завершения кампании\n' +
          'на балансе останется 1 000 ₽',
        buttonValue: 'Сохранить изменения'
      }
    }
  },
  components : {
    ControlBudget
  }
}
</script>

<style scoped>
.not__change {
  width: 100%;
}
</style>
