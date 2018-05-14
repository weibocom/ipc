import * as API from '../urls.js'

let count_result = {
  retCode: 200,
  msg: 'ok',
  data: {
    count: 112345
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
  }
]
