<template>
  <div
    ref="popup_wrapper"
    class="popup-wrapper-big"
  >
    <div class="popup-big">
       <div>
        <b-icon
          class="x-button"
          icon="x"
          @click="closePopup"
        />
      </div>
      <div>
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
  /*.popup-content-big{*/
  /*  max-height: 470px;*/
  /*  overflow-y: auto !important;*/
  /*}*/

  /*.popup-content-big:before {*/
  /*  content:'';*/
  /*  margin-top: 50px;*/
  /*  width:100%;*/
  /*  height:440px;*/
  /*  position:absolute;*/
  /*  left:0;*/
  /*  top:0;*/
  /*  background:linear-gradient(transparent 350px, white);*/
  /*}*/

  .popup-wrapper-big {
    /* //TODO */
    position: fixed;
    z-index: 1001;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.9);
    display: flex;
    justify-content: center;
    align-items: center;
    right: 0;
    left: 0;
    top: 0;
    bottom: 0;
  }

  /*.popup-big {*/
  /*  padding: 20px;*/
  /*  align-content: center;*/
  /*  text-align: center;*/
  /*  position:relative;*/
  /*  width:676px;*/
  /*  height: 1000px;*/
  /*  background-color: white;*/
  /*  border-radius: 25px;*/
  /*}*/

  @media (max-width: 600px) {
    /*.popup-big{*/
    /*  width:320px;*/
    /*  margin-left:-160px;   !* negative half of width above *!*/
    /*  padding: 15px;*/
    /*}*/

    /*.popup-footer-big {*/
    /*  padding-left: 30px;*/
    /*  padding-right: 30px;*/
    /*}*/
}
</style>
