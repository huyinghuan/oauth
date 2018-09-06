# Open Auth

## 注册应用

### API

```
URL: /app-register
Method: POST
Request Body: 

  {  
    name:  string 应用名称
    password: string 登陆密码，管理用户权限使用
    callback:  string 回调地址【应用必须拥有一个回调地址，用来接收 OpenAuth 服务的Token, 使用该Token获取用户信息】
  }
Content-Type: application/json


Respose：
  当 http status code = 200:
    Content-Type: application/json
    Body: 
      {
          "client_id": string 24位字符串，用来标示该应用
          "private_key": string 用来加密与OpenAuth的通讯数据
      }
  当 http status code != 200:
  注册失败
```

## Example:

### curl

```curl
curl -X POST \
  http://localhost:8000/app-register \
  -H 'Content-Type: application/json' \
  -d '{
	"name":"TestApp",
	"password":"123456",
	"callback":"http://localhost:8080/auth"
}'

```

### nodejs

```js
var request = require("request");

var options = { method: 'POST',
  url: 'http://localhost:8000/app-register',
  headers: 
   { 'Content-Type': 'application/json' },
  body: 
   { name: 'shorturl',
     password: '123456',
     callback: 'http://localhost:8080/auth' },
  json: true };

request(options, function (error, response, body) {
  if (error) throw new Error(error);

  console.log(body);
});
```

Respose：

```json
{
  "client_id":"6J1NMOBgyMrVRmiJzpZI1p4g",
  "private_key":"JzTDzVjPhQTbuhftD3qlYXFd"
}
```

## 获取登陆用户Token

### API


```
URL: /authorize?client_id=xxx
Method: GET
```

## 获取登陆用户信息


## 校验权限

```
URL: /authorize
Mehod: POST
Header:
  account=username
  client_id=client_id
Body:
  将需要校验的权限，按下面的格式填充
  data = {
    timestamp: 1970年到现在的秒数【js:  Date.now()/1000 】 服务器将校验该值，如果该值与服务器误差大于5分钟，请求将倍拒绝
    scope: {
      type: string
      name: string
      actions: []string
    }
  }

  将data 转成json格式的字符串： plainText = JSON.stringif(data) 得到待加密的字符串 
  将 plainText 用 private_key 进行 cfb 加密 得到字符串 encryptText.
  body就为该 encryptText

Respose:
  http status code: 200
  获取respose.body,  使用 private_key 解析 cfb 解密。 解密得到一个json 字符串。格式为
   {
    timestamp: 1970年到现在的秒数【js:  Date.now()/1000 】请校验服务器的时间戳与本地时间戳的误差，自行判断该结果是否可信。
    scope: {
      type: string
      name: string
      actions: []string
    }
  }
  http status code: 500
  服务器出错
  http status code: 其他
  权限校验未通过

```