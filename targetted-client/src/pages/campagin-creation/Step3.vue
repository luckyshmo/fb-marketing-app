<template>
   <div>
        <slot name="header"></slot>
        <b-form>
            <p>
                Мы самостоятельно подберем наиболее подходящие настройки целевой аудитории на основе анализа большого количество рекламных кампаний из вашей ниши. Но вы можете уточнить настройки, если считаете это важным.
            </p>

            <b-row align-h="between" no-gutters>

                <b-col cols="10">
                    <b-button type="button"
                            class="app-new-submit-button"
                            v-if="!showAdvancedClicked"
                            @click="sendData">
                    {{isEdit ? "Назад":"Продолжить"}}
                </b-button>
         
                    <b-button
                    type="button"
                       class="main-button-grey ml-0 ml-sm-2 ml-lg-2 ml-md-2 mt-lg-0 mt-md-0 mt-sm-0 mt-2"
                    @click="showAdvanced = true; showAdvancedClicked = true"
                    v-if="!showAdvancedClicked">Уточнить настройки</b-button>
                </b-col>
            </b-row>
      
      <br>

      <div v-if="showAdvanced">

          <h3> Уточнение аудитории</h3>

          <!-- <p>
              Укажите характеристики вашей аудитории, наши алгоритмы будут учитывать их при настройке рекламной кампании. 
          </p> -->

           <b-form-group
                label="Пол"
                :label-cols="label_cols"
                :content-cols="content_cols"
                id="input-group-main"
                label-for="input-horizontal"
                >
                    <b-form-radio-group
                        class="app-new-radio"
                        v-model="$v.companyData.auditory.gender.$model"
                        :options="auditory.gender"
                    ></b-form-radio-group>
                </b-form-group>

            <b-form-group label="Возраст"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                    id="input-group-main" label-for="input-horizontal">

                <b-row >
                    <b-col cols="6" sm="5" md="6" lg="3" xl="3">
                        <b-form-input
                            class="form-input app-new-form-input-small "
                            v-model="$v.companyData.auditory.age.from.$model"
                            placeholder="от"></b-form-input>
                                 <small
              v-if="$v.companyData.auditory.age.from.$dirty && !$v.companyData.auditory.age.from.required"
              class="error-message"
            >
              Пустое поле 
            </small>
                    </b-col>
            
                    <b-col cols="6" sm="5" md="6" lg="3" xl="3">
                        <b-form-input
                            class="form-input app-new-form-input-small"
                            v-model="$v.companyData.auditory.age.to.$model"
                            placeholder="до"></b-form-input>
                             <small
              v-if="$v.companyData.auditory.age.to.$dirty && !$v.companyData.auditory.age.to.required"
              class="error-message"
            >
              Пустое поле 
            </small>
                    </b-col>
                </b-row>
          </b-form-group>
           <b-form-group
                    class="input-group"
                    label="Местоположение"
                     id="input-group-main">
                    <b-form-input
                    id="tel"
                    v-model="$v.companyData.auditory.location.$model"
                    class="form-input"
                    placeholder="Введите адрес"
                    />
                    <small
              v-if="$v.companyData.auditory.location.$dirty && !$v.companyData.auditory.location.required"
              class="error-message"
            >
              Пустое поле 
            </small>
           </b-form-group>
           <b-form-group
                    class="input-group"
                    label="Интересы"
                >
                   <p class="app-new-info">
                        Коротко опишите интересы и прочие характеристики ваших клиентов.
                    </p>
                      <b-form-textarea
                        id="textarea"
                        class="form-input"
                        v-model="$v.companyData.auditory.interests.$model"
                        rows="3"
                        max-rows="6"
                        ></b-form-textarea>
                           <small
              v-if="$v.companyData.auditory.interests.$dirty && !$v.companyData.auditory.interests.required"
              class="error-message"
            >
              Пустое поле 
            </small>
           </b-form-group>

                      <b-row>
                <b-col cols="6" sm="8" md="3" lg="3" xl="3">
            <b-button type="button"
                        class="app-new-submit-button"
                        :class="{'disabled': !$v.$anyDirty || $v.$anyError}"
                        @click="sendData">
                {{isEdit ? "Назад":"Продолжить"}}
            </b-button>
 </b-col>
                <b-col cols="10" sm="8" md="6" lg="6" xl="6" >
            <b-button
                type="button"
                    class="main-button-grey mt-md-0 mt-sm-0 mt-2"
                @click="sendDataNoValidate"
                >Продолжить без уточнения</b-button>

                    </b-col>
          </b-row>
      </div>
      
    </b-form>
  </div>
</template>

<script>
import store from '@/store/store'
import { validationMixin } from 'vuelidate'
import { required, minValue, maxValue } from 'vuelidate/lib/validators'

export default {
  name: 'Step3',
    props: ['label_cols', 'content_cols', 'company', 'isEdit'],
    mixins: [validationMixin],
    validations: {
         companyData: {
              auditory: {
                  gender: {
                      required
                  },
                  location: {
                      required
                  },
                  interests: {
                      required
                  },
                  age: {
                      from: {
                          required,
                          minValue: minValue(0)
                      },
                      to: {
                          required,
                          maxValue: maxValue(99)
                      }
                  }
              }
          },
    },
    data: function () {
        return {
          store, //fixme
          showAdvanced: false,
          noValidate: false,
          showAdvancedClicked: false,
          companyData: {
              auditory: {
                  gender: 'any',
                  location: '',
                  interests: '',
                  age: {
                      from: '',
                      to: ''
                  }
              }
          },
          auditory: {
              gender: [
                    { text: 'Любой', value: 'any' },
                    { text: 'Мужской', value: 'male' },
                    { text: 'Женский', value: 'female' }
              ]
          }
        }
      },
    methods: {
        sendDataNoValidate(){
            this.noValidate = true;
            this.sendData();
        },
        sendData(){
            if(this.showAdvancedClicked && !this.noValidate) {
                this.$v.companyData.auditory.$touch()
                if (this.$v.companyData.auditory.$anyError) {
                    window.scrollTo(0, 100)
                    return
                }
            }

            this.$emit('next', this.companyData);
        }
    }
}
</script>