<template>
  <section class="unattached-account">
    <CampaginHeader
      :campaginHeaderData="unattachedData.campaginHeader"
    />
    <CampaginSection
      :campaginSectionData="unattachedData.campaginSection"
    />
    <CampaginBudget
      :campaginBudgetData="unattachedData.campaginBudget"
    />
    <CampaginControl
      :campaginControlData="unattachedData.campaginControl"
    />
  </section>
</template>

<script>
import CampaginHeader from "@src/components/unattached-account-components/CampaginHeader";
import CampaginSection from "@src/components/unattached-account-components/CampaginSection";
import CampaginBudget from "@src/components/unattached-account-components/CampaginBudget";
import CampaginControl from "@src/components/unattached-account-components/CampaginControl";

export default {
  name: "UnattachedAccount",
  data() {
    return {
      unattachedData: {
        campaginHeader: {
          imageUrl: '@src/assets/img/image_2.png',
          title: 'Рекламная кампания от 7 сентября',
          text: 'Аккаунт не привязан',
          activeColor: '#767676'
        },
        campaginSection: {
          title: 'Привязка страницы',
          text: 'Привяжите страницу Facebook, чтобы мы могли запустить рекламу от имени вашей\n' +
            '      компании. Это безопасно — мы не получим доступ к вашему логину, паролю, личному аккаунту и сообщениям.',
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
          }, {
            title: 'Привязать аккаунт',
            svg: true,
            colorRed: false,
            colorPurple: false
          }]
        }
      }
    }
  },
  components: {
    CampaginHeader,
    CampaginSection,
    CampaginBudget,
    CampaginControl
  }
}
</script>

<style lang="scss">
@import '@src/assets/styles/vars.scss';

.unattached-account {
  width: 580px;
}

@media (max-width: 600px) {
  .unattached-account {
    width: 100%;
  }
}
</style>
