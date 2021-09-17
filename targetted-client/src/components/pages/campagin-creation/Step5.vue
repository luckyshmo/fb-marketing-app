<template>
   <div>
        <slot name="header"></slot>
        <b-form>
            ...
        <b-button type="button"
                        class="app-new-submit-button"
                        @click="sendData">
        {{isEdit ? "Назад":"Продолжить"}}
      </b-button>
    </b-form>
  </div>
</template>

<script>
import store from '@/../store/store'

export default {
  name: 'Step5',
    props: ['label_cols', 'content_cols', 'company', 'isEdit'],
    data: function () {
        return {
          store, //fixme
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