<template>
   <div>
        <slot name="header"></slot>
        <b-form>
            ...
        <b-button class="app-new-submit-button" type="submit">
        {{isEdit ? "Назад":"Продолжить"}}
      </b-button>
    </b-form>
  </div>
</template>

<script>
import store from '@/../store/store'

export default {
  name: 'Step4',
    props: ['label_cols', 'content_cols', 'company', 'isEdit'],
    data: function () {
        return {
          store, //fixme
        }
      },
}
</script>

<style>

</style>