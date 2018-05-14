import * as API from '../urls.js'

let register_result = {
  retCode: 200,
  msg: 'ok',
  data: {
    uid: '11235813',
    company: 'weibo'
  }
}

export const register = [
  {
    method: 'get',
    path: /\/register.json\.*/,
    data: register_result
  }
]
