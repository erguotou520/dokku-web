<template>
  <div class="m-app-deploy">
    <el-row :gutter="20">
      <el-col :span="6" :offset="4">
        <h4>{{$t('app.deploy.repo')}}</h4>
      </el-col>
      <el-col :span="14">
        <p>Dokku Git</p>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="6" :offset="4">
        <h4>{{$t('app.deploy.tipLabel')}}</h4>
      </el-col>
      <el-col :span="14">
        <p>{{$t('app.deploy.dokkuGit')}}<el-input :value="repoUrl" readonly class="repo-url-input"></el-input></p>
        <p>{{$t('app.deploy.fromEmpty')}}</p>
        <pre><code>echo "# 123" >> README.md
git init
git add README.md
git commit -m "first commit"
git remote add dokku {{repoUrl}}
git push -u dokku master</code></pre>
        <p>{{$t('app.deploy.fromExisted')}}</p>
        <pre><code>git remote add origin {{repoUrl}}
git push -u dokku master</code></pre>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import locales from 'locales/apps'
export default {
  locales,
  data () {
    return {
    }
  },
  computed: {
    repoUrl () {
      return `dokku@${this.$store.getters.globalDomain}:${this.$route.params.name}`
    }
  }
}
</script>
<style lang="stylus">
.m-app-deploy
  .repo-url-input
    width auto
    margin-left .5rem
    > input
      padding 0 6px
      height 28px
      line-height @height
      border-radius 2px
</style>
