<template>
   <div>
        <slot name="header"></slot>
        <b-form>
            <p>
                Мы самостоятельно подберем наиболее подходящие настройки целевой аудитории на основе анализа большого количество рекламных кампаний из вашей ниши. Но вы можете уточнить настройки, если считаете это важным.
            </p>
        <b-button type="button"
                class="app-new-submit-button"
                @click="sendData">
        {{isEdit ? "Назад":"Продолжить"}}
      </b-button>
      <a @click="showAdvanced = true">Уточнить настройки</a>

      <div v-if="showAdvanced">

          <p>
              Укажите характеристики вашей аудитории, наши алгоритмы будут учитывать их при настройке рекламной кампании. 
          </p>
      </div>
    </b-form>
  </div>
</template>

<script>
import store from '@/../store/store'

export default {
  name: 'Step3',
    props: ['label_cols', 'content_cols', 'company', 'isEdit'],
    data: function () {
        return {
          store, //fixme
          showAdvanced
        }
      },
    methods: {
        sendData(){
            //todo
            this.$emit('next');
        }
    }
}
</script>

<style>

</style>