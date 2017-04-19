<template>
  <div class="m-app-deploy">
    <el-row :gutter="20">
      <el-col :span="4">
        <h4>{{$t('app.deploy.repo')}}</h4>
      </el-col>
      <el-col :span="8">
        <el-radio-group v-model="selectedRepo">
          <el-radio-button label="dokku">Dokku Git</el-radio-button>
          <el-radio-button label="github">GitHub</el-radio-button>
        </el-radio-group>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="4">
        <h4>{{$t('app.deploy.tipLabel')}}</h4>
      </el-col>
      <el-col :span="8">
        <p>{{$t('app.deploy.dokkuGit')}}<el-input :value="repoUrl" readonly></el-input></p>
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
export default {
  data () {
    return {
      selectedRepo: 'dokku'
    }
  },
  computed: {
    repoUrl () {
      return `dokku@dokku.me:${this.$route.params.name}`
    }
  }
}
</script>
<style lang="stylus">
/*.m-app-deploy*/
</style>
