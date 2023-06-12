# 使用原生yaml脚本部署
> 这里建议您把GoMessage服务部署成为一个StatefulSet服务；     
> 若您只是临时测试，不担心GoMessage的数据丢失，则部署为Deployment也未尝不可。
>
> 完整的投产，需要部署四个服务：
> - 部署一个StatefulSet服务
> - 创建一个PVC存储卷
> - 部署一个Service无头服务
> - 部署一个Ingress，把`GoMessage服务`开放到集群外部可以被访问到。

<br>

StatefulSet.yaml
```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: gomessage
  namespace: default
spec:
  serviceName: gomessage
  replicas: 2
  selector:
    matchLabels:
      app: gomessage
  template:
    metadata:
      labels:
        app: gomessage
    spec:
      containers:
        - name: gomessage
          image: gomessage/gomessage:latest
          ports:
            - containerPort: 7077
              name: web
          volumeMounts:
            - name: pvc-gomessage
              mountPath: /opt/gomessage/data
      #所有的pod都共用这一个pvc；需要单独创建好这个pvc
      volumes:
        - name: pvc-gomessage
          persistentVolumeClaim:
            claimName: pvc-gomessage

```

<br>

Pvc.yaml
```yaml
---
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: pvc-gomessage
  namespace: default
spec:
  accessModes:
    - ReadWriteMany    #由于要在多个pod之间共享元数据信息，这里必须选择"ReadWriteMany"，有的云厂商提供的存储类型不让选这个，那就找一个nfs类型的存储类，绝对是可以的。
  resources:
    requests:
      storage: 100Mi
  storageClassName: nfs-client  #这里最好选择nfs类型的存储类，因为各大云厂商的存储类都有最低容量限制，比如腾讯云cbs类型的存储类最低要求pvc至少得创建10Gi，但是GoMessage真的只需要100Mi就行了用不了那么多，只是存放元数据信息在多个pod之间共享。如果您存储资源比较充足那么选用其它类型的存储类也是可以的。

```

<br>

Service.yaml
```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: svc-gomessage
  namespace: default
  labels:
    app: svc-gomessage
spec:
  selector:
    app: gomessage
  ports:
    - port: 80
      targetPort: 7077
      name: svc-gomessage
      protocol: TCP
  clusterIP: None

```

<br>

Ingress.yaml
```yaml
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-gomessage
  namespace: default
spec:
  ingressClassName: nginx   #请注意低版本的k8s中没有"ingress classes"的概念，如果您的k8s是低版本，那去网上找找对应的yaml文件，其它的东西都大差不差。
  rules:
    - host: "gomessage.linux.com"
      http:
        paths:
          - pathType: Prefix
            path: "/"
            backend:
              service:
                name: svc-gomessage
                port:
                  number: 80

```
