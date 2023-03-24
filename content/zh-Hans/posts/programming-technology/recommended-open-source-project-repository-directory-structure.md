---
title: 推荐的开源项目存储库目录结构
date: 2021-03-07
lastmod: 2021-03-07
draft: false
description: ""
tags:
  - readme
  - git
  - github
categories:
  - 编程技术
---

# 推荐一个自用的开源项目 README 模板

## 内容目录

<details>
  <summary>点我 打开/关闭 目录列表</summary>

- [推荐一个自用的开源项目 README 模板](#推荐一个自用的开源项目-readme-模板)
  - [内容目录](#内容目录)
  - [1. README 部分](#1-readme-部分)
    - [1.1 简介](#11-简介)
    - [1.2 效果展示](#12-效果展示)
  - [2. Issue/PR 模板部分](#2-issuepr-模板部分)
    - [2.1 简介](#21-简介)
    - [2.2 效果展示](#22-效果展示)
  - [3. Github Actions CI 部分](#3-github-actions-ci-部分)
    - [3.1 简介](#31-简介)
    - [3.2 效果展示](#32-效果展示)

</details>

> 声明：我是 「[米司特包](http://misitebao.com)」 ，本篇文章 **首发** 于
> 「[米司博客](http://blog.misitebao.com)」 ，其他平台为同步推送，由于网站后续会
> 升级重构，很有可能导致文章路径规则变动，所以就没贴详细链接，可以进入博客根据标
> 签查找 🔍。

本来决定是要写一篇《如何写好开源项目的 README》的文章的，后来根据自己的一些项目
的情况，就把 Issue、PR 和 CI 部分也加了进来，各位看官姥爷可以根据自己的需要查看
。

由于作者本人也是一个常逛 Github 的人，偶尔也会在上面找一些别人实现好的轮子，在这
中间发现有的一些项目质量很好，但是有的没有提供文档或者 README 写的很简陋，使用者
要用的话，就需要自己去阅读源码，才知道怎么使用，我想这可能是很多技术人员的通病——
不喜欢写文档，因为我自己是个强迫症，我自己的项目一般都会有 README 或者专门的文档
，因为这个原因，我新建了一个 Github 仓库模板库，里面提供了基础的 README 要素 和
多语言范本，同时也提供了 Issue 和 PR 模板，以及基于 Github Actions 的 CI 脚本（
目前只有 NodeJS 项目和 Hugo 项目的示例）。基本上将大部分开源项目需要的东西包含了
进去。

代码仓库：

- [https://github.com/misitebao/standard-repository](https://github.com/misitebao/standard-repository)（Github）
- [https://gitee.com/misitebao/standard-repository](https://gitee.com/misitebao/standard-repository)（Gitee）

也可以直接点击链接在线上 VSCode 环境预览
：[VSCode 在线预览](https://github1s.com/misitebao/standard-repository)

<span id="nav-1"></span>

## 1. README 部分

<span id="nav-1-1"></span>

### 1.1 简介

README 是一个项目的简介，同时也是 Github 上别人访问你的项目第一眼要看的东西，所
以写好 README 不仅能为项目加分，同时也显示出你的专业，让别人可以通过 README 文件
，就能对你的项目有一个大致的了解。

本项目 README 文件提供以下内容：

- 项目——Logo 和标语 如果你的项目有一个明显的 logo 标识，那么会增加人们的记忆，特
  别是工具类软件，标语通常是一句简短的话
- 国际化——因为 Github 为全球社区，所以英文文档是必须的，在此基础上可以加上简体中
  文，方便英文基础薄弱的人
- 文档目录——以列表的形式展示了目录索引，让人可以一眼找到自己需要内容
- 项目介绍——这里可以详细介绍项目的背景与相关信息
- 图形展示——通常是项目的截图、视频，让他人可以预览项目
- 特色功能——以列表的形式列出项目的亮点，方便他人了解项目
- 软件架构——这里一般写软件架构图或者目录结构图，方便用户深入理解项目
- 使用方法——告诉他人怎么使用你的项目
- 贡献者——写上作者、贡献者或者赞助项目的人
- 版权声明——告诉别人你的项目版权

以上列出了本项目 README 包含的元素，基本都是一些通用的元素，你可以根据自己的需要
增加或者移除相关元素。项目里面带`tmpl.md`后缀的文件是两个模板文件，可以直接复制
使用，里面有写详细的使用说明。

<span id="nav-1-2"></span>

### 1.2 效果展示

![效果截图](https://cdn.jsdelivr.net/gh/misitebao/CDN@main/md/template-git-111.gif)

<span id="nav-2"></span>

## 2. Issue/PR 模板部分

Issue Template 部分，主要包含了 Bug 和 Features 两部分，也分别做了中英文双语支持
，算是一个简单的模板。

PR 模板就是一个简单的描述和更新列表

<span id="nav-2-1"></span>

### 2.1 简介

<span id="nav-1-2"></span>

### 2.2 效果展示

![效果截图](https://cdn.jsdelivr.net/gh/misitebao/CDN@main/md/template-git-222.gif)

<span id="nav-3"></span>

## 3. Github Actions CI 部分

<span id="nav-3-1"></span>

### 3.1 简介

Github Actions CI 部分，主要是我自己平时经常用的一套脚本， Github Actions 教程请
参考[官方文档](https://docs.github.com/cn/actions)，目录里面两个脚本分别是 Hugo
和 NodeJS 项目的配置。

Hugo 配置：

```yml
name: Deploy | 部署

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  build-and-deploy:
    name: Automatic deployment | 自动部署
    runs-on: ubuntu-latest
    if: github.repository == 'misitebao/standard-repository'

    steps:
      - name: Checkout | 切换到部署分支
        uses: actions/checkout@v2
        with:
          ref: "master"
          submodules: true
          fetch-depth: 0

      - name: Setup Hugo | 设置Hugo环境
        uses: peaceiris/actions-hugo@v2
        with:
          hugo-version: "0.81.0"
          extended: true

      - name: Build Site | 构建网站
        run: |
          hugo

      - name: Deploy to Server | 部署到服务器
        uses: hengkx/ssh-deploy@v1.0.1
        with:
          HOST: ${{ secrets.DEPLOY_HOST }}
          USERNAME: ${{ secrets.DEPLOY_HOST_USER }} # 为了用户信息安全对敏感数据可以在secrets中配置
          PASSWORD: ${{ secrets.DEPLOY_HOST_PASSWORD }}
          SOURCE: "public"
          TARGET: "/www/wwwroot/yoursite.com"
```

NodeJS 配置：

```yml
name: Deploy | 部署

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  build-and-deploy:
    name: Automatic deployment | 自动部署
    runs-on: ubuntu-latest
    if: github.repository == 'misitebao/standard-repository'

    steps:
      - name: Checkout | 切换到部署分支
        uses: actions/checkout@v2

      - name: Build Site | 构建网站
        run: |
          npm install && npm run build

      - name: Deploy to Server | 部署到服务器
        uses: hengkx/ssh-deploy@v1.0.1
        with:
          HOST: ${{ secrets.DEPLOY_HOST }}
          USERNAME: ${{ secrets.DEPLOY_HOST_USER }}
          PASSWORD: ${{ secrets.DEPLOY_HOST_PASSWORD }}
          SOURCE: "build"
          TARGET: "/www/wwwroot/yoursite.com"
```

<span id="nav-3-2"></span>

### 3.2 效果展示

![image-20210809024231277](https://cdn.jsdelivr.net/gh/misitebao/CDN@main/md/image-20210809024231277.png)
