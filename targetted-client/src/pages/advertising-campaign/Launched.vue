<template>
  <div class="launched">
    <CampaginHeader
      :campaginHeaderData="launchedData.campaginHeader"
    />
    <CampaginStatistics
    :campaginStatisticsData="launchedData.campaginStatistics"
    />
    <CampaginBudget
      :campaginBudgetData="launchedData.campaginBudget"
    />
    <CampaginControl
      :campaginControlData="launchedData.campaginControl"
    />
  </div>
</template>

<script>
import CampaginHeader from "@src/components/unattached-account-components/CampaginHeader";
import CampaginStatistics from "@src/components/unattached-account-components/CampaginStatistics";
import CampaginBudget from "@src/components/unattached-account-components/CampaginBudget";
import CampaginControl from "@src/components/unattached-account-components/CampaginControl";
export default {
  name: "Launched",
  data() {
    return {
      launchedData: {
        campaginHeader: {
          imageUrl: '@src/assets/img/image_2.png',
          title: 'Рекламная кампания от 7 сентября',
          text: 'Запущена',
          activeColor: '#767676'
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
  components: {
    CampaginHeader,
    CampaginStatistics,
    CampaginBudget,
    CampaginControl
  }
}
</script>

<style scoped>
.launched {
  width: 580px;
}

@media (max-width: 600px) {
  .launched {
    width: 100%;
  }
}
</style>
