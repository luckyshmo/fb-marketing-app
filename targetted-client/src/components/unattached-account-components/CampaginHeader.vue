<template>
  <div>
    <router-link :to="{name: 'mainPage'}">
      <p
        id="navigation-text"
        style="margin:52px 0 0 0; display: flex; align-items: center"
      >
        ← На главную
      </p>
    </router-link>
    <div class="unattached-account__block">
      <div class="unattached-account__img">
        <img src="@src/assets/img/image_2.png" alt="unattached-account__img">
      </div>
      <div class="unattached-account__block_right">
        <h2>
          {{campaginHeaderData.title}}
        </h2>
        <div>
          <div :style="{backgroundColor : campaginHeaderData.activeColor}"></div>
          <p>{{ campaginHeaderData.text }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "CampaginHeader",
  props : {
    campaginHeaderData : {
      type : Object,
      required : true
    }
  }
}
</script>

<style lang="scss">
@import '@src/assets/styles/vars.scss';
.unattached-account__block {
  width: 100%;
  display: flex;
  margin-top: 20px;
  border-radius: 16px;
  background-color: $light;
}
.unattached-account__img{
  width: 140px;
  height: 140px;
  border-radius: 12px;
  margin: 24px;
}
.unattached-account__img > img{
  width: 100%;
  height: 100%;
}
.unattached-account__block_right{
  max-width: 346px;
}
.unattached-account__block_right  > h2{
  margin-top: 38px !important;
  font-size: 32px !important;
  line-height: 38px;
  color: $black;
}
.unattached-account__block_right > div{
  width: 100%;
  display: flex;
  align-items: center;
  margin-top: 12px;
}
.unattached-account__block_right > div > p{
  font-size: 14px;
  line-height: 18px;
  color: $black;
  position: relative;
  margin: 0;
}
.unattached-account__block_right > div > div{
  width: 8px;
  height: 8px;
  //background-color: $gray;
  border-radius: 100%;
  margin-right: 7px;
}
@media (max-width: 600px) {
  .unattached-account__block{
    position: relative;
    justify-content: center;
    padding-bottom: 28px;
    margin-top: 70px;
  }
  .unattached-account__img{
    position: absolute;
    top: -50px;
    width: 100px;
    height: 100px;
    margin: 0;
  }
  .unattached-account__block_right{
    width: 100%;
    max-width: 254px;
    text-align: center;
  }
  .unattached-account__block_right > h2{
    margin-top: 66px !important;
    font-size: 24px !important;
  }
  .unattached-account__block_right > div{
    justify-content: center;
  }
}
</style>
