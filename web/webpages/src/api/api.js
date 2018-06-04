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
    if (response.data && response.data.code !== 200) {
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
          if (res.data.code && res.data.code !== 200) {
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
  createAccount(data) {
    return ajax(API.account_create, 'post', {}, data)
  },
  lookupAccount(params) {
    return ajax(API.account_lookup, 'get', params)
  },
  detailAccount(id, params) {
    return ajax(API.account_detail + id, 'get', params)
  },
  compare(params) {
    return ajax(API.compare, 'get', params)
  },
  compareText(params) {
    return ajax(API.compare_text, 'get', params)
  },
  createContent(data) {
    return ajax(API.content_create, 'post', {}, data)
  },
  lookupContent(params) {
    return ajax(API.content_create, 'get', params)
  },
  lookupUserContent(params) {
    return ajax(API.user_content_lookup, 'get', params)
  },
  lookupUserCount() {
    return ajax(API.account_count, 'get')
  },
  lookupMsgTs() {
    return ajax(API.msgts_fetch, 'get')
  },
  lookupSimilar(params) {
    return ajax(API.similar_lookup, 'get', params)
  }
}
