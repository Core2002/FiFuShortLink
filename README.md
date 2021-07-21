# 介绍
> [FiFuShortLink](https://github.com/Core2002/FiFuShortLink) 是一个使用[Golang](https://golang.google.cn/)开发的一款轻量级、零依赖、高性能的短链接服务

# 部署
## 源码部署
> 安装并配置Golang开发环境  
> `git clone https://github.com/Core2002/FiFuShortLink`  
> `cd FiFuShortLink`  
> `go build`  
> `FiFuShortLink`  

## 二进制部署
> 在Releases内下载好所在平台的二进制文件  
> 运行即可

## 配置文件及说明
```toml
# cli工具：直接将链接作为参数传入主程序即可注册映射，
# 多个链接用空格隔开，退出码即新增链接个数，
# 例： FiFuShortLink https://golang.google.cn https://www.bilibili.com
# 不传参数则直接启动服务
# 例： FiFuShortLink

# 内网地址+端口
Addr = ":80"

# / 目录或异常左值的跳转的地址
# 例： http://{Addr} -> {Core}
Core = "https://space.bilibili.com/30924239"

# 栈顶指针
Top = 2

# 字符串-字符串键值对，从左往右跳转
# 例： http://{Addr}/{i} -> {Links[i]}
[Links]
0 = "https://github.com/Core2002"
1 = "https://gitee.com/NekokeCore"
2 = "https://www.fifu.fun"
```