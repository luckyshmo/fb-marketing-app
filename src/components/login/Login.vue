<template>
    <div id="content">
        <h1 id="h1">Вход</h1>

        <b-form id="form-centred" @submit="onSubmit">
            <b-form-group
                id="input-group"
            >
                <b-form-input
                id="form-input"
                v-model="form.email"
                type="email"
                placeholder="Enter email"
                required
                ></b-form-input>
            </b-form-group>

            <b-form-group 
                id="input-group"
            >
                <b-form-input
                id="form-input"
                v-model="form.password"
                placeholder="Enter password"
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
        alert(JSON.stringify(this.form))
      },
    }
  }
</script>
<style>
#form-centred{
    text-align: center !important;
}
#content {
    width:600px;
    margin:0 auto !important;
}
#h1 {
    text-align: center;
}
#input-group{
    margin: 35px 20px 35px 20px;
}
#form-input{
    border-radius: 8px;
    border: none;
    background-color: #e4e4e4;
    color: #767676;
}
</style>