import Vue from 'vue'
// read localStorage stored data
import './stored'
// locale
import './locales'

// router and store
import store from './store'
import router, { hook as routerHook } from './router'
import { sync } from 'vuex-router-sync'
sync(store, router)

// ui library
import './ui'
// ajax
import './http'

// main component
import App from './App'

// import './socket'

store.dispatch('fetchInitialize').then(res => {
  const userPromise = res ? store.dispatch('initUserInfo') : store.dispatch('logout')
  routerHook(userPromise)

  userPromise.then(() => {
    const app = new Vue({
      router,
      store,
      ...App // Object spread copying everything from App.vue
    })
    // actually mount to DOM
    app.$mount('app')
  })
})
