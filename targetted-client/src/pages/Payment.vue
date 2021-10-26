<template>
    <div>
      <div id="content-wrapper">
        <div id="content">
          <b-row class="app-new-steps-header">
              <b-col>
                <h2>Оплата</h2>

                 <b-button type="button"
                    class="app-new-submit-button mt-5"
                    @click="$router.back();">
                  Назад
            </b-button>
              </b-col>
          </b-row>
        </div>
      </div>
    </div>
</template>

<script>
export default {

}
</script>

<style lang="scss">
  @import '@/assets/styles/vars.scss';

</style>