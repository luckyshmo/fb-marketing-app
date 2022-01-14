<template>
<div id="content-wrapper">
  <div id="content" class="step">
    <h1 class="app-header register__header">
      Регистрация
    </h1>

    <b-form
      @submit.prevent="register"
    >
        <div class="register__form_top">
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
        </div>

      <b-form-group
        label="Электронная почта"
        id="input-group-main"
        class="register__form_email"
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
      <div class="register__form_bottom">
        <b-form-group
          label="Пароль"
          id="input-group-main"
        >
          <b-form-input
            id="pas"
            v-model="form.password"
            class="form-input"
            type="password"
            placeholder="Введите пароль"
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
      </div>
      <b-form-group class="app-new-form-footer">
          <br class="d-none d-sm-block">
        <b-button
          class="app-new-submit-button"
          :class="{'disabled-look': !$v.$anyDirty || $v.$anyError}"
          @click="register()"
        >
          Зарегистрироваться
        </b-button>

        <p class="app-new-info register__form_bottom_text">
          Нажимая кнопку “Зарегистрироваться” вы соглашаетесь с условиями <a href="#">оферты</a> и <a href="#">политикой конфиденциальности</a>.
        </p>

        <p class="app-new-login-sub register__form_bottom_link">
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

//import router from '@src/router/router'
//import {instance as axios} from '@src/_services/index'
import {validationMixin} from 'vuelidate'
import {required, minLength, email} from 'vuelidate/lib/validators'

//const VUE_APP_API_URL = process.env.VUE_APP_API_URL
export default {
  mixins: [validationMixin],
  data() {
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
    resetErr() {
      this.emailExists = false
    },
    validateState(name) {
      const {$dirty, $error} = this.$v.form[name]
      return $dirty ? !$error : null
    },
    validatePassword(name) {
      const {$dirty, $error} = this.$v.form[name]
      // console.log($dirty, $error, this.form.password === this.form.passwordRepeat)
      if ($dirty ? !$error : null) {
        return this.form.password === this.form.passwordRepeat
      } else {
        return $dirty ? !$error : null
      }
    },
    register() {
      this.$v.form.$touch()
      if (this.$v.form.$anyError) {
        return
      }
      // const data = {
      //   name: this.form.name,
      //   email: this.form.email,
      //   password: this.form.password,
      //   phoneNumber: this.form.phoneNumber
      // }
      // const url = `${VUE_APP_API_URL}/auth/sign-up`
      // axios.post(url, data)
      //   .then(resp => {
      //     this.emailExists = false
      //     router.push('login')
      //   })
      //   .catch(err => {
      //       if (err.response.data.message === 'pq: duplicate key value violates unique constraint "user_tb_email_key"') {
      //         this.emailExists = true
      //       }
      //     }
      //   )
    }
  }
}
</script>
<style lang="scss">
.register__header{
  margin-top: 80px !important;
}
.register__form_top{
  width: 100%;
  display: flex;
  justify-content: space-between;
  margin-top: 46px;
  & > #input-group-main{
    max-width: 276px;
    }
}
.register__form_email{
  margin-top: -7px;
}
.register__form_bottom{
  width: 100%;
  display: flex;
  justify-content: space-between;
  margin-top: -7px;
  margin-bottom: -35px;
  & > #input-group-main{
    max-width: 276px;
  }
}
.register__form_bottom_text{
  font-style: normal;
  font-weight: normal;
  font-size: 14px !important;
  line-height: 18px !important;
  margin-top: 28px !important;
  margin-bottom: 39px !important;
}
.register__form_bottom_link{
  margin-bottom: 196px;
}
@media (max-width: 600px) {
  .register__header{
    margin-top: 55px !important;
  }
  .register__form_top{
    margin-top: 33px;
    flex-wrap: wrap;
    justify-content: flex-start;
  }
  .register__form_top > #input-group-main{
    max-width: 100%;
  }
  .register__form_top > #input-group-main:nth-child(2){
    margin-top: -6px;
  }
  .register__form_bottom{
    flex-wrap: wrap;
    margin-bottom: -13px;
  }
  .register__form_bottom > #input-group-main{
    width: 100%;
  }
  .register__form_bottom > #input-group-main:nth-child(2){
    margin-top: -7px;
  }
  .register__form_bottom_text{
    margin-bottom: 40px !important;
  }
  .register__form_bottom_link{
    margin-bottom: 80px;
  }
}
</style>
