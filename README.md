# GoMessage

### GoMessage的用途

**GoMessage是一款消息转发器，主要功能为：**

- 用更友好的方式补全`Prometheus + Alertmanager`报警链路中的最后一个环节，接收Alertmanager报警消息推送，`完成消息格式化`后转发到指定客户端上。`（参见：解释1）`
- 接收其它内网工具（例如Git仓库）的`WebHook回调`，`完成消息格式化`后转发到指定客户端上。
- 接收`任意Json结构`的消息推送，`完成消息格式化`后转发到指定客户端上。

> 解释1：
> - GoMessage实际上是提供了`更友好`的`消息转发能力`，在`友好度`和`操作便捷性`上有很多设计和思考，其目的也是为了解决`通用场景下`的消息转发问题~
> - 技术选型上，虽然选择Golang来完成底层逻辑的实现，但是性能瓶颈依然存在~
> - Prometheus官方生态圈中提供有`多种消息转发方案`，对`消息转发场景`有`复杂使用需求`的小伙伴，可以先从官方生态圈中找找解决方案，如果实在找不着，那可以联系作者构建新分支去帮你实现`非通用场景下的复杂需求`~



软件特色：

- 一次接收，多端推送
- 切割AlertManager消息中的数组，把`Alertmanager聚合组中`聚合起来的消息切割为`单条`然后逐一发送
- 元数据劫持，可以直观的观察到上游推送过来的json数据结构
- 模板支持`if...else`、`for循环`等逻辑门编写
- 模板支持`Markdown语法`编写
- 模板支持`文本`中插入`CSS样式标签`来控制字体颜色或其它样式渲染

### 投产架构：

![](https://img.taycc.com/2021-12-27-GoMessage的作用.png)

### 基本界面

![](https://img.taycc.com/2021-12-27-9HOAd2.png)

![](https://img.taycc.com/2021-12-27-UEgRNZ.png)

![](https://img.taycc.com/2021-12-27-rWvtmd.png)

![](https://img.taycc.com/2021-12-27-l2EmY0.png)

![](https://img.taycc.com/2021-12-27-NlUd9w.png)

### 接收端的消息样式

#### 飞书消息（分割）发送：

![](https://img.taycc.com/2021-12-27-nL4rRM.png)

#### 飞书消息（聚合）发送：

![](https://img.taycc.com/2021-12-27-kkFxHD.png)

#### 钉钉消息发送

![](https://img.taycc.com/2021-12-27-3hlvkE.png)

#### 企业微信消息发送

![](https://img.taycc.com/2021-12-27-DF6T9p.png)


### 体验地址

点击右侧地址进行体验：[http://106.15.51.55:7077](http://106.15.51.55:7077)

**（作者是个人开发者，体验服务器只有1MB的带宽，加载较慢请见谅哈~）**

> 小提示：
> - GoMessage的设计初衷就把它定位为`内网基础设施工具`，因此没有设计`账号密码及权限相关`的模块，广大用户可以借助Nginx的Base登录来控制相关权限。
> - 不排除以后会追加设计权限控制模块，但是至少`v2.0.0版本`的开发计划中暂时没有该模块的设计与追加。

### 安装包下载地址：

打开安装包下载界面：https://gitee.com/gomessage/gomessage/releases

### 安装步骤：

- [Mac环境下的安装](https://gitee.com/gomessage/gomessage/blob/master/docs/install.md#linux%E7%8E%AF%E5%A2%83%E5%AE%89%E8%A3%85)
- [Linux环境下的安装](https://gitee.com/gomessage/gomessage/blob/master/docs/install.md#linux%E7%8E%AF%E5%A2%83%E5%AE%89%E8%A3%85)
- [Windows环境下的安装](https://gitee.com/gomessage/gomessage/blob/master/docs/install.md#linux%E7%8E%AF%E5%A2%83%E5%AE%89%E8%A3%85)
