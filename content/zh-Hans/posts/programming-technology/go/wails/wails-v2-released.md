---
# 预定义
title: "Wails 2.0 正式发布，可以使用现代化 Web 技术开发桌面应用的 Go 框架" #文章标题
date: 2022-09-22 #文章创建时间
lastmod: 2022-09-22 #文章最后修改时间
draft: false #是否为草稿
weight: 0
slug: "wails-v2-released"
url: "posts/wails-v2-released"
keywords: ["wails", "go", "webview2", "desktop", "桌面应用"]
description: "Wails 2.0 正式发布，可以使用现代化 Web 技术开发桌面应用的 Go 框架" #文章内容描述
summary: "Wails 2.0 正式发布，可以使用现代化 Web 技术开发桌面应用的 Go 框架" # 文章摘要

# taxonomies
tags: ["go", "wails", "桌面应用", "webview", "webview2"] #文章标签
categories: ["编程技术", "Go", "Wails"] #文章分类

# 自定义
---

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/wailsapp/wails@2.0.0/website/static/img/blog/montage.png" alt="wails"/>
</div>

> 本文转载于 [Wails](https://github.com/wailsapp/wails) >
> [官方博客](https://wails.io/blog/wails-v2-released)，已获得官方授权。

## 内容目录

- [内容目录](#内容目录)
- [前言](#前言)
- [什么 _是_ Wails?](#什么-是-wails)
- [有哪些新东西？](#有哪些新东西)
- [致敬和感谢](#致敬和感谢)
- [吸取的教训](#吸取的教训)
  - [更小、更快、更聚焦的发布](#更小更快更聚焦的发布)
  - [鼓励参与](#鼓励参与)
  - [学会说不](#学会说不)
- [展望未来](#展望未来)
- [最后想说的](#最后想说的)

## 前言

今天标志着 [Wails](https://wails.io/) v2 的发布。 距离发布第一个 v2 alpha 版本大
约 18 个月，距离发布第一个 beta 版本大约一年。 我非常感谢参与项目发展的每一个人
。

花了这么长时间的部分原因是想要在正式将其称为 v2 之前对其进行一些完整性的定义。事
实是，从来没有一个完美的时间来标记一个版本 - 总是有突出的问题或 “只是一个” 功能
要加进去。 然而，标记一个不完美的主要版本的确是为项目的用户提供一些稳定性，并为
开发人员提供了一些重置。

这此发布超出了我的预期。 我希望它给您带来的乐趣与开发它给我们带来的乐趣一样多。

## 什么 _是_ Wails?

如果你对 Wails 不熟悉，它是一个让 Go 程序员能够使用熟悉的 Web 技术为其 Go 程序提
供丰富的前端的项目。 它是 Electron 的轻量级 Go 替代品。 更多信息可以在
[官方网站 ](https://wails.io/docs/introduction)上找到。

## 有哪些新东西？

V2 版本是该项目的巨大飞跃，解决了 v1 的许多痛点。 如果您尚未阅读任何有关
[macOS](https://wails.io/blog/wails-v2-beta-for-mac)、[Windows](https://wails.io/blog/wails-v2-beta-for-windows)
或 [Linux](https://wails.io/blog/wails-v2-beta-for-linux) Beta 版本的博客文章，
那么我鼓励您阅读，因为它更详细地涵盖了所有主要变更。 总结起来就是：

- 适用于 Windows 的 Webview2 组件，支持现代 Web 标准和调试功能。
- Windows 上的 [深色/浅色主题](https://wails.io/docs/reference/options#主题) 和
  [自定义主题](https://wails.io/docs/reference/options#自定义主题)。
- Windows 现在没有 CGO 要求。
- 支持对 Svelte、Vue、React、Preact、Lit 和 Vanilla 项目模板的开箱即用。
- 集成 [Vite](https://vitejs.dev/) 为您的应用程序提供热重载开发环境。
- 原生应用的
  [菜单](https://wails.io/docs/guides/application-development#应用程序菜单) 和
  [对话框](https://wails.io/docs/reference/runtime/dialog)。
- 适用于 [Windows](https://wails.io/docs/reference/options#窗口半透明) 和
  [macOS](https://wails.io/docs/reference/options#窗口半透明-1) 的原生窗口半透明
  效果。 支持 Mica 和 亚克力 背景。
- 为 Windows 部署轻松生成 [NSIS](https://wails.io/docs/guides/windows-installer)
  安装程序。
- 一个为窗口操作、事件、对话框、菜单和日志记录提供实用方法的丰富的
  [运行时库](https://wails.io/docs/reference/runtime/intro)。
- 支持使用 [garble](https://github.com/burrowers/garble)
  [混淆](https://wails.io/docs/guides/obfuscated) 您的应用程序。
- 支持使用 [UPX](https://upx.github.io/) 压缩您的应用程序。
- 自动为 Go 结构体的生成 Typescript。 更多信息在
  [这里](https://wails.io/docs/howdoesitwork#调用绑定的-go-方法)。
- 您的应用程序不需要附带额外的库或 DLL。 适用于任何平台。
- 无需打包前端资产。 只需像开发任何其他 Web 应用程序一样开发您的应用程序。

## 致敬和感谢

到达 v2 是一项巨大的努力。 从最初的 alpha 版本到今天的发布，已经有 89 位贡献者提
交了大约 2200 次提交，还有很多很多的贡献者在讨论论坛和 Issue 跟踪器上提供了翻译
、测试、反馈和帮助。 我非常感谢你们每一个人。 我还要特别感谢所有提供指导、建议和
反馈的项目赞助商。 非常感谢您所做的一切。

我想特别提及以下几个人：

首先，**非常** 感谢 [@stffabi](https://github.com/stffabi)，他提供了许多我们都从
中受益的贡献，并在许多 Issue 上提供了很多支持。 他提供了一些关键功能，例如外部开
发服务器的支持，这改变了我们的开发模式，允许我们与 [Vite](https://vitejs.dev/)
的超级能力挂钩。 客观地说，如果没有他令人
[难以置信的贡献](https://github.com/wailsapp/wails/commits?author=stffabi&since=2020-01-04)，Wails
v2 将是一个不那么令人兴奋的版本。 非常感谢 @stffabi！

我还想对 [@misitebao](https://github.com/misitebao) 表示热烈的祝贺，他一直不懈地
维护网站、提供中文翻译、管理 Crowdin 并帮助新的翻译人员跟上进度。 这是一项非常重
要的任务，我非常感谢为此付出的所有时间和精力！ 你太棒了！

最后但同样重要的是，非常感谢 Mat Ryer，他在 v2 的开发过程中提供了建议和支持。 使
用 v2 的早期 Alpha 一起编写 xBar 有助于塑造 v2 的方向，并让我了解早期版本中的一
些设计缺陷。 我很高兴地宣布，从今天开始，我们将开始将 xBar 移植到 Wails v2，它将
成为该项目的旗舰应用。 干杯！

## 吸取的教训

在到达 v2 的过程中吸取了许多经验教训，这些经验教训将影响未来的发展。

### 更小、更快、更聚焦的发布

在开发 v2 的过程中，有许多特性和错误修复是在临时基础上开发的。 这导致了更长的发
布周期并且更难调试。 未来，我们将更频繁地创建包含较少数量功能的发布。 发布一个版
本将包括文档的更新和完全的测试。 希望这些更小、更快、更集中的发布将导致更少的回
归和更好的文档质量。

### 鼓励参与

在开始这个项目时，我想立即帮助每个有问题的人。 问题是“个人的”，我希望它们尽快得
到解决。 这是不可持续的，最终会影响项目的寿命。 未来，我将为人们提供更多空间来参
与回答问题和对问题进行分类。 通过一些工具来帮助解决这个问题会很好，所以如果您有
任何建议，请在 [此处](https://github.com/wailsapp/wails/discussions/1855) 加入讨
论。

### 学会说不

参与开源项目的人越多，对大多数人可能有用也可能没用的附加功能的请求就越多。 这些
特性的开发和调试需要一段初始的时间，并从那时起产生持续的维护成本。 我本人对此负
有最大责任，往往想要 “好高骛远”，而不是提供最小的可行功能。 未来，我们需要对添加
核心功能多说一些 “不”，并将我们的精力集中在授权开发人员自己提供该功能的方式上。
我们正在认真研究这种情景的插件。 这将使任何人都能够按照他们认为合适的方式扩展该
项目，并提供一种为项目做出贡献的简单方法。

## 展望未来

在下一个主要开发周期中，我们正在考虑将许多核心功能添加到 Wails 中。
[路线图](https://github.com/wailsapp/wails/discussions/1484) 充满了令人感兴趣的
想法，我非常希望着手处理这些想法。 最多的要求之一是多窗口支持。 这是一个棘手的问
题，并且要正确执行，我们可能需要考虑提供替代 API，因为当前的 API 在设计时并未考
虑到这一点。 根据一些初步的想法和反馈，我想你会喜欢我们的预期。

我个人对在移动设备上运行 Wails 应用程序的前景非常感兴趣。 我们已经有一个演示项目
表明可以在 Android 上运行 Wails 应用程序，所以我真的很想探索我们可以用它做到什么
地步！

我想提出的最后一点是功能均等性。 长期以来，我们不会在项目中添加任何东西，除非有
完整的跨平台支持，这一直是一个核心原则。 虽然到目前为止这已被证明（主要）是可以
实现的，但它确实阻碍了项目发布新功能。 未来，我们将采用稍微不同的方法：任何无法
立即为所有平台发布的新功能都将在实验配置或 API 下发布。 这允许某些平台上的早期采
用者尝试该功能并提供反馈，这些反馈将反馈到该功能的最终设计中。 当然，这意味着在
所有可以支持它的平台完全支持之前，不能保证 API 的稳定性，但至少它会解除对开发的
阻碍。

## 最后想说的

我真的为我们在 V2 版本中所取得的成就感到自豪。 看到人们已经能够使用 beta 版本构
建的东西真是太神奇了。 像
[Varly](https://varly.app/)、[Surge](https://getsurge.io/) 和
[October](https://october.utf9k.net/) 等优质的应用程序。 我鼓励你们去看看。

这个版本是通过许多贡献者的辛勤工作实现的。 虽然它可以免费下载和使用，但它并不是
通过零成本实现的。 毫无疑问，这个项目的成本很大。 这不仅是我的时间和每一个贡献者
的时间，也是这些人的朋友和家人缺席的代价。 这就是为什么我非常感谢为实现这个项目
而付出的每一秒。 我们拥有的贡献者越多，这种努力就可以分散得越多，我们可以共同实
现的目标就越多。 我想鼓励大家选择一件可以贡献的东西，无论是确认某人的错误、提出
修复建议、更改文档还是帮助需要它的人。 所有这些小事都有如此巨大的影响！ 如果你也
能成为 v3 故事的一部分那就太棒了。

请尽情享用！

- Lea

PPS：如果您或您的公司发现 Wails 有用，可以考虑
[赞助该项目](https://github.com/sponsors/leaanthony)。 谢谢！