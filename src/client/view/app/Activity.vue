<template>
  <div class="m-app-activity">
    <div v-for="commit in histories" class="git-commit flex flex-cross-center">
      <octicon class="icon" name="upload" :w="32" :h="32" fill="#447498"></octicon>
      <img class="avatar" :src="commit.avatar" :alt="commit.avatar">
      <div class="flex-1">
        <p>
          <strong>{{commit.author}}:</strong>
          {{$t('app.activity.deployed')}}
          <code class="commit-sha">{{commit.sha}}</code>
        </p>
        <p>
          <span class="date">{{commit.date|dateFormat('yyyy-MM-dd hh:mm:ss')}}</span>
          <code class="commit-message">{{commit.message}}</code>
        </p>
      </div>
    </div>
  </div>
</template>
<script>
import { format } from 'shared/utils'
import { app as res } from 'resources'
import locales from 'locales/apps'
export default {
  locales,
  data () {
    return {
      histories: []
    }
  },
  filters: {
    dateFormat (val, _format) {
      return format(val, _format)
    }
  },
  created () {
    res.getAppGitLogs({ name: this.$route.params.name }).then(data => {
      this.histories = data.data.map(commit => {
        commit.avatar = commit.avatar + '?s=32&d=mm&r=g'
        return commit
      })
    }).catch(() => {})
  }
}
</script>
<style lang="stylus">
@import "../../assets/css/variable"
.m-app-activity
  .git-commit
    margin .5rem 0
    padding 0 0 .5rem
    border-bottom 1px solid $color-gray
    &:last-child
      border none
    .dokkuicon-upload
      color $color-primary
      font-size 24px
    .avatar
      min-width 2rem
      margin 0 .5rem
    .data
      font-size 12px
    p
      margin 0
      line-height 1.5rem
</style>
