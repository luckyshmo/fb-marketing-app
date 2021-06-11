<template>
    <div id="content-login">
        <h1 id="h1-centered">Регистрация</h1>

        <b-form id="form-centred" @submit="register">
          <b-form-group 
            id="input-group"
          >
            <b-form-input
            class="form-input"
            v-model="form.name"
            placeholder="Имя"
            required
            ></b-form-input>
          </b-form-group>

          <b-form-group 
            id="input-group"
          >
            <b-form-input
            class="form-input"
            type='tel'
            v-model="form.phoneNumber"
            placeholder="Номер телефона"
            required
            ></b-form-input>
          </b-form-group>

          <b-form-group
            id="input-group"
          >
            <b-form-input
            class="form-input"
            v-model="form.email"
            type="email"
            placeholder="Электронная почта"
            required
            ></b-form-input>
          </b-form-group>

          <b-form-group 
              id="input-group"
          >
              <b-form-input
              class="form-input"
              type='password'
              v-model="form.password"
              placeholder="Пароль"
              required
              ></b-form-input>
          </b-form-group>
          <b-form-group 
              id="input-group"
          >
              <b-form-input
              type='password'
              class="form-input"
              v-model="form.passwordRepeat"
              placeholder="Введите пароль повторно"
              required
              ></b-form-input>
          </b-form-group>

          <b-button type="submit" id="main-button">Зарегистрироваться</b-button>
          
          <p id="navigation-text">
            или <router-link style="color: #6C1BD2" :to="{name: 'login'}">войдите</router-link> в свой аккаунт
          </p>
          
        </b-form>
    </div>
</template>
<script>
  import router from '../../../router/router'
  import store from '../../../store/store'
  export default {
    data() {
      return {
        form: {
          name: '',
          phoneNumber: '',
          email: '',
          password: '',
          passwordrepeat: ''
        },
      }
    },
    methods: {
      register: function () {
        let data = {
            name: this.form.name,
            email: this.form.email,
            password: this.form.password,
            phoneNumber: this.form.phoneNumber
        }
        store.dispatch('register', data)
        .then((resp) => {
          console.log(resp)
          router.push('login')
        }
        )
        .catch(err => console.log(err))
      },
    }
  }
</script>
<style>
</style>