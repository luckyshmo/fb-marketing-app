<template>
<div id="content-wrapper">

  <div id="content" class="step">

    <h1 class="app-header login__header">
      Войти
    </h1>
<!--    <b-button   class="app-new-submit-button less-padding"-->
<!--                  @click="loginFB">-->
<!--                Есть аккаунт-->
<!--              </b-button>-->
<b-form
      @submit.prevent="login"
    >
  <div class="login__form_block">
    <b-form-group
      id="input-group-main"
      label="Электронная почта"
    >
      <b-form-input
        v-model="email"
        class="form-input"
        :state="$v.email.$dirty ? !$v.email.$error : null"
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
  </div>

      <b-form-group class="app-new-form-footer">
      <b-button
        type="submit"
        class="app-new-submit-button"
        :class="{'disabled-look': !$v.$anyDirty || $v.$anyError}"
      >
        Войти
      </b-button>
      <p class="app-new-login-sub login__form_block_text">
        Или <span @click="setRegisterPage" style="color: #6C1BD2;">
          зарегистрируйтесь
        </span>
      </p>
      </b-form-group>

    </b-form>

  </div>
  </div>
</template>
<script>

import { validationMixin } from 'vuelidate'
import { required, minLength, email } from 'vuelidate/lib/validators'
import store from '@src/store/store'
import accountService from '@src/_services/account.service'

export default {
  mixins: [validationMixin],
  props : ['windowInnerWidth'],
  data () {
    return {
      userNotExist: false,
      email: '',
      password: '',
      store
    }
  },
  methods: {
    setRegisterPage () {
      if (this.windowInnerWidth > 600){
        this.$router.push('/register')
      }
      if (this.windowInnerWidth <= 600){
        this.$modal.show('modal--register')
      }
    },
    loginFB () {
      accountService.login()
    },
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
      minLength: minLength(3)
    }
  }
}
</script>
<style lang="scss">
.login__header{
  margin-top: 81px !important;
}
.login__form_block{
  display: flex;
  justify-content: space-between;
  margin-top: 46px;
  margin-bottom: -12px;
  & > .form-group{
      max-width: 276px;
    }
}
.login__form_block_text{
  margin-top: 40px;
}
@media (max-width: 600px) {
  .login__header{
    margin-top: 55px !important;
  }
  .login__form_block{
    flex-wrap: wrap;
    margin-top: 34px;
    width: 100%;
    & > .form-group{
      max-width: 100%;
    }
  }
  .login__form_block > #input-group-main:nth-child(2){
    margin-top: -8px;
  }
}
</style>
