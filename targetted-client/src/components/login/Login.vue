<template>
    <div id="content-login">
        <h1 id="h1-centered">Вход</h1>

        <b-form id="form-centred" @submit.prevent="login">
            <b-form-group
                id="input-group"
            >
                <b-form-input
                class="form-input"
                v-model="email"
                type="email"
                placeholder="Электронная почта"
                required
                ></b-form-input>
            </b-form-group>

            <b-form-group 
                id="input-group"
            >
                <b-form-input
                type='password'
                class="form-input"
                v-model="password"
                placeholder="Пароль"
                required
                ></b-form-input>
            </b-form-group>
            <b-button type="submit" id="main-button">Войти</b-button>
            
            <p id="navigation-text">или <router-link style="color: #6C1BD2" :to="{name: 'register'}">заргестрироваться</router-link></p>
            
        </b-form>
    </div>
</template>
<script>
import {minLength} from 'vuelidate/lib/validators'
import router from '../../../router/router'
import store from '../../../store/store'
  export default {
    data() {
      return {
        email: '',
        password: '',
      }
    },
    methods: {
      login: function () {
        let email = this.email 
        let password = this.password
        store.dispatch('login', { email, password })
        .then(() => router.push('main'))
        .catch(err => console.log(err)) 
      },
    },
    validations: {
      password: {
        minLength: minLength(8)
      }
    }
  }
</script>
<style>
</style>