<template>
   <div>
        <slot name="header"></slot>
        <b-form>
            <p>
                Мы самостоятельно подберем наиболее подходящие настройки целевой аудитории на основе анализа большого количество рекламных кампаний из вашей ниши. Но вы можете уточнить настройки, если считаете это важным.
            </p>

            <b-row align-h="between" no-gutters>
                <b-col cols="5">
                    <b-button type="button"
                            style="margin: 0;padding: 10px;"
                            class="app-new-submit-button"
                            v-if="!showAdvancedClicked"
                            @click="sendData">
                    {{isEdit ? "Назад":"Продолжить"}}
                </b-button>
                </b-col>
        
                <b-col cols="7">
                    <p
                    class="app-new-action-text" 
                    style="margin-bottom: 0px;"
                    @click="showAdvanced = true; showAdvancedClicked = true"           v-if="!showAdvancedClicked">Уточнить настройки</p>
                </b-col>
            </b-row>
      
      <br>

      <div v-if="showAdvanced">

          <p>
              Укажите характеристики вашей аудитории, наши алгоритмы будут учитывать их при настройке рекламной кампании. 
          </p>

           <b-form-group
                label="Пол"
                :label-cols="label_cols"
                :content-cols="content_cols"
                id="input-group-main"
                label-for="input-horizontal"
                >
                    <b-form-radio-group
                        v-model="companyData.auditory.gender"
                        :options="auditory.gender"
                    ></b-form-radio-group>
                </b-form-group>

            <b-form-group label="Возраст"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                    id="input-group-main" label-for="input-horizontal">

                <b-row align-h="between">
                    <b-col cols="6">
                        <b-form-input
                            class="form-input app-new-form-input-small"
                            v-model="companyData.auditory.age.from" placeholder="От"></b-form-input>
                    </b-col>
            
                    <b-col cols="6">
                        <b-form-input
                            class="form-input app-new-form-input-small"
                            v-model="companyData.auditory.age.to" placeholder="До"></b-form-input>
                    </b-col>
                </b-row>
          </b-form-group>
           <b-form-group
                    class="input-group"
                    label="Местоположение"
                >
                    <b-form-input
                    id="tel"
                    v-model="companyData.auditory.location"
                    class="form-input"
                    placeholder="Введите адрес"
                    />
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
                        v-model="companyData.auditory.interests"
                        rows="3"
                        max-rows="6"
                        ></b-form-textarea>
           </b-form-group>
            
            <b-button type="button"
                        class="app-new-submit-button"
                        @click="sendData">
                {{isEdit ? "Назад":"Продолжить"}}
            </b-button>
      </div>
    </b-form>
  </div>
</template>

<script>
import store from '@/../store/store'

export default {
  name: 'Step3',
    props: ['label_cols', 'content_cols', 'company', 'isEdit'],
    data: function () {
        return {
          store, //fixme
          showAdvanced: false,
          showAdvancedClicked: false,
          companyData: {
              auditory: {
                  gender: '',
                  location: '',
                  interests: '',
                  age: {}
              }
          },
          auditory: {
              gender: [
                    { text: 'Мужской', value: 'male' },
                    { text: 'Женский', value: 'female' },
                    { text: 'Любой', value: 'any' },
              ]
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

<style>

</style>