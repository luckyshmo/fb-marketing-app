<template>
    <div id="content-login">
        <h1 id="h1-centered">Вход</h1>

        <b-form id="form-centred" @submit="onSubmit">
            <b-form-group
                id="input-group"
            >
                <b-form-input
                id="form-input"
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
                id="form-input"
                v-model="form.password"
                placeholder="Пароль"
                required
                ></b-form-input>
            </b-form-group>

            <b-button type="submit" id="main-button">Войти</b-button>
            <router-link :to="{name: 'register'}">
                <p id="navigation-text">Регистрация</p>
            </router-link>
        </b-form>
            <b-card style="color:black" class="mt-3" header="Form Data Result">
            <pre class="m-0">{{ form }}</pre>
        </b-card>
    </div>
</template>
<script>
import router from '../../../router/router'
import store from '../../../store/store'
  export default {
    data() {
      return {
        form: {
          email: '',
          password: '',
        },
      }
    },
    methods: {
      onSubmit(event) {
        event.preventDefault()
        store.dispatch('setAuth', true);
        console.log(store.getters.IS_AUTH)
        router.push('main')
        // alert(JSON.stringify(this.form))
      },
    }
  }
</script>
<style>
</style>