# 支持的功能
1. 私聊回复
2. 群里@回复
3. 好友请求处理

# 使用
建议先在本地运行程序扫码登录之后默认会生成`qq.token`和`device.json`，将这两个文件上传到服务器再运行程序。不然扫码会提示当前设备网络不稳定或处于复杂网络环境.

### 直接下载二进制文件运行
需要安装[依赖](https://github.com/yxw21/qqchatgpt/edit/master/README.md#%E4%BE%9D%E8%B5%96)
并且提供一些[环境变量](https://github.com/yxw21/qqchatgpt/edit/master/README.md#%E4%BE%9D%E8%B5%96)
然后运行
```
./qqchatgpt
```

### 使用docker
需要提供一些[环境变量](https://github.com/yxw21/qqchatgpt/edit/master/README.md#%E4%BE%9D%E8%B5%96)
```
docker run -dit -e QQ_CHAT_GPT_USERNAME=example@gmail.com -e QQ_CHAT_GPT_PASSWORD=password -e QQ_KEY=I-12312 -v ./qq.token:/qqchatgpt/qq.token yxw21/qqchatgpt
```
https://hub.docker.com/r/yxw21/qqchatgpt

# QQ登录流程
如果提供了`QQ_UIN`和`QQ_PASSWORD`会先尝试使用qq和密码登录，遇到有验证消息或其他原因不能成功登录的会使用二维码扫码登录。

没有提供`QQ_UIN`和`QQ_PASSWORD`会使用二维码扫码登录。

# CHATGPT登录流程
如果提供了`QQ_CHAT_GPT_USERNAME`和`QQ_CHAT_GPT_PASSWORD`会自动登录获取`AccessToken`。

`QQ_CHAT_GPT_USERNAME`、`QQ_CHAT_GPT_PASSWORD`和`QQ_CHAT_GPT_ACCESS_TOKEN`必须提供一项

# 依赖
### Xvfb （只有linux环境需要安装）
  
Ubuntu or Debian
```
apt update
apt install xvfb
```
CentOS
```
yum update
yum install xorg-x11-server-Xvfb
```
Alpine
```
apk update
apk add xvfb
```
### Key
登录需要谷歌验证码，引入了第三方破解，需要去网站`nopecha.com`购买key，价格很便宜

```
https://nopecha.com
```
### Chrome

Ubuntu or Debian
```
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
apt install ./google-chrome-stable_current_amd64.deb
```
CentOS
```
wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
yum localinstall -y google-chrome-stable_current_x86_64.rpm
```
Alpine
```
apk add chromium
```


# 环境变量
### QQ_UIN (可选)
QQ号
### QQ_PASSWORD (可选)
QQ密码
### QQ_MSG_RETRY (可选)
chatgpt请求失败重试的次数，次数越多回复消息就越慢(默认3)
### QQ_KEY (必填)
破解谷歌验证码需要的key，需要去网站`nopecha.com`购买
### QQ_CHAT_GPT_USERNAME (可选)
openai用户名
### QQ_CHAT_GPT_PASSWORD (可选)
openai密码
### QQ_CHAT_GPT_ACCESS_TOKEN (可选)
大概7天过期
1. 登录 https://chat.openai.com
2. 访问 https://chat.openai.com/api/auth/session
### QQ_CHAT_GPT_POLICY (可选)
好友添加策略

同意好友添加请求
```
QQ_CHAT_GPT_POLICY = agree
```
拒绝好友添加请求
```
QQ_CHAT_GPT_POLICY = reject
```
不处理好友添加请求
```
QQ_CHAT_GPT_POLICY = ignore
```
当验证消息与给定的正则表达式匹配时才会同意添加
```
QQ_CHAT_GPT_POLICY = agree,123456
```
远程验证（GET请求`http://example.com/验证信息`），响应状态码为201同意添加
```
QQ_CHAT_GPT_POLICY = agree,https://example.com
```
