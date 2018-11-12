// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import 'babel-polyfill'

import Vue from 'vue'
import App from './App'

import router from './router'
import Vuex from 'vuex'
import axios from 'axios'

import store from '@/store/store.js'

Vue.use(Vuex)
Vue.use(ElementUI)

axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'

// set request header add token
axios.interceptors.request.use(function (config) {
  // if (localStorage.token && router.currentRoute.fullPath !== '/login') {
  const token = store.state.userInfo.token ? store.state.userInfo.token : window.sessionStorage.getItem('token')
  const username = store.state.userInfo.username ? store.state.userInfo.username : window.sessionStorage.getItem('username')
  if (token) {
    config.headers.authorization = `${token}`
    config.headers.username = `${username}`
  }
  return config
}, function (err) {
  return Promise.reject(err)
})

// response 401 redirect login
axios.interceptors.response.use(function (res) {
  if (res.data.code === 401) {
    store.commit('setToken', '')
    router.replace({
      path: '/login',
      query: {redirect: router.currentRoute.fullPath}
    })
    this.$notify({
      title: 'waring',
      type: 'waring',
      message: 'login信息失效，请重新login'
    })
  }
  return res
}, function (error) {
  return Promise.reject(error)
}
)

router.beforeEach((to, from, next) => {
  if (to.meta.requireAuth) {
    const token = store.state.userInfo.token ? store.state.userInfo.token : window.sessionStorage.getItem('token')
    if (token) {
      next()
    } else {
      next({
        path: '/login',
        query: {redirect: to.fullPath}
      })
    }
  } else {
    next()
  }
})

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
  // components: { App },
  // template: '<App/>'
})
