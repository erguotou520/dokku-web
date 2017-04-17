const state = {
  icons: {}
}

const mutations = {
  REGISTER (state, data) {
    for (const name in data) {
      state.icons[name] = data[name]
    }
  }
}

const actions = {
  registerOcticon ({ commit }, data) {
    commit('REGISTER', data)
  }
}

const getters = {
  octiconIcons (state) {
    return state.icons
  }
}

export default {
  state,
  mutations,
  actions,
  getters
}
