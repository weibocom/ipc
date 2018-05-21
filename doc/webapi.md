# web restful api

## login/logout

### login

### logout

## Account

### 获取账号总数

- URL: http://127.0.0.1:8080/counts/account
- HTTP METHOD: GET
- 参数
  

示例:

**请求**:
```
curl http://127.0.0.1:8080/counts/account
```

**返回结果**:

```
{
    "code": 200, 
    "msg": "ok", 
    "data": {
        "count": 374
    }
}
```

### 创建账号

- URL: http://127.0.0.1:8080/accounts
- HTTP METHOD: POST
- 参数
  - company: 公司名英文简称
  - uid: 用户id
  

示例:

**请求**:
```
curl http://127.0.0.1:8080/accounts -d "company=weibo&uid=800800"
```

**返回结果**:

```
{
    "code": 200, 
    "msg": "ok", 
    "data": {
        "user": {
            "id": 800800, 
            "company": "weibo", 
            "created_at": "0001-01-01T00:00:00Z", 
            "private_key": "6eee991f99b7198e6d76c8957e5d80e2694264e1ee4232b92e59a8bd10121024", 
            "public_key": "STM6yBuUTbZsHR728opLYCbQfK5X4rFZGNBdCcX2pf2bNrcFxbSGB"
        }
    }
}
```


### 查询账号

- URL: http://127.0.0.1:8080/accounts
- HTTP METHOD: GET
- 参数
  - company: 公司名英文简称
  - page: 页码, 从1开始
  - pagesize: 每页账号数
  

示例:

**请求**:
```
curl "http://127.0.0.1:8080/accounts?company=weibo&page=1&pagesize=10"
```

**返回结果**:

```
{
    "code": 200,
    "data": {
        "users": [
            {
                "company": "weibo",
                "created_at": "2018-05-18T19:37:24+08:00",
                "id": 800800,
                "private_key": "6eee991f99b7198e6d76c8957e5d80e2694264e1ee4232b92e59a8bd10121024",
                "public_key": "STM6yBuUTbZsHR728opLYCbQfK5X4rFZGNBdCcX2pf2bNrcFxbSGB"
            },
            {
                "company": "weibo",
                "created_at": "2018-05-18T19:40:42+08:00",
                "id": 800801,
                "private_key": "24c55c1ec5809123740fd3ae0160239a56a37e8d7b6786569607060217e55929",
                "public_key": "STM61K4G3q7aqYcYN1GYvAB686RkCL853Rym4oUHAB3kh1S3uAcc2"
            }
        ]
    },
    "msg": "ok"
}
```

### 查询账号详情

- URL: http://127.0.0.1:8080/accounts/:uid
- HTTP METHOD: GET
- 参数
  - company: 公司名英文简称
  - :uid: 用户uid
  

示例:

**请求**:
```
curl "http://127.0.0.1:8080/accounts/800820?company=weibo"
```

**返回结果**:

```
{
    "code": 200,
    "data": {
        "user": {
            "company": "weibo",
            "created_at": "2018-05-18T19:40:42+08:00",
            "id": 800801,
            "private_key": "24c55c1ec5809123740fd3ae0160239a56a37e8d7b6786569607060217e55929",
            "public_key": "STM61K4G3q7aqYcYN1GYvAB686RkCL853Rym4oUHAB3kh1S3uAcc2"
        }
    },
    "msg": "ok"
}
```

## Post

### 内容发布

- URL: http://127.0.0.1:8080/posts
- HTTP METHOD: POST
- 参数
  - uid: 用户id
  - company: 公司名英文简称
  - mid: 内容id
  - content: 内容
  - contentType: 内容类型，0: 文章, 1: 图片, 2: 视频， 缺省是0
  - title: 文章标题， optional。
  


示例:

**请求**:
```
curl http://127.0.0.1:8080/posts -d "company=weibo&uid=800820&mid=400401&content=北京现在进入了雨季"
```

**返回结果**:

```
{
    "code": 200,
    "data": {
        "post": {
            "author": "weibo-800820",
            "content": "\u5317\u4eac\u73b0\u5728\u8fdb\u5165\u4e86\u96e8\u5b63",
            "created_at": "2018-05-21T11:24:45+08:00",
            "digest": "5fb7d18d6184bdb2e48982e4ee6afd95479516f668ef1b204a230cb5df63c19e",
            "dna": "201cc923a5df9d8d814ff48382bfbc6f9a8148fe9d20f9ac8c638d46990ec9aaff19086841be78a3eac0bf9056d0ef4c12e612bdb7890955ab414ab7ce7f210be5",
            "mid": 400401,
            "title": "weibo-800820-400401",
            "uri": "400401"
        }
    },
    "msg": "ok"
}
```

### 根据uid和mid查询内容

- URL: http://127.0.0.1:8080/posts
- HTTP METHOD: GET
- 参数
  - uid: 用户id
  - company: 公司名英文简称
  - mid: 内容id
  - queryType: user (按照uid和mid进行查询)


示例:

**请求**:
```
curl "http://127.0.0.1:8080/posts?queryType=user&company=weibo&uid=800820&mid=400401"
```

**返回结果**:

```
{
    "code": 200,
    "data": {
        "post": {
            "author": "weibo-800820",
            "content": "\u5317\u4eac\u73b0\u5728\u8fdb\u5165\u4e86\u96e8\u5b63",
            "created_at": "2018-05-21T11:24:45+08:00",
            "digest": "5fb7d18d6184bdb2e48982e4ee6afd95479516f668ef1b204a230cb5df63c19e",
            "dna": "201cc923a5df9d8d814ff48382bfbc6f9a8148fe9d20f9ac8c638d46990ec9aaff19086841be78a3eac0bf9056d0ef4c12e612bdb7890955ab414ab7ce7f210be5",
            "mid": 400401,
            "title": "weibo-800820-400401",
            "uri": "400401"
        }
    },
    "msg": "ok"
}
```

### 根据dna查询内容

- URL: http://127.0.0.1:8080/posts
- HTTP METHOD: GET
- 参数
  - dna: 内容的唯一签名
  - queryType: dna (按照uid和mid进行查询)


示例:

**请求**:
```
curl "http://127.0.0.1:8080/posts?queryType=dna&dna=201cc923a5df9d8d814ff48382bfbc6f9a8148fe9d20f9ac8c638d46990ec9aaff19086841be78a3eac0bf9056d0ef4c12e612bdb7890955ab414ab7ce7f210be5"
```

**返回结果**:

```
{
    "code": 200,
    "data": {
        "post": {
            "author": "weibo-800820",
            "content": "\u5317\u4eac\u73b0\u5728\u8fdb\u5165\u4e86\u96e8\u5b63",
            "created_at": "2018-05-21T11:24:45+08:00",
            "digest": "5fb7d18d6184bdb2e48982e4ee6afd95479516f668ef1b204a230cb5df63c19e",
            "dna": "201cc923a5df9d8d814ff48382bfbc6f9a8148fe9d20f9ac8c638d46990ec9aaff19086841be78a3eac0bf9056d0ef4c12e612bdb7890955ab414ab7ce7f210be5",
            "mid": 400401,
            "title": "weibo-800820-400401",
            "uri": "400401"
        }
    },
    "msg": "ok"
}
```

## dci