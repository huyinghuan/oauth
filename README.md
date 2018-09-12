# Open Auth


## API说明

### timestamp

以下出现的`timestamp`参数均为int值，为当前时间的以秒为单位的时间戳。
如nodejs: `Date.now()/1000`, golang: `time.Now().Unix()`, 

数据通讯过程中，传到服务器的`timestamp`， 服务器将校验该值，如果该值与服务器误差大于5分钟，请求将被拒绝
服务器API响应的`timestamp`, 请应用自行判断该值是否可信。


## 获取登陆用户Token

- 1.将网页URL重定向到`/authorize?client_id=xxx`, 该URL会自动引导用户进行登陆操作
- 2.用户在登陆以后，将会由该URL跳转到`your_app_server_callback_url?token=xxxx`，这样你到应用拿到了一个60秒内有效到token值。


## 用token 获取登陆用户信息

### API

```
URL: /resource/account
Method:  POST
Header:
  client_id=your_app_client_id

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
  client_id=client_id

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

## 如何加密解密数据

### nodejs

```js

'use strict';

const crypto = require('crypto');

const algorithm = 'aes-256-cfb';

function encryptText(keyStr, text) {
  const hash = crypto.createHash('sha256');
  hash.update(keyStr);
  const keyBytes = hash.digest();

  const iv = crypto.randomBytes(16);
  const cipher = crypto.createCipheriv(algorithm, keyBytes, iv);
  console.log('IV:', iv);
  let enc = [iv, cipher.update(text, 'utf8')];
  enc.push(cipher.final());
  return Buffer.concat(enc).toString('base64');
}

function decryptText(keyStr, text) {
  const hash = crypto.createHash('sha256');
  hash.update(keyStr);
  const keyBytes = hash.digest();

  const contents = Buffer.from(text, 'base64');
  const iv = contents.slice(0, 16);
  const textBytes = contents.slice(16);
  const decipher = crypto.createDecipheriv(algorithm, keyBytes, iv);
  let res = decipher.update(textBytes, '', 'utf8');
  res += decipher.final('utf8');
  return res;
} 

const encrypted = encryptText('privateKey', 'It works!');
console.log('Encrypted: ', encrypted);

const decrypted = decryptText('SecretKey', encrypted);
console.log('Decrypted: ', decrypted);
```

### golang

```golang

package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)


func CFBEncrypt(keyStr string, cryptoText string) (string, error) {
	keyBytes := sha256.Sum256([]byte(keyStr))
	return encrypt(keyBytes[:], []byte(cryptoText))
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, plaintext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func CFBDecrypt(keyStr string, cryptoText string) (string, error) {
	keyBytes := sha256.Sum256([]byte(keyStr))
	return decrypt(keyBytes[:], cryptoText)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext), nil
}

```