<template>
    <div class="footer">
      <div class="footer-t">
        © targetted.ru, 2020–2021. <a href="#">Оферта</a>.
      </div>
    </div>
</template>
<script>
export default {

}
</script>
<style lang="scss">
  @import '@/assets/styles/vars.scss';

.footer-t{
    position:relative;
    bottom: -38px;
    margin: 0 auto;
    text-align: left;

    max-width: 1060px;
}
@media (max-width: 600px) {
  .footer-t{
    margin-left: 24px;
  }
}
.footer-t a {
  color: $gray !important;
  text-decoration: underline !important;
}
.footer {
    height: 100px;

    font-style: normal;
    font-weight: normal;
    font-size: 14px;
    line-height: 18px;
    /* identical to box height, or 150% */
    text-align: center;
    color: $gray;
}
</style>
