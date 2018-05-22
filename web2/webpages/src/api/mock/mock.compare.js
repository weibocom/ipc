import * as API from '../urls.js'
import { param2Obj } from './utils.js'

export const compare = [
  {
    method: 'get',
    path: /\/dci\/content\.*/,
    data: (req, res) => {
      return {
        code: 200,
        data: {
          dst: '\u5317\u4eac\u73b0\u5728\u8fdb\u5165\u4e86\u96e8\u5b63',
          similarity: '100.00',
          src: '\u5317\u4eac\u73b0\u5728\u8fdb\u5165\u4e86\u96e8\u5b63'
        },
        msg: 'ok'
      }
    }
  },
  {
    method: 'get',
    path: /\/dci\/text\.*/,
    data: (req, res) => {
      return {
        code: 200,
        data: {
          dst: '\u4eca\u5929\u5317\u4eac\u4f1a\u4e0b\u96e8\u5417',
          similarity: '20.00',
          src: '\u5317\u4eac\u73b0\u5728\u8fdb\u5165\u4e86\u96e8\u5b63'
        },
        msg: 'ok'
      }
    }
  }
]
