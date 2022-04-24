+++
title= "Web Components入门示例" #文章标题
date= 2021-05-15 #文章创建时间
lastmod= 2021-05-15 #文章最后修改时间
draft= false #是否为草稿
description= "" #文章内容描述

tags= ["components"] #文章标签
categories= ["编程技术","大前端"] #文章分类

+++

# Web Components 入门示例

## 文章目录

- [1. 自定义元素](#nav-1)
- [2. window.customElements.define()](#nav-2)
- [3. 自定义元素内容](#nav-3)
- [4. template 标签](#nav-4)
- [5. 添加样式](#nav-5)
- [6. 自定义元素参数](#nav-6)
- [7. Shadow DOM](#nav-7)

组件化一直是前端的发展方向，准确的说不止前端，现阶段流行的 Vue、React 框架都是为组件化编程而生的。

Google 公司一直在推动浏览器的原生组件发展，即`Web Components API`。相比于第三方框架，原生组件更简单，不在依赖外部的框架和库，代码量更小。目前为止，已经在一些浏览器里面得到了原生支持，并且可以用于生产。

这篇文章不是全面的教程，只是简单展示一下 web components 的使用方法，以及需要注意的一些问题。

这篇文章的示例代码，放在仓库：[https://github.com/misitebao/demo-web-components](https://github.com/misitebao/demo-web-components)

<span id="nav-1"></span>

## 1. 自定义元素

```html
<custom-button text="我是自定义参数-按钮文字测试A"></custom-button>
<custom-button text="我是自定义参数-按钮文字测试B"></custom-button>
<custom-button text="我是自定义参数-按钮文字测试C"></custom-button>

<custom-link
  text="我是自定义参数-Link测试A"
  href="https://baidu.com"
></custom-link>
<custom-link
  text="我是自定义参数-Link测试B"
  href="https://github.com"
></custom-link>
```

这种自定义的 HTML 元素，称之为自定义元素。根据规范要求，自定义元素必须包含连接线，主要用于和原生的元素区分。所以不能像在 Vue 种那样使用驼峰写法或者不使用连接线。

<span id="nav-2"></span>

## 2. window.customElements.define()

自定义元素需要使用 JavaScript 定义一个类，所有自定义元素都会是这个类的实例。

```javascript
class Button extends HTMLElement {
  constructor() {
    super();
  }
}
```

上面代码，Button 就是自定义的类，这里使用什么名字都无所谓，但是最好不要声明已经有的类，它继承于 HTMLElement，因此继承了 HTML 元素的特性。接着使用浏览器原生的 window.customElements.define()方法，告诉浏览器中使用的自定义元素与这个类关联。这个方法是实现自定义组件的核心 API。

```
// 实现web components的核心API
window.customElements.define('custom-button', Button);
```

<span id="nav-3"></span>

## 3. 自定义元素内容

自定义元素<custom-button>目前还是空的，下面在类里面给出这个元素的内容。

```javascript
class Button extends HTMLElement {
  constructor() {
    super();

    var button = document.createElement("div");
    button.innerText = "我是自定义组件内容";
    button.classList.add("button");

    this.append(button);
  }
}
```

上面代码最后一行，this.append()的 this 表示自定义元素实例。

完成这一步以后，自定义元素内部的 DOM 结构就已经生成了。

<span id="nav-4"></span>

## 4. template 标签

```html
<!-- 定义模板内容 -->
<template id="buttonTemplate">
  <div class="button"></div>
</template>
```

<span id="nav-5"></span>

## 5. 添加样式

```html
<!-- 定义模板内容 -->
<template id="buttonTemplate">
  <div class="button"></div>
  <style>
    /* :host代表自定义元素本身 */
    :host {
      box-sizing: border-box;
    }
  </style>
</template>
```

<span id="nav-6"></span>

## 6. 自定义元素参数

```javascript
// 定义Button类
class Button extends HTMLElement {
  constructor() {
    super();
    // 获取模板内容
    // 因为复用（多实例）所以此处使用clone
    var templateElem = document.getElementById("buttonTemplate");
    var content = templateElem.content.cloneNode(true);

    // 参数传入
    content.querySelector(".button").innerText = this.getAttribute("text");

    // 将模板内容挂在到当前类
    shadow.appendChild(content);
  }
}
// 实现web components的核心API
window.customElements.define("custom-button", Button);
```

<span id="nav-7"></span>

## 7. Shadow DOM

我们不希望用户能够看到<custom-button>的内部代码，Web Component 允许内部代码隐藏起来，这叫做 Shadow DOM，即这部分 DOM 默认与外部 DOM 隔离，内部任何代码都无法影响外部。

自定义元素的 this.attachShadow()方法开启 Shadow DOM，详见下面的代码。

```javascript
// 定义Button类
class Button extends HTMLElement {
  constructor() {
    super();
    // 隐藏内部代码
    var shadow = this.attachShadow({ mode: "closed" });

    // 获取模板内容
    // 因为复用（多实例）所以此处使用clone
    var templateElem = document.getElementById("buttonTemplate");
    var content = templateElem.content.cloneNode(true);

    // 参数传入
    content.querySelector(".button").innerText = this.getAttribute("text");

    // 将模板内容挂在到当前类
    shadow.appendChild(content);
  }
}
// 实现web components的核心API
window.customElements.define("custom-button", Button);
```

上面代码中，this.attachShadow()方法的参数{ mode: 'closed' }，表示 Shadow DOM 是封闭的，不允许外部访问。

至此，这个 Web Component 组件就完成了，完整代码可以访问这里。可以看到，整个过程还是很简单的，不像第三方框架那样有复杂的 API。
