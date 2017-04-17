import Vue from 'vue'
import Vuex from 'vuex'
import routeLoading from './modules/route'
import config from './modules/global-config'
import user from './modules/user'
import octicon from './modules/octicon'
Vue.use(Vuex)

const store = new Vuex.Store({
  strict: process.env.NODE_ENV !== 'production',
  modules: {
    user,
    config,
    routeLoading,
    octicon
  }
})

export default store
