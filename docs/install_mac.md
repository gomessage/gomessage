# Mac电脑上如何使用GoMessage

### 以Docker的方式来运行GoMessage（强烈推荐）

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

### 以二进制文件的方式来运行GoMessage
> 作者此处默认选项：
> 
> 愿意使用"二进制文件运行"的方式来在"本地电脑"上运行GoMessage的小伙伴，拥有丰富的Linux命令行操作经验。因此不对此过程进行过多赘述~

#### 二进制包获取地址：

安装包下载地址（国内）：https://gitee.com/gomessage/gomessage/releases （这个地址国内访问速度快~）

安装包下载地址（国外）：https://github.com/gomessage/gomessage/releases


解压之后得到的文件目录结构为：
```bash
./
├── conf          #这里面存放了GoMessage服务的全局配置，比如服务端口修改什么的
├── gomessage     #这个是GoMessage服务的二进制可执行程序，未来我们要启动的就是这个软件
├── static        #GoMessage服务依赖的前端静态文件
└── views         #GoMessage存放html页面的目录
```

此时保持以上目录结构不变，为 ./gomessage 赋予可执行权限即可，然后执行它即可~

<br><br><br>