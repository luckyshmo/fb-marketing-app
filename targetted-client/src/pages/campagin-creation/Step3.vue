<template>
   <div class="step">
        <slot name="header"></slot>
        <b-form>
            <p style="margin-bottom: 20px">
              Наш сервис проанализирует рекламные кампании похожих бизнесов из вашей ниши и автоматически подберет оптимальные настройки целевой аудитории. Или вы можете самостоятельно указать настройки аудитории, мы используем эту информацию при настройке.
            </p>

            <b-row align-h="between" no-gutters>

                <b-col cols="10">
                  <div class="step2__button">
                    <b-button type="button"
                              class="app-new-submit-button"
                              style="margin-right: 12px"
                              v-if="!showAdvancedClicked"
                              @click="sendData">
                      {{isEdit ? "Назад":"Продолжить"}}
                    </b-button>
                  </div>

                    <b-button
                    type="button"
                       class="main-button-grey"
                    @click="showAdvanced = true; showAdvancedClicked = true"
                    v-if="!showAdvancedClicked">Уточнить настройки</b-button>
                </b-col>
            </b-row>

      <template v-if="showAdvanced">

          <h3 style="margin-top: 20px;
            font-style: normal;
            font-weight: bold;
            font-size: 24px;
            line-height: 30px;" class="step3__text">Уточнение аудитории</h3>

          <!-- <p>
              Укажите характеристики вашей аудитории, наши алгоритмы будут учитывать их при настройке рекламной кампании.
          </p> -->

           <b-form-group
                label="Пол"
                :label-cols="label_cols"
                :content-cols="content_cols"
                id="input-group-main"
                label-for="input-horizontal"
                style="margin-top: 10px; margin-bottom: 7px !important;"
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
                          style="margin-bottom: 33px !important;"
                    id="input-group-main" label-for="input-horizontal">
                <b-row >
                    <b-col class="step3__input">
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

                    <b-col class="step3__input">
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

<!--           <b-form-group-->
<!--                    label="Местоположение"-->
<!--                     id="input-group-main">-->
<!--                    <b-form-input-->
<!--                    :state="$v.campaignData.auditory.location.$dirty ? !$v.campaignData.auditory.location.$error : null"-->
<!--                    v-model="$v.campaignData.auditory.location.$model"-->
<!--                    class="form-input"-->
<!--                    placeholder="Введите адрес"-->
<!--                    />-->
<!--                    <small-->
<!--              v-if="$v.campaignData.auditory.location.$dirty && !$v.campaignData.auditory.location.required"-->
<!--              class="error-message"-->
<!--            >-->
<!--              Пустое поле-->
<!--            </small>-->
<!--           </b-form-group>-->

           <b-form-group
                    id="input-group-main"
                    label="Интересы"
                >
                   <p class="app-new-info" style="margin-top: -7px">
                     Коротко опишите интересы и другие характеристики ваших клиентов, которые вы считаете основными.
                    </p>
                      <b-form-textarea
                        id="textarea"
                        class="form-input"
                        :state="$v.campaignData.auditory.interests.$dirty ? !$v.campaignData.auditory.interests.$error : null"
                        v-model="$v.campaignData.auditory.interests.$model"
                        rows="3"
                        max-rows="6"
                        style="min-height: 132px ; margin-top: 12px"
                        placeholder="Пример: девушки с доходом выше среднего, занимаются спортом и следят за своей фигурой"
                        ></b-form-textarea>
                           <small
              v-if="$v.campaignData.auditory.interests.$dirty && !$v.campaignData.auditory.interests.required"
              class="error-message"
            >
              Пустое поле
            </small>
           </b-form-group>

        <div class="bottom__block">
          <div style="max-width: 145px; margin-right: 12px" class="step2__button">
            <b-button  type="button"
                       class="app-new-submit-button"
                       @click="sendData">
              {{isEdit ? "Назад":"Продолжить"}}
            </b-button>
          </div>
          <div  class="app-new-small-text-fit" style="max-width: 259px">
            <b-button
              type="button"
              class="main-button-grey"
              @click="sendDataNoValidate"
            >Продолжить без уточнения</b-button>
          </div>
        </div>
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
<style lang="scss">
.step2__button{
  display: inline-block;
}
.step3__input{
  max-width: 136px !important;
  width: 100% !important;
  & > input{
    max-width: 100%;
  }
}
.step3__input:nth-child(1){
  padding-right: 0 !important;
  margin-right: 12px;
}
.step3__input:nth-child(2){
  padding-left: 0 !important;
}
@media (max-width: 600px) {
  .app-new-small-text-fit > button{
    margin-left: 0 !important;
  }
  .step2__button{
    width: 100%;
    margin-bottom: 12px;
  }
  .step3__text{
    margin-top: -4px !important;
  }
  .step3__input{
    max-width: 100% !important;
  }
  .step3__input:nth-child(1){
    margin-right: 8px;
  }
  /*.bottom__block > div:nth-child(2){*/
  /*  margin-top: 12px;*/
  /*}*/
}
</style>
