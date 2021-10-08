<template>
  <section class="app-new-stats">
    <b-row>
      <b-col cols="12" sm="6">
        <div class="app-new-stats-type">{{data.seen.min}} – {{data.seen.max}} человек</div>
        <div class="app-new-stats-details">Увидят вашу рекламу</div>
      </b-col>
      <b-col cols="12" sm="6">
        <div class="app-new-stats-type">{{data.clicks.min}} – {{data.clicks.max}} человек</div>
        <div class="app-new-stats-details">Перейдут по вашей рекламе</div>
      </b-col>
      <b-col cols="12" sm="6">
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

</style>