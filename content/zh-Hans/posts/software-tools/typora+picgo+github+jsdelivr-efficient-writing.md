+++
title= "Typora+PicGo+Github+Jsdelivr实现高效图文写作"
draft= false
date= 2019-05-16
lastmod= 2019-07-13
license= "MIT"

tags= ["软件工具","typora","picgo","jsdelivr","cdn"]
categories= ["软件工具"]

+++

# Typora+PicGo+Github+Jsdelivr 实现高效图文写作

> 声明：我是 「[米司特包](http://misitebao.com)」 ，本篇文章 **首发** 于 「[米司博客](http://blog.misitebao.com)」 ，其他平台为同步推送（_因为可能文章路径会变动，所以就没放详细链接，进来的可以 [根据标题查找 🔍](http://blog.misitebao.com/posts/)_）。

![Typora+PicGo+Github+Jsdelivr实现高效图文写作](https://cdn.jsdelivr.net/gh/misitebao/CDN/cover/003.png)

这段时间有看到小伙伴儿讨论图床的问题，都想知道有没有好用的图床，答案肯定是有的，比如国内的：腾讯云、阿里云、七牛云、又拍云等等，基本都是提供的有免费的`CDN`供用户使用，大概都是 10 个 G 左右的空间（具体的我不清楚，毕竟自己只用过七牛的，别的只是知道而已。。。），对于一般需求量不大的用户来说肯定是没问题了，但是对于技术产出较高的大佬们来说，肯定不够，所以这里结合我自己使用的一些方法总结了这篇文章，跟大家分享我自己的解决方案。

## 工具简介

### typora

我现在自己写文章、文档什么的基本都使用`markdown`，这是一种比较好的文本标记语言，特别是相对于编程的同学来说，这里我个人使用的`Typora`，这个我很早就接触这个东西了，更早的时候是使用`vscode`自带的 markdown 或者`HBuildX`去写的，后来偶然看到了`typora`就去使用这个了，经过多个版本的迭代，之前用的不爽的地方基本都改的差不多了，现在使用起来已经很流畅了（除了每次更新的时候有点慢。。）,至于`markdown`怎么用，这个我就不写了，网上教程大把，自己去找吧，界面长这样子

![image-20200527191211267](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527191214.png)

### PicGo

`PicGo`就是一个图片上传工具，目前支持`SM.MS图床`、`腾讯云COS`、`GitHub`（待会儿我们用到的）、`七牛图床`、`Imgur图床`、`阿里云OSS`、`又拍云图床`，用法也很简单，在配置页面配置好各个图床的参数，如下图：

![image-20200527191846033](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527191847.png)

然后就可以使用这个工具将图片上传到你配置的那个图床上面去，上传方式目前有三种：

- 拖拽上传
- 快捷键上传（需要剪切板里面有图片，也就是你要先复制一张图片）
- 命令行上传

### Github

emmm，这个我就不说了吧，能看到这篇文章的应该没有不知道的，万一真不知道也没关系。

### Jsdelivr

![image-20200527192400115](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527192402.png)

如图所示，就是 开源的免费`CDN`，快速、可靠、自动化。他能够让我们快速的访问我们`github`、`npm`、`wordpress`上开放的资源。

## 使用方法

### 1.下载安装`typora +PicGo`

`Typora`下载地址：https://www.typora.io/

`PicGo`下载地址：https://github.com/Molunerfinn/PicGo/releases 找到图片所示位置

![image-20200527193219825](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527204036.png)

### 2.申请 github 令牌

如果你已经有`github`账号，直接登录，没有的话就注册在登录，怎么注册，我就不写了，登陆进去鼠标点击头像，找到`Settings`，点击

![image-20200527193744232](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527204021.png)

找到`Developer settings`，点击

![image-20200527193918319](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527193920.png)

![image-20200527194130744](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527194131.png)

期间会让你输入密码，直接输入你的`github`账号密码就行，然后填写名字，名字随便写（尽量让自己知道是什么意思），下面红框里面的必须勾上，其他的不用管

![image-20200527194321556](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527194407.png)

完成后点击页面左下方的绿色按钮，然后成功以后会回到之前的页面，显示令牌

![image-20200527194516074](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527194517.png)

这个令牌只显示一次，要复制令牌，自己找个地方保存起来，后面步骤要用到这个令牌

### 3.配置 PicGo

打开`PicGo`，找到 图床设置>`Github`图床

![image-20200527195543276](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527195544.png)

- 仓库名填写，你的`github`上面你要放图片的那个仓库，必须要要是公有库；

- 分支名，填写你要上传图片的仓库的那个分支，一般填写`master`

- `Token`填写刚刚让你复制下来找个地方的保存的那串字符

- 储存路径，这个是相对于你的仓库根路径来说的，切记前面不能加斜杠，后面必须要加斜杠。如`bg/`
- 自定义域名：`https://cdn.jsdelivr.net/gh/你的github用户名/你设置的仓库名@分支名`，这个主要是用来上传以后，他用来拼接路径存在你的剪切板里面，返回图片的路径给你（这里的分支名这个位置，除了填写分支名以外，还可以填写 release 名称和 commit tag，这个专业人士请自行研究，非专业人事就直接用分支名就可以）

填写完成以后，点击确认，设置为默认图床，然后就可以切换到上传区，上传一张图片试试，上传完成以后会返回路径给你，这个路径格式也是可以设置的

![image-20200527200833855](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527200835.png)

### 4.配置 Typora

打开`Typora` 找到设置里面的 图像

![image-20200527194826427](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527194827.png)

![image-20200527194911887](https://cdn.jsdelivr.net/gh/misitebao/CDN/md/20200527194912.png)

`PicGo`路径选你安装的位置的路径，然后点击验证图片上传选项，如果提示成功就大功告成，如果失败，请重复上述步骤。

如有问题，请私信https://blog.misitebao.com/contact.html

![wechat](https://cdn.jsdelivr.net/gh/misitebao/CDN/wechat_misitebao.jpg)
