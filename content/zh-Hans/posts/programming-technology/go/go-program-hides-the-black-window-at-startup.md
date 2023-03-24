---
title: Go应用隐藏启动时的黑色窗口
date: 2021-06-14
lastmod: 2021-06-14
draft: false
description: ""
tags:
  - go
categories:
  - 编程技术
  - Go
---

# Go 应用隐藏启动时的黑色窗口

<!-- ## 文章目录

- [1. ](#nav-1)
- [2. ](#nav-2)
  - [2.1 ](#nav-2-1)
  - [2.2 ](#nav-2-2)
  - [2.3 ](#nav-2-3)
- [3. ](#nav-3)
  - [3.1 ](#nav-3-1)

<span id="nav-1"></span>

## 导航 -->

通过 go 的标准库 exec 调用 cmd 命令时会闪弹黑窗口,此问题主要出现在带 GUI 或者无
控制台的程序调用 cmd 时。解决方案是在 go 编译时添加参数
：`go build -ldflags="-H windowsgui"`
