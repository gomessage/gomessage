> - v1版本的gomessage在当前仓库中不再继续更新；
> - 但是考虑到v1版本为单通道的消息转发架构，性能极为出色，因此独立出去成为一个单独的仓库；
> - 新仓库地址为：https://github.com/gomessage/gomessage-v1；
> - 独立出去的仓库不进行较大的架构变动，而是不断定时的修补一些补丁和小规模的功能改进；
> - 如果您想体验gomessage的最新功能，请关注master分之上的最新提交。

# GoMessage

投产架构：

![](https://img.taycc.com/2021-12-27-GoMessage的作用.png)

## 体验地址

点击右侧地址进行体验：[http://47.102.46.109:7077](http://47.102.46.109:7077)

**（作者是个人开发者，体验服务器只有1MB的带宽，加载较慢请稍稍见谅~）**

> 小提示：
> - GoMessage的设计初衷就把它定位为`内网基础设施工具`，因此没有设计`账号密码及权限相关`的模块，广大用户可以借助Nginx的Base登录来控制相关权限。
> - 不排除以后会追加设计权限控制模块，但近期更新的版本暂时不考虑`权限模块`的开发。


<br><br>

## 安装包下载：

下载地址（国内）：https://gitee.com/gomessage/gomessage/releases （国内用这个，速度很快~）

下载地址（国外）：https://github.com/gomessage/gomessage/releases

<br><br>

## 安装步骤：

### `Docker方式`·安装和使用（强烈推荐）：

docker image仓库地址：https://hub.docker.com/r/taycc/gomessage

快速启动：

```bash
docker run -d -p 7077:7077 taycc/gomessage 
```

只运行一次（容器停止后自动删除，不残留和污染本地环境）：
```bash
docker run -d \
    -p 7077:7077 \
    --rm \
    --name=gomessage \
    taycc/gomessage
```

稳定的运行（且设定为开机启动）：
```bash
docker run -d \
    -p 7077:7077 \
    --restart=always \
    --name=gomessage \
    taycc/gomessage
```

<br>

### `Linux服务器：`

- [Linux环境下的安装](./docs/install.md)



### `Windows服务器：`
- [Windows环境下的安装（文档暂未完成...）](#windows%E6%9C%8D%E5%8A%A1%E5%99%A8)



### `Mac本地电脑：`
- [Mac环境下的安装](./docs/install_mac.md)


<br><br>

## GoMessage用途说明

GoMessage是一款消息转发器，主要功能为：

- 用`更友好的图形界面的方式`补全`Prometheus + Alertmanager`报警链路中的最后一个环节，接收Alertmanager报警消息推送，`完成消息格式化`后转发到指定客户端上。
-
- 接收其它工具的消息推送（例如Git的`WebHook回调`），`完成消息格式化`后转发到指定客户端上。
-
- 接收`任意Json结构`的消息推送，`完成消息格式化`后转发到指定客户端上...

软件特色：

- 一次接收，多端推送

- 切割AlertManager消息中的数组，把`Alertmanager聚合组中`聚合起来的消息切割为`单条`然后逐一发送

- 原始数据劫持，可以直观的观察到上游推送过来的json数据结构

- 模板支持`if...else`、`for循环`等逻辑门编写

- 模板支持`Markdown语法`编写

- 模板支持`文本`中插入`CSS样式标签`来控制字体颜色或其它样式渲染

<br><br>

<br><br>

## 操作界面展示

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
