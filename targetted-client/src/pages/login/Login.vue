<template>
  <div id="content-login">
    <h1 class="app-header">
      Вход
    </h1>

    <b-form
      id="form-centred"
      @submit.prevent="login"
    >
      <b-form-group
        class="input-group"
        label="Электронная почта"
      >
        <b-form-input
          v-model="email"
          class="form-input"
          type="email"
          :state="validateState('email')"
          placeholder="Введите почту"
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
        label="Пароль"
      >
        <b-form-input
          id="pas"
          v-model="password"
          type="password"
          class="form-input"
          :state="validateState('password')"
          placeholder="Введите пароль"
        />
        <b-form-invalid-feedback
          id="pas"
          class="error-message"
        >
          Пароль должен быть минимум 8 символов
        </b-form-invalid-feedback>
      </b-form-group>

      <b-form-group class="app-new-form-footer">
      <b-button
        type="submit"
        class="app-new-button-main"
      >
        Войти
      </b-button>

      <p class="app-new-login-sub">
        Или <router-link :to="{name: 'register'}">
          зарегистрироваться
        </router-link>
      </p>

      </b-form-group>

    </b-form>
  </div>
</template>
<script>
import router from '@/router/router'
import { validationMixin } from 'vuelidate'
import { required, minLength, email } from 'vuelidate/lib/validators'
import store from '@/store/store'
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
