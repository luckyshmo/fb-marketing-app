<template>
  <div class="advertising">
    <CampaginComponent
      v-if="pageId === 1"
      headerActiveColor="#767676"
      headerText="Аккаунт не привязан"
      :statistics="false"
      statisticsImpressions=""
      statisticsTransitions=""
      statisticsWrote=""
      :section="true"
      sectionTitle="Привязка страницы "
      sectionText="Привяжите страницу Facebook, чтобы <br class='display__none'/>  мы могли запустить рекламу от имени вашей компании. Это безопасно —  <br class='display__none'/> мы не получим доступ к вашему логину, паролю, личному аккаунту и сообщениям."
      sectionButton="Привязать страницу"
      :budget="true"
      balance="1 498 ₽"
      duration="7 дней"
      budgetDay="214 ₽"
      budgetButton="Изменить"
      :management="{
      0 : {
        text: 'Как выглядит реклама',
        svg : true,
        colorRed : false,
        colorPurple: false
      },
      1 : {
        text: 'Привязать аккаунт',
        svg : true,
        colorRed : false,
        colorPurple: false
      }
    }"
    />
    <CampaginComponent
      v-if="pageId === 2"
      headerActiveColor="#000000"
      headerText="Запущена"
      :statistics="true"
      statisticsImpressions="150"
      statisticsTransitions="34"
      statisticsWrote="15"
      :section="false"
      sectionTitle=""
      sectionText=""
      sectionButton=""
      :budget="true"
      balance="1 498 ₽"
      duration="7 дней"
      budgetDay="214 ₽"
      budgetButton="Изменить"
      :management="{
      0 : {
        text: 'Как выглядит реклама',
        svg : true,
        colorRed : false,
        colorPurple: false
      },
      1 : {
        text: 'Остановить кампанию',
        svg : false,
        colorRed : true,
        colorPurple: false
      }
    }"
    />
    <CampaginComponent
      v-if="pageId === 3"
      headerActiveColor="#FFCB12"
      headerText="На модерации"
      :statistics="false"
      statisticsImpressions=""
      statisticsTransitions=""
      statisticsWrote=""
      :section="false"
      sectionTitle=""
      sectionText=""
      sectionButton=""
      :budget="true"
      balance="1 498 ₽"
      duration="7 дней"
      budgetDay="214 ₽"
      budgetButton="Изменить"
      :management="{
      0 : {
        text: 'Как выглядит реклама',
        svg : true,
        colorRed : false,
        colorPurple: false
      },
      1 : {
        text: 'Остановить кампанию',
        svg : false,
        colorRed : true,
        colorPurple: false
      }
    }"
    />
    <CampaginComponent
      v-if="pageId === 4"
      headerActiveColor="#FFCB12"
      headerText="Отклонена"
      :statistics="false"
      statisticsImpressions=""
      statisticsTransitions=""
      statisticsWrote=""
      :section="true"
      sectionTitle="Объявление не прошло модерацию"
      sectionText="Напишите нам в поддержку, чтобы мы помогли разобраться "
      sectionButton="Привязать страницу"
      :budget="true"
      balance="1 498 ₽"
      duration="7 дней"
      budgetDay="214 ₽"
      budgetButton="Изменить"
      :management="{
      0 : {
        text: 'Как выглядит реклама',
        svg : true,
        colorRed : false,
        colorPurple: false
      }
    }"
    />
    <CampaginComponent
      v-if="pageId === 5"
      headerActiveColor="#FFCB12"
      headerText="Завершена"
      :statistics="true"
      statisticsImpressions="150"
      statisticsTransitions="34"
      statisticsWrote="15"
      :section="true"
      sectionTitle="Реклама завершена"
      sectionText="Закончился срок отображения рекламной кампании. Пополните бюджет, чтобы возобновить рекламную кампанию."
      sectionButton="Привязать страницу"
      :budget="true"
      balance="0 ₽"
      duration="Завершена"
      budgetDay="214 ₽"
      budgetButton="Изменить"
      :management="{
      0 : {
        text: 'Как выглядит реклама',
        svg : true,
        colorRed : false,
        colorPurple: false
      }
    }"
    />
    <CampaginComponent
      v-if="pageId === 6"
      headerActiveColor="#FFCB12"
      headerText="На паузе"
      :statistics="true"
      statisticsImpressions="150"
      statisticsTransitions="34"
      statisticsWrote="15"
      :section="true"
      sectionTitle="Реклама остановлена"
      sectionText="Запустите рекламную кампанию, чтобы продолжить получать клиентов."
      sectionButton="Привязать страницу"
      :budget="true"
      balance="1 000 ₽"
      duration="Ещё 7 дней"
      budgetDay="214 ₽"
      budgetButton="Изменить"
      :management="{
      0 : {
        text: 'Как выглядит реклама',
        svg : true,
        colorRed : false,
        colorPurple: false
      },
      1 : {
        text: 'Изменить настройки аудитории',
        svg : true,
        colorRed : false,
        colorPurple: false
      },
      2 : {
        text: 'Запустить кампанию',
        svg : false,
        colorRed : false,
        colorPurple: true
      }
    }"
    />
    <CampaginComponent
      v-if="pageId === 7"
      headerActiveColor="#767676"
      headerText="Черновик"
      :statistics="false"
      statisticsImpressions=""
      statisticsTransitions=""
      statisticsWrote=""
      :section="true"
      sectionTitle="Продолжить настройку"
      sectionText="Осталость выбрать желаемый охват аудитории и длительность рекламной кампании. "
      sectionButton="Продолжить"
      :budget="false"
      balance=""
      duration=""
      budgetDay=""
      budgetButton=""
      :management="{
      0 : {
        text: 'Как выглядит реклама',
        svg : true,
        colorRed : false,
        colorPurple: false
      }
    }"
    />
  </div>
</template>

<script>
import CampaginComponent from "@src/components/unattached-account-components/CampaginComponent";

export default {
  name: "AdvertisingCampaign",
  data() {
    return {
      pageId: 1
    }
  },
  components: {
    CampaginComponent
  }
}
</script>

<style scoped>
.advertising {
  width: 600px;
}

@media (max-width: 600px) {
  .advertising {
    width: 100%;
  }
}
</style>
