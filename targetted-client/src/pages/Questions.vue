<template>
  <div class="questions__block">
    <h2 class="questions__header">
      Остались вопросы?
    </h2>

    <b-form
      @submit.prevent="sendRequest">
      <b-form-group
        id="input-group-main">
      <p>
          Оставьте свой номер телефона, если у вас <br/> возникли вопросы или предложения :)
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
        class="app-new-submit-button"
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
<style lang="scss">
.questions__block{
  padding: 40px 48px 56px;
  #input-group-main {
    margin-bottom: 12px;
    p {
      margin-bottom: 12px;
    }
  }
  .app-new-login-sub{
    margin: 40px 0 0;
  }
}
.questions__header{
  font-style: normal;
  font-weight: bold;
  font-size: 32px !important;
  line-height: 38px !important;
  margin-bottom: 12px;
}
@media (max-width: 470px){
  .questions__block{
    padding: 134px 24px 0;
  }
}
</style>
