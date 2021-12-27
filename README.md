# GoMessage

## 用途

GoMessage是一款消息转发器，主要功能为：

- 用`更友好的图形界面的方式`补全`Prometheus + Alertmanager`报警链路中的最后一个环节，接收Alertmanager报警消息推送，`完成消息格式化`后转发到指定客户端上。

- 接收其它内网工具（例如Git仓库）的`WebHook回调`，`完成消息格式化`后转发到指定客户端上。

- 接收`任意Json结构`的消息推送，`完成消息格式化`后转发到指定客户端上。



软件特色：

- 一次接收，多端推送

- 切割AlertManager消息中的数组，把`Alertmanager聚合组中`聚合起来的消息切割为`单条`然后逐一发送

- 原始数据劫持，可以直观的观察到上游推送过来的json数据结构

- 模板支持`if...else`、`for循环`等逻辑门编写

- 模板支持`Markdown语法`编写

- 模板支持`文本`中插入`CSS样式标签`来控制字体颜色或其它样式渲染

## 投产架构：

![](https://img.taycc.com/2021-12-27-GoMessage的作用.png)

## 操作界面

> - 自动识别和生成消息推送地址：

![](https://img.taycc.com/2021-12-27-9HOAd2.png)

> - 劫持json原始数据，声明模板变量，编写模板（模板中可以带逻辑控制）

![](https://img.taycc.com/2021-12-27-UEgRNZ.png)

> - 客户端的多种搭配与激活使用（也可以只创建不激活）

![](https://img.taycc.com/2021-12-27-rWvtmd.png)

> - 新增客户端：添加一个新的企业微信接收客户端

![](https://img.taycc.com/2021-12-27-l2EmY0.png)

> - 新增客户单：天界一个新的飞书接收客户端

![](https://img.taycc.com/2021-12-27-NlUd9w.png)





## 体验地址

点击右侧地址进行体验：[http://106.15.51.55:7077](http://106.15.51.55:7077)

**（作者是个人开发者，体验服务器只有1MB的带宽，加载较慢请见谅哈~）**

> 小提示：
> - GoMessage的设计初衷就把它定位为`内网基础设施工具`，因此没有设计`账号密码及权限相关`的模块，广大用户可以借助Nginx的Base登录来控制相关权限。
> - 不排除以后会追加设计权限控制模块，但是至少`v2.0.0版本`的开发计划中暂时没有该模块的设计与追加。

## 安装包下载地址：

打开安装包下载界面：https://gitee.com/gomessage/gomessage/releases

## 安装步骤：

- [Mac环境下的安装](https://gitee.com/gomessage/gomessage/blob/master/docs/install.md#linux%E7%8E%AF%E5%A2%83%E5%AE%89%E8%A3%85)
- [Linux环境下的安装](https://gitee.com/gomessage/gomessage/blob/master/docs/install.md#linux%E7%8E%AF%E5%A2%83%E5%AE%89%E8%A3%85)
- [Windows环境下的安装](https://gitee.com/gomessage/gomessage/blob/master/docs/install.md#linux%E7%8E%AF%E5%A2%83%E5%AE%89%E8%A3%85)
