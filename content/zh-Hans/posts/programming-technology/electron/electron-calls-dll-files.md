---
title: Electron 调用 DLL 文件
draft: false
date: 2018-10-22
lastmod: 2019-05-09
license: MIT
tags:
  - electron
categories:
  - 编程技术
  - 大前端
  - Electron
---

# Electron 调用 DLL 文件

> 本文由米司在开发 electron 应用，找资料时发现。
>
> 转载自 [Ke Haw 🇨🇳]
> (<[http://www.kehaw.com/electron/node.js/2019/11/11/Electron-%E4%B8%8B%E8%B0%83%E7%94%A8DLL%E6%96%87%E4%BB%B6.html](http://www.kehaw.com/electron/node.js/2019/11/11/Electron-下调用DLL文件.html)>)
>
> 已获得作者授权

Node.JS 调用 DLL 文件分两种方式，其一是通过 Node Addon 将 C++ 程序编译成 Addon
加载到 Chromium 引擎中，然后通过 JS 去调用，这是比较正规的做法，但是这种做法需要
开发人员有一定的 C++ 编程技能，故而可能比较难以实现。

还有一种做法是基于 node.js 的 ffi-napi 开源项目去加载 dll 文件。该 dll 文件必须
由 c++ 编写，而基于 c# 编写的 dll 文件，请参考另外一个开源项目：`edge`，本文主要
讲述如何使用 `ffi-napi` 来调用 C++ 编写的 dll 动态链接库。

> 请勿使用 node-ffi 而是使用 ffi-napi

## 首先编写一个 DLL

为了讲述如何基于 `ffi-napi` 调用 dll 文件，我们首先编写一个 dll 文件，在 VS 2019
中创建一个基于 C++ 的 dll 库项目，项目名为 `demo-dll`，然后在 `pch.h` 文件中加入
函数定义如下：

```c++
// pch.h: 这是预编译标头文件。
// 下方列出的文件仅编译一次，提高了将来生成的生成性能。
// 这还将影响 IntelliSense 性能，包括代码完成和许多代码浏览功能。
// 但是，如果此处列出的文件中的任何一个在生成之间有更新，它们全部都将被重新编译。
// 请勿在此处添加要频繁更新的文件，这将使得性能优势无效。

#ifndef PCH_H
#define PCH_H

// 添加要在此处预编译的标头
#include "framework.h"

#endif //PCH_H

#ifdef IMPORT_DLL
#else
#define IMPORT_DLL extern "C" _declspec(dllimport)
#endif // IMPORT_DLL

IMPORT_DLL int add(int a, int b);
IMPORT_DLL int minus(int a, int b);
IMPORT_DLL int multiply(int a, int b);
IMPORT_DLL double divide(int a, int b);
```

然后在 `pch.cpp` 文件中编写函数体：

```c++
// pch.cpp: 与预编译标头对应的源文件

#include "pch.h"

// 当使用预编译的头时，需要使用此源文件，编译才能成功。

int add(int a, int b)
{
	return a + b;
}

int minus(int a, int b)
{
	return a - b;
}

int multiply(int a, int b)
{
	return a * b;
}

double divide(int a, int b)
{
	double m = (double)a / b;
	return m;
}
```

我们可以看到，就是最简单的加减乘除四个函数，然后执行编译生成 DLL。

然后，重新建立一个控制台项目，来验证 dll 文件是否能够正常使用，在这里就不写具体
的例子了。

## 自行创建一个最简单的 Electron 项目调用 dll

空白 Electron 项目创建，首先在空白目录下执行：

```javascript
npm init
```

接下来执行 install 命令

```
npm install --save-dev electron
```

然后新建一个`main.js`：

```
const { app, BrowserWindow } = require('electron')

function createWindow () {
  // 创建浏览器窗口
  let win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      nodeIntegration: true
    }
  })

  // 加载index.html文件
  win.loadFile('index.html')
}

app.on('ready', createWindow)
```

然后新建一个`index.html`：

```
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>Hello World!</title>
    <!-- https://electronjs.org/docs/tutorial/security#csp-meta-tag -->
    <meta http-equiv="Content-Security-Policy" content="script-src 'self' 'unsafe-inline';" />
  </head>
  <body>
    <h1>Hello World!</h1>
    We are using node <script>document.write(process.versions.node)</script>,
    Chrome <script>document.write(process.versions.chrome)</script>,
    and Electron <script>document.write(process.versions.electron)</script>.
  </body>
</html>
```

最后修改 package.json 配置文件如下：

```
{
  "name": "your-app",
  "version": "0.1.0",
  "main": "main.js",
  "scripts": {
    "start": "electron ."
  }
}
```

最终的目录结构如下：

```
your-app/
├── package.json
├── main.js
└── index.html
```

运行程序：

```
npm run start
```

正确运行之后，为项目加入`ffi-napi`依赖：

```
npm install --save ffi-napi
```

安装成功之后，在`main.js`下加入代码：

```
const ffi = require('ffi-napi');

console.log(__dirname + "\\demo-dll.dll");//确认dll库的路径是否正确

const libm = ffi.Library(__dirname + '\\demo-dll.dll', {
	'add': ['int', ['int', 'int']]
});

const result = libm.add(1, 1);

console.log(result);//这里会打印出正确的的计算结果
```

然后将上面编译好的 dll 文件放入到`dist_electron`目录下。

## 第三方 DLL 调用注意事项

首先要注意对方的 dll 是基于 `c#` 还是 `c++` 编写的，因为两种动态链接库的是不一样
的调用方式。

其次要注意对方的 dll 库是 x86 还是 x64 的，不一样的软件架构会导致程序崩溃。

在调用动态链接库中的方式为：

```
const libm = ffi.Library(__dirname + '\\demo-dll.dll', {
	'add': ['int', ['int', 'int']]
});
```

`add`为函数名，数组中的第一个参数代表这个函数的返回类型，第二个参数代表所需要传
入给这个函数的参数。

然后直接使用`Library`对象来调用就可以了。

### 启动命令

```
npm run electron:serve
```

启动之后，会在控制台中发现打印了结果（非 Chrome 控制台）。
