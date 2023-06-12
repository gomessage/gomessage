# 使用原生yaml脚本部署
> 这里建议您把GoMessage服务部署成为一个StatefulSet服务；     
> 若您只是临时测试，不担心GoMessage的数据丢失，则部署为Deployment也未尝不可。
>
> 完整的投产，需要部署四个服务：
> - 部署一个StatefulSet服务
> - 创建一个PVC存储卷
> - 部署一个Service无头服务
> - 部署一个Ingress，把`GoMessage服务`开放到集群外部可以被访问到。

具体的脚本内容：`有空再补充...`
