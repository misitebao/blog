+++
title= "Go应用隐藏启动时的黑色窗口" #文章标题
date= 2021-06-14 #文章创建时间
lastmod= 2021-06-14 #文章最后修改时间
draft= false #是否为草稿
description= "" #文章内容描述

tags= ["go"] #文章标签
categories= ["编程技术","Go"] #文章分类

+++

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

通过 go 的标准库 exec 调用 cmd 命令时会闪弹黑窗口,此问题主要出现在带 GUI 或者无控制台的程序调用 cmd 时。
解决方案是在 go 编译时添加参数：`go build -ldflags="-H windowsgui"`
