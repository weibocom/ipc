import Mock from 'mockjs'

import { register } from './mock.register.js'
import { compare } from './mock.compare.js'
import { content } from './mock.content.js'
import { account } from './mock.account.js'
import { ts } from './mock.ts.js'

Mock.setup({
  timeout: 300
})

function add2Mock(list) {
  list.forEach(n => {
    Mock.mock(n.path, n.method, n.data)
  })
}

if (process.env.NODE_ENV === 'development') {
  add2Mock(register)
  add2Mock(compare)
  add2Mock(content)
  add2Mock(account)
  add2Mock(ts)
}

export default Mock
