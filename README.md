# Bilibili API for Go (Dynamic Version)

- 接口文档：[SocialSisterYi/bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect)
- 基于 [CuteReimu/bilibili](https://github.com/CuteReimu/bilibili) 二次开发，保留登录与鉴权操作，删除其他接口。使用时可根据接口文档自行构建请求，响应将以 simplejson 格式返回，可根据文档进行解析。

## 快速开始

### 安装

```bash
go get -u github.com/rinkurt/bilibili
```

在项目中引用即可使用

```go
import "github.com/rinkurt/bilibili"

var client = bilibili.New()
```

### 首次登录

> [!TIP]
> 下文为了篇幅更短，示例中把很多显而易见的`err`校验忽略成了`_`，实际使用请自行校验`err`。

#### 方法一：扫码登录

首先获取二维码：

```go
qrCode, _ := client.GetQRCode()
buf, _ := qrCode.Encode()
img, _ := png.Decode(buf) // 或者写入文件 os.WriteFile("qrcode.png", buf, 0644)
// 也可以调用 qrCode.Print() 将二维码打印在控制台
```

扫码并确认成功后，发送登录请求：

```go
result, err := client.LoginWithQRCode(bilibili.LoginWithQRCodeParam{
    QrcodeKey: qrCode.QrcodeKey,
})
if err == nil && result.Code == 0 {
    log.Println("登录成功")
}
```

#### 方法二：账号密码登录

首先获取人机验证参数：

```go
captchaResult, _ := client.Captcha()
```

将`captchaResult`中的`gt`和`challenge`值保存下来，自行使用 [手动验证器](https://kuresaru.github.io/geetest-validator/) 进行人机验证，并获得`validate`和`seccode`。然后使用账号密码进行登录即可：

```go
result, err := client.LoginWithPassword(bilibili.LoginWithPasswordParam{
    Username:  userName,
    Password:  password,
    Token:     captchaResult.Token,
    Challenge: captchaResult.Geetest.Challenge,
    Validate:  validate,
    Seccode:   seccode,
})
if err == nil && result.Status == 0 {
    log.Println("登录成功")
}
```

#### 方法三：使用短信验证码登录（不推荐）

首先用上述方法二相同的方式获取人机验证参数并进行人机验证。然后获取国际地区代码：

```go
countryCrownResult, _ := client.GetCountryCrown()
```

当然，如果你已经确定`cid`的值，这一步可以跳过。中国大陆的`cid`就是`86`。

然后发送短信验证码：*（[这个接口大概率返回86103错误](https://github.com/SocialSisterYi/bilibili-API-collect/issues/756)）*

```go
sendSMSResult, _ := client.SendSMS(bilibili.SendSMSParam{
    Cid:       cid,
    Tel:       tel,
    Source:    "main_web",
    Token:     captchaResult.Token,
    Challenge: captchaResult.Geetest.Challenge,
    Validate:  validate,
    Seccode:   seccode,
})
```

然后就可以使用手机验证码登录了：

```go
result, err := client.LoginWithSMS(bilibili.LoginWithSMSParam{
    Cid:        cid,
    Tel:        tel,
    Code:       123456, // 短信验证码
    Source:     "main_web",
    CaptchaKey: sendSMSResult.CaptchaKey,
})
if err == nil && result.Status == 0 {
    log.Println("登录成功")
}
```

### 储存Cookies

使用上述任意方式登录成功后，Cookies值就已经设置好了。你可以保存Cookies值方便下次启动程序时不需要重新登录。

```go
// 获取cookiesString，自行存储，方便下次启动程序时不需要重新登录
cookiesString := client.GetCookiesString()

// 下次启动时，把存储的cookiesString设置进来，就不需要登录操作了
client.SetCookiesString(cookiesString)

// 如果你是从浏览器request的header中直接复制出来的cookies，则改为调用SetRawCookies
client.SetRawCookies("cookie1=xxx; cookie2=xxx")
```

> [!NOTE]
> - `GetCookiesString`和`SetCookiesString`使用的字符串是`"cookie1=xxx; expires=xxx; domain=xxx.com; path=/\ncookie2=xxx; expires=xxx; domain=xxx.com; path=/"`，包含过期时间、domain等一些其它信息，以`"\n"`分隔多个cookie
> - `SetRawCookies`使用的字符串是`"cookie1=xxx; cookie2=xxx"`，只包含key=value，以`"; "`分隔多个cookie，这和在浏览器F12里复制的一样
>
> 请注意不要混用。

### 其它接口

根据接口参数类型，调用 `client.Send(...)Request` 接口

```go
resp, err := client.SendUrlRequest("GET", "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/space", map[string]string{
    "host_mid": "2",
    "features": "itemOpusStyle",
})
```

返回 resp 为 simplejson 格式，可以调用 Get 或 GetPath 按照格式路径获取节点，然后调用相应的类型转换函数获取数据。具体参见 simplejson 文档。

```go
resp.Get("key").MustString()
resp.GetPath("key1", "key2").MustInt()
```

### 对B站返回的错误码进行处理

因为B站的返回内容是这样的格式：

```json
{
   "code": 0,
   "message": "错误信息",
   "data": {}
}
```

而我们这个库的接口只会返回`data`数据和一个`error`，若`code`为`0`则`error`为`nil`，否则我们并不会把`code`和`message`字段直接返回。

在一般情况下，调用者不太需要关心`code`和`message`字段，只需要关心是否有`error`即可。
但如果你实在需要`code`和`message`字段，我们也提供了一个办法：

```go
videoInfo, err := client.GetVideoInfo(bilibili.VideoParam{
    Aid: 12345678,
})
if err != nil {
    var e bilibili.Error
    if errors.As(err, &e) { // B站返回的错误
        log.Printf("错误码: %d, 错误信息: %s", e.Code, e.Message)
    } else { // 其它错误
        log.Printf("%+v", err)
    }
}
```

> [!TIP]
> 我们的所有`error`都包含堆栈信息。如有需要，你可以用`log.Printf("%+v", err)`打印出堆栈信息，方便追踪错误。

### 可能用到的工具接口

```go
// 解析短连接
typ, id, err := client.UnwrapShortUrl("https://b23.tv/xxxxxx")

// 获取服务器当前时间
now, err := client.Now()

// av号转bv号
bvid := bilibili.AvToBv(111298867365120)

// bv号转av号
aid := bilibili.BvToAv("BV1L9Uoa9EUx")

// 通过ip确定地理位置
zoneLocation, err := client.GetZoneLocation()

// 获取分区当日投稿稿件数
regionDailyCount, err := client.GetRegionDailyCount()
```

### 设置*resty.Client的一些参数

调用`client.Resty()`就可以获取到`*resty.Client`，然后自行操作即可。**但是不要做一些离谱的操作**~~（比如把Cookies删了）~~

```go
client.Resty().SetTimeout(20 * time.Second) // 设置超时时间
client.Resty().SetLogger(logger) // 自定义logger
```
