#!/bin/bash

##############################
# 安装服务
##############################
install_go_message() {
    mkdir -p /opt/gomessage
    rm -rf /opt/gomessage/*

    mv ./conf /opt/gomessage/
    mv ./GoMessage /opt/gomessage/
    mv ./static /opt/gomessage/
    mv ./swagger /opt/gomessage/
    mv ./views /opt/gomessage/

    cat > /usr/lib/systemd/system/go_message.service <<EOF
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
EOF

    systemctl daemon-reload
    systemctl enable go_message.service
    systemctl restart go_message.service
}

##############################
# 显示帮助
##############################
help() {
    cat <<EOF

GoMessage服务安装成功，默认监听端口为：http://0.0.0.0:8080


服务常用命令如下：
systemctl start   go_message.service
systemctl status  go_message.service
systemctl stop    go_message.service
systemctl restart go_message.service

设置开机启动：
systemctl enable  go_message.service

##############################
# 提示
##############################
（√）1.服务设置为开机启动
（√）2.服务已经启动
（√）3.服务监听8080端口

EOF
}

##############################
# 程序入口
##############################
if (systemctl list-unit-files | grep go_message); then
    echo "go_message服务已安装；无需再次进行安装..."
else
    if (ss -antp | grep "\[::\]:8080"); then
        echo "本机8080端口已经被占用，退出安装程序，请检测端口配置..."
        exit 1
    else
        install_go_message
        help
    fi
fi
