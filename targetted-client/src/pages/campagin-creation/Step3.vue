<template>
   <div class="step">
        <slot name="header"></slot>
        <b-form>
            <p>
             Наш сервис проанализирует рекламные кампании похожих бизнесов из вашей ниши и автоматически подберет оптимальные настройки целевой аудитории. Или вы можете самостоятельно указать настройки аудитории, если считаете это важным.
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
                       class="main-button-grey ml-1 ml-sm-1 ml-lg-1 ml-md-1 mt-lg-0 mt-md-0 mt-sm-0 mt-0"
                    @click="showAdvanced = true; showAdvancedClicked = true"
                    v-if="!showAdvancedClicked">Уточнить настройки</b-button>
                </b-col>
            </b-row>

      <br>

      <template v-if="showAdvanced">

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
                        v-model="$v.campaignData.auditory.gender.$model"
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
                            v-model="$v.campaignData.auditory.age.from.$model"
                            :state="$v.campaignData.auditory.age.from.$dirty ? !$v.campaignData.auditory.age.from.$error : null"
                            placeholder="от"></b-form-input>
                                 <small
              v-if="$v.campaignData.auditory.age.from.$dirty && !$v.campaignData.auditory.age.from.required"
              class="error-message"
            >
              Пустое поле
            </small>
                      <small
                  v-if="$v.campaignData.auditory.age.from.$dirty && !$v.campaignData.auditory.age.from.numeric"
                  class="error-message"
                >
                  Только цифры
                </small>
                    </b-col>

                    <b-col cols="6" sm="5" md="6" lg="3" xl="3">
                        <b-form-input
                            class="form-input app-new-form-input-small"
                            v-model="$v.campaignData.auditory.age.to.$model"
                            :state="$v.campaignData.auditory.age.to.$dirty ? !$v.campaignData.auditory.age.to.$error : null"
                            placeholder="до"></b-form-input>
                             <small
              v-if="$v.campaignData.auditory.age.to.$dirty && !$v.campaignData.auditory.age.to.required"
              class="error-message"
            >
              Пустое поле
            </small>
                  <small
                  v-if="$v.campaignData.auditory.age.to.$dirty && !$v.campaignData.auditory.age.to.numeric"
                  class="error-message"
                >
                  Только цифры
                </small>
                    </b-col>
                </b-row>
          </b-form-group>

           <b-form-group
                    label="Местоположение"
                     id="input-group-main">
                    <b-form-input
                    :state="$v.campaignData.auditory.location.$dirty ? !$v.campaignData.auditory.location.$error : null"
                    v-model="$v.campaignData.auditory.location.$model"
                    class="form-input"
                    placeholder="Введите адрес"
                    />
                    <small
              v-if="$v.campaignData.auditory.location.$dirty && !$v.campaignData.auditory.location.required"
              class="error-message"
            >
              Пустое поле
            </small>
           </b-form-group>

           <b-form-group
                    id="input-group-main"
                    label="Интересы"
                >
                   <p class="app-new-info">
                        Коротко опишите интересы и прочие характеристики ваших клиентов.
                    </p>
                      <b-form-textarea
                        id="textarea"
                        class="form-input"
                        :state="$v.campaignData.auditory.interests.$dirty ? !$v.campaignData.auditory.interests.$error : null"
                        v-model="$v.campaignData.auditory.interests.$model"
                        rows="3"
                        max-rows="6"
                        ></b-form-textarea>
                           <small
              v-if="$v.campaignData.auditory.interests.$dirty && !$v.campaignData.auditory.interests.required"
              class="error-message"
            >
              Пустое поле
            </small>
           </b-form-group>

                      <b-row class="bottom__block">
                <b-col cols="6" sm="8" md="3" lg="3" xl="3">
            <b-button type="button"
                        class="app-new-submit-button"
                        :class="{'disabled-look': !$v.$anyDirty || $v.$anyError}"
                        @click="sendData">
                {{isEdit ? "Назад" : "Продолжить"}}
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
      </template>

    </b-form>
  </div>
</template>

<script>
import store from '@src/store/store'
import { validationMixin } from 'vuelidate'
import { required, minValue, maxValue, numeric } from 'vuelidate/lib/validators'

export default {
  name: 'Step3',
  props: ['label_cols', 'content_cols', 'campaign', 'isEdit'],
  mixins: [validationMixin],
  validations: {
    campaignData: {
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
            numeric,
            minValue: minValue(0)
          },
          to: {
            required,
            numeric,
            maxValue: maxValue(99)
          }
        }
      }
    }
  },
  data: function () {
    return {
      store, // fixme
      showAdvanced: false,
      noValidate: false,
      showAdvancedClicked: false,
      campaignData: {
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
  mounted () {
    // todo copy from Step1

    // auditory
  },
  methods: {
    sendDataNoValidate () {
      this.noValidate = true
      this.sendData()
    },
    sendData () {
      if (this.showAdvancedClicked && !this.noValidate) {
        this.$v.campaignData.auditory.$touch()
        if (this.$v.campaignData.auditory.$anyError) {
          window.scrollTo(0, 100)
          return
        }
      }
      this.$router.push('/campaign-step-5')

      // this.$emit('next', this.campaignData)
    }
  }
}
</script>
<style>
.bottom__block{
  margin-bottom: 200px;
}
</style>
