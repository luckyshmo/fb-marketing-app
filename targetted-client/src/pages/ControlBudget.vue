<template>
<div class="control_budget">
  <ControlBudgetComponent
    v-if="pageId === 1"
    budgetPerDay="500"
    duration="3"
    willSeeMin="3 000"
    willSeeMax="5 000"
    willGoMin="1 500"
    willGoMax="2 000"
    totalCostsMoney="1 500"
    totalCostsDay="3"
    campaignBalance="2 500"
    :buttonBlock="false"
    buttonBlockText=""
    buttonValue=""
  />
  <ControlBudgetComponent
    v-if="pageId === 2"
    budgetPerDay="500"
    duration="3"
    willSeeMin="3 000"
    willSeeMax="5 000"
    willGoMin="1 500"
    willGoMax="2 000"
    totalCostsMoney="1 500"
    totalCostsDay="3"
    campaignBalance="2 500"
    :buttonBlock="true"
    buttonBlockText="После завершения кампании на балансе останется 1 000 ₽"
    buttonValue="Сохранить изменения"
  />
  <ControlBudgetComponent
    v-if="pageId === 3"
    budgetPerDay="500"
    duration="7"
    willSeeMin="3 000"
    willSeeMax="5 000"
    willGoMin="1 500"
    willGoMax="2 000"
    totalCostsMoney="3 500"
    totalCostsDay="7"
    campaignBalance="2 500"
    :buttonBlock="true"
    buttonBlockText="Недостаточно средств для изменения бюджета и длительности"
    buttonValue="Пополнить на 1 000 ₽"
  />
</div>
</template>

<script>
import ControlBudgetComponent from "@src/components/control-budget-components/ControlBudgetComponent";

export default {
  name: "ControlBudget",
  components: {ControlBudgetComponent},
  data() {
    return {
      pageId: 2
    }
  }

}
</script>

<style scoped>
.control_budget{

}
</style>
