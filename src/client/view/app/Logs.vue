<template>
  <div class="m-app-logs">
    <pre v-html="logs"></pre>
  </div>
</template>
<script>
import { app as res } from 'resources'
export default {
  data () {
    return {
      logs: ''
    }
  },
  created () {
    res.getAppLogs({ name: this.$route.params.name }).then(data => {
      this.logs = data.data.replace(/\[36m/g, '<span class="cyan">').replace(/\[0m/g, '</span>')
    }).catch(() => {})
  }
}
</script>
<style lang="stylus">
.m-app-logs
  pre
    background #111111
    color #fff
    white-space pre-wrap
    word-wrap break-word
    overflow auto
    overflow-y hidden
    font-size 12px
    .cyan
      color cyan
</style>
