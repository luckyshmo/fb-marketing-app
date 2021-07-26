<template>
  <div id="content-login">
    <h1 id="h1-centered">
      Регистрация
    </h1>

    <b-form
      id="form-centred"
      @submit="register"
    >
      <b-form-group
        class="input-group"
      >
        <b-form-input
          v-model="form.name"
          class="form-input"
          placeholder="Имя"
          :state="validateState('name')"
        />
        <small
          v-if="$v.form.name.$dirty && !$v.form.name.required"
          class="error-message"
        >
          Пустое поле Имя
        </small>
      </b-form-group>

      <b-form-group
        class="input-group"
      >
        <b-form-input
          id="tel"
          v-model="form.phoneNumber"
          class="form-input"
          type="tel"
          :state="validateState('phoneNumber')"
          placeholder="Номер телефона"
        />
        <b-form-invalid-feedback
          id="tel"
          class="error-message"
        >
          Номер телефона должен содержать как минимум 10 символов
        </b-form-invalid-feedback>
      </b-form-group>

      <b-form-group
        class="input-group"
      >
        <b-form-input
          v-model.trim="form.email"
          class="form-input"
          placeholder="Электронная почта"
          :state="validateState('email')"
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
          Некрорректный email
        </small>
        <small
          v-if="emailExists"
          class="error-message"
        >
          Пользователь с таким email уже существует
        </small>
      </b-form-group>

      <b-form-group
        class="input-group"
      >
        <b-form-input
          id="pas"
          v-model="form.password"
          class="form-input"
          type="password"
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
      <b-form-group
        class="input-group"
      >
        <b-form-input
          v-model="form.passwordRepeat"
          type="password"
          class="form-input"
          aria-describedby="password-repeat-feedback"
          :state="validatePassword('passwordRepeat')"
          placeholder="Введите пароль повторно"
        />
        <small
          v-if="form.password != form.passwordRepeat"
          class="error-message"
        >
          Пароли не совпадают
        </small>
      </b-form-group>

      <b-button
        class="main-button-big"
        @click="register()"
      >
        Зарегистрироваться
      </b-button>

      <p id="navigation-text">
        или <router-link
          style="color: #6C1BD2"
          :to="{name: 'login'}"
        >
          войдите
        </router-link> в свой аккаунт
      </p>
    </b-form>
  </div>
</template>
<script>
import router from '../../../router/router'
import axios from 'axios'
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
      console.log($dirty, $error, this.form.password === this.form.passwordRepeat)
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
      axios({ url: `${VUE_APP_API_URL}/auth/sign-up`, data: data, method: 'POST' })
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
<style>
.error-message{
  position: absolute;
  color: red;
  font-family: Montserrat;
  font-style: normal;
}
.form-control{
  width: 520px !important;
}
@media(max-width: 600px){
  .form-control{
    width: 265px !important;
  }
}
.input-group{
  text-align: left;
  margin: 32px 20px 32px 20px !important;
}
</style>
