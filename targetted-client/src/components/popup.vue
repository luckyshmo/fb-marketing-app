<template>
  <div class="popup_wrapper" ref="popup_wrapper">
    <div class='popup'>
      <div 
      style="position: absolute;
      margin-left: 570px;
      margin-top: -5px;">
          <b-icon
          @click="closePopup"
          style="height: 35px;
                  width: 35px;
                  background: #e4e4e4;
                  border-radius: 17.5px;
                  cursor: pointer;"
          icon="x"></b-icon>
      </div>
      <div class="popup-header">
        <span><h1 id="h1">{{popupTitle}}</h1></span>
        <span>
      </span>
      </div>
      <div class="popup-content">
        <slot></slot>
      </div>
      <div class="popup-footer">
        <p>Whatsup</p>
        <p>Telegram</p>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "v-popup",
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
    data() {
      return {}
    },
    methods: {
      rightBtnAction() {
        this.$emit('rightBtnAction')
      },
      closePopup() {
        this.$emit('closePopup')
      }
    },
    mounted() {
      let vm = this;
      document.addEventListener('click', function (item) {
        if (item.target === vm.$refs['popup_wrapper']) {
          vm.closePopup()
        }
      })
    },
  }
</script>

<style>
  .popup_wrapper {
    cursor: grab;
    position: fixed;
    z-index: 1;
    background: rgba(0, 0, 0, 0.9);
    display: flex;
    justify-content: center;
    align-items: center;
    right: 0;
    left: 0;
    top: 0;
    bottom: 0;
    display: flex;
  }
  .popup {
    padding: 20px;
    align-content: center;
    text-align: center;
    position:absolute;
    top:50%;
    left:50%;
    width:640px;
    height:470px;
    margin-left:-320px;   /* negative half of width above */
    margin-top:-235px;   /* negative half of height above */
    background-color: white;
    border-radius: 25px;
  }
  .popup-header {
    /* display: flex; */
    padding-top: 50px;
      justify-content: space-between;
      align-items: center;
  }
  .popup-footer {
    color: violet;
    padding-left: 150px;
    padding-right: 150px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  .popup-content {
    padding: 40px;
  }
</style>