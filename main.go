/**
 *  Copyright (c) 2021 NekokeCore
 *  FiFuShortLink is licensed under Mulan PSL v2.
 *  You can use this software according to the terms and conditions of the Mulan PSL v2.
 *  You may obtain a copy of Mulan PSL v2 at:
 *           http://license.coscl.org.cn/MulanPSL2
 *  THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 *  See the Mulan PSL v2 for more details.
 */

package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/encoding/gtoml"
)

type data struct {
	Addr  string
	Core  string
	Top   int64
	Links map[string]string
}

var toml_str string = `# cli工具：直接将链接作为参数传入主程序即可注册映射，
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
`

var d data

func init() {
	t := time.Now()
	log.Println("初始化中...")
	log.Println("正在加载数据...")

	_, err := os.Stat("./data.toml")
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create("./data.toml")
			if err != nil {
				log.Fatalf("配置文件有误，示例配置文件[data.toml]：\n%v\n更多细节请学习Toml语法", toml_str)
				log.Fatalln(err)
			} else {
				file.WriteString(toml_str)
				log.Println("未检测到配置文件，已自动创建")
			}
		}
	}
	if _, err := toml.DecodeFile("./data.toml", &d); err != nil {
		log.Fatalln("配置文件解析失败！请检查！详细：")
		log.Fatalln(err)
	}

	log.Println(d)
	log.Println("数据加载完毕，耗时", time.Since(t))
}

func main() {
	t := time.Now()
	args := os.Args
	if largs := len(args); largs > 1 {
		log.Println("正在注册映射...")
		for _, v := range args[1:] {
			d.Top++
			i := strconv.FormatInt(d.Top, 36)
			d.Links[i] = v
			log.Println(i, " = ", v)
		}
		byt, err := gtoml.Encode(d)
		if err != nil {
			log.Println(err)
		}
		file, _ := os.Create("./data.toml")
		file.Write(byt)
		log.Println("映射注册完毕，耗时", time.Since(t))
		os.Exit(largs - 1)
	}

	t = time.Now()
	log.Println("正在启动服务...")
	gin.SetMode(gin.ReleaseMode)

	log.Println("正在加载中间件...")
	r := gin.Default()
	log.Println("中间件加载完毕")

	log.Println("正在注册路由...")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, d.Core)
	})
	r.GET("/:url", func(c *gin.Context) {
		url := c.Param("url")
		link := d.Links[url]
		c.Redirect(http.StatusMovedPermanently, link)
		log.Printf(" もぺもぺ 『%v ➤ %v』\n", url, link)
	})
	log.Println("路由注册完毕")

	log.Println("服务已启动，耗时", time.Since(t))
	r.Run(d.Addr)
}
