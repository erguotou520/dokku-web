<template>
  <div class="m-app-settings">
    <el-row class="setting-row" :gutter="20">
      <el-col :span="10">
        <h4>{{$t('app.settings.name')}}</h4>
      </el-col>
      <el-col :span="14">
        <el-row :gutter="20">
          <el-col :span="18">
            <el-input icon="edit" v-model="appName" :on-icon-click="()=>{nameEditable=true}" :readonly="!nameEditable"/>
          </el-col>
          <el-col :span="6">
            <el-button v-if="nameEditable" type="primary" @click="saveName">{{$t('operation.save')}}</el-button>
            <el-button v-if="nameEditable" type="default" @click="nameEditable=false;appName=bak.appName">{{$t('operation.cancel')}}</el-button>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
    <el-row class="setting-row" :gutter="20">
      <el-col :span="10">
        <h4>{{$t('app.settings.environment')}}</h4>
      </el-col>
      <el-col :span="14">
        <el-row v-for="e in envs" class="env-row" :key="e.key" :gutter="20">
          <el-col :span="9">
            <el-input v-model="e.key" readonly></el-input>
          </el-col>
          <el-col :span="9">
            <el-input v-model="e.value" readonly></el-input>
          </el-col>
          <el-col :span="6">
            <el-button type="warning" @click="removeEnvironment(e)">{{$t('operation.remove')}}</el-button>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="9">
            <el-input v-model="env.key" placeholder="KEY"></el-input>
          </el-col>
          <el-col :span="9">
            <el-input v-model="env.value" placeholder="VALUE"></el-input>
          </el-col>
          <el-col :span="6">
            <el-button type="primary" @click="addEnvironment">{{$t('operation.add')}}</el-button>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
    <el-row class="setting-row" :gutter="20">
      <el-col :span="10">
        <h4 class="text-danger">{{$t('app.settings.deleteApp')}}</h4>
      </el-col>
      <el-col :span="14">
        <el-button type="danger" @click="deleteApp" :loading="pending.del">{{$t('app.settings.deleteApp')}}</el-button>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import locales from 'locales/apps'
import { app as res } from 'resources'
export default {
  locales,
  data () {
    return {
      appName: this.$route.params.name,
      nameEditable: false,
      env: {
        key: '',
        value: ''
      },
      envs: [],
      bak: {
        appName: this.$route.params.name
      },
      pending: {
        rename: false,
        del: false
      }
    }
  },
  methods: {
    saveName () {
      if (!this.appName) {
        return this.$message.error('App name required')
      }
      this.pending.rename = true
      res.update({ name: this.bak.appName }, { newName: this.appName }).then(() => {
        this.bak.appName = this.appName
        this.$router.replace(`/apps/${this.appName}/settings`)
      }).catch(e => {
        this.$message.error(JSON.stringify(e) || 'Rename failed')
      }).then(() => {
        this.pending.renmae = false
      })
    },
    addEnvironment () {
      if (!this.env.key || !this.env.value) {
        return this.$message.error('Environment key and value are required')
      }
      res.addEnvironment({ name: this.bak.appName }, this.env).then(() => {
        this.envs.push({ key: this.env.key, value: this.env.value })
        this.env.key = ''
        this.env.value = ''
      }).catch(e => {
        this.$message.error(JSON.stringify(e) || 'Add env failed')
      })
    },
    removeEnvironment (env) {
      res.removeEnvironment({ name: this.bak.appName }, { key: env.key }).then(() => {
        this.envs.splice(this.envs.indexOf(env))
      }).catch(e => {
        this.$message.error(JSON.stringify(e) || 'Remove env failed')
      })
    },
    deleteApp () {
      this.$confirm(this.$t('app.settings.removeAppConfirm'), this.$t('constant.warning'), {
        type: 'danger'
      }).then(() => {
        this.pending.del = true
        res.delete({ name: this.bak.appName }).then(() => {
          this.$message.success(this.$t('app.settings.removed'))
          this.$router.replace('/apps')
        }).catch(e => {
          this.$message.error(JSON.stringify(e) || 'Remove app failed')
        }).then(() => {
          this.pending.del = false
        })
      }).catch(() => {})
    }
  },
  created () {
    res.getEnvironment({ name: this.$route.params.name }).then(data => {
      this.envs = data.data
    }).catch(() => {})
  }
}
</script>
<style lang="stylus">
.m-app-settings
  .setting-row
    margin 1rem 0 2rem
  .env-row
    margin .5rem .75rem
</style>
