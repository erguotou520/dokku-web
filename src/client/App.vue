<template>
  <div id="app-main">
    <x-header></x-header>
    <div class="app-container">
      <top-nav></top-nav>
      <!-- <router-loading></router-loading> -->
      <router-view></router-view>
    </div>
  </div>
</template>
<script>
import Vue from 'vue'
import XHeader from './components/Header'
import TopNav from './components/TopNav'
import ContentModule from './components/ContentModule'
Vue.component('ContentModule', ContentModule)
export default {
  components: {
    XHeader, TopNav
  }
}
</script>
<style lang="stylus">
@import "assets/css/variable"
@import "assets/css/animate"
@import "assets/fonts/iconfont.css"
@import "assets/css/flex"
::-webkit-scrollbar
  width 4px
  height 4px
  background-color $color-gray
::-webkit-scrollbar-button
::-webkit-scrollbar-track
  display none
::-webkit-scrollbar-thumb
  background-color $color-silver-light
::-webkit-scrollbar-thumb:hover
  background-color $color-black-exact-light

html
  font-size 16px
  font-family "Helvetica Neue", Helvetica, Tahoma, "PingFang SC", "Hiragino Sans GB", "Heiti SC", "Microsoft YaHei", "微软雅黑", "WenQuanYi Micro Hei", Arial, sans-serif
body
  margin 0
  padding 0
  font-size .75rem
  background-color #f0f0f0
/**
*::before
*::after
  box-sizing border-box*/
a
  text-decoration none
#app-main
  min-height 100vh
  .app-container
    position relative
    min-height calc(100vh - 3.5rem)
    background-color #fff
    overflow auto
.limit-width
  width 100%
  max-width 75rem
  margin 0 auto
.el-dropdown-link
  cursor pointer
</style>
