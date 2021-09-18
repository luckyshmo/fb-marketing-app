<template>
  <div id="content-login">
    <h1 class="app-header">
      Остались воросы?
    </h1>

    <b-form
      id="form-centred"
      @submit.prevent="sendRequest"
    >
      <b-form-group
        class="input-group"
      >
      <p>
          Оставьте свой номер телефона, если у вас возникли вопросы или предложения :)
      </p>
        <b-form-input
          v-model="phone"
          class="form-input"
        :state="validateState('phoneNumber')"
          placeholder="Телефон"
        />
        
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
import router from '../../../router/router'
import { validationMixin } from 'vuelidate'
import { required, minLength } from 'vuelidate/lib/validators'
import store from '../../../store/store'
export default {
    name: 'Questions',
  mixins: [validationMixin],
  data () {
    return {
      phone: false,
    }
  },
    validations: {
    form: {
      phoneNumber: {
        required,
        minLength: minLength(10)
      },
    }
    },
  methods: {
    validateState (name) {
    //   const { $dirty, $error } = this.$v[name]
    //   return $dirty ? !$error : null
    },
    sendRequest: function () {
      this.$v.$touch()
      if (this.$v.$anyError) {
        return
      }
        //TODO
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
  },
}
</script>
<style>
</style>
