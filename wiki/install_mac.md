# Mac电脑上如何使用GoMessage

### 以Docker的方式来运行GoMessage（强烈推荐）

快速启动：

```bash
docker run -d -p 7077:7077 gomessage/gomessage:latest
```

只运行一次（容器停止后自动删除，不残留和污染本地环境）：
```bash
docker run -d \
    -p 7077:7077 \
    --rm \
    --name=gomessage \
    gomessage/gomessage:latest
```

稳定的运行（且设定为开机启动）：
```bash
docker run -d \
    -p 7077:7077 \
    --restart=always \
    --name=gomessage \
    gomessage/gomessage:latest
```

<br>

### 以二进制文件的方式来运行GoMessage
> 作者在此处默认：
>
> 愿意使用<二进制文件运行>的方式来在<本地电脑>上运行GoMessage的玩家，拥有丰富的Linux命令行操作使用经验，具备识别<命令行危险操作>的基本常识，因此不对该方式的部署进行过多赘述~

<br>

#### 二进制包的获取地址：

下载地址（国内）：https://gitee.com/gomessage/gomessage/releases （国内用这个，速度很快~）

下载地址（国外）：https://github.com/gomessage/gomessage/releases


解压之后得到的文件目录结构为：

```bash
./
├── assets        #GoMessage服务依赖的前端静态文件
├── config        #这里面存放了GoMessage服务的全局配置，比如服务端口修改什么的
├── gomessage     #这个是GoMessage服务的二进制可执行程序，未来我们要启动的就是这个软件
```

<br>

- 首先，保持以上目录结构不变
- 接着，为 ./gomessage 赋予可执行权限
- 然后，在终端中执行 ./gomessage 命令即可启动服务~ （本地访问端口默认为：7077）
- 最后，封装成Mac可以理解的守护进程，就算大功告成啦~

<br><br><br>
