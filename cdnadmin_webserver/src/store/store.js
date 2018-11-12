import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    userInfo: {
      token: '',
      username: '',
      roles: ''
    }
  },
  mutations: {
    setToken (state, data) {
      state.userInfo.token = data.token
      window.sessionStorage.setItem('token', data.token)
    },
    setUserName (state, data) {
      state.userInfo.username = data.username
      window.sessionStorage.setItem('username', data.username)
    },
    setRoles (state, data) {
      if (data.is_admin === 0) {
        state.userInfo.roles = 'admin'
      } else {
        state.userInfo.roles = ''
      }
      window.sessionStorage.setItem('roles', state.userInfo.roles)
    },
    logout (state) {
      state.userInfo.username = ''
      window.sessionStorage.removeItem('username')
      state.userInfo.roles = []
      window.sessionStorage.removeItem('roles')
    }
  }
})

export default store
