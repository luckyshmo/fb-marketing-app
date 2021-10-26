<template>
  <div
    ref="popup_wrapper"
    class="popup-wrapper-big"
  >
    <div class="popup-big">
      <!-- <div
        style="position: absolute;
      margin-left: 740px;
      margin-top: -5px;"
      >
        <b-icon
          class="x-button"
          icon="x"
          @click="closePopup"
        />
      </div> -->
      <div
        style="position: absolute;
      margin-left: 370px;
      margin-top: 475px;"
      >
        <b-icon
          style="height: 30px; width: 30px;"
          icon="arrow-down-circle"
        >
          huw
        </b-icon>
      </div>
      <div class="popup-content-big">
        <slot />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Popup',
  props: {
    popupTitle: {
      type: String,
      default: 'Popup name'
    },
    rightBtnTitle: {
      type: String,
      default: 'Ok'
    }
  },
  data () {
    return {}
  },
  mounted () {
    const vm = this

    vm.renderAction()
    document.addEventListener('click', function (item) {
      if (item.target === vm.$refs.popup_wrapper) {
        vm.closePopup()
      }
    })
  },
  methods: {
    rightBtnAction () {
      this.$emit('rightBtnAction')
    },
    renderAction () {
      this.$emit('renderAction')
    },
    closePopup () {
      console.log('huy')
      this.$emit('closePopup')
    }
  }
}
</script>

<style>
  .popup-content-big{
    max-height: 470px;
    overflow-y: auto !important;
  }

  .popup-content-big:before {
    content:'';
    margin-top: 50px;
    width:100%;
    height:440px;
    position:absolute;
    left:0;
    top:0;
    background:linear-gradient(transparent 350px, white);
  }

  .popup-wrapper-big {
    /* //TODO */
    cursor: grab;
    position: fixed;
    z-index: 1001;
    background: rgba(0, 0, 0, 0.9);
    display: flex;
    justify-content: center;
    align-items: center;
    right: 0;
    left: 0;
    top: 0;
    bottom: 0;
  }

  .popup-big {
    padding: 20px;
    align-content: center;
    text-align: center;
    position:absolute;
    top:50%;
    left:50%;
    width:820px;
    height:530px;
    margin-left:-410px;   /* negative half of width above */
    margin-top:-235px;   /* negative half of height above */
    background-color: white;
    border-radius: 25px;
  }

  @media (max-width: 600px) {
    .popup-big{
      width:320px;
      margin-left:-160px;   /* negative half of width above */
      padding: 15px;
    }

    .popup-footer-big {
      padding-left: 30px;
      padding-right: 30px;
    }
}
</style>
