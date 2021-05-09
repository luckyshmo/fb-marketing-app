<template>
    <div id="content-login">
        <h1 id="h1-centered">Регистрация</h1>

        <b-form id="form-centred" @submit="onSubmit">
          <b-form-group 
            id="input-group"
          >
            <b-form-input
            id="form-input"
            v-model="form.name"
            placeholder="Имя"
            required
            ></b-form-input>
          </b-form-group>

          <b-form-group 
            id="input-group"
          >
            <b-form-input
            id="form-input"
            v-model="form.phoneNumber"
            placeholder="Номер телефона"
            required
            ></b-form-input>
          </b-form-group>

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

          <b-button type="submit" id="main-button">Регистрация</b-button>
          <router-link :to="{name: 'login'}">
              <p id="navigation-text">Войти</p>
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
          name: '',
          phoneNumber: '',
          email: '',
          password: '',
        },
      }
    },
    methods: {
      onSubmit(event) {
        event.preventDefault()
        alert(JSON.stringify(this.form))
        store.dispatch('setAuth', false);
        router.push({ name: 'login', params: { email: this.form.email } }) //TODO
        // router.push('login')
      },
    }
  }
</script>
<style>
</style>