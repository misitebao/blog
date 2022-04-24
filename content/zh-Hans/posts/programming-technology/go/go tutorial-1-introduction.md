+++
title= "Go语言教程（1）介绍" #文章标题
date= 2021-03-09 #文章创建时间
lastmod= 2021-03-09 #文章最后修改时间
draft= false #是否为草稿
description= "" #文章内容描述

tags= [] #文章标签
categories= ["编程技术","Go"] #文章分类

+++

# Go 语言教程（1）介绍

## 文章目录

- [1. Go 语言是什么？](#nav-1)
  - [1.1 为并发而生](#nav-1-1)
  - [1.2 `goroutine`的特点](#nav-1-2)
- [2. 简单易学](#nav-2)
- [3. 代码风格统一](#nav-3)
- [4. 开发效率高](#nav-4)

<span id="nav-1"></span>

## 1. Go 语言是什么？

`Go`语言（或 `Golang`）是`Google`开发的开源编程语言，诞生于 2006 年 1 月 2 日下午 15 点 4 分 5 秒，于 2009 年 11 月开源，2012 年发布`Go稳定版`。`Go`语言在多核并发上拥有原生的设计优势，`Go`语言从底层原生支持并发，无须第三方库、开发者的编程技巧和开发经验。

`Go`是非常年轻的一门语言，它的主要目标是“兼具`Python`等动态语言的开发速度和`C/C++`等编译型语言的性能与安全性”

很多公司，特别是中国的互联网公司，即将或者已经完成了使用 Go 语言改造旧系统的过程。经过 Go 语言重构的系统能使用更少的硬件资源获得更高的并发和 I/O 吞吐表现。充分挖掘硬件设备的潜力也满足当前精细化运营的市场大环境。

<span id="nav-1-1"></span>

### 1.1 为并发而生

`Go`语言的并发是基于 `goroutine` 的，`goroutine` 类似于线程，但并非线程。可以将 `goroutine` 理解为一种虚拟线程。`Go`语言运行时会参与调度 `goroutine`，并将 `goroutine` 合理地分配到每个 `CPU` 中，最大限度地使用`CPU`性能。开启一个`goroutine`的消耗非常小（大约 2KB 的内存），你可以轻松创建数百万个`goroutine`。

<span id="nav-1-2"></span>

### 1.2 `goroutine`的特点

- `goroutine`具有可增长的分段堆栈。这意味着它们只在需要时才会使用更多内存。
- `goroutine`的启动时间比线程快。
- `goroutine`原生支持利用 channel 安全地进行通信。
- `goroutine`共享数据结构时无需使用互斥锁。

<span id="nav-2"></span>

## 2. 简单易学

`Go`语言简单易学，学习曲线平缓，不需要像 `C/C++ `语言动辄需要两到三年的学习期。`Go` 语言被称为“互联网时代的`C`语言”。`Go` 语言的风格类似于`C`语言。其语法在`C`语言的基础上进行了大幅的简化，去掉了不需要的表达式括号，循环也只有 `for` 一种表示方法，就可以实现数值、键值等各种遍历。

<span id="nav-3"></span>

## 3. 代码风格统一

`Go` 语言提供了一套格式化工具——`go fmt`。一些 `Go` 语言的开发环境或者编辑器在保存时，都会使用格式化工具进行修改代码的格式化，这样就保证了不同开发者提交的代码都是统一的格式。(吐槽下：再也不用担心那些看不懂的黑魔法了…)

<span id="nav-4"></span>

## 4. 开发效率高

`Go`语言实现了开发效率与执行效率的完美结合，让你像写`Python`代码（效率）一样编写`C`代码（性能）。