<template>
        <section class="app-new-stats">
        <div class="app-new-stats-type">{{data.seen.min}} – {{data.seen.max}} человек</div>
        <div class="app-new-stats-details">Увидят вашу рекламу</div>
        <div class="app-new-stats-type">{{data.clicks.min}} – {{data.clicks.max}} человек</div>
        <div class="app-new-stats-details">Перейдут по вашей рекламе</div>
        <div class="app-new-stats-type">{{data.costs}} ₽ </div>
        <div class="app-new-stats-details">Общие затраты</div>

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