# Open Auth

## docker部署说明

见 `github.com/huyinghuan/oauth/docker-compose.yaml`

将里面的 `services.oauth.image`修改为`huyinghuan/oauth:latest`，

运行 `docker-compose up`即可

##  nginx 配置说明

```
server {
        listen 80;
        server_name d.imgo.tv;
        
        underscores_in_headers on; #需增加
        location / {
                 proxy_pass_request_headers      on; #需增加
                proxy_pass http://127.0.0.1:8000;
        }
}
```

## API说明

### timestamp

以下出现的`timestamp`参数均为int值，为当前时间的以秒为单位的时间戳。
如nodejs: `Date.now()/1000`, golang: `time.Now().Unix()`, 

数据通讯过程中，传到服务器的`timestamp`， 服务器将校验该值，如果该值与服务器误差大于5分钟，请求将被拒绝
服务器API响应的`timestamp`, 请应用自行判断该值是否可信。

### 加密方式

以下所有用到加密解密的方式均采用`aes-cfb-256`

## 获取登陆用户Token

- 1.将网页URL重定向到`/authorize?client_id=xxx&redirect=xxxxx`, 该URL会自动引导用户进行登陆操作
- 2.用户在登陆以后，将会由该URL跳转到`your_app_server_callback_url?token=xxxx&redirect=xxx&t=timestamp`，这样你到应用拿到了一个60秒内有效到token值。


## 用token 获取登陆用户信息

### API

```
URL: /resource/account
Method:  POST
Header:
  client-id=your_app_client_id

Body为将数据格式 
  {
    timestamp: xxx,
    token: your_token #从callback获取到的token
  }
  的字符串 用 private_key 进行加密后得到的密文

Response:
  http status code: 200
  reponse body: 为用private_key进行加密后的密文。解密后将得到:
  {
    timestamp: xxxx,
    username: xxx
  }

  http status code != 200

  获取失败
```


## 校验权限

```
URL: /authorize
Mehod: POST
Header:
  account=username
  client-id=client_id

Body为将数据格式:
  {
    timestamp:  xxx
    scope: {
      name: string # URL
      type: string # Method
      actions: []string #其他一些标示
    }
  }
  的字符串 用 private_key 进行加密后得到的密文

Respose:
  http status code: 200
  respose body为用private_key进行加密后的密文。解密后将得到:
   {
    timestamp: xxx
    scope: {
      type: string
      name: string
      actions: []string
    }
  }
  scope内容将与传过来的scope内容一致。
  http status code: 500
  服务器出错
  http status code: 其他
  权限校验未通过

```

## sdk

### golang

`github.com/huyinghuan/oauth_sdk`

### nodejs

`github.com/huyinghuan/oauth-sdk-nodejs`