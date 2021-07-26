<template>
  <div
    ref="popup_wrapper"
    class="loading_wrapper"
  >
    <b-spinner
      style="width: 3rem; height: 3rem;"
      type="grow"
    />
  </div>
</template>

<script>
export default {
  name: 'Lading',
  props: {
  },
  data () {
    return {}
  },
  mounted () {
  },
  methods: {
  }
}
</script>

<style>
  .loading_wrapper {
    /* //TODO */
    cursor: grab;
    position: fixed;
    z-index: 1;
    background: rgba(0, 0, 0, 0.3);
    display: flex;
    justify-content: center;
    align-items: center;
    right: 0;
    left: 0;
    top: 0;
    bottom: 0;
    display: flex;
  }
  .loading {
    padding: 20px;
    align-content: center;
    text-align: center;
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
</style>
