---
title: 安卓andriod开发者数字证书生成教程
date: 2019-03-17
lastmod: 2019-03-17
draft: false
description: ""
tags: []
categories:
  - 编程技术
---

# 安卓 andriod 开发者数字证书生成

## 文章目录

- [1. 安装 Java 环境](#nav-1)
- [2. 生成签名证书](#nav-2)
- [3. 查看证书信息](#nav-3)

<span id="nav-1"></span>

## 1. 安装 Java 环境

> 自行搜索教程

<span id="nav-2"></span>

## 2. 生成签名证书

使用 keytool -genkey 命令生成证书：

```
keytool -genkey -alias aliasname -keyalg RSA -keysize 2048 -validity 36500 -keystore test.keystore
```

- testalias 是证书别名，可修改为自己想设置的字符，建议使用英文字母和数字
- test.keystore 是证书文件名称，可修改为自己想设置的文件名称，也可以指定完整文件
  路径

回车后会提示：

```
Enter keystore password:  //输入证书文件密码，输入完成回车
Re-enter new password:   //再次输入证书文件密码，输入完成回车
What is your first and last name?
  [Unknown]:  //输入名字和姓氏，输入完成回车
What is the name of your organizational unit?
  [Unknown]:  //输入组织单位名称，输入完成回车
What is the name of your organization?
  [Unknown]:  //输入组织名称，输入完成回车
What is the name of your City or Locality?
  [Unknown]:  //输入城市或区域名称，输入完成回车
What is the name of your State or Province?
  [Unknown]:  //输入省/市/自治区名称，输入完成回车
What is the two-letter country code for this unit?
  [Unknown]:  //输入国家/地区代号（两个字母），中国为CN，输入完成回车
Is CN=XX, OU=XX, O=XX, L=XX, ST=XX, C=XX correct?
  [no]:  //确认上面输入的内容是否正确，输入y，回车

Enter key password for <testalias>
        (RETURN if same as keystore password):  //确认证书密码与证书文件密码一样（HBuilder|HBuilderX要求这两个密码一致），直接回车就可以
```

<span id="nav-3"></span>

## 3. 查看证书信息

```
keytool -list -v -keystore test.keystore
Enter keystore password: //输入密码，回车
```

会输出以下格式信息：

```
Keystore type: PKCS12
Keystore provider: SUN

Your keystore contains 1 entry

Alias name: test
Creation date: 2019-10-28
Entry type: PrivateKeyEntry
Certificate chain length: 1
Certificate[1]:
Owner: CN=Tester, OU=Test, O=Test, L=HD, ST=BJ, C=CN
Issuer: CN=Tester, OU=Test, O=Test, L=HD, ST=BJ, C=CN
Serial number: 7dd12840
Valid from: Fri Jul 26 20:52:56 CST 2019 until: Sun Jul 02 20:52:56 CST 2119
Certificate fingerprints:
         MD5:  F9:F6:C8:1F:DB:AB:50:14:7D:6F:2C:4F:CE:E6:0A:A5
         SHA1: BB:AC:E2:2F:97:3B:18:02:E7:D6:69:A3:7A:28:EF:D2:3F:A3:68:E7
         SHA256: 24:11:7D:E7:36:12:BC:FE:AF:2A:6A:24:BD:04:4F:2E:33:E5:2D:41:96:5F:50:4D:74:17:7F:4F:E2:55:EB:26
Signature algorithm name: SHA256withRSA
Subject Public Key Algorithm: 2048-bit RSA key
Version: 3
```

其中证书指纹信息（Certificate fingerprints）：

- MD5 证书的 MD5 指纹信息（安全码 MD5）
- SHA1 证书的 SHA1 指纹信息（安全码 SHA1）
- SHA256 证书的 SHA256 指纹信息（安全码 SHA245）
