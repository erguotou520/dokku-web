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
  }
}
</script>
<style lang="stylus">
/*.m-app-deploy*/
</style>
