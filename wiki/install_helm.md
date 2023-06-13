# 使用Helm部署

如果您的k8s版本低于1.26.x；

那么作者已经封装好的helm chart可能您无法使用;

因为在1.26.x版本中Kubernetes官方调整了部分api内容；

此时您可以自己封装一个老版本的helm chart（或）喊作者帮您封装（作者微信：`SPE3SRU3STAY`）；

或者您也可以直接查看`方式二：使用原生yaml脚本部署`的方式来进行操作，同样也可以快速的投产GoMessage。

```bash
#添加 GoMessage Helm Chart 仓库
helm repo add gomessage "https://occos-helm.pkg.coding.net/repos/gomessage"

#更新helm索引
helm repo update

#安装gomessage到您的集群中
helm install \
    gomessage-service \
    gomessage/gomessage \
    --namespace="default"
    
```

- `gomessage-service`：代表GoMessage服务，安装到您的k8s集群里时的名字，您可以自定义。
- `gomessage/gomessage`：是GoMessage服务的官方仓库。
- `--namespace`：代表您这个将会把服务安装到哪个命名空间下（如果不需要这个参数，那么会默认安装到`default`命名空间下）。
