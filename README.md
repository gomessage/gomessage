# GoMessage：报警信息转发器

# 安装

> (以Centos7.x环境安装为例进行演示)

### 下载安装包

打开安装包下载界面：https://gitee.com/tay3223/go-message/releases

```bash
# 下载安装包
wget https://gitee.com/tay3223/go-message/attach_files/912104/download/GoMessage-v1.0.0-linux-amd64.tar.gz

# 新建服务运行时的工作目录；
# 作者推荐把服务安装到：/opt/gomessage/目录下；
# 这样GoMessage服务所有内容都只在这一个目录中，不会污染其它目录。
mkdir -p /opt/gomessage

# 把安装包解压到上面创建好的工作目录中
tar - zxvf ./GoMessage-v1.0.0-linux-amd64.tar.gz -C /opt/gomessage/
```

### 安装

创建systemd守护进程，以该方式运行服务有诸多好处，小伙伴们可以去查阅了解一下~

新建以下文件：

```bash
/usr/lib/systemd/system/go_message.service
```

追加如如下内容：

```bash
[Unit]
Description=GoMessage
After=network.target


[Service]
Type=simple
User=root
Group=root
LimitNOFILE=65535
LimitNPROC=65535
WorkingDirectory=/opt/gomessage/
ExecStart=/opt/gomessage/GoMessage
Restart=no


[Install]
WantedBy=multi-user.target
```

刷新守护进程列表：

```bash
systemctl daemon-reload
```

到此为止就算安装成功了，常用的启动命令如下：

```bash
systemctl start   go_message.service
systemctl stop    go_message.service
systemctl restart go_message.service
systemctl status  go_message.service
```

# 如何使用

(有空了再写...)