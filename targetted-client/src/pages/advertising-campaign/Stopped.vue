<template>
<div class="stopped">
  <CampaginHeader
    :campaginHeaderData="stoppedData.campaginHeader"
  />
  <CampaginSection
    :campaginSectionData="stoppedData.campaginSection"
  />
  <CampaginStatistics
    :campaginStatisticsData="stoppedData.campaginStatistics"
  />
  <CampaginBudget
    :campaginBudgetData="stoppedData.campaginBudget"
  />
  <CampaginControl
    :campaginControlData="stoppedData.campaginControl"
  />
</div>
</template>

<script>
import CampaginHeader from "@src/components/unattached-account-components/CampaginHeader";
import CampaginSection from "@src/components/unattached-account-components/CampaginSection";
import CampaginBudget from "@src/components/unattached-account-components/CampaginBudget";
import CampaginControl from "@src/components/unattached-account-components/CampaginControl";
import CampaginStatistics from "@src/components/unattached-account-components/CampaginStatistics";
export default {
  name: "Stopped",
  data (){
    return{
      stoppedData : {
        campaginHeader: {
          imageUrl: '@src/assets/img/image_2.png',
          title: 'Рекламная кампания от 7 сентября',
          text: 'На паузе',
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
          title: 'Реклама остановлена',
          text: 'Запустите рекламную кампанию, чтобы продолжить получать клиентов.',
          buttonValue: 'Привязать страницу'
        },
        campaginBudget: {
          title: 'Бюджет и длительность',
          campaginBudgetValue: [{
            title: 'Баланс',
            value: '1 000 ₽'
          }, {
            title: 'Длительность',
            value: 'Ещё 7 дней'
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
            title: 'Изменить настройки аудитории',
            svg: true,
            colorRed: false,
            colorPurple: false
          }, {
            title: 'Запустить кампанию',
            svg: false,
            colorRed: false,
            colorPurple: true
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
.stopped {
  width: 580px;
}

@media (max-width: 600px) {
  .stopped {
    width: 100%;
  }
}
</style>
