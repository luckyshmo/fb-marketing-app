<template>
<div class="not-passed-moderation">
  <CampaginHeader
    :campaginHeaderData="notPassedModerationData.campaginHeader"
  />
  <CampaginSection
    :campaginSectionData="notPassedModerationData.campaginSection"
  />
  <CampaginBudget
    :campaginBudgetData="notPassedModerationData.campaginBudget"
  />
  <CampaginControl
    :campaginControlData="notPassedModerationData.campaginControl"
  />
</div>
</template>

<script>

import CampaginHeader from "@src/components/unattached-account-components/CampaginHeader";
import CampaginSection from "@src/components/unattached-account-components/CampaginSection";
import CampaginBudget from "@src/components/unattached-account-components/CampaginBudget";
import CampaginControl from "@src/components/unattached-account-components/CampaginControl";

export default {
  name: "NotPassedModeration",
  data () {
    return {
      notPassedModerationData : {
        campaginHeader: {
          imageUrl: '@src/assets/img/image_2.png',
          title: 'Рекламная кампания от 7 сентября',
          text: 'Отклонена',
          activeColor: '#FF2F12'
        },
        campaginSection: {
          title: 'Объявление не прошло модерацию',
          text: 'Напишите нам в поддержку, чтобы мы помогли разобраться',
          buttonValue: 'Привязать страницу'
        },
        campaginBudget: {
          title: 'Бюджет и длительность',
          campaginBudgetValue: [{
            title: 'Баланс',
            value: '1 498 ₽'
          }, {
            title: 'Длительность',
            value: '7 дней'
          }, {
            title: 'Бюджет в день',
            value: '214 ₽'
          }],
          buttonValue: 'Изменить'
        },
        campaginControl: {
          title: 'Управление кампанией',
          campaginControlValue: [{
            title: 'Как выглядит реклама',
            svg: true,
            colorRed: false,
            colorPurple: false
          }]
        }
      }
    }
  },
  components : {
    CampaginHeader,
    CampaginSection,
    CampaginBudget,
    CampaginControl
  }
}
</script>

<style scoped>
.not-passed-moderation {
  width: 580px;
}

@media (max-width: 600px) {
  .not-passed-moderation {
    width: 100%;
  }
}
</style>
