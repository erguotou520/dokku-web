<template>
  <content-module name="dashboard">
    <div class="apps">
      <div v-for="n in 4" class="app-wrapper">
        <div class="app limit-width flex flex-cross-center flex-between">
          <div class="flex flex-cross-center">
            <octicon class="app-status" name="app" :w="24" :h="24"></octicon>
            <span class="app-name">Vue-fullstack</span>
          </div>
          <div class="flex flex-cross-center">
            <img class="app-lang-icon" src="/static/nodejs.png">
            <span class="app-lang">Node.js</span>
          </div>
        </div>
      </div>
    </div>
  </content-module>
</template>
<script>
export default {

}
</script>
<style lang="stylus">
@import "../assets/css/variable.styl"
.m-dashboard
  margin 20px 0
  .app-wrapper
    border-top 1px solid $color-gray
    transition all .3s
    &:first-child
      border-top none
    &:hover
      background-color $color-gray-light
  .app
    padding 15px 0
    font-size 1rem
    .app-status
      margin-right .75rem
    .app-lang
      margin-left .75rem
</style>
