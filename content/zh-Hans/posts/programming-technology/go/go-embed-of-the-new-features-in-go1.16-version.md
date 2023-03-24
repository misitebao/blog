---
title: Go1.16新特性之Go embed
subtitle: ""
date: 2021-02-24
lastmod: 2021-02-28
draft: false
description: Go embed教程
tags:
  - go
  - embed
categories:
  - 编程技术
  - Go
---

# Go1.16 新特性之 Go embed 教程

大家好，我是正在沉迷学习 Unity 3D 的 米司特包。

最近，Go1.16 版本发布，此次更新带来了一些新特性，这篇文章主要说一下其中的 Go
embed 功能。

在 Go embed 出来之前，我们的 Go 项目编译打包是不包含非 Go 文件的，其他的文件要另
外和打包后的文件放在一起，当然社区也出现了一些静态文件打包到二进制文件的方法。比
如 [github.com/markbates/pkger](github.com/markbates/pkger)
[github.com/gobuffalo/packr](github.com/gobuffalo/packr)等。不过这些都是第三方提
供的功能，并且有一些使用相对比较复杂，此次更新 Go1.16 以后，就可以直接使用 Go
embed 来解决这类需求了。

## 嵌入（Embed）

Go embed 支持嵌入为三种数据类型

| 数据类型 | 说明                                                                                            |
| :------: | :---------------------------------------------------------------------------------------------- |
|  []byte  | 标识存储为二进制格式，只能存储单个文件，使用时需要以 import (\_ "embed")的形式引入 embed 标准库 |
|  string  | 表示数据被编码成 utf8 编码的字符串，只能存储单个文件，引入 embed 的规则同[]byte                 |
| embed.FS | 表示存储多个文件和目录的结构                                                                    |

## 代码演示

### 嵌入为[]byte

```go
package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed hello.txt
var b []byte

func main() {
	fmt.Printf("字节数组表示为：%#v\n", b)
	fmt.Printf("转换为字符串表示为：%#v\n", string(b))

	// 避免编译之后运行的程序黑窗口一闪而逝
	time.Sleep(time.Second * 5)
}

```

运行结果：

```shell
misitebao@MISITEBAO-BOOK MINGW64 /e/WorkSpace/Git/misitebao/demo-goembed
$ go run main.go
字节数组表示为：[]byte{0x74, 0x78, 0x74, 0x3a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0xd, 0xa, 0xe7, 0xb1, 0xb3, 0xe5, 0x8f, 0xb8, 0xe7, 0x89, 0xb9, 0xe5, 0x8c, 0x85, 0x20, 0x4d, 0x69, 0x73, 0x69, 0x74, 0x65, 0x62, 0x61, 0x6f}
转换为字符串表示为："txt:hello\r\n米司特包 Misitebao"
```

### 嵌入为字符串

```go
package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed hello.txt
var s string

func main() {
	fmt.Printf("字符串表示为：%#v\n", s)

	// 避免编译之后运行的程序黑窗口一闪而逝
	time.Sleep(time.Second * 5)
}
```

运行结果：

```shell
misitebao@MISITEBAO-BOOK MINGW64 /e/WorkSpace/Git/misitebao/demo-goembed
$ go run main.go
字符串表示为："txt:hello\r\n米司特包 Misitebao"
```

### embed.FS

```go
package main

import (
	"embed"
	"fmt"
	"time"
)

//go:embed hello.txt demo.txt
var efm embed.FS

//go:embed template
var eft embed.FS

func main() {

	data, err := efm.ReadFile("hello.txt")
	if err != nil {
		fmt.Printf("读取文件错误：%#v", err)
	}

	fmt.Printf("读取出来字节表示为：%#v\n", data)
	fmt.Printf("转化为字符串表示为：%#v\n\n", string(data))

	data2, _ := efm.ReadFile("demo.txt")
	fmt.Printf("读取出来字节表示为：%#v\n", data2)
	fmt.Printf("转化为字符串表示为：%#v\n\n", string(data2))

	// 文件夹读取

	data3, _ := eft.ReadFile("template/index.html")
	data4, _ := eft.ReadFile("template/about.html")

	fmt.Printf("转化为字符串表示为：%#v\n", string(data3))
	fmt.Printf("转化为字符串表示为：%#v\n", string(data4))

	// 避免编译之后运行的程序黑窗口一闪而逝
	time.Sleep(time.Second * 5)
}


```

运行结果：

```shell

misitebao@MISITEBAO-BOOK MINGW64 /e/WorkSpace/Git/misitebao/demo-goembed
$ go run main.go
读取出来字节表示为：[]byte{0x74, 0x78, 0x74, 0x3a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0xd, 0xa, 0xe7, 0xb1, 0xb3, 0xe5, 0x8f, 0xb8, 0xe7, 0x89, 0xb9, 0xe5, 0x8c, 0x85, 0x20, 0x4d, 0x69, 0x73, 0x69, 0x74, 0x65, 0x62, 0x61, 0x6f}
转化为字符串表示为："txt:hello\r\n米司特包 Misitebao"

读取出来字节表示为：[]byte{0x74, 0x78, 0x74, 0x3a, 0x64, 0x65, 0x6d, 0x6f, 0xd, 0xa, 0xe7, 0xb1, 0xb3, 0xe5, 0x8f, 0xb8, 0xe7, 0x89, 0xb9, 0xe5, 0x8c, 0x85, 0x20, 0x4d, 0x69, 0x73, 0x69, 0x74, 0x65, 0x62, 0x61, 0x6f}
转化为字符串表示为："txt:demo\r\n米司特包 Misitebao"

转化为字符串表示为："<!DOCTYPE html>\r\n<html lang=\"en\">\r\n<head>\r\n  <meta charset=\"UTF-8\">\r\n  <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\r\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n  <title>index</title>\r\n</head>\r\n<body>\r\n  \r\n</body>\r\n</html>"
转化为字符串表示为："<!DOCTYPE html>\r\n<html lang=\"en\">\r\n<head>\r\n  <meta charset=\"UTF-8\">\r\n  <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\r\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n  <title>about</title>\r\n</head>\r\n<body>\r\n  \r\n</body>\r\n</html>"

```

## 视频介绍

针对 Go embed 的相关代码我有录制视频，上传在 B 站，新手可以上去对着视频敲。
{{< bilibili BV1LU4y1p7XZ >}}

## 总结

通过 Go1.16 正式提供的 Go embed 特性，可以原生支持静态资源文件的嵌入。整体如下：

- 在功能上：能够将静态资源嵌入二进制文件中，在运行时可以打开并读取打包后的静态文
  件，并且是线程安全的。

- 在安全上：是在编译期编译嵌入，在运行时不支持修改。

- 在使用上：支持单文件、多文件、目录。

总的来讲，Go1.16 embed 特性很好的填补了 Go 语言在打包静态文件资源的一块原生空白
领域。同时也说明了 Go 官方的确在不断地吸收社区的一些良好的想法和经验。
