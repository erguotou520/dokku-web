<template>
  <div class="m-app-deploy">
    <el-row class="line-row" :gutter="20">
      <el-col :span="10">
        <h4>{{$t('app.deploy.repo')}}</h4>
      </el-col>
      <el-col :span="14">
        <el-radio-group v-model="origin">
          <el-radio-button label="git">{{$t('app.deploy.origin.git')}}</el-radio-button>
          <el-radio-button label="tar">{{$t('app.deploy.origin.tar')}}</el-radio-button>
        </el-radio-group>
      </el-col>
    </el-row>
    <el-row class="line-row" :gutter="20" v-if="origin==='git'">
      <el-col :span="10">
        <h4>{{$t('app.deploy.label.git')}}</h4>
      </el-col>
      <el-col :span="14">
        <p>
          {{$t('app.deploy.dokkuGit')}}
          <el-input :value="repoUrl" readonly class="repo-url-input">
            <img slot="append" v-clipboard="repoUrl" @success="$message.success($t('message.copied'))"
              @error="$message.error($t('message.copyFailed'))" src="../../assets/images/clippy.svg" style="width:18px;vertical-align:middle">
          </el-input>
        </p>
        <span ref="copy" id="demo">123</span>
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
    <el-row class="line-row" :gutter="20" v-if="origin==='tar'">
      <el-col :span="10">
        <h4>{{$t('app.deploy.label.tar')}}</h4>
      </el-col>
      <el-col :span="14">
        <el-upload drag :action="`/api/apps/${$route.params.name}/tar-deploy`"
          :headers="{Authorization: `Bearer ${$store.getters.accessToken}`}"
          accept=".tar,.tar.gz"
          :on-success="() => { $message.success('Uploaded')}"
          :on-error="() => { $message.error('Upload failed') }">
          <i class="el-icon-upload"></i>
          <div class="el-upload__text" v-html="$t('app.deploy.uploadText')"></div>
        </el-upload>
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
      origin: 'git'
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
    width 16rem
    margin-left .5rem
    > input
      padding 0 6px
      height 28px
      line-height @height
  .line-row + .line-row
    margin-top 1rem
</style>
