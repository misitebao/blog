+++
title= "Go语言实现系统托盘Tray" #文章标题
date= 2021-05-17 #文章创建时间
lastmod= 2021-05-17 #文章最后修改时间
draft= false #是否为草稿
description= "" #文章内容描述

tags= ["go","系统托盘","tray"] #文章标签
categories= ["编程技术","Go"] #文章分类

+++

# Go 语言实现系统托盘 Tray

## 文章目录

- [1. 背景](#nav-1)
- [2. 使用方法](#nav-2)
- [3. 效果展示](#nav-3)

<span id="nav-1"></span>

## 背景

最近在开发一个 Go GUI 应用的时候，需要在应用运行的时候在操作系统托盘显示应用图标，目前 Go 官方库并没有这种库能满足我们的要求，所以在 github 找到一个可以满足我们的要求的并且跨平台的库：[github.com/riftbit/go-systray](github.com/riftbit/go-systray)

这篇文章的示例代码，放在仓库：[https://github.com/misitebao/demo-go-tray](https://github.com/misitebao/demo-go-tray)

<span id="nav-2"></span>

## 使用方法

代码如下：

```go
package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/riftbit/go-systray"
)

var (
	timezone string
)

func main() {
	//systray.SetCustomLeftClickAction()
	//systray.SetCustomRightClickAction()
	systray.Run(onReady, onExit)
}

func onReady() {
	timezone = "默认文字"
	systray.SetIcon(getIcon("./favicon_misitebao.ico"))

	submenu := systray.AddSubMenu("子菜单")
	_ = submenu.AddSubMenuItem("开始", "", 0)
	_ = submenu.AddSubMenuItem("结束", "", 0)

	localTime := systray.AddMenuItem("炎龙", "炎龙", 0)
	hcmcTime := systray.AddMenuItem("风鹰", "风鹰", 0)
	sydTime := systray.AddMenuItem("黑犀", "黑犀", 0)
	gdlTime := systray.AddMenuItem("地虎", "地虎", 0)
	sfTime := systray.AddMenuItem("雪獒", "雪獒", 0)

	fmt.Printf("%#v", localTime)

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出", 0)

	go func() {
		for {
			systray.SetTitle(timezone)
			systray.SetTooltip(timezone)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-localTime.OnClickCh():
				timezone = "炎龙"
			case <-hcmcTime.OnClickCh():
				timezone = "风鹰"
			case <-sydTime.OnClickCh():
				timezone = "黑犀"
			case <-gdlTime.OnClickCh():
				timezone = "地虎"
			case <-sfTime.OnClickCh():
				timezone = "雪獒"
			case <-mQuit.OnClickCh():
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// 清除销毁
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}


```

项目根目录下运行：`go build -ldflags="-H windowsgui"`

> 关于编译命令中的参数，请参考：[Go应用隐藏启动时的黑色窗口的方法](https://blog.misitebao.com/posts/编程技术/go/go应用隐藏启动时的黑色窗口的方法/)

<span id="nav-3"></span>

## 效果展示

![](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20210721062440.gif)
