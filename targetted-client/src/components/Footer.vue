<template>
  <section class="footer">
    <span class="footer-text">
      © targetted.ru, 2020–2021. <a class="footer-text-link" href="#">Оферта</a>.
    </span>
  </section>
</template>
<style lang="scss">
  @import '@src/assets/styles/vars.scss';

  .footer {
    width: 100%;
    align-self: flex-end;
    padding-bottom: 29px;
  }

  .footer-text {
    font-style: normal;
    font-weight: normal;
    font-size: 14px;
    line-height: 18px;
    color: $gray;
  }

  .footer-text-link {
    color: $gray;
    text-decoration: underline !important;
  }
</style>
