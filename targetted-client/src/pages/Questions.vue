<template>
  <div>
    <h1 class="app-header">
      Остались вопросы?
    </h1>

    <b-form
      @submit.prevent="sendRequest">
      <b-form-group
        id="input-group-main">
      <p>
          Оставьте свой номер телефона, если у вас возникли вопросы или предложения :)
      </p>
        <b-form-input
          v-model="$v.form.phoneNumber.$model"
          :state="$v.form.phoneNumber.$dirty ? !$v.form.phoneNumber.$error : null"
          class="form-input"
          placeholder="Телефон"
        />

            <small
                  v-if="$v.form.phoneNumber.$dirty && !$v.form.phoneNumber.required"
                  class="error-message"
                >
                  Телефон не указан
                </small>
                  <small
                  v-if="$v.form.phoneNumber.$dirty && !$v.form.phoneNumber.numeric"
                  class="error-message"
                >
                  Только цифры в телефоне
                </small>
      </b-form-group>

      <b-form-group class="app-new-form-footer">
      <b-button
        type="submit"
        class="app-new-button-main"
      >
        Заказать звонок
      </b-button>

      <p class="app-new-login-sub">
        Или напишите нам в <a href="#">Whatsapp</a> и <a href="#">Telegram</a>
      </p>

      </b-form-group>

    </b-form>
  </div>
</template>
<script>
import { validationMixin } from 'vuelidate'
import { required, numeric } from 'vuelidate/lib/validators'

export default {
  name: 'Questions',
  mixins: [validationMixin],
  data () {
    return {
      form: {
        phoneNumber: ''
      }
    }
  },
  validations: {
    form: {
      phoneNumber: {
        required,
        numeric
        // minLength: minLength(10)
      }
    }
  },
  methods: {
    sendRequest: function () {
      this.$v.$touch()
      // eslint-disable-next-line no-empty
      if (this.$v.$anyError) {

      }
      // TODO
      //   store.dispatch('login', { email, password })
      //     .then(resp => {
      //       this.userNotExist = false
      //       console.log(resp)
      //       router.push('main')
      //     })
      //     .catch(err => {
      //       console.log(err)
      //       if (err.response.data.message === 'sql: no rows in result set') {
      //         this.userNotExist = true
      //       }
      //     }
      //     )
    }
  }
}
</script>
<style>
</style>
