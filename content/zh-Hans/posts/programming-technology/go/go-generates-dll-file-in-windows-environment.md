---
title: Go语言在windows环境下生成dll文件
date: 2021-03-31
lastmod: 2021-03-31
draft: false
description: ""
tags:
  - dll
categories:
  - 编程技术
  - Go
---

# Go 语言 windows 环境下生成 dll 文件

## 文章目录

- [Go 语言 windows 环境下生成 dll 文件](#go-语言-windows-环境下生成-dll-文件)
	- [文章目录](#文章目录)
	- [1. 运行环境](#1-运行环境)
	- [2. 实现](#2-实现)
		- [2.1 Go 代码](#21-go-代码)
		- [2.2 编译](#22-编译)
		- [2.3 结果](#23-结果)
	- [3. 调用](#3-调用)
		- [3.1 Go 语言调用](#31-go-语言调用)

> 声明：我是 「[米司特包](http://misitebao.com)」 ，本篇文章 **首发** 于
> 「[米司博客](http://blog.misitebao.com)」 ，其他平台为同步推送（_因为可能文章
> 路径会变动，所以就没放详细链接，进来的可以
> [根据标题查找 🔍](http://blog.misitebao.com/posts/)_）。

因为之前开发一个 Electron 项目，里面有一个功能用到了外部的 dll 文件，当时自己不
会开发，还是找的其他项目组的同事帮忙写的一个 dll 文件，最近忽然想到这个，而且现
在用 Go 也比较多，所以就想着用 Go 实现一下（因为之前了解过 cgo，但是没实际应用过
，主要是我接触的项目并没有应用场景，但是我知道 Go 可以实现。），如果可以的话，以
后写项目如果再需要用到，就不用麻烦别人了。

**文章里面的示例代码已经放在 Github 了,下载地址 👇**

- [https://github.com/misitebao/demo-go-dll](https://github.com/misitebao/demo-go-dll)

<span id="nav-1"></span>

## 1. 运行环境

- windows 10：版本 Dev （OS 内部版本 21343.1000）
- go `go version go1.16 windows/amd64`
- gcc `gcc.exe (tdm64-1) 9.2.0`

windows 查看版本信息命令：`winver`

<span id="nav-2"></span>

## 2. 实现

<span id="nav-2-1"></span>

### 2.1 Go 代码

```go
// go_dll.go
package main

import "C"

//export Add
func Add(a C.int, b C.int) C.int {
	return a + b
}

//export Print
func Print(s *C.char) {

	// 函数参数可以用 string, 但是用*C.char更通用一些。
	// 由于string的数据结构，是可以被其它go程序调用的，
	// 但其它语言（如 nodejs）就不行了

	print("Hello ", C.GoString(s)) //这里不能用fmt包，会报错，调了很久...
}

func main() {
}

```

<span id="nav-2-2"></span>

### 2.2 编译

```shell
go build -ldflags "-s -w" -buildmode=c-shared -o demo_go_dll.dll go_dll.go
```

<span id="nav-2-3"></span>

### 2.3 结果

生成 .dll 和 .h 文件

![image-20210331125052548](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/image-20210331125052548.png)

<span id="nav-3"></span>

## 3. 调用

<span id="nav-3-1"></span>

### 3.1 Go 语言调用

![image-20210331133106311](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/image-20210331133106311.png)

输出:

![image-20210331141251300](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/image-20210331141251300.png)

这只是一个简单的例子，实现基本的运算和输出，而且也只是在 Go 里面调用测试，后续还
要测试在 node 里面调用测试，还要测试更加复杂的功能。

关于 node 调用 dll 的文章，可以在 node 分类里面查找
