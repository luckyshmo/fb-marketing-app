<template>
<section class="unattached-account">
  <router-link :to="{name: 'mainPage'}">
    <p
      id="navigation-text"
      style="margin:52px 0 0 0;"
    >
      ← На главную
    </p>
  </router-link>
  <div class="unattached-account__block">
    <div class="unattached-account__img">
      <img src="../../assets/img/image 2.png" alt="unattached-account__img">
    </div>
    <div class="unattached-account__block_right">
      <h2>
        Рекламная кампания от 7 сентября
      </h2>
      <div>
        <div></div>
        <p>Аккаунт не привязан</p>
      </div>
    </div>
  </div>
  <div class="page_anchor">
    <h2 >Привязка страницы</h2>
    <p >Привяжите страницу Facebook, чтобы мы могли запустить рекламу от имени вашей
      компании. Это безопасно — мы не получим доступ к вашему логину, паролю, личному аккаунту и сообщениям.</p>
    <button class="app-new-submit-button">Привязать страницу</button>
  </div>
  <div class="budget__duration">
    <h2>Бюджет и длительность</h2>
    <p>Баланс: <span>1 498 ₽</span></p>
    <p>Длительность: <span>7 дней</span></p>
    <p>Бюджет в день: <span>214 ₽</span></p>
    <button class="app-new-submit-button budget__button">Изменить</button>
  </div>
  <div class="campaign__management">
    <h2>Управление кампанией</h2>
      <div>
        <p>
          Как выглядит реклама
          <svg viewBox="0 0 7 12" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M1 1L5.5 6L1 11" stroke="black" stroke-width="2" stroke-linecap="round"/>
          </svg>
        </p>
        <div></div>
        <p>
          Привязать аккаунт
          <svg viewBox="0 0 7 12" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M1 1L5.5 6L1 11" stroke="black" stroke-width="2" stroke-linecap="round"/>
          </svg>
        </p>
      </div>
    </div>
</section>
</template>

<script>
export default {
  name: "UnattachedAccount"
}
</script>

<style lang="scss">
 @import '@src/assets/styles/vars.scss';
  .unattached-account{
    width: 580px;
  }
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
    background-color: $gray;
   border-radius: 100%;
   margin-right: 7px;
 }
 .budget__button{
   background: $light !important;
   margin-top: 20px;
 }
 .page_anchor{
   padding-top: 40px;
 }
 .page_anchor > p{
  margin: 8px 0 20px;
 }
 .budget__duration{
   margin-top: 40px;
 }
 .budget__duration > p{
   margin: 12px 0 0;
 }
 .budget__duration > p > span{
   margin-left: 6px;
   font-weight: bold;
 }
 .campaign__management{
   margin: 40px 0 200px;
 }
 .campaign__management > div{
   margin-top: 16px;
   border-radius: 16px;
   padding: 0 16px;
   background-color: $light;
 }
 .campaign__management > div >div{
   width: 100%;
   height: 1px;
   background-color: $rectangle;
 }
 .campaign__management > div > p{
   padding: 16px 0 ;
   margin: 0;
   display: flex;
   justify-content: space-between;
   align-items: center;
 }
 .campaign__management > div > p > svg{
   width: 8px;
   height: 8px;
   fill: $black;
   cursor: pointer;
 }
 @media (max-width : 600px) {
   .unattached-account{
     width: 100%;
   }
   .unattached-account__block{
    position: relative;
    justify-content: center;
     padding-bottom: 28px;
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
