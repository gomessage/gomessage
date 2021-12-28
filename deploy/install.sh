#!/usr/bin/env bash

##############################
# 安装服务
##############################
install_gomessage() {
    systemctl stop gomessage.service
    mkdir -p /opt/gomessage
    rm -rf /opt/gomessage/*
    tar zxvf ./gomessage.tar.gz -C /opt/gomessage/
    chown -R root:root /opt/gomessage


    cat >/usr/lib/systemd/system/gomessage.service <<EOF
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
ExecStart=/opt/gomessage/gomessage
Restart=no


[Install]
WantedBy=multi-user.target
EOF

    #如果有检测到用户的备份数据库存在，也直接使用备份数据库启动新版本的gomessage，否则啥也不干
    if (ls /tmp/gomessage_cache | grep db.sqlite3); then
        cp -rf /tmp/gomessage_cache/db.sqlite3 /opt/gomessage/conf/
        echo "复用已有的数据库..."
    else
        #如果没有备份数据库的存在，则啥也不干
        echo ""
    fi

    systemctl daemon-reload
    systemctl enable gomessage.service
    systemctl restart gomessage.service
}

##############################
# 显示帮助
##############################
help() {
    cat <<EOF

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

EOF
}

##############################
# 程序入口
##############################
if (systemctl list-unit-files | grep gomessage); then
    echo "gomessage服务已安装；无需再次进行安装..."
else
    if (ss -antp | grep "\[::\]:7077"); then
        echo "本机7077端口已经被占用，退出安装程序，请检测端口配置..."
        exit 1
    else
        install_gomessage
        help
    fi
fi
