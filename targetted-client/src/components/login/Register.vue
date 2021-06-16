<template>
    <div id="content-login">
        <h1 id="h1-centered">Регистрация</h1>

        <b-form id="form-centred" @submit="register">
          <b-form-group 
            class="input-group"
          >
            <b-form-input
            class="form-input"
            v-model="form.name"
            placeholder="Имя"
            :state="validateState('name')"
            ></b-form-input>
            <small class="error-message" 
            v-if="$v.form.name.$dirty && !$v.form.name.required">
              Пустое поле Имя
            </small>
          </b-form-group>

          <b-form-group 
            class="input-group"
          >
            <b-form-input
            class="form-input"
            id="tel"
            type='tel'
            v-model="form.phoneNumber"
            :state="validateState('phoneNumber')"
            placeholder="Номер телефона"
            ></b-form-input>
            <b-form-invalid-feedback 
            class="error-message" 
            id="tel">
              Номер телефона должен содержать как минимум 10 символов
            </b-form-invalid-feedback>
          </b-form-group>

          <b-form-group
            class="input-group"
          >
            <b-form-input
            @click="resetErr()"
            class="form-input"
            v-model.trim="form.email"
            placeholder="Электронная почта"
            :state="validateState('email')"
            ></b-form-input>
            <small class="error-message" 
            v-if="$v.form.email.$dirty && !$v.form.email.required">
              Пустое поле email
            </small>
            <small class="error-message"
              v-if="$v.form.email.$dirty && !$v.form.email.email">
              Некрорректный email
            </small>
            <small class="error-message"
              v-if="emailExists">
              Пользователь с таким email уже существует
            </small>
          </b-form-group>

          <b-form-group 
              class="input-group"
          >
              <b-form-input
              class="form-input"
              type='password'
              id="pas"
              v-model="form.password"
              :state="validateState('password')"
              placeholder="Пароль"
              ></b-form-input>
              <b-form-invalid-feedback 
              class="error-message" 
              id="pas">
                Пароль должен быть минимум 8 символов
              </b-form-invalid-feedback>
          </b-form-group>
          <b-form-group 
              class="input-group"
          >
              <b-form-input
              type='password'
              class="form-input"
              aria-describedby="password-repeat-feedback"
              v-model="form.passwordRepeat"
              :state="validatePassword('passwordRepeat')"
              placeholder="Введите пароль повторно"
              ></b-form-input>
              <small class="error-message"
                v-if="form.password != form.passwordRepeat">
                Пароли не совпадают
              </small>
          </b-form-group>

          <b-button type="submit" class="main-button-big">Зарегистрироваться</b-button>
          
          <p id="navigation-text">
            или <router-link style="color: #6C1BD2" :to="{name: 'login'}">войдите</router-link> в свой аккаунт
          </p>
          
        </b-form>
    </div>
</template>
<script>
  import router from '../../../router/router'
  import axios from 'axios'
  const VUE_APP_API_URL = process.env.VUE_APP_API_URL;
  import { validationMixin } from "vuelidate";
  import { required, minLength, email } from "vuelidate/lib/validators";
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
        },
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
      resetErr(){
        this.emailExists = false;
      },
      validateState(name) {
        const { $dirty, $error } = this.$v.form[name];
        return $dirty ? !$error : null;
      },
      validatePassword(name){
        const { $dirty, $error } = this.$v.form[name];
        console.log($dirty, $error, this.form.password === this.form.passwordRepeat)
        if ($dirty ? !$error : null != null){
          console.log("some sta")
          return this.form.password === this.form.passwordRepeat
        } else {
          return $dirty ? !$error : null;
        }
      },
      register: function () {

        this.$v.form.$touch();
        if (this.$v.form.$anyError) {
          return;
        }

        let data = {
            name: this.form.name,
            email: this.form.email,
            password: this.form.password,
            phoneNumber: this.form.phoneNumber
        }
        console.log("rData", data)
        axios({url: `${VUE_APP_API_URL}/auth/sign-up`, data: data, method: 'POST'})
        .then(resp => {
            this.emailExists = false;
            console.log(resp)
            router.push('login')
        })
        .catch(err => {
            console.log(err.response)
            if (err.response.data.message === 'pq: duplicate key value violates unique constraint "user_tb_email_key"') {
              this.emailExists = true;
            }
          }
        )
      },
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