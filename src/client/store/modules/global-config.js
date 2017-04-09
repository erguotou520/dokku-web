import Vue from 'vue'
import { lang, pageLimit } from '../../stored'
import { save } from '../../storage'
import { STORE_KEY_CONFIG_LANG, STORE_KEY_CONFIG_PAGE_LIMIT } from '../../constants'

const state = {
  lang: lang,
  inited: false,
  // value see http://stackoverflow.com/questions/5580876/navigator-language-list-of-all-languages
  langs: [{
    label: '中文',
    value: 'zh-CN'
  }, {
    label: 'English',
    value: 'en'
  }],
  pageLimit: pageLimit
}

const mutations = {
  UPDATE (state, config) {
    state.lang = config.lang || state.lang
    state.pageLimit = config.pageLimit || state.pageLimit
  },
  UPDATE_LANG (state, lang) {
    state.lang = lang || state.lang
  },
  CHANGE_INITED (state, inited) {
    state.inited = inited
  }
}

const actions = {
  changeLang ({ commit }, lang) {
    Vue.config.lang = lang
    commit('UPDATE_LANG', lang)
    save(STORE_KEY_CONFIG_LANG, lang)
  },
  updateGlobalConfig ({ commit, state, dispatch }, config) {
    if (config.lang !== state.lang) {
      Vue.config.lang = config.lang
      save(STORE_KEY_CONFIG_LANG, config.lang)
    }
    commit('UPDATE', config)
    save(STORE_KEY_CONFIG_LANG, state.lang)
    save(STORE_KEY_CONFIG_PAGE_LIMIT, state.pageLimit)
  },
  fetchInitialize ({ commit }) {
    return Vue.http.get('initialize').then(res => {
      commit('CHANGE_INITED', res.data)
      return res.data
    })
  },
  initialize ({ commit }, form) {
    return Vue.http.post('initialize', form).then(res => {
      commit('CHANGE_INITED', true)
      return res.data
    })
  }
}

const getters = {
  globalConfig (state) {
    return state
  },
  systemInited (state) {
    return state.inited
  }
}

export default {
  state,
  mutations,
  actions,
  getters
}
