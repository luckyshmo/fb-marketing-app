<template>
  <div class="page_anchor">
    <h2>{{campaginSectionData.title}}</h2>
    <p>{{campaginSectionData.text}}</p>
    <button class="app-new-submit-button">{{campaginSectionData.buttonValue}}</button>
  </div>
</template>

<script>
export default {
  name: "CampaginSection",
  props: {
    campaginSectionData: {
      type: Object,
      required: true
    }
  }
}
</script>

<style scoped>
.page_anchor{
  padding-top: 40px;
}
.page_anchor > p{
  margin: 8px 0 20px;
}
</style>
