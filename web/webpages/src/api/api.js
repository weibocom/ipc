import Vue from 'vue'
import axios from 'axios'
import router from '@/router'
import * as API from './urls.js'
import storage from '@/utils/storage'
import { STORAGE_KEY } from '@/utils/constants'

// http request 拦截器
axios.interceptors.request.use(
  config => {
    return config
  },
  err => {
    return Promise.reject(err)
  }
)

// http response 拦截器
axios.interceptors.response.use(
  response => {
    if (response.data && response.data.retCode !== 200) {
      if (response.data.msg) {
        Vue.prototype.$error(response.data.msg)
      } else {
        console.log('http 200 error data without msg: ' + response.data)
        Vue.prototype.$error('系统繁忙，请重试')
      }
    }
    return response
  },
  error => {
    Vue.prototype.$error('系统繁忙，请重试')
    console.log(error)
    // 返回接口返回的错误信息
    return Promise.reject(error)
  }
)

Vue.prototype.$http = axios
axios.defaults.baseURL = ''

/**
 * @param url
 * @param method get|post|put|delete...
 * @param params like queryString. if a url is index?a=1&b=2, params = {a: '1', b: '2'}
 * @param data post data, use for method put|post
 * @returns {Promise}
 */
function ajax(url, method, params, data) {
  if (params === undefined) {
    params = {}
  }
  if (data === undefined) {
    data = {}
  }
  return new Promise((resolve, reject) => {
    axios({
      url,
      method,
      params,
      data
    })
      .then(
        res => {
          // API正常返回(status=20x), 是否错误通过有无error判断
          if (res.data.retCode && res.data.retCode !== 200) {
            reject(res)
          } else {
            resolve(res)
          }
        },
        res => {
          // API请求异常
          reject(res)
        }
      )
      .catch(res => {
        window.location.reload()
      })
  })
}

export default {
  // 注册
  register(params) {
    return ajax(API.register, 'get', params)
  },
  search(params) {
    return ajax(API.search, 'get', params)
  },
  compare(params) {
    return ajax(API.compare, 'get', params)
  },
  compare_resource(params) {
    return ajax(API.compare_resource, 'get', params)
  },
  submit(params) {
    return ajax(API.submit, 'get', params)
  },
  lookupUserCount() {
    return ajax(API.account_count, 'get')
  },
  lookupMsgTs() {
    return ajax(API.msgts_fetch, 'get')
  }
}
