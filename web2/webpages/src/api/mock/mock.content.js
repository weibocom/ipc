import * as API from '../urls.js'
import { param2Obj } from './utils.js'

export const content = [
  {
    method: 'post',
    path: API.content_create,
    data: (req, res) => {
      return {
        code: 200,
        data: {
          post: {
            author: 'weibo-800820',
            content: '\u5317\u4eac\u73b0\u5728\u8fdb\u5165\u4e86\u96e8\u5b63',
            created_at: '2018-05-21T11:24:45+08:00',
            digest:
              '5fb7d18d6184bdb2e48982e4ee6afd95479516f668ef1b204a230cb5df63c19e',
            dna:
              '201cc923a5df9d8d814ff48382bfbc6f9a8148fe9d20f9ac8c638d46990ec9aaff19086841be78a3eac0bf9056d0ef4c12e612bdb7890955ab414ab7ce7f210be5',
            mid: 400401,
            title: 'weibo-800820-400401',
            uri: '400401'
          }
        },
        msg: 'ok'
      }
    }
  },
  {
    method: 'get',
    path: /\/posts\.*/,
    data: (req, res) => {
      return {
        code: 200,
        data: {
          post: {
            author: 'weibo-800820',
            content: '\u5317\u4eac\u73b0\u5728\u8fdb\u5165\u4e86\u96e8\u5b63',
            created_at: '2018-05-21T11:24:45+08:00',
            digest:
              '5fb7d18d6184bdb2e48982e4ee6afd95479516f668ef1b204a230cb5df63c19e',
            dna:
              '201cc923a5df9d8d814ff48382bfbc6f9a8148fe9d20f9ac8c638d46990ec9aaff19086841be78a3eac0bf9056d0ef4c12e612bdb7890955ab414ab7ce7f210be5',
            mid: 400401,
            title: 'weibo-800820-400401',
            uri: '400401'
          }
        },
        msg: 'ok'
      }
    }
  },
  {
    method: 'get',
    path: /\/account_posts\.*/,
    data: (req, res) => {
      return {
        code: 200,
        data: {
          post_count: 1,
          posts: [
            {
              author: 'weibo-800820',
              content: '\u5317\u4eac\u73b0\u5728\u8fdb\u5165\u4e86\u96e8\u5b63',
              created_at: '2018-05-21T11:24:45+08:00',
              digest:
                '5fb7d18d6184bdb2e48982e4ee6afd95479516f668ef1b204a230cb5df63c19e',
              dna:
                '201cc923a5df9d8d814ff48382bfbc6f9a8148fe9d20f9ac8c638d46990ec9aaff19086841be78a3eac0bf9056d0ef4c12e612bdb7890955ab414ab7ce7f210be5',
              mid: 400401,
              title: 'weibo-800820-400401',
              uri: '400401'
            }
          ]
        },
        msg: 'ok'
      }
    }
  }
]
