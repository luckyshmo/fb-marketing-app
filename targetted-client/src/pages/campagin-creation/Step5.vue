<template>
   <div>
        <slot name="header"></slot>

     <p>Выберите желаемый охват аудитории вашей рекламной кампании.</p>

      <br class="d-none d-sm-block">
      <br class="d-none d-sm-block">

        <b-form>

        <CampaginStats :data="stats"/>

    <b-row>
      <b-col cols="12" sm="6">
      <b-form-group
            class="input-group-range"
            :label="`Бюджет в день: ${campaginData.budget} ₽`">
        <!-- <p class="app-label-right">{{campaginData.budget}} ₽ </p> -->
            <b-form-input id="range-1" v-model="campaginData.budget" type="range" min="5" max="5000" step="5"></b-form-input>

           </b-form-group>
      </b-col>
      <b-col cols="12" sm="6">
            <b-form-group
              class="input-group-range"
              :label="`Продолжительность: ${campaginData.timeLength} д`">
          <!-- <p class="app-label-right">{{campaginData.timeLength}} д</p> -->
              <b-form-input id="range-2" v-model="campaginData.timeLength" type="range" min="1" max="31" step="1"></b-form-input>
           </b-form-group>
      </b-col>
    </b-row>

      <section v-if="!isRegistered">
        <h3 style="font-weight: bold;">Зарегистрируйтесь</h3>
        <p>Осталось только зарегистрироваться, чтобы запустить рекламу.</p>

            <b-form-group
             id="input-group-main"
                    label="Номер телефона">
                    <b-form-input
                    :state="$v.userData.phone.$dirty ? !$v.userData.phone.$error : null"
                    v-model="$v.userData.phone.$model"
                    class="form-input width-1-2"
                    placeholder="Введите телефон"
                    />
                        <small
                  v-if="$v.userData.phone.$dirty && !$v.userData.phone.required"
                  class="error-message"
                >
                  Телефон на указан
                </small>
                        <small
                  v-if="$v.userData.phone.$dirty && !$v.userData.phone.numeric"
                  class="error-message"
                >
                  Только цифры в телефоне
                </small>
           </b-form-group>

  <b-form-group
                 id="input-group-main"
                    label="Имя">
                    <b-form-input
                    v-model="$v.userData.name.$model"
                    class="form-input"
                    placeholder="Введите ваше имя"
                    />
                     <small
                      v-if="$v.userData.name.$dirty && !$v.userData.name.required"
                      class="error-message"
                    >
                      Имя не указанно
                    </small>
                             <small
                  v-if="$v.userData.name.$dirty && !$v.userData.name.alpha"
                  class="error-message"
                >
                  Только буквы в имени
                </small>
           </b-form-group>

               <b-form-group
                id="input-group-main"
                    label="Электронная почта">
                    <b-form-input
                    v-model="$v.userData.email.$model"
                    :state="$v.userData.email.$dirty ? !$v.userData.email.$error : null"
                    class="form-input"
                    placeholder="Введите вашу почту"
                    />
                    <small
                  v-if="$v.userData.email.$dirty && !$v.userData.email.required"
                  class="error-message"
                >
                  Пустое поле email
                </small>
                <small
                  v-if="$v.userData.email.$dirty && !$v.userData.email.email"
                  class="error-message"
                >
                  Некорректный email
                </small>
           </b-form-group>

          <b-form-group
           id="input-group-main"
                    label="Пароль">
                    <b-form-input
                    v-model="$v.userData.password.$model"
                    :state="$v.userData.password.$dirty ? !$v.userData.password.$error : null"
                    class="form-input width-1-2"
                    type="password"
                    placeholder=""
                    />
                      <small
                      v-if="$v.userData.password.$dirty && !$v.userData.password.required"
                      class="error-message"
                    >
                      Пароль не указан
                    </small>
                      <small
                      v-if="$v.userData.password.$dirty && !$v.userData.password.minLength"
                      class="error-message"
                    >
                      Пароль менее 8 символов
                    </small>
           </b-form-group>
            <b-form-group
                    label="Введите пароль повторно"
                     id="input-group-main"
                >
                    <b-form-input
                    v-model="$v.userData.passwordCheck.$model"
                    :state="$v.userData.passwordCheck.$dirty ? !$v.userData.passwordCheck.$error : null"
                    class="form-input width-1-2"
                    type="password"
                    placeholder=""
                    />
                    <small
                      v-if="$v.userData.passwordCheck.$dirty && !$v.userData.passwordCheck.required"
                      class="error-message"
                    >
                      Пароль не указан
                    </small>
                      <small
                      v-if="$v.userData.passwordCheck.$dirty && !$v.userData.passwordCheck.minLength"
                      class="error-message"
                    >
                      Пароль менее 8 символов
                    </small>
           </b-form-group>
            <b-button type="button"
                         v-if="!isRegistered"
                         :class="{'disabled-look': !$v.$anyDirty || $v.$anyError}"
                        class="app-new-submit-button"
                        @click="sendData">
        Зарегистрироваться
      </b-button>
         <p class="app-new-info mt-5 mt-sm-2">
          Нажимая кнопку “Зарегистрироваться” вы соглашаетесь с условиями <a href="#">оферты</a> и <a href="#">политикой конфиденциальности</a>.
        </p>
      </section>

        <b-button type="button"
        :class="{'disabled-look': !$v.$anyDirty || $v.$anyError}"
                         v-if="isRegistered"
                        class="app-new-submit-button"
                        @click="sendData">
        {{isEdit ? "Назад":"Перейти к оплате"}}
      </b-button>
    </b-form>
  </div>
