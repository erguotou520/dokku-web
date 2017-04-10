import Vue from 'vue'
// things resource
export const thing = Vue.resource('things{/id}')
// users resource
export const user = Vue.resource('users{/id}', {}, {
  changePassword: { method: 'put', url: 'users{/id}/password' }
})
