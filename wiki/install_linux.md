# Linux环境安装

<br><br>

GoMessage当前支持在`Linux裸机`上进行部署，部署在`Linux裸机上的GoMessage服务`其稳定性经历过过生产环境高并发验证。

<br><br>

# 安装包下载：

安装包下载地址（国内）：https://gitee.com/gomessage/gomessage/releases （这个地址国内访问速度快~）

安装包下载地址（国外）：https://github.com/gomessage/gomessage/releases

> 目前提供有Linux、Mac、Windows三个版本的安装包（x86_64架构），如需支持其它版本，可以自行下载源码编译，也可以直接联系作者编译出指定版本的安装包。


<br><br>

## 手动部署（详细操作步骤）

> 本文以`Centos7.x`环境来举例     


### 安装包下载：

```bash
# 进入到linux命令行的家目录中
cd ~/

# 下载安装包
wget https://github.com/gomessage/gomessage/releases/download/版本号/gomessage-版本号-linux-x64.tar.gz
```

<br>

### 创建工作目录

（作者推荐您将工作目录设置在 /opt/gomessage/ 目录中，可以干净清爽一些，不会污染系统上干的其它目录）

```bash
# 创建一个新目录/opt/gomessage
mkdir -p /opt/gomessage
```

<br>

### 解压安装包

解压前面下载好的安装包

```bash
# 解压gomessage压缩包的内容到/opt/gomessage/目录下
tar -zxvf ./gomessage-x.x.x-linux-x64.tar.gz -C /opt/gomessage/
```

解压完成后，GoMessage就算是部署好了，直接运行一下命令就可以运行起来GoMessage服务：

```bash
# 授予可执行权限
sudo chmod +x /opt/gomessage/gomessage

# 对，您没看错，就是这么一行命令，什么参数都不需要
/opt/gomessage/gomessage
```

此时GoMessage默认会启动和监听本机的`7077`端口;

您只需要用浏览器访问：`http://<服务器地址>:7077/` 就可以访问到GoMessage服务了。

假设此时我的服务器地址为：`192.168.33.201`，那么我访问GoMessage的地址就应该是：`http://192.168.33.201:7077`


<br>


> 虽然GoMessage服务已经可以正常启动了；     
> 但是用上面挂载前台的方式启动GoMessage服务是不可靠的。
> 
> Linux会有一套自己的QOS机制来管理自身的资源分配和回收；     
> 因此作者推荐您将GoMessage服务封装为`Systemd守护进程`的方式来运行。

<br><br>

## 封装systemd守护进程（以守护进程的方式来运行GoMessage）

将下面这段配置内容写入到`/usr/lib/systemd/system/gomessage.service`这个文件内容，就算是把systemd守护进程封装好啦；

如果您没有把`/opt/gomessage/目录`设定为您的解压目录，那么一定要记得把下文中`追加有注释的那两行内容中的"目录路径"`替换为您自己的`实际路径`；

别的部分您都可以直接抄写作者已经写好的默认配置。

（您可以直接粘贴下面的命令，去命令行执行，如果`目录路径`没错，那么一切的配置参数都是可靠且值得信赖的，可以放心使用）

```bash
cat > /usr/lib/systemd/system/gomessage.service << \EOF

[Unit]
Description=GoMessage
After=network.target


[Service]
Type=simple
User=root
Group=root
LimitNOFILE=65535
LimitNPROC=65535
WorkingDirectory=/opt/gomessage/       #此处是您GoMessage服务的解压目录（可以替换成您自己的实际目录位置）
ExecStart=/opt/gomessage/gomessage     #此处是您GOMessage二进制文件的绝对路径（可以替换成您自己的实际文件位置）
Restart=no


[Install]
WantedBy=multi-user.target

EOF
```

以上文件写入并保存好之后，执行如下命令重新加载守护进程列表：

```bash
systemctl daemon-reload
```

到此为止，您就已经为GoMessage服务封装好了一个健壮的守护进程，接下来您可以通过如下命令来操作和管理GoMessage服务：

```bash
systemctl start   gomessage.service   #启动gomessage服务
systemctl stop    gomessage.service   #停止gomessage服务
systemctl restart gomessage.service   #重启gomessage服务
systemctl status  gomessage.service   #查看gomessage服务的运行状态
```

设置GoMessage服务的开机启动：

```bash
systemctl enable  gomessage.service   #设置gomessage服务为开机启动
systemctl disable gomessage.service.  #取消gomessage服务的开机启动
```

<br><br>

## 重启服务

执行如下命令启动GoMessage服务

```bash
systemctl restart gomessage.service
```

假定此时您部署GoMessage的服务器地址为`192.168.33.201`；

您打开浏览器访问 `http://192.168.33.201:7077` 这个地址地址，就可以成功的访问GoMessage的服务啦~

<br><br>

## 访问界面

第一次访问GoMessage的页面显示如下（这是v1.0版的界面，不排除以后的界面可能会有变化）

![image-20211223161137899](https://img.taycc.com/picgo/image-20211223161137899.png)

(手动部署的内容到此结束)

<br><br><br><br>
