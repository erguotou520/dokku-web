<template>
  <div class="ui-top-nav flex">
    <div class="limit-width flex flex-1 flex-cross-center flex-between">
      <div class="breadcrumb flex flex-cross-center">
        <octicon class="icon" name="apps" :w="16" :h="16" fill="#07cbcf"></octicon>Apps
      </div>
      <el-button name="plus" @click.native="$router.push('/app-create')">New</el-button>
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
  width 100%
  height 3.75rem
  background-color $color-gray-exact-light
  .icon
    margin-right .5rem
</style>
