<template>
  <div class="draft">
    <CampaginHeader
      :campaginHeaderData="draftData.campaginHeader"
    />
    <CampaginSection
    :campaginSectionData="draftData.campaginSection"
    />
    <CampaginControl
    :campaginControlData="draftData.campaginControl"
    />
  </div>
</template>

<script>
import CampaginHeader from "@src/components/unattached-account-components/CampaginHeader";
import CampaginSection from "@src/components/unattached-account-components/CampaginSection";
import CampaginControl from "@src/components/unattached-account-components/CampaginControl";
export default {
  name: "Draft",
  data () {
    return {
      draftData : {
        campaginHeader: {
          imageUrl: '@src/assets/img/image_2.png',
          title: 'Рекламная кампания от 7 сентября',
          text: 'Аккаунт не привязан',
          activeColor: '#767676'
        },
        campaginSection: {
          title: 'Продолжить настройку',
          text: 'Осталость выбрать желаемый охват аудитории и длительность рекламной кампании. ',
          buttonValue: 'Продолжить'
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
  components: {
    CampaginHeader,
    CampaginSection,
    CampaginControl
  }

}
</script>

<style scoped>
.draft {
  width: 580px;
}
.page_anchor{
  padding-right: 24px;
}
@media (max-width: 600px) {
  .draft {
    width: 100%;
  }
  .page_anchor{
    padding-right: 0;
  }
}
</style>
