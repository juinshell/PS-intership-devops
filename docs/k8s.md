# Service
## 现状
每个pod都有各自的ip。Pod 是非永久性资源。
Deployment 中，在同一时刻运行的 Pod 集合可能与稍后运行该应用程序的 Pod 集合不同。
## 创建和查看Service
有选择算符，直接访问pod
```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: MyApp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
```
![](service1.png)
* 当每个 Service 创建时，会被分配一个**唯一的 IP 地址（也称为 clusterIP）**。 这个 IP 地址与一个 Service 的生命周期绑定在一起，当 Service 存在的时候它也不会改变。 可以配置 Pod 使它与 Service 进行通信，Pod 知道与 Service 通信将被自动地负载均衡到该 Service 中的某些 Pod 上。
* targetPort：容器接收流量的端口；
* port：抽象的 Service 端口，可以使任何其它 Pod 访问该 Service 的端口
* Service 能够将一个接收 port 映射到任意的 targetPort。 默认情况下，targetPort 将被设置为与 port 字段相同的值。
# Deployment（Pod模板之一）
## 动机
描述 Deployment 中的 目标状态，而 Deployment 控制器（Controller） 以受控速率更改实际状态， 使其变为期望状态。
## 更新部署
当且仅当 Deployment 的 Pod 模板（即.spec.template）发生更改时，才会触发 Deployment 的推出，例如，如果模板的标签或容器映像已更新。其他更新，例如扩展部署，不会触发推出。
## 查看部署推出状态
```
kubectl rollout status deployment
```
## 部署会控制pod的数量
部署确保在更新时只有一定数量的 Pod 关闭。默认情况下，它确保至少 75% 的所需 Pod 数量已启动（最多 25% 不可用）。

部署还确保仅创建超过所需数量的 Pod 的特定数量的 Pod。默认情况下，它确保最多 125% 的所需 Pod 数量增加（最大激增 25%）。
## 获取部署的详细信息
```
kubectl describe deployments xxx
```
# Pod
## 通信
在同一个 Pod 内，所有容器共享一个 IP 地址和端口空间，并且可以通过 localhost 发现对方。
# 端口转发
只能转发pod的端口(包括pod模板)，对于Service是针对于内部互访(可用域名访问)