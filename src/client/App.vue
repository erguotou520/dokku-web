<template>
  <div id="app-main">
    <x-header v-if="loggedIn"></x-header>
    <div class="app-container flex flex-column flex-main-center flex-cross-center" :class="{active:loggedIn}">
      <top-nav v-if="loggedIn&&!$route.meta[404]"></top-nav>
      <!-- <router-loading></router-loading> -->
      <router-view></router-view>
    </div>
    <footer-bar v-if="loggedIn"></footer-bar>
  </div>
</template>
<script>
import Vue from 'vue'
import { mapGetters } from 'vuex'
import XHeader from './components/Header'
import TopNav from './components/TopNav'
import FooterBar from './components/FooterBar'
import ContentModule from './components/ContentModule'
import Octicon from './components/Octicon'
import './components/octicons'
Vue.component('ContentModule', ContentModule)
Vue.component('Octicon', Octicon)
export default {
  computed: {
    ...mapGetters(['loggedIn'])
  },
  components: {
    XHeader, TopNav, FooterBar
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
    min-height 100vh
    background-color #fff
    overflow auto
    &.active
      min-height calc(100vh - 6.625rem)
.limit-width
  width 100%
  max-width 75rem
  margin 0 auto
.el-dropdown-link
  cursor pointer
</style>
