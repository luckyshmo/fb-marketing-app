<template>
  <section class="app-new-stats stats-section">
    <b-row>
      <b-col cols="12" sm="6" class="stats-col">
        <div class="app-new-stats-type">{{data.seen.min}} – {{data.seen.max}} человек</div>
        <div class="app-new-stats-details">Увидят вашу рекламу</div>
      </b-col>
      <b-col cols="12" sm="6" class="stats-col">
        <div class="app-new-stats-type">{{data.clicks.min}} – {{data.clicks.max}} человек</div>
        <div class="app-new-stats-details">Перейдут по вашей рекламе</div>
      </b-col>
      <b-col cols="12" sm="6" class="stats-col">
        <div class="app-new-stats-type">{{data.costs}} ₽ </div>
        <div class="app-new-stats-details">Общие затраты</div>
      </b-col>
    </b-row>
  </section>
</template>

<script>
export default {
  name: 'CampaginStats',
  props: ['data']
}
</script>
<style>
.stats-section{
  margin-top: 40px;
}
.stats-col:nth-child(3){
  margin-top: -7px;
}
@media (max-width: 600px) {
  .stats-section{
    margin-top: 28px;
  }
  .stats-col:nth-child(2){
    margin-top: -7px;
  }
}
</style>
