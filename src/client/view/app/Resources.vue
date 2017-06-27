<template>
  <div class="m-app-resources flex flex-column flex-main-center">
    <el-alert class="plugin-tip" :title="$t('app.resources.tip')" type="info"></el-alert>
    <p>{{$t('app.plugin')}}</p>
    <div class="flex">
      <el-autocomplete v-model="search" icon="search" class="flex-1"
        :fetch-suggestions="querySearch" :placeholder="$t('app.resources.searchTip')"
        :trigger-on-focus="false" @select="handleSelect" custom-item="plugin-item">
      </el-autocomplete>
      <el-button type="primary" v-if="selectedPlugin" @click="install" style="margin-left:1rem">{{$t('app.resources.install')}}</el-button>
    </div>
    <el-table :data="installed" border :empty-text="$t('app.resources.empty')" style="margin-top:1rem;width: 100%">
      <el-table-column align="center" :label="$t('app.resources.plugin.name')" prop="name"></el-table-column>
      <el-table-column align="center" :label="$t('app.resources.plugin.version')" prop="version"></el-table-column>
      <el-table-column align="center" :label="$t('app.resources.plugin.author')" prop="author"></el-table-column>
      <el-table-column align="center" :label="$t('app.resources.plugin.status')" prop="enabled">
        <template scope="scope">
          {{scope.row.enabled?$t('status.enabled'):$t('status.disabled')}}
          (<el-button @click="toggleStatus(scope.row)" type="text">{{scope.row.enabled?$t('operation.enable'):$t('operation.disable')}}</el-button>)
        </template>
      </el-table-column>
      <el-table-column align="center" :label="$t('app.resources.plugin.description')" prop="description"></el-table-column>
      <el-table-column align="center" :label="$t('operation.operation')">
        <template scope="scope">
          <el-button v-if="scope.row.author==='dokku'" @click="create(scope.row)" type="text" size="small">{{$t('operation.create')}}</el-button>
          <el-button v-if="scope.row.author==='dokku'" @click="link(scope.row)" type="text" size="small">{{$t('app.resources.link')}}</el-button>
          <el-button v-if="scope.row.author==='dokku'" @click="unlink(scope.row)" type="text" size="small">{{$t('app.resources.unlink')}}</el-button>
          <el-popover placement="top" width="160">
            <el-button size="mini" type="text">取消</el-button>
            <el-button type="primary" size="mini" @click="uninstall(scope.row)">确定</el-button>
            <el-button slot="reference" v-if="scope.row.author==='dokku'" type="text" size="small">{{$t('app.resources.uninstall')}}</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>
    <div class="plugin-link flex flex-cross-center" v-show="selectedPlugin">
      <div class="img-wrapper t-c flex-1">
        <div class="flex flex-main-center flex-cross-center" style="height:96px">
          <img class="plugin-img" :src="pluginImage" :alt="pluginImage" width="96">
        </div>
        <p class="plugin-name">{{selectedPlugin?selectedPlugin.name:''}}</p>
      </div>
      <octicon name="arrowR" class="link-line" :w="32" :h="32" fill="#a98"></octicon>
      <div class="plugin-item t-c flex-1">
        <octicon name="hexagon" :w="96" :h="96" fill="#07cbcf"></octicon>
        <p>{{$route.params.name}}</p>
      </div>
    </div>
    <el-button class="link-btn" type="primary" v-show="selectedPlugin" @click="linkPlugin">{{$t('app.resources.link')}}</el-button>
  </div>
</template>
<script>
import Vue from 'vue'
import { merge, search } from 'vtc'
import { plugin as res, app as aRes } from 'resources'
import locales from 'locales/apps'

const SUPPORTED_IMAGES = ['couchdb', 'elasticsearch', 'mariadb', 'memcached', 'mongodb', 'mysql', 'nats', 'postgresql', 'rabbitmq', 'redis', 'rethinkdb']
const DEFAULT_IMAGE = '/static/dokku.png'

