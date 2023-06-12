# GoMessage

### 用途说明：

GoMessage是一款消息转发器，主要功能为：

- 用`更友好的图形界面的方式`补全`Prometheus + Alertmanager`
  报警链路中的最后一个环节，接收Alertmanager报警消息推送，`完成消息格式化`后转发到指定客户端上。
- 接收其它工具的消息推送（例如Git的`WebHook回调`），`完成消息格式化`后转发到指定客户端上。
- 接收`任意Json结构`的消息推送，`完成消息格式化`后转发到指定客户端上...

软件特色：

- 一次接收，多端推送

- 切割AlertManager消息中的数组，把`Alertmanager聚合组中`聚合起来的消息切割为`单条`然后逐一发送

- 原始数据劫持，可以直观的观察到上游推送过来的json数据结构

- 模板支持`if...else`、`for循环`等逻辑门编写

- 模板支持`Markdown语法`编写

- 模板支持`文本`中插入`CSS样式标签`来控制字体颜色或其它样式渲染

<br><br>

### 投产架构：

![](https://img.taycc.com/2021-12-27-GoMessage.png)


<br>

### 版本说明：

> 假设当前版本为：`v2.3.6`
>  - 2代表大版本：表示底层架构发生了`向前不兼容`的更新或变动。
>  - 3代表中版本：表示底层架构依然`向前兼容`，但是`增加了新的功能`。
>  - 6代表小版本：表示`没有增加`任何新功能，只是`对已知的问题和BUG进行了修复`。

 - [查看所有的版本](https://github.com/gomessage/gomessage/releases)

<br><br>

### 体验地址：

点击右侧地址进行体验：[http://47.102.46.109:7077](http://47.102.46.109:7077)

**（作者是个人开发者，体验服务器只有1MB的带宽，加载较慢请稍稍见谅~）**

> 小提示：
> - GoMessage的设计初衷就把它定位为`内网基础设施工具`，因此没有设计`账号密码及权限相关`的模块，广大用户可以借助Nginx的Base登录来控制相关权限。
> - 不排除以后会追加设计权限控制模块，但近期更新的版本暂时不考虑`权限模块`的开发。
> - 加作者微信 `SPE3SRU3STAY` 进入GoMessage技术交流群。




<br><br>

### 安装步骤：


- [使用Docker部署（强烈推荐，满足日常内网使用场景）](wiki/install_docker.md)


- [在Linux裸机上部署（稳定性经过生产验证，部署过程略微繁琐）](wiki/install_linux.md)


- [使用Helm Chart在Kubernetes集群中部署（完美支持）](wiki/install_helm.md)


- [使用Yaml脚本在Kubernetes集群中部署（完美支持）](wiki/install_yaml.md)


- [在Windows机器上启动](wiki/install_windows.md)


- [在Mac机器上启动](wiki/install_mac.md)




<br><br>

### 交流和问题反馈：

微信扫码加入GoMessage技术交流群：
> 如果二维码失效可以添加作者微信号 `SPE3SRU3STAY` 邀请您进群。

<img src="./assets/images/wechat.png" width="300" />

<br><br>

### 操作界面展示

（下图是v1.0版的界面，不排除后面版本中，会对界面UI进行调整）

> 自动识别和生成消息推送地址：

![](https://img.taycc.com/2021-12-27-9HOAd2.png)

<br>

> 劫持json原始数据，声明模板变量，编写模板（模板中可以带逻辑控制）

![](https://img.taycc.com/2021-12-27-UEgRNZ.png)

<br>

> 客户端的多种搭配与激活使用（也可以只创建不激活）

![](https://img.taycc.com/2021-12-27-rWvtmd.png)

<br>

> 新增客户端：添加一个新的企业微信接收客户端

![](https://img.taycc.com/2021-12-27-l2EmY0.png)

<br>

> 新增客户单：添加一个新的飞书接收客户端

![](https://img.taycc.com/2021-12-27-NlUd9w.png)
