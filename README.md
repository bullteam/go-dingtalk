# DingTalk Golang SDK

DingTalk Golang SDK fork https://github.com/icepy

由于原SDK，停止更新，因此特意为了Zeus fork该SDK，感谢作者`icepy`的开源贡献。

# Feature Overview

- 支持ISV，企业，SSO，SNS（H5 Web App）免登
- 支持个人小程序免登
- 支持对access_token自动续期过期管理（已使用独占锁，请勿再加锁）
- 支持注册钉钉事件回调
- 支持对钉钉事件回调消息签名的加解密
- 支持全部 Open api
- 支持全部 Top api，并且自动处理生成加密签名

# Test

- Test get auth scopes
- Test get company acess_token
- Test get company ticket
- Test upload file
- Test download file

```bash
~ ᐅ cd __test__
~ ᐅ go test
```