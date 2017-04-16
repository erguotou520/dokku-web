<template>
  <div class="ui-top-nav flex">
    <div class="limit-width flex flex-1 flex-cross-center flex-between">
      <div class="breadcrumb flex flex-cross-center">
        <span class="iconfont dokkuicon-apps"></span>Apps
      </div>
      <el-button name="plus">New</el-button>
    </div>
  </div>
</template>
<script>
export default {
  data () {
    return {

    }
  }
}
</script>
<style lang="stylus">
@import "../assets/css/variable.styl"
.ui-top-nav
  height 3.75rem
  background-color $color-gray-exact-light
  .iconfont
    margin-right .5rem
    color $color-primary
</style>
