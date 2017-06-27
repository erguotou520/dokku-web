import Vue from 'vue'
export const user = Vue.resource('users{/id}', {}, {
  changePassword: { method: 'put', url: 'users{/id}/password' }
})
// app api
export const app = Vue.resource('apps{/name}', {}, {
  operate: { method: 'PUT', url: 'apps/ps{/name}' },
  getEnvironment: { method: 'GET', url: 'apps{/name}/envs' },
  addEnvironment: { method: 'PUT', url: 'apps{/name}/addEnv' },
  removeEnvironment: { method: 'PUT', url: 'apps{/name}/removeEnv' },
  getAppLogs: { method: 'GET', url: 'apps{/name}/logs' },
  getAppGitLogs: { method: 'GET', url: 'apps{/name}/git-logs' },
  linkPlugin: { method: 'GET', url: 'apps{/name}/link{/pluginName}' }
})

// plugin api
export const plugin = Vue.resource('plugins{/name}', {}, {
  getInstalled: { method: 'GET', url: 'installed-plugins' },
  togglePluginStatus: { method: 'PUT', url: 'plugins{/name}/status' }
})