</template>

<script>
import store from '@src/store/store'
import CampaginStats from '@src/components/CampaginStats'

import { required, email, minLength, numeric, alpha } from 'vuelidate/lib/validators'
import { validationMixin } from 'vuelidate'

// todo: load from backend
const stats = {
  seen: { min: 5, max: 5000 },
  clicks: { min: 67, max: 1300 },
  costs: 5000
}

export default {
  name: 'Step5',
  props: ['label_cols', 'content_cols', 'company', 'isEdit'],
  components: {
    CampaginStats
  },
  mixins: [validationMixin],
  validations: {
    userData: {
      email: {
        email,
        required
      },
      password: {
        required,
        minLength: minLength(8)
      },
      passwordCheck: {
        required,
        minLength: minLength(8)
      },
      name: {
        required,
        alpha,
        minLength: minLength(2)
      },
      phone: {
        required,
        numeric
        // minLength: minLength(9)
      }
    }
  },
  data: function () {
    return {
      stats,
      store, // fixme
      isRegistered: false,
      campaginData: {
        budget: 100,
        timeLength: 6
      },
      // todo: mv to separate component ?
      userData: {
        phone: '',
        email: '',
        name: '',
        password: '',
        passwordCheck: ''
      }
    }
  },

  mounted () {
    // todo copy from Step1

    // userData
    // campaginData
  },
  methods: {
    sendData () {
      this.$v.userData.$touch()
      if (this.$v.userData.$anyError) {
        window.scrollTo(0, 100)
        return
      }

      this.$emit('next', this.campaginData)
    }
  }
}
</script>

<style lang="scss">
  @import '@src/assets/styles/vars.scss';

  .app-label-right {
    font-family: IBM Plex Sans;
    font-style: normal;
    font-weight: bold;
    font-size: $baseFont;
    line-height: $baseLH;
    position: absolute;
    top: 0;
    right: 0;
  }
  .app-new-stats {
     &-details {
        font-size: $baseFont;
        line-height: $baseLH;
        color: $gray;
        margin-bottom: 27px;
      }
      &-type {
        font-weight: bold;
        font-size: $baseLH;
        line-height: 30px;
      }
  }
  .input-group.input-group-range > div{
    width: 100%;
  }
  .input-group-range {
    label {
      width: 80%;
    }
  }

@media (min-width: 550px) {
  .input-group.input-group-range {
    margin-top: 14px !important;
  }
  .app-label-right {
    display: inline-block;
  }
  legend {
    width: auto;
  }
}

  input[type=range] {
  height: 28px;
  -webkit-appearance: none;
  margin: 10px 0;
  width: 100%;
}

input[type=range]:focus {
  outline: none;
}

input[type=range]::-webkit-slider-runnable-track {
  width: 100%;
  height: 5px;
  cursor: pointer;
  animate: 0.2s;
  box-shadow: 0 0 0 $black;
  background: $light;
  border-radius: 2px;
  border: 0 solid  $black;
}

input[type=range]::-webkit-slider-thumb {
  box-shadow: 0 0 0 $white;
  border: 1px solid $light;
  height: 32px;
  width: 32px;
  border-radius: 25px;
  background: $black;
  cursor: pointer;
  -webkit-appearance: none;
  margin-top: -13px;
}

input[type=range]:focus::-webkit-slider-runnable-track {
  background: $light;
}

input[type=range]::-moz-range-track {
  width: 100%;
  height: 5px;
  cursor: pointer;
  animate: 0.2s;
  box-shadow: 0 0 0 $black;
  background: $light;
  border-radius: 2px;
  border: 0 solid $black;
}

input[type=range]::-moz-range-thumb {
  box-shadow: 0 0 0 $white;
  border: 1px solid $light;
  height: 32px;
  width: 32px;
  border-radius: 32px;
  background: $black;
  cursor: pointer;
}

input[type=range]::-ms-track {
  width: 100%;
  height: 5px;
  cursor: pointer;
  animate: 0.2s;
  background: transparent;
  border-color: transparent;
  color: transparent;
}

input[type=range]::-ms-fill-lower {
  background: $light;
  border: 0 solid $black;
  border-radius: 4px;
  box-shadow: 0 0 0 $black;
}

input[type=range]::-ms-fill-upper {
  background: $light;
  border: 0 solid $black;
  border-radius: 4px;
  box-shadow: 0 0 0 $black;
}

input[type=range]::-ms-thumb {
  margin-top: 1px;
  box-shadow: 0 0 0 $white;
  border: 1px solid #E3E3E3;
  height: 32px;
  width: 32px;
  border-radius: 32px;
  background: $black;
  cursor: pointer;
}

input[type=range]:focus::-ms-fill-lower {
  background: $light;
}

input[type=range]:focus::-ms-fill-upper {
  background: $light;
}
</style>
