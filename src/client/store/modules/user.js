import { merge } from 'lodash'
import { saveMulti, clearMulti } from '../../storage'
import { login, getUserInfo } from './user.api'
// eslint-disable-next-line camelcase
import { username, access_token, refresh_token } from '../../stored'
import { STORE_KEY_USERNAME, STORE_KEY_ACCESS_TOKEN, STORE_KEY_REFRESH_TOKEN } from '../../constants'

const state = {
  id: '',
  role: 'guest',
  username: username,
  access_token, // eslint-disable-line
  refresh_token // eslint-disable-line
}

const mutations = {
  // set user info
  SET_USER_INFO (state, userInfo) {
    merge(state, userInfo)
  },
  // after logout
  LOGOUT (state) {
    state.id = ''
    state.username = ''
    state.role = 'guest'
    state.access_token = '' // eslint-disable-line
    state.refresh_token = '' // eslint-disable-line
  }
}

const actions = {
  // set user info
  setUserInfo ({ commit }, user) {
    commit('SET_USER_INFO', user)
    saveMulti([{
      key: STORE_KEY_USERNAME,
      value: user.username
    }, {
      key: STORE_KEY_ACCESS_TOKEN,
      value: user.access_token // eslint-disable-line
    }, {
      key: STORE_KEY_REFRESH_TOKEN,
      value: user.refresh_token // eslint-disable-line
    }])
  },
  // init user info
  initUserInfo ({ commit, state }) {
    return new Promise((resolve, reject) => {
      // token
      if (username) {
        getUserInfo(state.access_token).then(data => { // eslint-disable-line
          if (data.id) {
            commit('SET_USER_INFO', data)
          }
          resolve(data)
        }).catch(err => { reject(err) })
      } else {
        resolve()
      }
    })
  },
  // login action
  login ({ commit }, payload) {
    return new Promise((resolve, reject) => {
      login(payload.username, payload.password).then(data => {
        getUserInfo(data.token).then(user => {
          const userInfo = merge({}, user, {
            username: payload.username,
            access_token: data.token, // eslint-disable-line
            refresh_token: '' // eslint-disable-line
          })
          commit('SET_USER_INFO', userInfo)
          saveMulti([{
            key: STORE_KEY_USERNAME,
            value: userInfo.username
          }, {
            key: STORE_KEY_ACCESS_TOKEN,
            value: userInfo.access_token // eslint-disable-line
          }, {
            key: STORE_KEY_REFRESH_TOKEN,
            value: userInfo.refresh_token // eslint-disable-line
          }])
          resolve()
        }).catch(() => {})
      }).catch(err => { reject(err) })
    })
  },
  // refresh token action
  refreshToken ({ commit }, payload) {
    commit('REFERE_TOKEN', payload)
    saveMulti([{
      key: STORE_KEY_ACCESS_TOKEN,
      value: payload.access_token // eslint-disable-line
    }, {
      key: STORE_KEY_REFRESH_TOKEN,
      value: payload.refresh_token // eslint-disable-line
    }])
  },
  // logout action
  logout ({ commit }, payload) {
    commit('LOGOUT')
    clearMulti([STORE_KEY_USERNAME, STORE_KEY_ACCESS_TOKEN, STORE_KEY_REFRESH_TOKEN])
  }
}

const getters = {
  userId (state) {
    return state.id
  },
  userRole (state) {
    return state.role
  },
  accessToken (state) {
    return state.access_token // eslint-disable-line
  },
  username (state) {
    return state.username
  },
  loggedIn (state) {
    return !!(state.username && state.access_token) // eslint-disable-line
  }
}

export default {
  state,
  mutations,
  actions,
  getters
}
