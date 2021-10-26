<template>
<div id="content-wrapper">
  <div id="content">
    <h1 class="app-header">
      Регистрация
    </h1>

    <b-form
      @submit="register"
    >

      <b-form-group
        label="Номер телефона"
         id="input-group-main"
      >

        <b-form-input
          id="tel"
          v-model="form.phoneNumber"
          class="form-input"
          :state="$v.form.phoneNumber.$dirty ? !$v.form.phoneNumber.$error : null"
          type="tel"
          placeholder="Введите телефон"
        />
        <b-form-invalid-feedback
          class="error-message"
        >
          Минимум 10 символов
        </b-form-invalid-feedback>
      </b-form-group>

      <b-form-group
        label="Электронная почта"
         id="input-group-main"
      >
        <b-form-input
          v-model.trim="form.email"
          class="form-input"
          placeholder="Введите почту"
          :state="$v.form.email.$dirty ? !$v.form.email.$error : null"
          @click="resetErr()"
        />
        <small
          v-if="$v.form.email.$dirty && !$v.form.email.required"
          class="error-message"
        >
          Пустое поле email
        </small>
        <small
          v-if="$v.form.email.$dirty && !$v.form.email.email"
          class="error-message"
        >
          Некорректный email
        </small>
        <small
          v-if="emailExists"
          class="error-message"
        >
          Пользователь с таким email уже существует
        </small>
      </b-form-group>

       <b-form-group
        label="Имя"
         id="input-group-main"
      >
        <b-form-input
          v-model="form.name"
          class="form-input"
          placeholder="Введите имя"
          :state="$v.form.name.$dirty ? !$v.form.name.$error : null"
        />
        <small
          v-if="$v.form.name.$dirty && !$v.form.name.required"
          class="error-message"
        >
          Пустое поле Имя
        </small>
      </b-form-group>

      <b-form-group
        label="Пароль"
         id="input-group-main"
      >
        <b-form-input
          id="pas"
          v-model="form.password"
          class="form-input"
          type="password"
          placeholder="Пароль"
        />
        <b-form-invalid-feedback
          id="pas"
          class="error-message"
        >
          Пароль должен быть минимум 8 символов
        </b-form-invalid-feedback>
      </b-form-group>
      <b-form-group
         id="input-group-main"
        label="Введите пароль повторно"
      >
        <b-form-input
          v-model="form.passwordRepeat"
          type="password"
          class="form-input"
          aria-describedby="password-repeat-feedback"
          :state="validatePassword('passwordRepeat')"
          placeholder="Введите пароль"
        />
        <small
          v-if="form.password !== form.passwordRepeat"
          class="error-message"
        >
          Пароли не совпадают
        </small>
      </b-form-group>

      <b-form-group class="app-new-form-footer">
<br class="d-none d-sm-block">
        <b-button
          class="app-new-submit-button"
          :class="{'disabled-look': !$v.$anyDirty || $v.$anyError}"
          @click="register()"
        >
          Зарегистрироваться
        </b-button>

        <p class="app-new-info">
          Нажимая кнопку “Зарегистрироваться” вы соглашаетесь с условиями <a href="#">оферты</a> и <a href="#">политикой конфиденциальности</a>.
        </p>

        <p class="app-new-login-sub">
          Или <router-link :to="{name: 'login'}">
            войдите
          </router-link> в свой аккаунт
        </p>
      </b-form-group>
    </b-form>
  </div>
  </div>
</template>
<script>
import router from '@src/router/router'
import { instance as axios } from '@src/_services/index'
import { validationMixin } from 'vuelidate'
import { required, minLength, email } from 'vuelidate/lib/validators'
const VUE_APP_API_URL = process.env.VUE_APP_API_URL
export default {
  mixins: [validationMixin],
  data () {
    return {
      emailExists: false,
      form: {
        name: '',
        phoneNumber: '',
        email: '',
        password: '',
        passwordRepeat: ''
      }
    }
  },
  validations: {
    form: {
      email: {
        email,
        required
      },
      phoneNumber: {
        required,
        minLength: minLength(10)
      },
      password: {
        required,
        minLength: minLength(8)
      },
      passwordRepeat: {
        required,
        minLength: minLength(8)
      },
      name: {
        required,
        minLength: minLength(3)
      }
    }
  },
  methods: {
    resetErr () {
      this.emailExists = false
    },
    validateState (name) {
      const { $dirty, $error } = this.$v.form[name]
      return $dirty ? !$error : null
    },
    validatePassword (name) {
      const { $dirty, $error } = this.$v.form[name]
      // console.log($dirty, $error, this.form.password === this.form.passwordRepeat)
      if ($dirty ? !$error : null) {
        return this.form.password === this.form.passwordRepeat
      } else {
        return $dirty ? !$error : null
      }
    },
    register () {
      this.$v.form.$touch()
      if (this.$v.form.$anyError) {
        return
      }

      const data = {
        name: this.form.name,
        email: this.form.email,
        password: this.form.password,
        phoneNumber: this.form.phoneNumber
      }
      console.log('rData', data)
      axios.post({
        url: `${VUE_APP_API_URL}/auth/sign-up`,
        data
      })
        .then(resp => {
          this.emailExists = false
          console.log(resp)
          router.push('login')
        })
        .catch(err => {
          console.log(err.response)
          if (err.response.data.message === 'pq: duplicate key value violates unique constraint "user_tb_email_key"') {
            this.emailExists = true
          }
        }
        )
    }
  }
}
</script>
