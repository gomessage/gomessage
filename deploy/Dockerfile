#############################################################################
#
#   如果您想以Docker的方式运行GoMessage服务，只需在Dockerfile文件所在的当前目录中执行以下命令：
#  （所有的目录层级和文件位置都是被预先设计好的，您无需担心会出现不可预知的报错内容）
#
#   # <封装image> 把gomessage服务封装为一个docker image，命令如下
#   docker build -t gomessage:latest .
#
#   # <使用image> 启动这个docker image，命令如下
#   docker run -d -p 7077:7077 gomessage:latest
#
#   ---------------------- 分割线 ----------------------
#
#   # <使用作者的image> 如果您不想重新封装docker image，可以使用作者提供的 <已经封装好的> docker image来使用，命令如下：
#   docker run -d -p 7077:7077 taycc/gomessage
#
#############################################################################

FROM golang:1.15.15-alpine

RUN mkdir -p /opt/gomessage && rm -rf /opt/gomessage/*

WORKDIR /opt/gomessage

ADD ./gomessage.tar.gz /opt/gomessage/

EXPOSE 7077

ENTRYPOINT ["/opt/gomessage/gomessage"]