<template>
  <div class="m-app flex-1">
    <el-tabs v-model="path" @tab-click="$router.push(`/apps/${$route.params.name}/${path}`)" ref="tabs" :style="{'margin-left':marginLeft}">
      <el-tab-pane :label="$t('app.tab.overview')" name="overview"></el-tab-pane>
      <el-tab-pane :label="$t('app.tab.resources')" name="resources"></el-tab-pane>
      <el-tab-pane :label="$t('app.tab.deploy')" name="deploy"></el-tab-pane>
      <el-tab-pane :label="$t('app.tab.activity')" name="activity"></el-tab-pane>
      <el-tab-pane :label="$t('app.tab.settings')" name="settings"></el-tab-pane>
    </el-tabs>
    <div class="hack-line"></div>
    <div class="app-wrapper limit-width">
      <router-view></router-view>
    </div>
  </div>
</template>
<script>
import locales from 'locales/apps'
export default {
  locales,
  data () {
    return {
      path: this.$route.path.split('/')[3],
      marginLeft: '0'
    }
  },
  mounted () {
    this.$nextTick(() => {
      const $nav = this.$refs.tabs.$el.querySelector('.el-tabs__nav')
      this.marginLeft = `-${$nav.clientWidth / 2}px`
    })
  }
}
</script>
<style lang="stylus">
.m-app
  .el-tabs
    position relative
    display inline-block
    left 50%
  .hack-line
    margin-top -19px
    border-bottom 1px solid #d1dbe5
  .el-tabs__item
    padding 0 1.5rem
  .app-wrapper
    margin-top 1rem
  h4
    margin 0
    line-height 2.25rem
</style>
