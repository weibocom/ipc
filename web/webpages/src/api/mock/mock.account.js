import * as API from '../urls.js'

let count_result = {
  code: 200,
  msg: 'ok',
  data: {
    count: 374
  }
}

export const account = [
  {
    method: 'get',
    path: API.account_count,
    data: (req, res) => {
      count_result.data.count += 1
      return count_result
    }
  },
  {
    method: 'get',
    path: /\/accounts\/\d*\.*/,
    data: (req, res) => {
      return {
        code: 200,
        data: {
          user: {
            company: 'wb',
            created_at: '2018-05-18T19:40:42+08:00',
            id: 800801,
            private_key:
              '24c55c1ec5809123740fd3ae0160239a56a37e8d7b6786569607060217e55929',
            public_key: 'STM61K4G3q7aqYcYN1GYvAB686RkCL853Rym4oUHAB3kh1S3uAcc2',
            post_count: 3
          }
        },
        msg: 'ok'
      }
    }
  },
  {
    method: 'get',
    path: /\/accounts\.*/,
    data: (req, res) => {
      return {
        code: 200,
        data: {
          count: 1234567,
          users: [
            {
              company: 'wb',
              created_at: '2018-05-18T19:37:24+08:00',
              id: 800800,
              private_key:
                '6eee991f99b7198e6d76c8957e5d80e2694264e1ee4232b92e59a8bd10121024',
              public_key:
                'STM6yBuUTbZsHR728opLYCbQfK5X4rFZGNBdCcX2pf2bNrcFxbSGB',
              post_count: 3
            },
            {
              company: 'wb',
              created_at: '2018-05-18T19:40:42+08:00',
              id: 800801,
              private_key:
                '24c55c1ec5809123740fd3ae0160239a56a37e8d7b6786569607060217e55929',
              public_key:
                'STM61K4G3q7aqYcYN1GYvAB686RkCL853Rym4oUHAB3kh1S3uAcc2',
              post_count: 3
            }
          ]
        },
        msg: 'ok'
      }
    }
  }
]
