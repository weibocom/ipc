import Vue from 'vue'
import App from './App'
import mock from './api/mock'
import router from './router'
import store from './store'
import axios from 'axios'
import ElementUI from 'element-ui'
import 'vue-awesome/icons'
import Icon from 'vue-awesome/components/Icon'
import 'element-ui/lib/theme-chalk/index.css'
import 'element-ui/lib/theme-chalk/display.css'
import 'babel-polyfill'

Vue.use(ElementUI)
Vue.component('icon', Icon)

Vue.prototype.$axios = axios

// 注册全局消息提示
Vue.prototype.$error = msg => {
  Vue.prototype.$message({ message: msg, type: 'error' })
}

Vue.prototype.$warning = msg => {
  Vue.prototype.$message({ message: msg, type: 'warning' })
}

Vue.prototype.$success = msg => {
  if (!msg) {
    Vue.prototype.$message({ message: 'Succeeded', type: 'success' })
  } else {
    Vue.prototype.$message({ message: msg, type: 'success' })
  }
}

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
