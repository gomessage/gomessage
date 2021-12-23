#!/bin/bash


systemctl stop gomessage.service
echo "检测与停止GoMessage服务..."

systemctl disable gomessage.service
echo "从Daemon中注销GoMessage服务..."

systemctl daemon-reload
echo "刷新守护进程列表..."

rm -rf /opt/gomessage
rm -rf /usr/lib/systemd/system/gomessage.service
echo "删除GoMessage服务所需的全部依赖文件..."

echo "GoMessage服务卸载完成..."

