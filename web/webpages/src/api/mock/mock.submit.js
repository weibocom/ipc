import * as API from '../urls.js'
import { param2Obj } from './utils.js'

let submit_result = {
  retCode: 200,
  msg: 'ok',
  data: {
    blockInfo: {
      blockCreateTime: '2018-01-01 12:33:21',
      blockTitle: 'thisistitle',
      authorName: 'wb-testauthor',
      meta: '{\"tags\":[\"test\",\"dontvote\"],\"app\":\"SteemJ-Core/0.5.0-SNAPSHOT\",\"title\":\"thisistitle\", \"time\":1522145063295,\"format\":\"markdown\"}',
      url: 'linktonowhere',
      content: 'message content detail'
    }
  }
}

let search_result = {
  retCode: 200,
  msg: 'ok',
  data: {
    blockInfo: {
      blockCreateTime: '2018-01-01 12:33:21',
      blockTitle: 'thisistitle',
      authorName: 'wb-testauthor',
      meta: '{\"tags\":[\"test\",\"dontvote\"],\"app\":\"SteemJ-Core/0.5.0-SNAPSHOT\",\"title\":\"thisistitle\", \"time\":1522145063295,\"format\":\"markdown\"}',
      url: 'linktonowhere',
      content: 'message content detail'
    }
  }
}

export const submit = [
  {
    method: 'get',
    path: /\/submit.json\.*/,
    data: (req, res) => {
      let params = param2Obj(req.url)
      switch (params.action) {
        case 'write':
          return submit_result
        case 'search':
          return search_result
        default:
          return { msg: 'invalid action', retCode: 404 }
      }
    }
  }
]
