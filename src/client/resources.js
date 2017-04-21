import Vue from 'vue'
export const user = Vue.resource('users{/id}', {}, {
  changePassword: { method: 'put', url: 'users{/id}/password' }
})
// app api
export const app = Vue.resource('apps{/name}', {}, {
  operate: { method: 'PUT', url: 'apps/ps{/name}' }
})

// plugin api
export const plugin = Vue.resource('plugins{/name}')
