import * as API from '../urls.js'

let ts_result = {
  code: 200,
  msg: 'ok',
  data: { timestamp: '2018-05-22T11:47:42+08:00' }
}

export const ts = [
  {
    method: 'get',
    path: API.msgts_fetch,
    data: (req, res) => {
      return ts_result
    }
  }
]
