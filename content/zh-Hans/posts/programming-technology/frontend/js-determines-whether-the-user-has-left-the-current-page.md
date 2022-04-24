+++
title= "JS判断用户是是否离开当前页面"
subtitle= ""
date= 2018-01-22
lastmod= 2018-01-22
draft= false
description= ""

tags= ["javascript"]
categories= ["编程技术","大前端","JavaScript"]

+++

# JS 判断用户是是否离开当前页面

## 文章目录

- [1. 介绍](#nav-1)
- [2. 代码](#nav-2)
- [3. 效果](#nav-3)
- [4. `document.visibilityState` 监听浏览器最小化](#nav-4)

<span id="nav-1"></span>

## 1. 介绍

`visibilitychange`事件：用于判断用户是否离开当前页面

![](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200519162115.png)

<span id="nav-2"></span>

## 2. 代码

```html
<!DOCTYPE html>
<html lang="zh">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>初始化</title>
  </head>

  <body>
    <script>
      !(function () {
        // 添加事件监听
        document.addEventListener("visibilitychange", function () {
          let pageVisibility = document.visibilityState;
          // 页面变为不可见时触发
          if (pageVisibility == "hidden") {
            document.title = "主人，你去哪儿了？";
          }
          // 页面变为可见时触发
          if (pageVisibility == "visible") {
            document.title = "主人，你终于回来了？";
          }
        });
      })();
    </script>
  </body>
</html>
```

<span id="nav-3"></span>

## 3. 效果

![GIF 2020-5-19 16-34-55.gif](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/2020-5-1916-44-15.gif)

<span id="nav-4"></span>

## 4. `document.visibilityState` 监听浏览器最小化

`document.hidden`表示页面是否隐藏的布尔值。页面隐藏包括 页面在后台标签页中 或者 浏览器最小化 （注意，页面被其他软件遮盖并不算隐藏，比如打开的 `vscode` 遮住了浏览器）。

`document.visibilityState`值：

- hidden：页面在后台标签页中或者浏览器最小化
- visible：页面在前台标签页中
- prerender：页面在屏幕外执行预渲染处理 `document.hidden` 的值为 true
- unloaded：页面正在从内存中卸载

visibilitychange 事件：当文档从可见变为不可见或者从不可见变为可见时，会触发该事件

这样，我们可以监听 `visibilitychange` 事件，当该事件触发时，获取 `document.hidden` 的值，根据该值进行页面一些事件的处理
