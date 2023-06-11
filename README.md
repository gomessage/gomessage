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

### 安装包下载：

下载地址（国内）：https://gitee.com/gomessage/gomessage/releases （国内用这个，速度很快~）

下载地址（国外）：https://github.com/gomessage/gomessage/releases

<br><br>

### 安装步骤：

<br>

#### Docker方式（强烈推荐）：

前台启动，只运行一次 (容器停止后自动删除，不残留和污染本地环境)：

```bash
docker run --rm \
    -p 7077:7077 \
    gomessage/gomessage:latest 
```

后台运行，只运行一次 (容器停止后自动删除，不残留和污染本地环境)：

```bash
docker run -d \
    -p 7077:7077 \
    --rm \
    --name=gomessage \
    gomessage/gomessage:latest
```

稳定的运行 (且设定为开机启动)：

```bash
docker run -d \
    -p 7077:7077 \
    --restart=always \
    --name=gomessage \
    gomessage/gomessage:latest
```

稳定的运行 (且设定为开机启动)；同时保存数据文件到宿主机上，下次使用新版镜像时，原来的数据不丢失

```bash
docker run -d \
    -p 7077:7077 \
    -v /data:/opt/gomessage/data \
    --restart=always \
    --name=gomessage \
    gomessage/gomessage:latest
```

<br>

#### Kubernetes内部署（完美支持）:

方式一：使用helm部署
```bash
#添加 GoMessage Helm Chart 仓库
helm repo add gomessage "https://occos-helm.pkg.coding.net/repos/gomessage"

#更新helm索引
helm repo update

#安装gomessage到您的集群中
helm install gomessage-service gomessage/gomessage \
  --namespace="default" \                           #指定namespace
  --set ingress.domain="gomessage.taycc.top" \      #指定gomessage暴露至k8s外的访问域名
  --set ingress.ingressClassName="nginx"            #指定k8s中的 Ingress Classes 名称（低版本的k8s可能没有这个东西）
```
> - 如果您的k8s版本较低，那么作者已经封装好的helm chart包您可能无法使用；     
> - 因此可能需要您自己手动封装一个属于自己的helm chart；     
> - 您可以参考 /helm 目录内的文件，尝试封装自己的helm chart；     
> - 如果您不想操作helm，则可以查看`方式二：使用原生yaml脚本部署`的方式来进行部署。

方式二：使用原生yaml脚本部署     
> 这里建议您把GoMessage服务部署成为一个StatefulSet服务；     
> 若您只是临时测试，暂不担心GoMessage的数据丢失，则部署为Deployment也未尝不可。      
> 完整的投产，需要部署四个服务：
> - 部署一个StatefulSet服务
> - 创建一个PVC存储卷
> - 部署一个Service无头服务
> - 部署一个Ingress，把`GoMessage服务`开放到集群外部可以被访问到。

具体的脚本内容：`有空再补充...`


<br><br>

#### Linux服务器上进行安装：

- [Linux环境下的安装](wiki/install.md)

#### Windows服务器上进行安装：

- [Windows环境下的安装（文档暂未完成...）](#windows%E6%9C%8D%E5%8A%A1%E5%99%A8)

#### Mac本地电脑安装和启动：

- [Mac环境下的安装](wiki/install_mac.md)

<br><br>

### 交流和反馈：

微信扫码加入GoMessage技术交流群：
> 如果二维码失效可以添加作者微信号 `SPE3SRU3STAY` 邀请您进群。

<img src="./assets/images/wechat.png" width="300" />

<br><br>

### 操作界面展示

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
