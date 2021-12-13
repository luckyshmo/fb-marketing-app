<template>
  <div class="money">
    <CampaginHeader
      :campaginHeaderData="moneyData.campaginHeader"
    />
    <CampaginSection
      :campaginSectionData="moneyData.campaginSection"
    />
    <CampaginStatistics
      :campaginStatisticsData="moneyData.campaginStatistics"
    />
    <CampaginBudget
      :campaginBudgetData="moneyData.campaginBudget"
    />
    <CampaginControl
      :campaginControlData="moneyData.campaginControl"
    />
  </div>
</template>

<script>
import CampaginHeader from "@src/components/unattached-account-components/CampaginHeader";
import CampaginStatistics from "@src/components/unattached-account-components/CampaginStatistics";
import CampaginSection from "@src/components/unattached-account-components/CampaginSection";
import CampaginBudget from "@src/components/unattached-account-components/CampaginBudget";
import CampaginControl from "@src/components/unattached-account-components/CampaginControl";
export default {
  name: "Money",
  data (){
    return{
      moneyData : {
        campaginHeader: {
          imageUrl: '@src/assets/img/image_2.png',
          title: 'Рекламная кампания от 7 сентября',
          text: 'Завершена',
          activeColor: '#FFCB12'
        },
        campaginStatistics: {
          title: 'Статистика',
          campaginStatisticsValue: [{
            title: 'Количество показов',
            value: '150'
          }, {
            title: 'Количество переходов',
            value: '34'
          }, {
            title: 'Количество написавших',
            value: '15'
          }],
          buttonValue: 'Вся статистика'
        },
        campaginSection: {
          title: 'Реклама завершена',
          text: 'Закончился срок отображения рекламной кампании. Пополните бюджет, чтобы возобновить рекламную кампанию.',
          buttonValue: 'Привязать страницу'
        },
        campaginBudget: {
          title: 'Бюджет и длительность',
          campaginBudgetValue: [{
            title: 'Баланс',
            value: '0 ₽'
          }, {
            title: 'Длительность',
            value: 'Завершена'
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
    CampaginStatistics,
    CampaginSection,
    CampaginBudget,
    CampaginControl
  }
}
</script>

<style scoped>
.money {
  width: 580px;
}

@media (max-width: 600px) {
  .money {
    width: 100%;
  }
}
</style>
