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
@import '~vtc/stylus/flex.styl'
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
p
  font-size 14px
pre
code
  font-family "Source Code Pro", Consolas, Monaco, Menlo, monospace
pre
  padding 1rem
  background-color $color-gray-exact-light
  border-radius 4px
code
  padding 2px
  background-color $color-gray-exact-light
  color #666
  border 1px solid $color-gray-light
  border-radius 3px
  font-size 12px
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
.t-l
  text-align left
.t-c
  text-align center
.t-r
  text-align right
</style>
