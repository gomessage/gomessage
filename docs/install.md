# Linux环境安装
GoMessage当前支持`Linux`和`Mac`系统的部署

<br><br>

## 安装包下载：

安装包下载地址（国内）：https://gitee.com/gomessage/gomessage/releases （这个地址国内访问速度快~）

安装包下载地址（国外）：https://github.com/gomessage/gomessage/releases

> 目前提供有Linux、Mac、Windows三个版本的安装包，如需支持其它版本，可以自行下载源码编译，也可以直接联系作者编译出指定版本的安装包。

<br><br>

## 一键自动化部署（推荐）

本项目提供Liunx裸机运行环境上的`一键自动化部署脚本`，强烈推荐小伙伴们使用脚本进行服务安装。

```bash
# 下载安装包到本机（安装包下载链接可以从上文中的下载页面获取）
wget https://github.com/gomessage/gomessage/releases/download/版本号/gomessage-版本号-linux-amd64.tar.gz

# 解压安装包
tar -zxvf gomessage-版本号-linux-amd64.tar.gz

# 进入aaa目录
cd ./gomessage-版本号-linux-amd64/

# 可以看到目录结构体如下
./
├── gomessage.tar.gz     #安装包文件
├── install.sh           #一键安装自动化部署脚本
└── uninstall.sh         #一键卸载自动化部署脚本

# 执行自动化安装脚本
bash ./install.sh
```

./install.sh脚本运行完成后，出现以下内容代表一键自动化安装成功

```bash
Created symlink from /etc/systemd/system/multi-user.target.wants/gomessage.service to /usr/lib/systemd/system/gomessage.service.

服务常用命令如下：
systemctl restart gomessage.service
systemctl start   gomessage.service
systemctl stop    gomessage.service
systemctl status  gomessage.service

设置开机启动：
systemctl enable  gomessage.service

##############################
# 提示
##############################
（√）1.服务设置为开机启动
（√）2.服务已经启动
（√）3.服务监听7077端口
（√）4.服务安装目录：/opt/gomessage/

GoMessage服务安装成功，默认监听端口为：http://0.0.0.0:7077

```

假定此时您部署GoMessage的服务器地址为`192.168.88.33`，您打开浏览器访问`http://192.168.88.33:7077`地址，就可以成功访问GoMessage的服务啦~



第一次访问GoMessage的页面显示如下：

![image-20211223161137899](https://img.taycc.com/picgo/image-20211223161137899.png)

(自动化一键部署的内容到此结束)

<br><br><br><br><br>

## 手动部署（详细操作步骤）

> 以Centos7.x环境来举例



## 安装包下载：

安装包下载地址（国内）：https://gitee.com/gomessage/gomessage/releases （这个地址国内访问速度快~）

安装包下载地址（国外）：https://github.com/gomessage/gomessage/releases

```bash
# 进入到linux命令行的家目录中
cd ~/

# 下载安装包
wget https://github.com/gomessage/gomessage/releases/download/版本号/gomessage-版本号-linux-amd64.tar.gz
```



#### 创建工作目录

> （作者推荐您将工作目录设置在 /opt/gomessage/ 目录中，可以干干清爽一些）

```bash
# 创建一个新目录/opt/gomessage
mkdir -p /opt/gomessage
```



#### 解压安装包

解压前面下载好的安装包

```bash
tar -zxvf ./gomessage-1.0.9-linux-amd64.tar.gz


# 解压之后可以看到的目录结构如下：
./
├── gomessage.tar.gz     #安装包文件
├── install.sh           #一键安装自动化部署脚本
└── uninstall.sh         #一键卸载自动化部署脚本
```



此时因为咱们要`手动安装`GoMessage服务，因此可以直接忽略掉作者预留的`install.sh`和`uninstall.sh`这两个脚本文件。直接解压`gomessage.tar.gz`这个压缩文件来得到真正的安装包：

```bash
# 解压gomessage.tar.gz压缩包的内容到/opt/gomessage/目录下
tar -zxvf ./gomessage.tar.gz -C /opt/gomessage/

#解压之后的目录结构如下：
/opt/gomessage/
├── conf          #这里面存放了GoMessage服务的全局配置，比如服务端口修改什么的
├── gomessage     #这个是GoMessage服务的二进制可执行程序，未来我们要启动的就是这个软件
├── static        #GoMessage服务依赖的前端静态文件
├── swagger       #GoMessage对外提供的API接口相关的静态文件
└── views         #GoMessage存放html页面的目录
```



到这一步为止，GoMessage就算是部署好了，直接运行一下命令就可以运行起来GoMessage服务：

```bash
# 对，您没看错，就是这么一行命令，什么参数都不需要
/opt/gomessage/gomessage
```

此时GoMessage默认会启动和监听本机的`7077`端口，您只需要用浏览器访问：`http://<服务器地址>:7077/`就可以访问到GoMessage服务了。



> 但是用上面挂载前台的方式启动GoMessage服务（或其它任意二进制命令形式的服务）都是不可靠的，Linux会有一套自己的QOS机制来管理自身的资源分配和回收。因此作者推荐您将GoMessage服务封装为Systemd守护进程的方式来运行。



### 封装systemd守护进程（以守护进程的方式来运行GoMessage）

新建一个文件：

```bash
vim  /usr/lib/systemd/system/gomessage.service



#把以下内容写入到该文件中

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

#只有上文这一小段写入到指定的文件中之后，别的就不用再额外写入了，除了目录修为为您自己的实际路径之后，别的部分您都可以直接抄写作者提供的默认配置。
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



#### 启动服务

执行如下命令启动GoMessage服务

```bash
systemctl restart gomessage.service
```

假定此时您部署GoMessage的服务器地址为`192.168.88.33`，您打开浏览器访问`http://192.168.88.33:7077`地址，就可以成功访问GoMessage的服务啦~



第一次访问GoMessage的页面显示如下：

![image-20211223161137899](https://img.taycc.com/picgo/image-20211223161137899.png)

(手动部署的内容到此结束)


<br><br><br><br>
