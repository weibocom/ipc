import * as API from '../urls.js'

let register_result = {
  code: 200,
  msg: 'ok',
  data: {
    user: {
      id: 800800,
      company: 'wb',
      created_at: '0001-01-01T00:00:00Z',
      private_key:
        '6eee991f99b7198e6d76c8957e5d80e2694264e1ee4232b92e59a8bd10121024',
      public_key: 'STM6yBuUTbZsHR728opLYCbQfK5X4rFZGNBdCcX2pf2bNrcFxbSGB'
    }
  }
}

export const register = [
  {
    method: 'post',
    path: API.account_create,
    data: register_result
  },
  {
    method: 'post',
    path: API.account_create + '?batch=true',
    data: (req, res) => {
      return {
        code: 200,
        msg: 'ok'
      }
    }
  }
]
