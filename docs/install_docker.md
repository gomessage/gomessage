# Docker方式（强烈推荐）：


拉取指定版本的镜像到本地：
```bash
docker pull gomessage/gomessage:latest
```

<br>


前台启动，只运行一次 (容器停止后自动删除，不残留和污染本地环境)：

```bash
docker run --rm \
    -p 7077:7077 \
    gomessage/gomessage:latest 
```

<br>

后台运行，只运行一次 (容器停止后自动删除，不残留和污染本地环境)：

```bash
docker run -d \
    -p 7077:7077 \
    --rm \
    --name=gomessage \
    gomessage/gomessage:latest
```

<br>

稳定的运行 (且设定为开机启动)：

```bash
docker run -d \
    -p 7077:7077 \
    --restart=always \
    --name=gomessage \
    gomessage/gomessage:latest
```

<br>

稳定的运行 (且设定为开机启动)；同时保存数据文件到宿主机上，下次使用新版镜像时，原来的数据不丢失

```bash
mkdir -p /opt/gomessage/data

docker run -d \
    -p 7077:7077 \
    -v /opt/gomessage/data:/opt/gomessage/data \
    --restart=always \
    --name=gomessage \
    gomessage/gomessage:latest
```
