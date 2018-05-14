import * as API from '../urls.js'
import { param2Obj } from './utils.js'

let text_compare_result = {
  retCode: 200,
  msg: 'ok',
  data: { dstContent: '测试test', similarity: '0.00%', srcContent: 'hahahah' }
}

let resource_compare_result = {
  retCode: 200,
  msg: 'ok',
  data: {
    dstContent: '测试test',
    similarity: '23.14%',
    srcContent: '测一下123'
  }
}

export const compare = [
  {
    method: 'get',
    path: /\/compare.json\.*/,
    data: (req, res) => {
      let params = param2Obj(req.url)
      if (params.srcCompany !== undefined && params.srcUid !== undefined && params.srcMid !== undefined) {
        return resource_compare_result
      }
      return text_compare_result
    }
  }
]
