<template>
  <div
    ref="popup_wrapper"
    class="popup_wrapper"
  >
   <b-icon
          class="x-button"
          icon="x"
          @click="closePopup"
        />
    <div class="popup">
      <div style="position: fixed;right: 20px;top: 35px;">
          <b-icon
          class="x-button"
          icon="x"
          @click="closePopup"
        />
      </div>
      <div class="popup-content">
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
      this.$emit('closePopup')
    }
  }
}
</script>

<style>
  .popup_wrapper {
    /* //TODO */
    cursor: grab;
    position: fixed;
    z-index: 1;
    background: #fff;
    display: flex;
    justify-content: center;
    align-items: center;
    right: 0;
    left: 0;
    top: 0;
    bottom: 0;
  }

  .popup {
    padding: 20px;
    align-content: center;
    position:absolute;
    top:50%;
    left:50%;
    width:640px;
    min-height:470px;
    margin-left:-320px;   /* negative half of width above */
    margin-top:-235px;   /* negative half of height above */
    background-color: white;
    border-radius: 25px;
  }

  @media (max-width: 600px) {
    .popup{
      width:320px;
      margin-left:-160px;   /* negative half of width above */
      padding: 0;
    }
    .popup-footer {
      padding-left: 30px;
      padding-right: 30px;
    }
}
</style>
