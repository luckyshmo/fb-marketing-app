<template>
  <div class="moderation">
    <CampaginHeader
      :campaginHeaderData="moderationData.campaginHeader"
    />
    <CampaginBudget
      :campaginBudgetData="moderationData.campaginBudget"
    />
    <CampaginControl
      :campaginControlData="moderationData.campaginControl"
    />
  </div>
</template>

<script>
import CampaginHeader from "@src/components/unattached-account-components/CampaginHeader";
import CampaginBudget from "@src/components/unattached-account-components/CampaginBudget";
import CampaginControl from "@src/components/unattached-account-components/CampaginControl";
export default {
  name: "Moderation",
  data () {
    return {
      moderationData : {
        campaginHeader: {
          imageUrl: '@src/assets/img/image_2.png',
          title: 'Рекламная кампания от 7 сентября',
          text: 'На модерации',
          activeColor: '#FFCB12'
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
          }, {
            title: 'Остановить кампанию',
            svg: false,
            colorRed: true,
            colorPurple: false
          }]
        }
      }
    }
  },
  components : {
    CampaginHeader,
    CampaginBudget,
    CampaginControl
  }
}
</script>

<style scoped>
.moderation {
  width: 580px;
}

@media (max-width: 600px) {
  .moderation {
    width: 100%;
  }
}
</style>
