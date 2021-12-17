<template>
<div>
  <section class="steps">
    <b-row class="app-new-steps-header">
      <b-col class="align-items-start">
        <router-link id="app-link-back" :to="{name: 'mainPage'}">
          ← Назад
        </router-link>
          <div class="text-muted" style="line-height: 20px">&nbsp;∙&nbsp;Шаг 1 из 4</div>
      </b-col>
    </b-row>
    <b-row align-h="between">
      <b-col>
        <Step1
          label_cols="12"
          content_cols="12"
          :isEdit="isEdit"
          @logout="logout"
          @next="saveAndNext">
          <template v-slot:header>
            <h1>О бизнесе</h1>
          </template>
        </Step1>
      </b-col>
      <b-col class="steps-process">
        <aside>
          <h2>
            Заполнено 25%
          </h2>
          <ul class="app-new-progress-text">
            <li class="active">
              О бизнесе
            </li>
            <li >
              Изображения
            </li>
            <li>
              Привязка аккаунта
            </li>
            <li>
              Охват рекламы
            </li>
            <li>
              Аудитория
            </li>
          </ul>
        </aside>
      </b-col>
    </b-row>
  </section>
</div>
</template>

<script>
// import { validationMixin } from 'vuelidate'
// import { required, maxLength, minLength } from 'vuelidate/lib/validators'
// import loading from '@src/components/Loading.vue'

import Step1 from './campagin-creation/Step1.vue'
export default {
  name: "CreateCampaignStep-1",
  components:{Step1},
  data(){
    return{

    }
  },
  methods : {
    saveAndNext(){

    }
  }
}
</script>

<style lang="scss">
@import '@src/assets/styles/vars.scss';

#app-link-back {
  // float: left;
  margin-right: 5px;
  float:left;
  line-height:20px;
  color: #6c757d ;

  &:hover {
    color:$black;
  }
}
.form-control{
  width: 100% !important;
}
.app-new-progress-text  {
  list-style: none;
  margin-left: -20px;
  padding-left: 0;
}

.app-new-progress-text li {
  color: $gray;
  font-size: $baseFont;
  line-height: 46px;

  &::before {
    margin-right: 10px;
    color: $gray;
    content: "•";
  }

  &.active {
    color: $black;

    &::before {
      color: $lime;
      content: "•";
    }
  }
}

.app-new-steps-header {
  font-size: 14px;
  line-height: 18px;
  margin: 52px 0 12px !important;
}

#icon-div-image{
  position: absolute;
  margin-left: 140px;
  margin-top: -15px;
}

#primary-under{
  margin-left: 20px;
}

@media (max-width: 722px) {
  #c-status-text-u{
    margin-left: 30px;
  }
  #primary-under{
    margin-left: 0;
    margin-top: 20px;
  }
}

aside {
  margin-top: 35px;
}

.custom-control-input{
  margin-right: 3px;
}

@media (max-width: 1008px) {
  .steps-process {
    display: none;
  }
}
</style>
