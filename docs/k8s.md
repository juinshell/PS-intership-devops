# Service
## 现状
每个pod都有各自的ip。Pod 是非永久性资源。
Deployment 中，在同一时刻运行的 Pod 集合可能与稍后运行该应用程序的 Pod 集合不同。
## 创建和查看Service
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
kubectl describe deployments
```
