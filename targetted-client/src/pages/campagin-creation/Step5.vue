<template>
   <div>
        <slot name="header"></slot>
        <b-form>
          
          <p>
            Выберите желаемый охват аудитории вашей рекламной кампании.</p>

        <CampaginStats/>

      <b-form-group
                    class="input-group input-group-range"
                    label="Бюджет в день"
                >
                <p class="app-label-right">{{campaginData.budget}} ₽ </p>
                    <b-form-input id="range-1" v-model="campaginData.budget" type="range" min="5" max="5000" step="5"></b-form-input>

     

           </b-form-group>
                 <b-form-group
                    class="input-group input-group-range"
                    label="Продолжительность"
                >
                <p class="app-label-right">{{campaginData.timeLength}} д</p>
                    <b-form-input id="range-2" v-model="campaginData.timeLength" type="range" min="1" max="31" step="1"></b-form-input>

    
           </b-form-group>



      <section v-if="!isRegistered">
        <h1 id="h2">Зарегистрируйтесь</h1>
        <p>Осталось только зарегистрироваться, чтобы запустить рекламу.</p>

            <b-form-group
                    class="input-group"
                    label="Номер телефона"
                >
                    <b-form-input
                    v-model="userData.phone"
                    class="form-input"
                    placeholder="Введите телефон"
                    />
           </b-form-group>
               <b-form-group
                    class="input-group"
                    label="Электронная почта"
                >
                    <b-form-input
                    v-model="userData.email"
                    class="form-input"
                    placeholder="Введите вашу почту"
                    />
           </b-form-group>
                <b-form-group
                    class="input-group"
                    label="Имя"
                >
                    <b-form-input
                    v-model="userData.name"
                    class="form-input"
                    placeholder="Введите ваше имя"
                    />
           </b-form-group>
    <b-form-group
                    class="input-group"
                    label="Пароль"
                >
                    <b-form-input
                    v-model="userData.password"
                    class="form-input"
                    type="password"
                    placeholder=""
                    />
           </b-form-group>
               <b-form-group
                    class="input-group"
                    label="Введите пароль повторно"
                >
                    <b-form-input
                    v-model="userData.passwordCheck"
                    class="form-input"
                    type="password"
                    placeholder=""
                    />
           </b-form-group>
            <b-button type="button"
                         v-if="!isRegistered"
                        class="app-new-submit-button"
                        @click="sendData">
        Зарегистрироваться
      </b-button>
         <p class="app-new-info">
          Нажимая кнопку “Зарегистрироваться” вы соглашаетесь с условиями <a href="#">оферты</a> и <a href="#">политикой конфиденциальности</a>.
        </p>
      </section>


        <b-button type="button"
                         v-if="isRegistered"
                        class="app-new-submit-button"
                        @click="sendData">
        {{isEdit ? "Назад":"Перейти к оплате"}}
      </b-button>
    </b-form>
  </div>
</template>

<script>
import store from '@/../store/store'
import CampaginStats from '@/components/CampaginStats'

export default {
  name: 'Step5',
    props: ['label_cols', 'content_cols', 'company', 'isEdit'],
    components: {
      CampaginStats
    },
    data: function () {
        return {
          store, //fixme
          isRegistered: false,
          campaginData: {
            budget: 100,
            timeLength: 6
          },
          userData: {
            phone:'',
            email:'',
            name: '',
            password: '',
            passwordCheck:''
          }
        }
      },
    methods: {
        sendData(){
            //todo
            this.$emit('next');
        }
    }
}
</script>

<style lang="scss">
  @import '@/../../../assets/vars.scss';
  .app-label-right {
    font-family: IBM Plex Sans;
    font-style: normal;
    font-weight: bold;
    font-size: 16px;
    line-height: 24px;
    position: absolute;
    top: 0px;
    right: 0px;
  }
  .app-new-stats {
     &-details {
        font-size: 16px;
        line-height: 24px;
        color: $gray;
      }
      &-type {
        font-weight: bold;
        font-size: 24px;
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
  box-shadow: 0px 0px 0px #000000;
  background: #F3F3F3;
  border-radius: 2px;
  border: 0px solid #000000;
}
input[type=range]::-webkit-slider-thumb {
  box-shadow: 0px 0px 0px #FFFFFF;
  border: 1px solid #E3E3E3;
  height: 21px;
  width: 21px;
  border-radius: 25px;
  background: #050608;
  cursor: pointer;
  -webkit-appearance: none;
  margin-top: -8.5px;
}
input[type=range]:focus::-webkit-slider-runnable-track {
  background: #F3F3F3;
}
input[type=range]::-moz-range-track {
  width: 100%;
  height: 5px;
  cursor: pointer;
  animate: 0.2s;
  box-shadow: 0px 0px 0px #000000;
  background: #F3F3F3;
  border-radius: 2px;
  border: 0px solid #000000;
}
input[type=range]::-moz-range-thumb {
  box-shadow: 0px 0px 0px #FFFFFF;
  border: 1px solid #E3E3E3;
  height: 21px;
  width: 21px;
  border-radius: 25px;
  background: #050608;
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
  background: #F3F3F3;
  border: 0px solid #000000;
  border-radius: 4px;
  box-shadow: 0px 0px 0px #000000;
}
input[type=range]::-ms-fill-upper {
  background: #F3F3F3;
  border: 0px solid #000000;
  border-radius: 4px;
  box-shadow: 0px 0px 0px #000000;
}
input[type=range]::-ms-thumb {
  margin-top: 1px;
  box-shadow: 0px 0px 0px #FFFFFF;
  border: 1px solid #E3E3E3;
  height: 21px;
  width: 21px;
  border-radius: 25px;
  background: #050608;
  cursor: pointer;
}
input[type=range]:focus::-ms-fill-lower {
  background: #F3F3F3;
}
input[type=range]:focus::-ms-fill-upper {
  background: #F3F3F3;
}
</style>