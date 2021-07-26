<template>
  <div id="content-login">
    <h1 id="h1-centered">
      Вход
    </h1>

    <b-form
      id="form-centred"
      @submit.prevent="login"
    >
      <b-form-group
        class="input-group"
      >
        <b-form-input
          v-model="email"
          class="form-input"
          type="email"
          :state="validateState('email')"
          placeholder="Электронная почта"
          @click="resetErr()"
        />
        <small
          v-if="$v.email.$dirty && !$v.email.required"
          class="error-message"
        >
          Пустое поле email
        </small>
        <small
          v-if="$v.email.$dirty && !$v.email.email"
          class="error-message"
        >
          Некрорректный email
        </small>
        <small
          v-if="userNotExist"
          class="error-message"
        >
          Пользователь с такими почтой/паролем не существует
        </small>
      </b-form-group>

      <b-form-group
        class="input-group"
      >
        <b-form-input
          id="pas"
          v-model="password"
          type="password"
          class="form-input"
          :state="validateState('password')"
          placeholder="Пароль"
        />
        <b-form-invalid-feedback
          id="pas"
          class="error-message"
        >
          Пароль должен быть минимум 8 символов
        </b-form-invalid-feedback>
      </b-form-group>
      <b-button
        type="submit"
        class="main-button-big"
      >
        Войти
      </b-button>

      <p id="navigation-text">
        или <router-link
          style="color: #6C1BD2"
          :to="{name: 'register'}"
        >
          зарегистрироваться
        </router-link>
      </p>
    </b-form>
  </div>
</template>
<script>
import router from '../../../router/router'
import { validationMixin } from 'vuelidate'
import { required, minLength, email } from 'vuelidate/lib/validators'
import store from '../../../store/store'
export default {
  mixins: [validationMixin],
  data () {
    return {
      userNotExist: false,
      email: '',
      password: ''
    }
  },
  methods: {
    resetErr () {
      this.userNotExist = false
    },
    validateState (name) {
      const { $dirty, $error } = this.$v[name]
      return $dirty ? !$error : null
    },
    login: function () {
      this.$v.$touch()
      if (this.$v.$anyError) {
        return
      }
      const email = this.email
      const password = this.password
      store.dispatch('login', { email, password })
        .then(resp => {
          this.userNotExist = false
          console.log(resp)
          router.push('main')
        })
        .catch(err => {
          console.log(err)
          if (err.response.data.message === 'sql: no rows in result set') {
            this.userNotExist = true
          }
        }
        )
    }
  },
  validations: {
    email: {
      email,
      required
    },
    password: {
      required,
      minLength: minLength(8)
    }
  }
}
</script>
<style>
</style>
