# web restful api

## login/logout

### login

### logout

## Account

账户(account/user/author)管理。

router.GET("/counts/account", auth(accountCount))
	router.POST("/accounts", auth(createAccount))
	router.GET("/accounts", auth(queryAccount))
	router.GET("/accounts/:account", auth(accountDetail))

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
curl "http://127.0.0.1:8080/accounts/800801?company=weibo"
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


## dci