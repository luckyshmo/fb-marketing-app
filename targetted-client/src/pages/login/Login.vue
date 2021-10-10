<template>
<div id="content-wrapper">
  <div id="content">
    
    <h1 class="app-header">
      Вход
    </h1>
<b-form
      id="form-centred"
      @submit.prevent="login"
    >
      <b-form-group
        class="input-group"
        id="input-group-main"
        label="Электронная почта"
      >
        <b-form-input
          v-model="email"
          class="form-input"
          type="email"
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
          Некорректный email
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
        id="input-group-main"
      >
        <b-form-input
          id="pas"
          v-model="password"
          type="password"
          class="form-input"
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
        <br class="d-none d-sm-block">
      <b-button
        type="submit"
        class="app-new-submit-button"
        :class="{'disabled': !$v.$anyDirty || $v.$anyError}"
      >
        Войти
      </b-button>

      <p class="app-new-login-sub mt-5">
        Или <router-link :to="{name: 'register'}">
          зарегистрируйтесь
        </router-link>
      </p>

      </b-form-group>

    </b-form>
    
  </div>
  </div>
</template>
<script>

import { validationMixin } from 'vuelidate'
import { required, minLength, email } from 'vuelidate/lib/validators'
import store from '@/store/store'

export default {
  mixins: [validationMixin],
  data () {
    return {
      userNotExist: false,
      email: '',
      password: '',
      store
    }
  },
  methods: {
    resetErr () {
      this.userNotExist = false
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
          this.$router.push('main')
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
