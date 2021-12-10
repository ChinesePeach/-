# 广西大学程序设计竞赛邮件群发系统

## 前言

广西大学程序设计竞赛组委会邮箱:

|        邮箱       |    授权码 严禁外传  |
|------------------|------------------|
|  gxucpc@163.com  |                   |

## 项目介绍

广西大学程序设计竞赛邮件群发系统是为了减少竞赛报名后工作人员群发邮件的工作量而编写的

仅需要参赛人员的邮箱，即可完成一键群发

本篇README.md面向使用者

## 目录

- [xlsx文件书写规范](#xlsx文件书写规范)
- [txt文件书写规范](#txt文件书写规范)
- [使用说明](#使用说明)
- [作者](#作者)
- [鸣谢](#鸣谢)

### xlsx文件书写规范

首先Excel文件格式必须使用xlsx，名称为 `email.xlsx`

将参赛人员邮箱放到A列

不要添加其余无关元素

![示例图](https://i.bmp.ovh/imgs/2021/12/1438527849697448.png)

### txt文件书写规范

文件名称为：`ProjectHtml.txt`

txt的第一行为邮件标题

第二行开始为HTML模板，只需要修改其中部分内容

Ctrl + F 搜索文字 “搜索标记1”

之后按需修改p标签的内容就可以

Ctrl + F 搜索文字 “搜索标记2”

改一下年份

### 使用说明
1.  **主流邮件运营商服务器地址与端口**
- 网易163邮箱（mail.163.com）:

    POP3服务器地址：pop.163.com（端口：110）

    SMTP服务器地址：smtp.163.com（端口：25）


- 网易126邮箱（mail.126.com）：

    POP3服务器地址：pop.126.com（端口：110）

    SMTP服务器地址：smtp.126.com（端口：25）


- 移动139邮箱（mail.10086.cn）：

    POP3服务器地址：POP.139.com（端口：110）

    SMTP服务器地址：smtp.139.com(端口：25)



- 腾讯QQ邮箱（mail.qq.com）：

    POP3服务器地址：pop.qq.com（端口：110）

    SMTP服务器地址：smtp.qq.com （端口：25）


- 腾讯QQ企业邮箱（exmail.qq.com） ：

    POP3服务器地址：pop.exmail.qq.com （SSL启用 端口：995）

    SMTP服务器地址：smtp.exmail.qq.com（SSL启用 端口：587/465）


- 谷歌Gmail邮箱（mail.google.com）：

    POP3服务器地址：pop.gmail.com（SSL启用 端口：995）

    SMTP服务器地址：smtp.gmail.com（SSL启用 端口：587）


- 腾讯Foxmail邮箱（mail.qq.com）：

    POP3服务器地址：pop.foxmail.com（端口：110）

    SMTP服务器地址：smtp.foxmail.com（端口：25）


2. 邮件From格式：发件署名+[空格]+”<电子邮件>“

    例如：广西大学程序设计竞赛组委会 <gxucpc@163.com>


3. 附件
    
    由于未知原因，中文文件名的附件发出去会乱码（文件名称），建议采用`ASCII`字符
    
    输入附件名称要写全，包括名称+格式
    
    例如123.pdf


4. 关于防止被封禁

    一般运营商要求15min内不得发出超过200封邮件
    
    根据这个要求，每次发送邮件都暂停5s `time.Sleep(5 * time.Second)`



### 作者

sct

邮箱:1272607918@qq.com

qq:1272607918

### 鸣谢

- [GitHub excelize](https://github.com/qax-os/excelize)
- [GitHub jordan-wright/email](https://github.com/jordan-wright/email)
- [email-templates](https://wordpress.org/plugins/email-templates/)



