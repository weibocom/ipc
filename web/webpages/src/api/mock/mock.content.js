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
  },
  {
    method: 'get',
    path: /\/similar\/post\.*/,
    data: (req, res) => {
      let response = {
        code: 200,
        msg: 'ok',
        data: {
          posts: []
        }
      }

      let pagesize = 10
      let params = param2Obj(req.url)
      if (params.page == 3) {
        pagesize = 5
      }
      for (let i = 0; i < pagesize; i++) {
        response.data.posts.push({
          similarity: '98.76',
          mid: 4246160954700587,
          dna:
            '2031ce9ad9a7898247af5dc08a81b009e7eead0e2e45a148dc5ed021e84078e2e517b69e17a6f5d2a67f9d3c24f42d9d16251725234ba25dcd9708d01a5af5ff8b',
          author: '5579703430',
          title: '4246160954700587',
          content:
            '午安，六一儿童节，节日快乐哟[爱你][爱你][爱你] ​ 各位宝宝们，六一怎么可以没礼物呢[挤眼]你们的福利来了~\\(≧▽≦)/~\n还是老规矩①转发#怦然心动漫画[超话]#的套装链接到朋友圈并截图②关注、转发本微博在评论留下你的截图 ③ @ 三名有联系的好友\n就可以参与本次的抽奖活动啦。\n \n活动截止到6月11日届时会抽取一名幸运粉丝赠送怦然心动套装哦o(≧v≦)o宝宝们快点行动起来吧！',
          keywords: '六一,套装,宝宝,怦然心动,截图,转发',
          uri: '4246160954700587',
          digest:
            '051ae08f531a95dea250bdde87b854386ebeb4223a97686b173482b36bb1dd7d',
          created_at: '2018-06-01T17:58:48+08:00'
        })
      }

      return response
    }
  }
]
