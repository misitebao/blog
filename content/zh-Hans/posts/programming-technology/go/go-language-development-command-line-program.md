---
title: Go语言开发命令行（cli）程序
date: 2020-10-22
lastmod: 2020-10-22
draft: false
description: ""
tags: []
categories:
  - 编程技术
  - Go
---

# Go 语言编写命令行（cli）程序

## 文章目录

- [Go 语言编写命令行（cli）程序](#go-语言编写命令行cli程序)
  - [文章目录](#文章目录)
  - [1. 问题](#1-问题)
  - [2. 解决](#2-解决)
    - [2.1 实现查看版本和切换语言的程序功能](#21-实现查看版本和切换语言的程序功能)
    - [2.2 效果](#22-效果)
    - [2.3 扩展](#23-扩展)
      - [2.3.1 命令行的形式](#231-命令行的形式)
      - [2.3.2 关于 go 标准库当中的 flag 包以及 os. Args 方法](#232-关于-go-标准库当中的-flag-包以及-os-args-方法)
      - [2.3.3 设置环境变量在任意文件夹使用 cli 程序](#233-设置环境变量在任意文件夹使用-cli-程序)

<span id="nav-1"></span>

## 1. 问题

如何使用 golang 创建可以在命令行当中传递参数的程序？go 程序如何带参数执行？

比如我们期望使用*hello -version*来查看 hello 程序的版本号码。或者输入 hello
language chinese 来切换 hello 程序的语言版本。

<span id="nav-2"></span>

## 2. 解决

在网络上大家可以经常搜索到的解决方案是使用 go 标准库当中的 os. Args 或者 flag 方
法获取命令行参数再执行程序。可以很容易 google 和 baidu 的这里就不重复说明了。 这
里介绍一个开源的 cli 库的解决方案。

`github.com/urfave/cli`

该 cli 库集成了很多创建 cli 程序的工具和方法，官方 github 库文档功能案例介绍比较
全面很容易学习使用。

截止到 2018 年 2 月 19 日，该开源库有 875 次 commit，33 次 releases，7462 个
star。经过长期迭代和社区验证功能比较完整，程序稳定性也相对会可靠很多。

<span id="nav-2-1"></span>

### 2.1 实现查看版本和切换语言的程序功能

注：本案例代码只是简单的打印出 language 而已。

```
package main

import (
    "os"
    "github.com/urfave/cli"
    "fmt"
)

func main() {
    // 实例化cli
    app := cli.NewApp()
    // Name可以设定应用的名字
    app.Name = "hello"
    // Version可以设定应用的版本号
    app.Version = "1.0.0"
    // Commands用于创建命令
    app.Commands = []cli.Command{
        {
            // 命令的名字
            Name:    "language",
            // 命令的缩写，就是不输入language只输入lang也可以调用命令
            Aliases: []string{"lang"},
            // 命令的用法注释，这里会在输入 程序名 -help的时候显示命令的使用方法
            Usage:   "change language",
            // 命令的处理函数
            Action:  func(c *cli.Context) error {
                language := c.Args().First()
                if language == "chinese"{
                    fmt.Println("Language is 中文")
                }else {
                    fmt.Println("Language is English")
                }
                return nil
            },
        },

    }
    // 接受os.Args并启动程序
    app.Run(os.Args)
}
```

<span id="nav-2-2"></span>

### 2.2 效果

```
liuMacBook-Pro:gowork liuyunpeng$ go run hello.go -help
NAME:
   hello - A new cli application

USAGE:
   hello [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     language, lang  change language
     help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

liuMacBook-Pro:gowork liuyunpeng$ go run hello.go -version
hello version 1.0.0
liuMacBook-Pro:gowork liuyunpeng$ go run hello.go lang chinese
Language is 中文
liuMacBook-Pro:gowork liuyunpeng$ go run hello.go lang french
Language is English
liuyMacBook-Pro:gowork liuyunpeng$ go run hello.go lang english
Language is English
```

更多的使用方法可以[github](github.com/urfave/cli)查看

<span id="nav-2-3"></span>

### 2.3 扩展

<span id="nav-2-3-1"></span>

#### 2.3.1 命令行的形式

值得大家注意的是，-flag 前面带减号（实际上应该是破折号为了方便就缩减成减号了）这
种带标志方式的使用方法和 command 方法是两个概念。个人理解是前者代表参数的传递比
如给 -name=ppmoon 就是给 name 变量传递 ppmoon 参数。后者是执行一段命令，比如
hello add 1 1 返回结果就是 2，就是通过 hello 程序执行了一个 1+1 的运算，参数通过
空格的方式依次传递到程序当中了。 从使用体验的角度来说个人认为并没有什么大的区别
，因为本质上来说 cli 程序只是没有图形化的 UI 界面，我们只是在命令提示符当中操作
程序罢了，不管你使用什么样子的方式传递都只是形式上的变化。

<span id="nav-2-3-2"></span>

#### 2.3.2 关于 go 标准库当中的 flag 包以及 os. Args 方法

flag 是官方提供的一个不错的 cli 程序扩展包，里面也包含了很多常用的方法。不管是
flag 还是 urfave/cli 都需要依赖 OS 包当中的 os. Args 方法来获取命令行参数，这一
点  是值得注意的便于理解 cli 程序的原理。

<span id="nav-2-3-3"></span>

#### 2.3.3 设置环境变量在任意文件夹使用 cli 程序

在 Mac Windows Linux 我们都会使用到  命令行程序，如果想在全局任何一个文件夹下面
都使用命令行程序，我们只需要将 cli 程序的可执行程序文件所在的文件夹路径添加到环
境变量里就可以直接使用程序的名字全局调用 cli 程序了。具体  网络上可以搜索到诸多
教程这里就不赘述了。