Vue.component('plugin-item', {
  functional: true,
  render: (h, ctx) => {
    var item = ctx.props.item
    return h('li', Object.assign({ staticClass: 'flex', style: { paddingRight: '2rem' }}, ctx.data), [
      h('span', { attrs: { class: 'flex-1' }}, [item.name]),
      h('span', { attrs: { class: 'flex-1 t-c' }}, [item.compatibility]),
      h('span', { attrs: { class: 'flex-1 t-r' }}, [item.author])
    ])
  },
  props: {
    item: { type: Object, required: true }
  }
})

export default {
  locales,
  data () {
    return {
      search: '',
      plugins: {
        official: [],
        community: [],
        others: []
      },
      installed: [],
      selectedPlugin: null
    }
  },
  computed: {
    computedPlugins () {
      return [].concat(this.plugins.official).concat(this.plugins.community).concat(this.plugins.others)
    },
    pluginImage () {
      if (this.selectedPlugin) {
        const name = this.selectedPlugin.name.toLowerCase()
        const filtered = SUPPORTED_IMAGES.filter(img => name.indexOf(img) > -1)
        return filtered.length ? `/static/${filtered[0]}.png` : DEFAULT_IMAGE
      } else {
        return DEFAULT_IMAGE
      }
    }
  },
  methods: {
    fetch () {
      res.getInstalled().then(data => {
        this.installed = data.data
      }).catch(() => {})
    },
    install () {
      res.save(null, { url: this.selectedPlugin.url }).then(() => {
        this.$message.success(this.$t('app.resources.installSuccess'))
        this.fetch()
      })
    },
    toggleStatus (plugin) {
      res.togglePluginStatus({ name: plugin.name }, {
        action: plugin.enabled ? 'disable' : 'enable'
      }).then(data => {
        if (data.data) {
          plugin.enabled = !plugin.enabled
        } else {
          this.$message.error(this.$t('app.resources.toggleFailed'))
        }
      })
    },
    querySearch (query, cb) {
      if (!query) {
        cb([])
      } else {
        cb(search(this.computedPlugins, query))
      }
    },
    handleSelect (plugin) {
      this.selectedPlugin = plugin
      this.search = plugin.name
    },
    linkPlugin () {
      if (this.selectedPlugin) {
        const pluginName = this.pluginName !== DEFAULT_IMAGE ? this.pluginName.match(/^\/static\/(\w+).png$/)[1] : encodeURIComponent(this.pluginName)
        aRes.linkPlugin({ name: this.$route.params.name, pluginName }, {
          url: this.selectedPlugin.url
        }).then(() => {
          this.$message.success(this.$t('app.resources.linkSuccess', this.selectedPlugin.name))
          this.search = ''
          this.selectedPlugin = null
        }).catch(() => {
          this.$message.error(this.$t('app.resources.linkFailed', this.selectedPlugin.name))
        })
      }
    },
    uninstall (plugin) {
      const nameObj = { name: plugin.name }
      res.remove(nameObj).then(() => {
        this.$message.success(this.$t('app.resources.uninstallSuccess', nameObj))
        this.plugins.splice(this.plugins.indexOf(plugin), 1)
      }).catch(() => {
        this.$message.error(this.$t('app.resources.uninstallFailed', nameObj))
      })
    }
  },
  created () {
    res.query().then(data => {
      merge(this.plugins, data.data)
    }).catch(() => {})
    this.fetch()
  }
}
</script>
<style lang="stylus">
.m-app-resources
  .plugin-tip
    background-color #f3eb97
    color #756e22
    .el-alert__closebtn
      color #756e22
  .plugin-link
    margin 2rem auto 0
    width 24rem
  .link-line
    margin -1rem .75rem 0
  .link-btn
    width 24rem
    margin 2rem auto 0
</style>
