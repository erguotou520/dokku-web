<template>
  <div class="m-dashboard flex-1">
    <div class="apps">
      <div v-if="!apps.length" class="empty-app flex flex-main-center flex-cross-center">{{$t('app.empty')}}</div>
      <router-link v-for="app in apps" tag="div" :to="`/apps/${app.name}`" :key="app" class="app-wrapper">
        <div class="app limit-width flex flex-cross-center flex-between">
          <div class="flex flex-cross-center">
            <octicon class="app-status" name="app" :w="24" :h="24"></octicon>
            <span class="app-name">{{app.name}}</span>
          </div>
          <div class="flex flex-cross-center">
            <img class="app-lang-icon" src="/static/nodejs.png">
            <span class="app-lang">Node.js</span>
          </div>
        </div>
      </router-link>
    </div>
  </div>
</template>
<script>
import { app as res } from 'resources'
import locales from 'locales/apps'
export default {
  locales,
  data () {
    return {
      apps: []
    }
  },
  created () {
    res.query().then(data => {
      this.apps = data.data
    }).catch(() => {})
  }
}
</script>
<style lang="stylus">
@import "../assets/css/variable.styl"
.m-dashboard
  margin 1.5rem 0
  .app-wrapper
    border-top 1px solid $color-gray
    transition all .3s
    cursor pointer
    &:first-child
      border-top none
    &:hover
      background-color $color-gray-light
  .app
    padding .875rem 0
    font-size 1rem
    .app-status
      margin-right .75rem
      margin-left 1rem
    .app-lang
      margin-left .75rem
      margin-right 1rem
</style>
